package rest

import (
	"context"
	cosmoserrors "cosmossdk.io/errors"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	tmtypes "github.com/cometbft/cometbft/proto/tendermint/types"
	cosmosclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	"github.com/curtis0505/bridge/libs/client/chain/conf"
	cosmostypes "github.com/curtis0505/bridge/libs/client/chain/cosmos/types"
	clienttypes "github.com/curtis0505/bridge/libs/client/chain/types"
	cosmoscommon "github.com/curtis0505/bridge/libs/common/cosmos"
	"github.com/curtis0505/bridge/libs/types"
	tokentypes "github.com/curtis0505/bridge/libs/types/token"
	"github.com/curtis0505/bridge/libs/util"
	"github.com/curtis0505/grpc-idl/finschia/foundation"
	"github.com/shopspring/decimal"
	"math/big"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	gogoproto "github.com/cosmos/gogoproto/proto"
)

var (
	_ clienttypes.CosmosClient  = &client{}
	_ clienttypes.Client        = &client{}
	_ clienttypes.StakingClient = &client{}
	_ clienttypes.IBCClient     = &client{}
	_ clienttypes.BridgeClient  = &client{}
)

const (
	ModuleFinschia = "lbm"
	ModuleCosmos   = "cosmos"
	ModuleWasm     = "cosmwasm"

	v1Beta1 = "v1beta1"
	v1      = "v1"
)

type client struct {
	*http.Client
	*sync.RWMutex
	url       string
	chain     string
	chainName string
	chainId   string

	// subscribe 관련 변수
	finalizedBlockCount int
	executedTxHash      map[string]int64

	interfaceRegistry codectypes.InterfaceRegistry
	cdc               *codec.ProtoCodec
	txConfig          cosmosclient.TxConfig
	txDecoder         cosmossdk.TxDecoder
}

func NewClient(config conf.ClientConfig) (clienttypes.CosmosClient, error) {
	c := &client{
		RWMutex: &sync.RWMutex{},

		url:            config.Url,
		chain:          config.Chain,
		Client:         http.DefaultClient,
		executedTxHash: map[string]int64{},
	}

	if c.chain == types.ChainFNSA || c.chain == types.ChainTFNSA {
		c.interfaceRegistry = cosmostypes.NewInterfaceRegistry(cosmostypes.WithCosmosRegistry(), cosmostypes.WithFinschiaRegistry(), cosmostypes.WithIBCRegistry(), cosmostypes.WithWasmRegistry())
	} else {
		c.interfaceRegistry = cosmostypes.NewInterfaceRegistry(cosmostypes.WithCosmosRegistry(), cosmostypes.WithIBCRegistry(), cosmostypes.WithWasmRegistry())
	}
	c.txConfig = authtx.NewTxConfig(codec.NewProtoCodec(c.interfaceRegistry), authtx.DefaultSignModes)
	c.txDecoder = c.txConfig.TxDecoder()
	c.cdc = codec.NewProtoCodec(c.interfaceRegistry)

	block, err := c.GetBlockByNumber(context.Background(), big.NewInt(0))
	if err != nil {
		return nil, err
	}
	c.chainId = block.GetHeader().ChainID
	return c, nil
}

func (c *client) ChainName() string          { return c.chainName }
func (c *client) Chain() string              { return c.chain }
func (c *client) ChainType() types.ChainType { return types.GetChainType(c.chain) }

func (c *client) Invoke(ctx context.Context, v gogoproto.Message, path ...string) error {
	call, err := url.JoinPath(c.url, path...)
	if err != nil {
		return err
	}
	bz, err := util.GetRetry(ctx, call, nil, nil, 6)
	if err != nil {
		return err
	}
	if err = c.cdc.UnmarshalJSON(bz, v); err != nil {
		return c.CaptureError(bz, err)
	}
	return nil
}

func (c *client) InvokeWithParams(ctx context.Context, v gogoproto.Message, param url.Values, path ...string) error {
	call, err := url.JoinPath(c.url, path...)
	if err != nil {
		return err
	}

	bz, err := util.GetRetry(ctx, fmt.Sprintf("%s?%s", call, param.Encode()), nil, nil, 6)
	if err != nil {
		return err
	}
	if err = c.cdc.UnmarshalJSON(bz, v); err != nil {
		return c.CaptureError(bz, err)
	}
	return nil
}

func (c *client) InvokeWithOption(ctx context.Context, v gogoproto.Message, opt func([]byte) []byte, path ...string) error {
	call, err := url.JoinPath(c.url, path...)
	if err != nil {
		return err
	}

	bz, err := util.GetRetry(ctx, call, nil, nil, 6)
	if err != nil {
		return err
	}
	bz = opt(bz)
	if err = c.cdc.UnmarshalJSON(bz, v); err != nil {
		return c.CaptureError(bz, err)
	}
	return nil
}

