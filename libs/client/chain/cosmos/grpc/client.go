package cosmos

import (
	"context"
	cosmoserrors "cosmossdk.io/errors"
	"encoding/hex"
	"encoding/json"
	"fmt"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	kmultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/gogoproto/proto"
	cosmostypes "github.com/curtis0505/bridge/libs/client/chain/cosmos/types"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
	"math/big"
	"sync"
	"time"

	// common
	"github.com/curtis0505/bridge/libs/client/chain/conf"
	clienttypes "github.com/curtis0505/bridge/libs/client/chain/types"
	cosmoscommon "github.com/curtis0505/bridge/libs/common/cosmos"
	"github.com/curtis0505/bridge/libs/types"
	tokentypes "github.com/curtis0505/bridge/libs/types/token"
	"github.com/curtis0505/bridge/libs/util"

	// finschia
	"github.com/curtis0505/grpc-idl/finschia/foundation"
	ostraconservice "github.com/curtis0505/grpc-idl/finschia/tmservice"

	// cosmos
	mintv1beta1 "cosmossdk.io/api/cosmos/mint/v1beta1"
	tmtypes "github.com/cometbft/cometbft/proto/tendermint/types"
	cosmosclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	// ibc
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
)

var (
	_ clienttypes.Client        = &client{}
	_ clienttypes.StakingClient = &client{}
	_ clienttypes.IBCClient     = &client{}
	_ clienttypes.CosmosClient  = &client{}
	_ clienttypes.BridgeClient  = &client{}
)

type client struct {
	cc *grpc.ClientConn

	*sync.RWMutex

	chain              string
	chainName          string
	chainId            string
	nodeService        node.ServiceClient            // cosmos service
	tmService          tmservice.ServiceClient       // cosmos service
	txService          txtypes.ServiceClient         // cosmos service
	authClient         authtypes.QueryClient         // cosmos module
	bankClient         banktypes.QueryClient         // cosmos module
	stakingClient      stakingtypes.QueryClient      // cosmos module
	distributionClient distributiontypes.QueryClient // cosmos module
	mintClient         mintv1beta1.QueryClient       // cosmos module pulsar
	foundationClient   foundation.QueryClient        // finschia module
	ibcTransferClient  ibctransfertypes.QueryClient  // ibc module
	ibcChannelClient   ibcchanneltypes.QueryClient   // ibc module
	wasmClient         wasmtypes.QueryClient         // wasm module

	interfaceRegistry codectypes.InterfaceRegistry
	txConfig          cosmosclient.TxConfig
	txDecoder         cosmossdk.TxDecoder

	// subscribe 관련 변수
	finalizedBlockCount int
	executedTxHash      map[string]int64
}

