package klay

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"github.com/curtis0505/bridge/libs/client/chain/conf"
	clienttypes "github.com/curtis0505/bridge/libs/client/chain/types"
	"github.com/curtis0505/bridge/libs/common"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/types/bridge"
	klayabi "github.com/kaiachain/kaia/accounts/abi"
	klaytypes "github.com/kaiachain/kaia/blockchain/types"
	klayclient "github.com/kaiachain/kaia/client"
	klaycommon "github.com/kaiachain/kaia/common"
	klayhexutil "github.com/kaiachain/kaia/common/hexutil"
	"github.com/kaiachain/kaia/networks/rpc"
	"math/big"
	"strings"
)

var (
	_ clienttypes.Client         = &client{}
	_ clienttypes.EVMClient      = &client{}
	_ clienttypes.EVMClient      = &client{}
	_ clienttypes.BridgeClient   = &client{}
	_ clienttypes.FxPortalClient = &client{}
)

type client struct {
	c         *klayclient.Client
	cc        *rpc.Client
	networkID *big.Int
	feepayer  common.Address
	chain     string

	// subscribe 관련 변수
	finalizedBlockCount int
}

func NewClient(config conf.ClientConfig) (clienttypes.KlayClient, error) {
	c := &client{
		chain:               config.Chain,
		finalizedBlockCount: config.FinalizedBlockCount,
	}

	var err error
	c.cc, err = rpc.DialContext(context.Background(), config.Url)
	if err != nil {
		return nil, types.WrapError("DialContext", err)
	}
	c.c = klayclient.NewClient(c.cc)

	return c, nil
}

func (c *client) Chain() string              { return c.chain }
func (c *client) ChainType() types.ChainType { return types.GetChainType(c.chain) }

func (c *client) NewMinter(address string, abi []map[string]interface{}) (bridge.Minter, error) {
	return newMinter(c, address), nil
}
func (c *client) NewVault(address string, abi []map[string]interface{}) (bridge.Vault, error) {
	return newVault(c, address), nil
}

func (c *client) NewFxERC20RootTunnel(address string, abi []map[string]interface{}) (bridge.FxERC20RootTunnel, error) {
	return nil, clienttypes.NotImplemented
}

func (c *client) NewFxERC20ChildTunnel(address string, abi []map[string]interface{}) (bridge.FxERC20ChildTunnel, error) {
	return nil, clienttypes.NotImplemented
}

func (c *client) GetBlockByNumber(ctx context.Context, blockNumber *big.Int) (*klaytypes.Block, error) {
	return c.c.BlockByNumber(ctx, blockNumber)
}

func (c *client) GetChainID(ctx context.Context) (*big.Int, error) {
	return c.c.ChainID(ctx)
}

func (c *client) NetworkId(ctx context.Context) (*big.Int, error) {
	if c.networkID != nil {
		return c.networkID, nil
	}
	var err error
	c.networkID, err = c.c.NetworkID(ctx)
	return c.networkID, err
}

func (c *client) CallContract(ctx context.Context, msg types.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return c.c.CallContract(ctx, msg.Klaytn(), blockNumber)
}

func (c *client) PendingNonceAt(ctx context.Context, address string) (uint64, error) {
	return c.c.PendingNonceAt(ctx, klaycommon.HexToAddress(address))
}

func (c *client) NonceAt(ctx context.Context, address string) (uint64, error) {
	return c.c.NonceAt(ctx, klaycommon.HexToAddress(address), nil)
}

func (c *client) GasPrice(ctx context.Context) (*big.Int, error) {
	return c.c.SuggestGasPrice(ctx)
}

func (c *client) RawTxAsync(ctx context.Context, rawTx []byte, rawProxyRequest []byte) (*clienttypes.SendTxAsyncResult, error) {
	tx, err := types.NewTransactionFromRLP(klayhexutil.Encode(rawTx), c.chain)
	if err != nil {
		return nil, err
	}

	_tx, err := tx.KlaytnTransaction()
	if err != nil {
		return nil, err
	}
	err = c.c.SendTransaction(ctx, _tx)
	if err != nil {
		return nil, err
	}
	result := clienttypes.NewSendTxAsyncResult(clienttypes.SendTxResultType_Success, "", common.HexToHash(_tx.Hash().String()))
	return result, nil
}

func (c *client) RawTxAsyncByTx(ctx context.Context, tx *types.Transaction) (*clienttypes.SendTxAsyncResult, error) {
	_tx, err := tx.KlaytnTransaction()
	if err != nil {
		return nil, err
	}

	err = c.c.SendTransaction(ctx, _tx)
	if err != nil {
		return nil, err
	}

	result := clienttypes.NewSendTxAsyncResult(clienttypes.SendTxResultType_Success, "", common.HexToHash(_tx.Hash().String()))
	return result, nil
}

func (c *client) TxAsync(ctx context.Context, rlpTx string, proxyRequest clienttypes.ProxyRequest) (*clienttypes.SendTxAsyncResult, error) {
	tx, err := klayhexutil.Decode(rlpTx)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err = encoder.Encode(proxyRequest)
	if err != nil {
		return nil, err
	}

	return c.RawTxAsync(ctx, tx, buffer.Bytes())
}