func (c *client) CaptureError(bz []byte, err error) error {
	errResponse := struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{}

	json.Unmarshal(bz, &errResponse)
	if errResponse.Message != "" {
		return fmt.Errorf("code (%d): desc: (%s)", errResponse.Code, errResponse.Message)
	}

	return err
}

func (c *client) PendingNonceAt(ctx context.Context, address string) (uint64, error) {
	_, seq, err := c.GetAccountNumberAndSequence(ctx, address)
	if err != nil {
		return 0, err
	}
	return seq, nil
}

func (c *client) NonceAt(ctx context.Context, address string) (uint64, error) {
	return c.PendingNonceAt(ctx, address)
}

func (c *client) EstimateGas(ctx context.Context, callMsg types.CallMsg) (*big.Int, error) {
	return big.NewInt(0), clienttypes.NotImplemented
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

	raw, err := gogoproto.Marshal(cosmosTx)
	if err != nil {
		return nil, err
	}

	return c.BroadcastRawTx(ctx, raw, txtypes.BroadcastMode_BROADCAST_MODE_SYNC)
}

func (c *client) BroadcastRawTx(ctx context.Context, rawTx []byte, mode txtypes.BroadcastMode) (*clienttypes.SendTxAsyncResult, error) {
	bz, err := util.PostRetry(
		c.url+"/cosmos/tx/v1beta1/txs",
		map[string]any{
			"tx_bytes": base64.StdEncoding.EncodeToString(rawTx),
			"mode":     mode.String(),
		},
		nil, nil, 6)
	if err != nil {
		return nil, err
	}

	resp := txtypes.BroadcastTxResponse{}
	err = c.cdc.UnmarshalJSON(bz, &resp)
	if err != nil {
		return nil, err
	}

	var result clienttypes.SendTxAsyncResult
	if resp.GetTxResponse().Code == 0 {
		result.Result = clienttypes.SendTxResultType_Success
		result.Hash = resp.GetTxResponse().TxHash
		return &result, nil
	}

	return nil, cosmoserrors.ABCIError(resp.GetTxResponse().Codespace, resp.GetTxResponse().Code, resp.TxResponse.RawLog)
}

func (c *client) GetHeaderByHash(ctx context.Context, txHash string) (*types.Header, error) {
	return nil, clienttypes.NotImplemented
}

func (c *client) BlockOptionFinshcia(s []byte) []byte {
	const blockKey = "block"
	const entropy = "entropy"

	var rawBlock map[string]map[string]any

	err := json.Unmarshal(s, &rawBlock)
	if err != nil {
		return s
	}
	delete(rawBlock[blockKey], entropy)

	bz, err := json.Marshal(rawBlock)
	if err != nil {
		return s
	}
	return bz
}

func (c *client) BlockOptionCosmos(s []byte) []byte {
	// https://github.com/cometbft/cometbft/blob/v0.34.2/proto/tendermint/types/types.pb.go
	//type BlockID struct {
	//	Hash          []byte        `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	//	PartSetHeader PartSetHeader `protobuf:"bytes,2,opt,name=part_set_header,json=partSetHeader,proto3" json:"part_set_header"`
	//}
	//
	// https://github.com/cometbft/cometbft/blob/v0.34.29/types/block.go#L1166
	//type BlockID struct {
	//	Hash          []byte        `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	//	PartSetHeader PartSetHeader `protobuf:"bytes,2,opt,name=part_set_header,json=partSetHeader,proto3" json:"parts"`
	//}
	const parts = "parts"
	const partSetHeader = "part_set_header"

	return []byte(strings.ReplaceAll(string(s), parts, partSetHeader))
}

func (c *client) BlockNumber(ctx context.Context) (*big.Int, error) {
	var block tmservice.GetLatestBlockResponse
	if c.chain == types.ChainFNSA || c.chain == types.ChainTFNSA {
		if err := c.InvokeWithOption(ctx, &block, c.BlockOptionFinshcia, ModuleFinschia, "base", "ostracon", v1, "blocks", "latest"); err != nil {
			return big.NewInt(0), err
		}
	} else {
		if err := c.InvokeWithOption(ctx, &block, c.BlockOptionCosmos, ModuleCosmos, "base", "tendermint", v1Beta1, "blocks", "latest"); err != nil {
			return big.NewInt(0), err
		}
	}

	return big.NewInt(block.GetBlock().GetHeader().Height), nil
}