func NewClient(config conf.ClientConfig) (clienttypes.CosmosClient, error) {
	c := &client{
		RWMutex:        &sync.RWMutex{},
		chain:          config.Chain,
		executedTxHash: map[string]int64{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var err error
	c.cc, err = grpc.DialContext(ctx, config.Url, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, types.WrapError("DialContext", err)
	}

	// tendermint
	if c.chain == types.ChainFNSA || c.chain == types.ChainTFNSA {
		c.interfaceRegistry = cosmostypes.NewInterfaceRegistry(
			cosmostypes.WithCosmosRegistry(), cosmostypes.WithFinschiaRegistry(), cosmostypes.WithIBCRegistry(), cosmostypes.WithWasmRegistry(),
		)
		c.tmService = ostraconservice.NewServiceClient(c.cc)
		c.foundationClient = foundation.NewQueryClient(c.cc)
	} else {
		c.interfaceRegistry = cosmostypes.NewInterfaceRegistry(
			cosmostypes.WithCosmosRegistry(), cosmostypes.WithIBCRegistry(), cosmostypes.WithWasmRegistry(),
		)
		c.tmService = tmservice.NewServiceClient(c.cc)
	}

	c.nodeService = node.NewServiceClient(c.cc)
	c.txService = txtypes.NewServiceClient(c.cc)

	c.authClient = authtypes.NewQueryClient(c.cc)
	c.bankClient = banktypes.NewQueryClient(c.cc)
	c.stakingClient = stakingtypes.NewQueryClient(c.cc)
	c.distributionClient = distributiontypes.NewQueryClient(c.cc)
	c.mintClient = mintv1beta1.NewQueryClient(c.cc)

	c.ibcTransferClient = ibctransfertypes.NewQueryClient(c.cc)
	c.ibcChannelClient = ibcchanneltypes.NewQueryClient(c.cc)
	c.wasmClient = wasmtypes.NewQueryClient(c.cc)

	c.txConfig = authtx.NewTxConfig(codec.NewProtoCodec(c.interfaceRegistry), authtx.DefaultSignModes)
	c.txDecoder = c.txConfig.TxDecoder()

	block, err := c.tmService.GetLatestBlock(ctx, &tmservice.GetLatestBlockRequest{})
	if err != nil {
		return nil, types.WrapError("GetLatestBlock", err)
	}
	c.chainId = block.GetBlock().GetHeader().ChainID

	return c, nil
}

// common

func ProxyClient(proxy clienttypes.Proxy, chain string) (clienttypes.CosmosClient, error) {
	c := proxy.ProxyClient(chain)
	if c == nil {
		return nil, fmt.Errorf("not found proxy")
	}

	client, ok := c.(clienttypes.CosmosClient)
	if !ok {
		return nil, fmt.Errorf("failed to casting client")
	}

	return client, nil
}

func (c *client) ChainName() string          { return c.chainName }
func (c *client) Chain() string              { return c.chain }
func (c *client) ChainType() types.ChainType { return types.GetChainType(c.chain) }

func (c *client) PendingNonceAt(ctx context.Context, address string) (uint64, error) {
	in := authtypes.QueryAccountRequest{
		Address: address,
	}
	resp, err := c.authClient.Account(ctx, &in)
	if err != nil {
		return 0, err
	}
	var account authtypes.AccountI
	if err != nil {
		return 0, err
	}
	err = c.interfaceRegistry.UnpackAny(resp.Account, &account)
	if err != nil {
		return 0, err
	}

	return account.GetSequence(), nil
}

func (c *client) NonceAt(ctx context.Context, address string) (uint64, error) {
	return c.PendingNonceAt(ctx, address)
}
func (c *client) GetChainID(ctx context.Context) (*big.Int, error) {
	return big.NewInt(0), nil
}
func (c *client) NetworkId(ctx context.Context) (*big.Int, error) {
	return nil, clienttypes.NotImplemented
}
func (c *client) GasPrice(ctx context.Context) (*big.Int, error) {
	return big.NewInt(0), clienttypes.NotImplemented
}
func (c *client) RawTxAsync(ctx context.Context, rawTx []byte, rawProxyRequest []byte) (*clienttypes.SendTxAsyncResult, error) {
	return c.BroadcastRawTx(ctx, rawTx, txtypes.BroadcastMode_BROADCAST_MODE_SYNC)
}

func (c *client) RawTxAsyncByTx(ctx context.Context, tx *types.Transaction) (*clienttypes.SendTxAsyncResult, error) {
	cosmosTx, err := tx.CosmosTransaction()
	if err != nil {
		return nil, err
	}

	raw, err := proto.Marshal(cosmosTx)
	if err != nil {
		return nil, err
	}

	return c.BroadcastRawTx(ctx, raw, txtypes.BroadcastMode_BROADCAST_MODE_SYNC)
}

func (c *client) BroadcastRawTx(ctx context.Context, rawTx []byte, mode txtypes.BroadcastMode) (*clienttypes.SendTxAsyncResult, error) {
	in := txtypes.BroadcastTxRequest{
		TxBytes: rawTx,
		Mode:    mode,
	}
	resp, err := c.txService.BroadcastTx(ctx, &in)
	if err != nil {
		return nil, err
	}

	var result clienttypes.SendTxAsyncResult
	if resp.GetTxResponse().Code == 0 {
		result.Result = clienttypes.SendTxResultType_Success
		result.Hash = resp.GetTxResponse().TxHash
		return &result, nil
	}

	return nil, cosmoserrors.ABCIError(resp.GetTxResponse().Codespace, resp.GetTxResponse().Code, "RawTxAsync")
}

func (c *client) GetHeaderByHash(ctx context.Context, txHash string) (*types.Header, error) {
	return nil, clienttypes.NotImplemented
}

func (c *client) BlockNumber(ctx context.Context) (*big.Int, error) {
	in := tmservice.GetLatestBlockRequest{}
	block, err := c.tmService.GetLatestBlock(ctx, &in)
	if err != nil {
		return nil, err
	}

	return big.NewInt(block.GetBlock().GetHeader().Height), nil
}

func (c *client) GetTransaction(ctx context.Context, txHash string) (*types.Transaction, error) {
	in := txtypes.GetTxRequest{
		Hash: txHash,
	}
	tx, err := c.txService.GetTx(ctx, &in)
	if err != nil {
		return nil, err
	}

	err = tx.GetTx().UnpackInterfaces(c.interfaceRegistry)
	if err != nil {
		return nil, err
	}

	err = tx.GetTxResponse().UnpackInterfaces(c.interfaceRegistry)
	if err != nil {
		return nil, err
	}
	return types.NewTransactionWithReceipt(tx.GetTx(), tx.GetTxResponse(), c.chain), nil
}

func (c *client) GetTransactionWithReceipt(ctx context.Context, txHash string) (*types.Transaction, bool, error) {
	transaction, err := c.GetTransaction(ctx, txHash)
	return transaction, false, err
}

func (c *client) GetTransactionReceipt(ctx context.Context, txHash string) (*types.Receipt, error) {
	in := txtypes.GetTxRequest{
		Hash: txHash,
	}
	tx, err := c.txService.GetTx(ctx, &in)
	if err != nil {
		return nil, err
	}

	err = tx.GetTx().UnpackInterfaces(c.interfaceRegistry)
	if err != nil {
		return nil, err
	}

	err = tx.GetTxResponse().UnpackInterfaces(c.interfaceRegistry)
	if err != nil {
		return nil, err
	}

	return types.NewReceipt(tx.GetTxResponse(), c.chain), nil
}

func (c *client) BalanceAt(ctx context.Context, address string, blockNumber *big.Int) (*big.Int, error) {
	return c.Balance(ctx, address, tokentypes.DenomByChain(c.chain))
}

func (c *client) Balances(ctx context.Context, address string) (cosmossdk.Coins, error) {
	in := banktypes.QueryAllBalancesRequest{
		Address: address,
	}
	resp, err := c.bankClient.AllBalances(ctx, &in)
	if err != nil {
		return nil, err
	}
	return resp.Balances, nil
}

func (c *client) Balance(ctx context.Context, address, denom string) (*big.Int, error) {
	in := banktypes.QueryBalanceRequest{
		Address: address,
		Denom:   denom,
	}
	resp, err := c.bankClient.Balance(ctx, &in)
	if err != nil {
		return nil, err
	}

	return resp.GetBalance().Amount.BigInt(), nil
}

func (c *client) TxAsync(ctx context.Context, rlpTx string, proxyRequest clienttypes.ProxyRequest) (*clienttypes.SendTxAsyncResult, error) {
	b, err := hex.DecodeString(rlpTx)
	if err != nil {
		return nil, err
	}

	return c.RawTxAsync(ctx, b, nil)
}

// ----------
// COSMOS depends
// ----------

func (c *client) TxDecoder() cosmossdk.TxDecoder                  { return c.txDecoder }
func (c *client) TxConfig() cosmosclient.TxConfig                 { return c.txConfig }
func (c *client) InterfaceRegistry() codectypes.InterfaceRegistry { return c.interfaceRegistry }
func (c *client) ChainId() string                                 { return c.chainId }

// GetMinimumGasPrice returns minimum gas price from node
func (c *client) GetMinimumGasPrice(ctx context.Context) (float64, error) {
	resp, err := c.nodeService.Config(ctx, &node.ConfigRequest{})
	if err != nil {
		return 0.0, err
	}

	if resp.GetMinimumGasPrice() == "" {
		return 0.0, nil
	}

	coin, err := cosmossdk.ParseDecCoin(resp.GetMinimumGasPrice())
	if err != nil {
		return 0.0, err
	}

	return coin.Amount.Float64()
}

// GetBlockByNumber
// 0 == latest
// 0 < number = by number
func (c *client) GetBlockByNumber(ctx context.Context, number *big.Int) (*tmtypes.Block, error) {
	if number.Int64() == 0 {
		resp, err := c.tmService.GetLatestBlock(ctx, &tmservice.GetLatestBlockRequest{})
		if err != nil {
			return nil, err
		}
		return resp.GetBlock(), nil
	} else {
		resp, err := c.tmService.GetBlockByHeight(ctx, &tmservice.GetBlockByHeightRequest{
			Height: number.Int64(),
		})
		if err != nil {
			return nil, err
		}

		return resp.GetBlock(), nil
	}
}

// GetBlockByNumberWithTxs returns block with decoded txs
func (c *client) GetBlockByNumberWithTxs(ctx context.Context, number int64) (*tmtypes.Block, []*txtypes.Tx, error) {
	if c.chain == types.ChainFNSA || c.chain == types.ChainTFNSA {
		return nil, nil, fmt.Errorf("not supported")
	}

	resp, err := c.txService.GetBlockWithTxs(ctx, &txtypes.GetBlockWithTxsRequest{
		Height: number,
	})
	if err != nil {
		return nil, nil, err
	}

	return resp.GetBlock(), resp.GetTxs(), nil
}

func (c *client) GetAccountNumberAndSequence(ctx context.Context, address string) (uint64, uint64, error) {
	in := authtypes.QueryAccountRequest{
		Address: address,
	}
	resp, err := c.authClient.Account(ctx, &in)
	if err != nil {
		return 0, 0, err
	}
	var account authtypes.AccountI
	if err != nil {
		return 0, 0, err
	}
	err = c.interfaceRegistry.UnpackAny(resp.Account, &account)
	if err != nil {
		return 0, 0, err
	}

	return account.GetAccountNumber(), account.GetSequence(), nil
}

func (c *client) GetReward(ctx context.Context, address string) (int64, error) {
	in := distributiontypes.QueryDelegationTotalRewardsRequest{
		DelegatorAddress: address,
	}

	resp, err := c.distributionClient.DelegationTotalRewards(ctx, &in)
	if err != nil {
		return 0, err
	}

	reward := int64(0)
	for _, coin := range resp.GetTotal() {
		if coin.GetDenom() == tokentypes.DenomByChain(c.chain) {
			reward += util.ToBigInt(util.ToEtherWithDecimal(coin.Amount.BigInt(), 18)).Int64()
		}
	}

	return reward, nil
}

func (c *client) GetStakeAmount(ctx context.Context, address string) (*big.Int, error) {
	in := stakingtypes.QueryDelegatorDelegationsRequest{
		DelegatorAddr: address,
	}

	resp, err := c.stakingClient.DelegatorDelegations(ctx, &in)
	if err != nil {
		return big.NewInt(0), err
	}

	amount := int64(0)
	for _, delegation := range resp.GetDelegationResponses() {
		if delegation.GetBalance().Denom == tokentypes.DenomByChain(c.chain) {
			amount += delegation.GetBalance().Amount.Int64()
		}
	}
	return big.NewInt(amount), nil
}

func (c *client) GetWithdrawal(ctx context.Context, address string) ([]clienttypes.Withdrawal, error) {
	in := stakingtypes.QueryDelegatorUnbondingDelegationsRequest{
		DelegatorAddr: address,
	}

	withdrawal := make([]clienttypes.Withdrawal, 0)
	resp, err := c.stakingClient.DelegatorUnbondingDelegations(ctx, &in)
	if err != nil {
		return nil, err
	}

	for _, unbondingDelegations := range resp.GetUnbondingResponses() {
		for _, entries := range unbondingDelegations.Entries {
			withdrawal = append(withdrawal, clienttypes.Withdrawal{
				Chain:          c.chain,
				Amount:         entries.Balance.BigInt(),
				CompletionTime: entries.CompletionTime,
			})
		}
	}

	return withdrawal, nil
}

func (c *client) GetStakingParams(ctx context.Context) (stakingtypes.Params, error) {
	resp, err := c.stakingClient.Params(ctx, &stakingtypes.QueryParamsRequest{})
	if err != nil {
		return stakingtypes.Params{}, err
	}

	return resp.Params, nil
}

func (c *client) GetTxs(ctx context.Context, events ...string) (*txtypes.GetTxsEventResponse, error) {
	return c.txService.GetTxsEvent(ctx, &txtypes.GetTxsEventRequest{
		Events: events,
		Limit:  1000,
	})
}

func (c *client) GetStaking(ctx context.Context, address string) (*clienttypes.Staking, error) {
	reward, err := c.GetReward(ctx, address)
	if err != nil {
		return nil, err
	}

	withdrawal, err := c.GetWithdrawal(ctx, address)
	if err != nil {
		return nil, err
	}

	stakeAmount, err := c.GetStakeAmount(ctx, address)
	if err != nil {
		return nil, err
	}

	staking := &clienttypes.Staking{
		Amount:       stakeAmount,
		AmountLegacy: big.NewInt(0),
		Withdrawal:   withdrawal,
		Reward: clienttypes.Reward{
			Chain:  c.chain,
			Amount: big.NewInt(reward),
		},
	}

	return staking, nil
}

func (c *client) GetValidatorApr(ctx context.Context, validator string) (float64, error) {
	supplyOfResp, err := c.bankClient.SupplyOf(ctx, &banktypes.QuerySupplyOfRequest{
		Denom: tokentypes.DenomByChain(c.chain),
	})
	if err != nil {
		return 0, types.WrapError("SupplyOf", err)
	}

	poolResp, err := c.stakingClient.Pool(ctx, &stakingtypes.QueryPoolRequest{})
	if err != nil {
		return 0, types.WrapError("Pool", err)
	}

	inflationResp, err := c.mintClient.Inflation(ctx, &mintv1beta1.QueryInflationRequest{})
	if err != nil {
		return 0, types.WrapError("Inflation", err)
	}

	validatorResp, err := c.stakingClient.Validator(ctx, &stakingtypes.QueryValidatorRequest{
		ValidatorAddr: validator,
	})
	if err != nil {
		return 0, types.WrapError("Validator", err)
	}

	distributionParamResp, err := c.distributionClient.Params(context.Background(), &distributiontypes.QueryParamsRequest{})
	if err != nil {
		return 0, types.WrapError("Params", err)
	}

	var totalSupply, totalBonded, inflation decimal.Decimal
	var commissionRate, communityTax, foundationTax decimal.Decimal
	totalSupply = util.ToDecimal(supplyOfResp.Amount.Amount.BigInt(), 0)
	totalBonded = util.ToDecimal(poolResp.Pool.BondedTokens.BigInt(), 0)
	if i, err := validatorResp.Validator.Commission.Rate.Float64(); err != nil {
		return 0, err
	} else {
		commissionRate = util.ToDecimal(i, 0)
	}

	inflationDec := cosmossdk.Dec{}
	err = inflationDec.Unmarshal(inflationResp.Inflation)
	if err != nil {
		return 0, err
	}
	if i, err := inflationDec.Float64(); err != nil {
		return 0, err
	} else {
		inflation = util.ToDecimal(i, 0)
	}

	if i, err := distributionParamResp.Params.CommunityTax.Float64(); err != nil {
		return 0, err
	} else {
		communityTax = util.ToDecimal(i, 0)
	}

	if c.chain == types.ChainFNSA || c.chain == types.ChainTFNSA {
		paramResp, err := c.foundationClient.Params(context.Background(), &foundation.QueryParamsRequest{})
		if err != nil {
			return 0, err
		} else {
			if i, err := paramResp.Params.FoundationTax.Float64(); err == nil {
				foundationTax = util.ToDecimal(i, 0)
			} else {
				return 0, err
			}
		}
	} else {
		foundationTax = decimal.NewFromInt(0)
	}

	return util.EstimateCosmosStakingAnnualPercentageRate(
		totalSupply, totalBonded, inflation,
		commissionRate, communityTax, foundationTax,
	), nil
}

func (c *client) GetValidatorTvl(ctx context.Context, validator string) (*big.Int, error) {
	validatorResp, err := c.stakingClient.Validator(ctx, &stakingtypes.QueryValidatorRequest{
		ValidatorAddr: validator,
	})
	if err != nil {
		return big.NewInt(0), types.WrapError("Validator", err)
	}

	return validatorResp.Validator.BondedTokens().BigInt(), nil
}

func (c *client) GetValidatorDelegations(ctx context.Context, validator string) (stakingtypes.DelegationResponses, error) {
	var delegations stakingtypes.DelegationResponses
	var delegationsResp *stakingtypes.QueryValidatorDelegationsResponse
	var err error
	p := cosmostypes.DefaultPageRequest()
	for {
		delegationsResp, err = c.stakingClient.ValidatorDelegations(ctx, &stakingtypes.QueryValidatorDelegationsRequest{
			ValidatorAddr: validator,
			Pagination:    p,
		})
		if err != nil {
			return nil, err
		}

		delegations = append(delegations, delegationsResp.DelegationResponses...)
		if len(delegationsResp.GetPagination().GetNextKey()) == 0 {
			break
		}
		p.Key = delegationsResp.GetPagination().GetNextKey()
	}

	return delegations, nil
}

func (c *client) SetAccountPrefix() {
	c.Lock()
	defer c.Unlock()

	prefix := cosmoscommon.GetAddressPrefixByChain(c.chain)
	cosmossdk.GetConfig().SetBech32PrefixForAccount(prefix, prefix+"pub")
}

func (c *client) CallWasm(ctx context.Context, address string, msg types.CallWasm, v interface{}) error {
	bz, _ := cosmostypes.MarshalCallWasm(msg)

	resp, err := c.wasmClient.SmartContractState(ctx, &wasmtypes.QuerySmartContractStateRequest{
		Address:   address,
		QueryData: bz,
	})
	if err != nil {
		return err
	}

	if err = json.Unmarshal(resp.Data, &v); err != nil {
		return err
	}

	return nil
}

func (c *client) IBCDenomTrace(ctx context.Context, hash string) (*ibctransfertypes.DenomTrace, error) {
	resp, err := c.ibcTransferClient.DenomTrace(ctx, &ibctransfertypes.QueryDenomTraceRequest{
		Hash: hash,
	})
	if err != nil {
		return nil, err
	}
	return resp.DenomTrace, nil
}

func (c *client) SendTransaction(ctx context.Context, account *types.Account, opts ...cosmostypes.TxOptionFunc) (*clienttypes.SendTxAsyncResult, error) {
	return c.SendTransactionPrivKey(ctx, account.Secp256k1(), opts...)
}

func (c *client) SendTransactionPrivKey(ctx context.Context, priv *secp256k1.PrivKey, opts ...cosmostypes.TxOptionFunc) (*clienttypes.SendTxAsyncResult, error) {
	_, txBytes, err := c.BuildTransaction(ctx, priv, opts...)
	if err != nil {
		return nil, err
	}
	txOption := cosmostypes.NewTxOption().Apply(opts...)

	return c.BroadcastRawTx(ctx, txBytes, txOption.BroadcastMode)
}

func (c *client) SimulateTransaction(ctx context.Context, priv *secp256k1.PrivKey, opts ...cosmostypes.TxOptionFunc) (*cosmossdk.GasInfo, error) {
	_, txBytes, err := c.BuildTransaction(ctx, priv, opts...)
	if err != nil {
		return nil, err
	}

	resp, err := c.txService.Simulate(ctx, &txtypes.SimulateRequest{
		TxBytes: txBytes,
	})
	if err != nil {
		return nil, err
	}

	return resp.GetGasInfo(), nil
}

func (c *client) BuildTransaction(ctx context.Context, priv *secp256k1.PrivKey, opts ...cosmostypes.TxOptionFunc) (xauthsigning.Tx, []byte, error) {
	c.SetAccountPrefix()

	txOption := cosmostypes.NewTxOption().Apply(opts...)

	num, seq, err := c.GetAccountNumberAndSequence(ctx, cosmoscommon.FromPublicKeyUnSafe(c.chain, priv.PubKey().Bytes()).String())
	if err != nil {
		return nil, nil, err
	}
	txBuilder := c.txConfig.NewTxBuilder()
	err = txBuilder.SetMsgs(txOption.Msgs...)
	if err != nil {
		return nil, nil, err
	}
	txBuilder.SetFeeAmount(txOption.FeeAmount)
	txBuilder.SetGasLimit(txOption.GasLimit)

	if txOption.FeeGranter != "" {
		feeGranter, err := cosmossdk.AccAddressFromBech32(txOption.FeeGranter)
		if err != nil {
			return nil, nil, err
		}
		txBuilder.SetFeeGranter(feeGranter)
	}

	signRound1 := signing.SignatureV2{
		PubKey: priv.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  c.txConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: seq,
	}

	err = txBuilder.SetSignatures(signRound1)
	if err != nil {
		return nil, nil, err
	}

	signRound2, err := tx.SignWithPrivKey(
		c.txConfig.SignModeHandler().DefaultMode(),
		xauthsigning.SignerData{
			ChainID:       c.chainId,
			AccountNumber: num,
			Sequence:      seq,
		},
		txBuilder,
		priv,
		c.txConfig,
		seq,
	)
	err = txBuilder.SetSignatures(signRound2)
	if err != nil {
		return nil, nil, err
	}

	txBytes, err := c.txConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return nil, nil, err
	}

	return txBuilder.GetTx(), txBytes, nil
}

func (c *client) EstimateGas(ctx context.Context, callMsg types.CallMsg) (*big.Int, error) {
	panic("implement me")
}

func (c *client) SignMultiSig(ctx context.Context, account *types.Account, multiSig string, msg *wasmtypes.MsgExecuteContract) (signing.SignatureV2, error) {
	panic("implement me")
}

func (c *client) SendMultiSigTransaction(ctx context.Context, multiSigPubKey *kmultisig.LegacyAminoPubKey, msg *wasmtypes.MsgExecuteContract, singedTxs ...signing.SignatureV2) (*clienttypes.SendTxAsyncResult, error) {
	panic("implement me")
}

func (c *client) GetTransactionOption(ctx context.Context, from string) (*types.TransactionOption, error) {
	return nil, clienttypes.NotImplemented
}

func (c *client) GetTransactionData(data *types.RequestTransaction) (*types.Transaction, error) {
	return nil, clienttypes.NotImplemented
}
