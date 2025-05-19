package types

import (
	"context"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	kmultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	arbcommon "github.com/curtis0505/arbitrum/common"
	arbtypes "github.com/curtis0505/arbitrum/core/types"
	basetypes "github.com/curtis0505/base/core/types"
	cosmostypes "github.com/curtis0505/bridge/libs/client/chain/cosmos/types"
	ethertypes "github.com/ethereum/go-ethereum/core/types"
	klaytypes "github.com/klaytn/klaytn/blockchain/types"
	"math/big"

	tmtypes "github.com/cometbft/cometbft/proto/tendermint/types"
	cosmosclient "github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	tronapi "github.com/curtis0505/grpc-idl/tron/api"
	troncore "github.com/curtis0505/grpc-idl/tron/core"
	ethercommon "github.com/ethereum/go-ethereum/common"
	"google.golang.org/protobuf/proto"

	"github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/types/bridge"
)

type Proxy interface {
	GetChains() []string
	ProxyClient(chain string) Client
}

type Client interface {
	Chain() string
	ChainType() types.ChainType
	NetworkId(ctx context.Context) (*big.Int, error)
	GetChainID(ctx context.Context) (*big.Int, error)

	NonceAt(ctx context.Context, address string) (uint64, error)
	PendingNonceAt(ctx context.Context, address string) (uint64, error)
	GasPrice(ctx context.Context) (*big.Int, error)
	RawTxAsync(ctx context.Context, rawTx []byte, rawProxyRequest []byte) (*SendTxAsyncResult, error)
	RawTxAsyncByTx(ctx context.Context, tx *types.Transaction) (*SendTxAsyncResult, error)
	TxAsync(ctx context.Context, rlpTx string, proxyRequest ProxyRequest) (*SendTxAsyncResult, error)
	BalanceAt(ctx context.Context, address string, blockNumber *big.Int) (*big.Int, error)
	EstimateGas(ctx context.Context, msg types.CallMsg) (*big.Int, error)
	BlockNumber(ctx context.Context) (*big.Int, error)
	GetTransactionReceipt(ctx context.Context, txHash string) (*types.Receipt, error)
	GetHeaderByHash(ctx context.Context, txHash string) (*types.Header, error)
	GetTransaction(ctx context.Context, txHash string) (*types.Transaction, error)
	GetTransactionWithReceipt(ctx context.Context, txHash string) (*types.Transaction, bool, error)
	Subscribe(ctx context.Context, cb func(eventLog types.Log), addresses ...string) error
	GetTransactionOption(ctx context.Context, from string) (*types.TransactionOption, error)
	GetTransactionData(data *types.RequestTransaction) (*types.Transaction, error)
}

type CosmosClient interface {
	Client
	StakingClient
	IBCClient
	BridgeClient
	WasmClient

	SetAccountPrefix()
	TxDecoder() cosmossdk.TxDecoder
	TxConfig() cosmosclient.TxConfig
	InterfaceRegistry() codectypes.InterfaceRegistry
	ChainId() string

	Balance(ctx context.Context, address, denom string) (*big.Int, error)
	Balances(ctx context.Context, address string) (cosmossdk.Coins, error)

	GetMinimumGasPrice(ctx context.Context) (float64, error)
	GetBlockByNumber(ctx context.Context, blockNumber *big.Int) (*tmtypes.Block, error)
	GetBlockByNumberWithTxs(ctx context.Context, number int64) (*tmtypes.Block, []*txtypes.Tx, error)
	GetAccountNumberAndSequence(ctx context.Context, address string) (uint64, uint64, error)

	GetValidatorApr(ctx context.Context, validator string) (float64, error)
	GetValidatorTvl(ctx context.Context, validator string) (*big.Int, error)
	GetValidatorDelegations(ctx context.Context, validator string) (stakingtypes.DelegationResponses, error)
	GetStakingParams(ctx context.Context) (stakingtypes.Params, error)

	BroadcastRawTx(ctx context.Context, rawTx []byte, mode txtypes.BroadcastMode) (*SendTxAsyncResult, error)

	GetTxs(ctx context.Context, events ...string) (*txtypes.GetTxsEventResponse, error)
	SendTransaction(ctx context.Context, account *types.Account, opts ...cosmostypes.TxOptionFunc) (*SendTxAsyncResult, error)
	SendTransactionPrivKey(ctx context.Context, priv *secp256k1.PrivKey, opts ...cosmostypes.TxOptionFunc) (*SendTxAsyncResult, error)
	SimulateTransaction(ctx context.Context, priv *secp256k1.PrivKey, opts ...cosmostypes.TxOptionFunc) (*cosmossdk.GasInfo, error)

	SignMultiSig(ctx context.Context, priv *types.Account, multiSig string, msg *wasmtypes.MsgExecuteContract) (signing.SignatureV2, error)
	SendMultiSigTransaction(ctx context.Context, multiSigPubKey *kmultisig.LegacyAminoPubKey, msg *wasmtypes.MsgExecuteContract, singedTxs ...signing.SignatureV2) (*SendTxAsyncResult, error)
}