func (c *client) BalanceAt(ctx context.Context, address string, blockNumber *big.Int) (*big.Int, error) {
	return c.c.BalanceAt(ctx, klaycommon.HexToAddress(address), blockNumber)
}

func (c *client) BlockNumber(ctx context.Context) (*big.Int, error) {
	return c.c.BlockNumber(ctx)
}

func (c *client) GetTransactionReceipt(ctx context.Context, txHash string) (*types.Receipt, error) {
	receipt, err := c.c.TransactionReceipt(ctx, klaycommon.HexToHash(txHash))
	return types.NewReceipt(receipt, c.chain), err
}

func (c *client) GetHeaderByHash(ctx context.Context, blockHash string) (*types.Header, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	result := &klaytypes.Header{}
	//klaytn의 경우 receipt에 blockHash(logs엔 존재), blockNumber가 없음.
	//klay(coin)의 경우 logs가 존재 하지 않아 blockHash 조회 불가.
	//거의 대부분 klay base fee는 25ston 이므로 고정.
	if strings.Compare(blockHash, klaycommon.Hash{}.Hex()) == 0 {
		result.BaseFee = big.NewInt(2500000000) //조회 불가 시, 25ston(klay)
		result.Number = big.NewInt(0)
		result.Time = big.NewInt(0)
		return types.NewHeader(result), nil
	}
	header, err := c.c.HeaderByHash(ctx, klaycommon.HexToHash(blockHash))
	return types.NewHeader(header), err
}

func (c *client) GetTransaction(ctx context.Context, txHash string) (*types.Transaction, error) {
	tx, _, err := c.c.TransactionByHash(ctx, klaycommon.HexToHash(txHash))
	return types.NewTransaction(tx, c.chain), err
}

func (c *client) GetTransactionWithReceipt(ctx context.Context, txHash string) (*types.Transaction, bool, error) {
	hash := klaycommon.HexToHash(txHash)
	tx, pending, err := c.c.TransactionByHash(ctx, hash)
	if err != nil {
		return nil, false, err
	}

	receipt, err := c.c.TransactionReceipt(ctx, hash)
	if err != nil {
		return types.NewTransaction(tx, c.chain), pending, err
	}

	return types.NewTransactionWithReceipt(tx, receipt, c.chain), pending, nil
}

func (c *client) CallMsg(ctx context.Context, from, to, methodName string, abi []map[string]interface{}, args ...interface{}) ([]interface{}, error) {
	b, err := json.Marshal(abi)
	if err != nil {
		return nil, err
	}

	abiParsed, err := klayabi.JSON(strings.NewReader(string(b)))
	if err != nil {
		return nil, err
	}

	input, err := abiParsed.Pack(methodName, args...)
	if err != nil {
		return nil, err
	}

	msg := types.CallMsg{From: from, To: to, Data: input}
	out, err := c.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}

	result, err := abiParsed.Methods[methodName].Outputs.UnpackValues(out)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) callMsg(abi klayabi.ABI, to, methodName string, args ...interface{}) ([]interface{}, error) {
	ctx := context.Background()

	input, err := abi.Pack(methodName, args...)
	if err != nil {
		return nil, err
	}

	msg := types.CallMsg{To: to, Data: input}
	out, err := c.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}

	result, err := abi.Methods[methodName].Outputs.UnpackValues(out)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) HeaderByNumber(ctx context.Context, blockNumber *big.Int) (*types.Header, error) {
	header, err := c.c.HeaderByNumber(ctx, blockNumber)
	if err != nil {
		return nil, err
	}

	return types.NewHeader(header), nil
}

func (c *client) EstimateGas(ctx context.Context, callMsg types.CallMsg) (*big.Int, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	var err error
	resultBig := new(big.Int)
	var result uint64
	result, err = c.c.EstimateGas(ctx, callMsg.Klaytn())
	if err == nil {
		resultBig.SetUint64(result)
	}
	return resultBig, err
}

func (c *client) NewMultiSigWallet(address string) (bridge.MultiSigWallet, error) {
	panic("implement me")
}

func (c *client) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return nil, types.NotSupported
}

func (c *client) GetTransactionOption(ctx context.Context, from string) (*types.TransactionOption, error) {
	gasPrice, err := c.c.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	nonce, err := c.PendingNonceAt(ctx, from)
	if err != nil {
		return nil, err
	}

	return &types.TransactionOption{
		Nonce: nonce, GasPrice: gasPrice.Mul(gasPrice, big.NewInt(2)),
	}, nil
}

func (c *client) GetTransactionData(request *types.RequestTransaction) (*types.Transaction, error) {
	txData := map[klaytypes.TxValueKeyType]interface{}{
		klaytypes.TxValueKeyFrom:     klaycommon.HexToAddress(request.From),
		klaytypes.TxValueKeyTo:       klaycommon.HexToAddress(request.To),
		klaytypes.TxValueKeyNonce:    request.Nonce,
		klaytypes.TxValueKeyGasPrice: request.GasPrice,
		klaytypes.TxValueKeyGasLimit: request.GasLimit,
		klaytypes.TxValueKeyAmount:   request.Value,
		klaytypes.TxValueKeyData:     request.Data,
	}

	tx := types.NewTx(c.chain, txData)

	return tx, nil
}