func (c *client) GetTransaction(ctx context.Context, txHash string) (*types.Transaction, error) {
	var tx txtypes.GetTxResponse
	if err := c.Invoke(ctx, &tx, ModuleCosmos, "tx", v1Beta1, "txs", txHash); err != nil {
		return nil, err
	}

	err := tx.GetTx().UnpackInterfaces(c.interfaceRegistry)
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
	var tx txtypes.GetTxResponse
	if err := c.Invoke(ctx, &tx, ModuleCosmos, "tx", v1Beta1, "txs", txHash); err != nil {
		return nil, err
	}

	err := tx.GetTx().UnpackInterfaces(c.interfaceRegistry)
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
func (c *client) Balance(ctx context.Context, address, denom string) (*big.Int, error) {
	balances, err := c.Balances(ctx, address)
	if err != nil {
		return big.NewInt(0), err
	}
	if ok, coin := balances.Find(denom); ok {
		return coin.Amount.BigInt(), nil
	} else {
		return big.NewInt(0), nil
	}
}

func (c *client) Balances(ctx context.Context, address string) (cosmossdk.Coins, error) {
	var resp banktypes.QueryAllBalancesResponse

	if err := c.Invoke(ctx, &resp, ModuleCosmos, banktypes.ModuleName, v1Beta1, "balances", address); err != nil {
		return cosmossdk.NewCoins(), err
	}

	return resp.Balances, nil
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

func (c *client) GetMinimumGasPrice(ctx context.Context) (float64, error) {
	var resp node.ConfigResponse

	if err := c.Invoke(ctx, &resp, ModuleCosmos, "base", "node", v1Beta1, "config"); err != nil {
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

func (c *client) GetBlockByNumber(ctx context.Context, number *big.Int) (*tmtypes.Block, error) {
	if number.Int64() == 0 {
		var block tmservice.GetLatestBlockResponse
		if c.chain == types.ChainFNSA || c.chain == types.ChainTFNSA {
			if err := c.InvokeWithOption(ctx, &block, c.BlockOptionFinshcia, ModuleFinschia, "base", "ostracon", v1, "blocks", "latest"); err != nil {
				return nil, err
			}
		} else {
			if err := c.InvokeWithOption(ctx, &block, c.BlockOptionCosmos, ModuleCosmos, "base", "tendermint", v1Beta1, "blocks", "latest"); err != nil {
				return nil, err
			}
		}
		return block.GetBlock(), nil
	} else {
		var block tmservice.GetBlockByHeightResponse
		if c.chain == types.ChainFNSA || c.chain == types.ChainTFNSA {
			if err := c.InvokeWithOption(ctx, &block, c.BlockOptionFinshcia, ModuleFinschia, "base", "ostracon", v1, "blocks", fmt.Sprintf("%d", number)); err != nil {
				return nil, err
			}
		} else {
			if err := c.InvokeWithOption(ctx, &block, c.BlockOptionCosmos, ModuleCosmos, "base", "tendermint", v1Beta1, "blocks", fmt.Sprintf("%d", number)); err != nil {
				return nil, err
			}
		}
		return block.GetBlock(), nil
	}
}

func (c *client) GetBlockByNumberWithTxs(ctx context.Context, number int64) (*tmtypes.Block, []*txtypes.Tx, error) {
	return nil, nil, fmt.Errorf("not supported")
}

func (c *client) GetAccountNumberAndSequence(ctx context.Context, address string) (uint64, uint64, error) {
	var resp authtypes.QueryAccountResponse
	if err := c.Invoke(ctx, &resp, ModuleCosmos, authtypes.ModuleName, v1Beta1, "accounts", address); err != nil {
		return 0, 0, err
	}

	var account authtypes.AccountI
	err := c.interfaceRegistry.UnpackAny(resp.Account, &account)
	if err != nil {
		return 0, 0, err
	}

	return account.GetAccountNumber(), account.GetSequence(), nil
}

func (c *client) GetReward(ctx context.Context, address string) (int64, error) {
	var resp distributiontypes.QueryDelegationTotalRewardsResponse
	if err := c.Invoke(ctx, &resp, ModuleCosmos, distributiontypes.ModuleName, v1Beta1, "delegators", address, "rewards"); err != nil {
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
	var resp stakingtypes.QueryDelegatorDelegationsResponse
	if err := c.Invoke(ctx, &resp, ModuleCosmos, stakingtypes.ModuleName, v1Beta1, "delegations", address); err != nil {
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
	var resp stakingtypes.QueryDelegatorUnbondingDelegationsResponse
	withdrawal := make([]clienttypes.Withdrawal, 0)
	if err := c.Invoke(ctx, &resp, ModuleCosmos, stakingtypes.ModuleName, v1Beta1, "delegators", address, "unbonding_delegations"); err != nil {
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

func (c *client) GetTxs(ctx context.Context, events ...string) (*txtypes.GetTxsEventResponse, error) {
	var resp txtypes.GetTxsEventResponse
	query := url.Values{}
	query.Add("pagination.limit", "5")
	query.Add("order_by", "ORDER_BY_DESC")
	for _, event := range events {
		query.Add("events", event)
	}
	if err := c.InvokeWithParams(ctx, &resp, query, ModuleCosmos, "tx", v1Beta1, "txs"); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *client) GetStakingParams(ctx context.Context) (stakingtypes.Params, error) {
	var resp stakingtypes.QueryParamsResponse
	if err := c.Invoke(ctx, &resp, ModuleCosmos, stakingtypes.ModuleName, v1Beta1, "params"); err != nil {
		return stakingtypes.Params{}, err
	}

	return resp.Params, nil
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
	var supplyOfResp banktypes.QuerySupplyOfResponse
	if err := c.Invoke(ctx, &supplyOfResp, ModuleCosmos, banktypes.ModuleName, v1Beta1, "supply", tokentypes.DenomByChain(c.chain)); err != nil {
		return 0, err
	}

	var poolResp stakingtypes.QueryPoolResponse
	if err := c.Invoke(ctx, &poolResp, ModuleCosmos, stakingtypes.ModuleName, v1Beta1, "pool"); err != nil {
		return 0, err
	}

	var inflationResp minttypes.QueryInflationResponse
	if err := c.Invoke(ctx, &inflationResp, ModuleCosmos, minttypes.ModuleName, v1Beta1, "inflation"); err != nil {
		return 0, err
	}

	var validatorResp stakingtypes.QueryValidatorResponse
	if err := c.Invoke(ctx, &validatorResp, ModuleCosmos, stakingtypes.ModuleName, v1Beta1, "validators", validator); err != nil {
		return 0, err
	}

	var distributionParamResp distributiontypes.QueryParamsResponse
	if err := c.Invoke(ctx, &distributionParamResp, ModuleCosmos, distributiontypes.ModuleName, v1Beta1, "params"); err != nil {
		return 0, err
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

	if i, err := inflationResp.Inflation.Float64(); err != nil {
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
		var paramResp foundation.QueryParamsResponse
		if err := c.Invoke(ctx, &paramResp, ModuleFinschia, foundation.ModuleName, v1, "params"); err != nil {
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
	var validatorResp stakingtypes.QueryValidatorResponse
	if err := c.Invoke(ctx, &validatorResp, ModuleCosmos, stakingtypes.ModuleName, v1Beta1, "validators", validator); err != nil {
		return big.NewInt(0), err
	}

	return validatorResp.Validator.BondedTokens().BigInt(), nil
}

func (c *client) GetValidatorDelegations(ctx context.Context, validator string) (stakingtypes.DelegationResponses, error) {
	var delegations stakingtypes.DelegationResponses
	var delegationsResp stakingtypes.QueryValidatorDelegationsResponse
	p := cosmostypes.DefaultPageRequest()
	for {
		query := url.Values{}
		query.Add("pagination.key", base64.StdEncoding.EncodeToString(p.GetKey()))
		query.Add("pagination.limit", fmt.Sprintf("%d", p.Limit))

		if err := c.InvokeWithParams(ctx, &delegationsResp, query, ModuleCosmos, stakingtypes.ModuleName, v1Beta1, "validators", validator, "delegations"); err != nil {
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

func (c *client) CallWasm(ctx context.Context, address string, msg types.CallWasm, v any) error {
	bz, err := cosmostypes.MarshalCallWasm(msg)
	if err != nil {
		return err
	}
	resp := wasmtypes.QuerySmartContractStateResponse{}
	err = c.Invoke(ctx, &resp, ModuleWasm, wasmtypes.ModuleName, v1, "contract", address, "smart", base64.StdEncoding.EncodeToString(bz))
	if err != nil {
		return err
	}

	return json.Unmarshal(resp.Data, &v)
}

func (c *client) SetAccountPrefix() {
	c.Lock()
	defer c.Unlock()

	prefix := cosmoscommon.GetAddressPrefixByChain(c.chain)
	cosmossdk.GetConfig().SetBech32PrefixForAccount(prefix, prefix+"pub")
}

func (c *client) IBCDenomTrace(ctx context.Context, hash string) (*ibctransfertypes.DenomTrace, error) {
	return nil, nil
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

	var resp txtypes.SimulateResponse
	bz, err := util.PostRetry(
		c.url+"/cosmos/tx/v1beta1/simulate",
		map[string]any{
			"tx_bytes": base64.StdEncoding.EncodeToString(txBytes),
		},
		nil, nil, 6)
	if err != nil {
		return nil, err
	}

	err = c.cdc.UnmarshalJSON(bz, &resp)
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

func (c *client) GetTransactionOption(ctx context.Context, from string) (*types.TransactionOption, error) {
	return nil, clienttypes.NotImplemented
}

func (c *client) GetTransactionData(data *types.RequestTransaction) (*types.Transaction, error) {
	return nil, clienttypes.NotImplemented
}