type TronClient interface {
	Client

	EVMClient
	StakingClient

	GetTronStaking(ctx context.Context, address string, validatorAddress string) (*Staking, error)
	GetBlock(ctx context.Context) (*tronapi.BlockExtention, error)
	GetAccount(ctx context.Context, address string) (*troncore.Account, error)
	CreateTransaction(ctx context.Context, contract proto.Message) (*types.Transaction, error)
	CreateTransferTransaction(ctx context.Context, from, to string, amount *big.Int) (*types.Transaction, error)
}

type KlayClient interface {
	Client

	EVMClient
	BridgeClient
	GetBlockByNumber(ctx context.Context, blockNumber *big.Int) (*klaytypes.Block, error)
}

type EtherClient interface {
	Client

	EVMClient
	BridgeClient
	FxPortalClient

	QueryRootHash(ctx context.Context, startBlock, endBlock int64) (ethercommon.Hash, error)
	GetBlockByNumber(ctx context.Context, blockNumber *big.Int) (*ethertypes.Block, error)
}

type ArbClient interface {
	Client

	EVMClient

	QueryRootHash(ctx context.Context, startBlock, endBlock int64) (arbcommon.Hash, error)
	GetBlockByNumber(ctx context.Context, blockNumber *big.Int) (*arbtypes.Block, error)
}

type BaseClient interface {
	Client
	GetBlockByNumber(ctx context.Context, blockNumber *big.Int) (*basetypes.Block, error)
}

type WasmClient interface {
	CallWasm(ctx context.Context, address string, msg types.CallWasm, v any) error
}

type IBCClient interface {
	IBCDenomTrace(ctx context.Context, hash string) (*ibctransfertypes.DenomTrace, error)
}

// EVMClient
// chain: ETH, MATIC, KLAY, TRX
type EVMClient interface {
	HeaderByNumber(ctx context.Context, blockNumber *big.Int) (*types.Header, error)
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	CallContract(ctx context.Context, msg types.CallMsg, blockNumber *big.Int) ([]byte, error)
	CallMsg(ctx context.Context, from, to, methodName string, abi []map[string]any, args ...any) ([]any, error)
}

// BridgeClient
// chain: ETH, MATIC, KLAY
type BridgeClient interface {
	NewMinter(address string, abi []map[string]any) (bridge.Minter, error)
	NewVault(address string, abi []map[string]any) (bridge.Vault, error)
	NewMultiSigWallet(address string) (bridge.MultiSigWallet, error)
}

// FxPortalClient
// chain: MATIC, ETH
type FxPortalClient interface {
	NewFxERC20RootTunnel(address string, abi []map[string]any) (bridge.FxERC20RootTunnel, error)
	NewFxERC20ChildTunnel(address string, abi []map[string]any) (bridge.FxERC20ChildTunnel, error)
}

// StakingClient
// chain: TRX, COSMOS, FNSA ...
type StakingClient interface {
	GetReward(ctx context.Context, address string) (int64, error)
	GetStaking(ctx context.Context, address string) (*Staking, error)
}
