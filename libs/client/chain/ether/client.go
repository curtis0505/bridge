package ether

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	ethertypes "github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"

	etherabi "github.com/ethereum/go-ethereum/accounts/abi"
	ethercommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/curtis0505/bridge/libs/client/chain/conf"
	clienttypes "github.com/curtis0505/bridge/libs/client/chain/types"
	"github.com/curtis0505/bridge/libs/common"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/types/bridge"
)

var (
	_ clienttypes.Client         = &client{}
	_ clienttypes.EVMClient      = &client{}
	_ clienttypes.BridgeClient   = &client{}
	_ clienttypes.FxPortalClient = &client{}
	_ clienttypes.EtherClient    = &client{}
)

type client struct {
	c         *ethclient.Client
	cc        *rpc.Client
	networkId *big.Int
	chain     string

	// subscribe 관련 변수
	finalizedBlockCount int
}

func NewClient(config conf.ClientConfig) (clienttypes.EtherClient, error) {
	c := &client{
		chain: config.Chain,
	}

	var err error
	c.cc, err = rpc.DialContext(context.Background(), config.Url)
	if err != nil {
		return nil, types.WrapError("DialContext", err)
	}
	c.c = ethclient.NewClient(c.cc)

	return c, nil
}

// Deprecated: ProxyClient
func ProxyClient(proxy clienttypes.Proxy, chain string) (clienttypes.EtherClient, error) {
	c := proxy.ProxyClient(chain)
	if c == nil {
		return nil, fmt.Errorf("not found proxy")
	}

	client, ok := c.(clienttypes.EtherClient)
	if !ok {
		return nil, fmt.Errorf("failed to casting client")
	}

	return client, nil
}

func (c *client) Chain() string { return c.chain }

func (c *client) ChainType() types.ChainType { return types.GetChainType(c.chain) }

func (c *client) NewMinter(address string, abi []map[string]interface{}) (bridge.Minter, error) {
	return newMinter(c, address), nil
}

func (c *client) NewVault(address string, abi []map[string]interface{}) (bridge.Vault, error) {
	return newVault(c, address), nil
}

func (c *client) NewFxERC20RootTunnel(address string, abi []map[string]interface{}) (bridge.FxERC20RootTunnel, error) {
	return newFxERC20RootTunnel(c, address), nil
}

func (c *client) NewFxERC20ChildTunnel(address string, abi []map[string]interface{}) (bridge.FxERC20ChildTunnel, error) {
	return newFxERC20ChildTunnel(c, address), nil
}

func (c *client) GetChainID(ctx context.Context) (*big.Int, error) {
	return c.c.ChainID(ctx)
}

func (c *client) NetworkId(ctx context.Context) (*big.Int, error) {
	if c.networkId != nil {
		return c.networkId, nil
	}
	var err error
	c.networkId, err = c.c.NetworkID(ctx)
	return c.networkId, err
}

func (c *client) CallContract(ctx context.Context, msg types.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return c.c.CallContract(ctx, msg.Ethereum(), blockNumber)
}

func (c *client) PendingNonceAt(ctx context.Context, address string) (uint64, error) {
	return c.c.PendingNonceAt(ctx, ethercommon.HexToAddress(address))
}

func (c *client) NonceAt(ctx context.Context, address string) (uint64, error) {
	return c.c.NonceAt(ctx, ethercommon.HexToAddress(address), nil)
}

func (c *client) GasPrice(ctx context.Context) (*big.Int, error) {
	return c.c.SuggestGasPrice(ctx)
}

func (c *client) RawTxAsync(ctx context.Context, rawTx []byte, rawProxyRequest []byte) (*clienttypes.SendTxAsyncResult, error) {
	tx, err := types.NewTransactionFromRLP(hexutil.Encode(rawTx), c.chain)
	if err != nil {
		return nil, err
	}

	_tx, err := tx.EthereumTransaction()
	if err != nil {
		return nil, err
	}
	err = c.c.SendTransaction(ctx, _tx)
	if err != nil {
		return nil, err
	}
	result := clienttypes.NewSendTxAsyncResult(clienttypes.SendTxResultType_Success, "", common.HexToHash(_tx.Hash().String()))
	//var result *clienttypes.SendTxAsyncResult
	//result.Result = clienttypes.SendTxResultType_Success
	//result.TxHash = common.HexToHash(_tx.Hash().String())
	return result, nil
}

func (c *client) RawTxAsyncByTx(ctx context.Context, tx *types.Transaction) (*clienttypes.SendTxAsyncResult, error) {
	_tx, err := tx.EthereumTransaction()
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
	tx, err := hexutil.Decode(rlpTx)
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
	return c.c.BalanceAt(ctx, ethercommon.HexToAddress(address), blockNumber)
}

func (c *client) BlockNumber(ctx context.Context) (*big.Int, error) {
	blockNumber, err := c.c.BlockNumber(ctx)
	if err != nil {
		return big.NewInt(0), err
	}
	return big.NewInt(int64(blockNumber)), nil
}

func (c *client) GetTransactionReceipt(ctx context.Context, txHash string) (*types.Receipt, error) {
	receipt, err := c.c.TransactionReceipt(ctx, ethercommon.HexToHash(txHash))
	return types.NewReceipt(receipt, c.chain), err
}

func (c *client) GetHeaderByHash(ctx context.Context, blockHash string) (*types.Header, error) {
	header, err := c.c.HeaderByHash(ctx, ethercommon.HexToHash(blockHash))
	return types.NewHeader(header), err
}

func (c *client) GetTransaction(ctx context.Context, txHash string) (*types.Transaction, error) {
	tx, _, err := c.c.TransactionByHash(ctx, ethercommon.HexToHash(txHash))
	return types.NewTransaction(tx, c.chain), err
}

func (c *client) GetTransactionWithReceipt(ctx context.Context, txHash string) (*types.Transaction, bool, error) {
	hash := ethercommon.HexToHash(txHash)
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

	abiParsed, err := etherabi.JSON(strings.NewReader(string(b)))
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

func (c *client) callMsg(abi etherabi.ABI, to, methodName string, args ...interface{}) ([]interface{}, error) {
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

// QueryRootHash : for pos client
func (c *client) QueryRootHash(ctx context.Context, startBlock, endBlock int64) (ethercommon.Hash, error) {
	var payload interface{}
	err := c.cc.CallContext(ctx, &payload, "eth_getRootHash",
		big.NewInt(startBlock),
		big.NewInt(endBlock),
	)
	if err != nil {
		return ethercommon.Hash{}, err
	}

	return ethercommon.HexToHash(payload.(string)), nil
}

// GetBlockByNumber : for pos client
func (c *client) GetBlockByNumber(ctx context.Context, blockNumber *big.Int) (*ethertypes.Block, error) {
	block, err := c.c.BlockByNumber(ctx, blockNumber)
	if err != nil {
		return nil, err
	}

	return block, nil
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
	result, err = c.c.EstimateGas(ctx, callMsg.Ethereum())
	if err == nil {
		resultBig.SetUint64(result)
	}
	return resultBig, err
}

// SuggestGasTipCap retrieves the currently suggested gas tip cap after 1559 to
// allow a timely execution of a transaction.
// NOT implemented in proxy server
func (c *client) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	return c.c.SuggestGasTipCap(ctx)
}

func (c *client) NewMultiSigWallet(address string) (bridge.MultiSigWallet, error) {
	panic("implement me")
}

func (c *client) GetTransactionOption(ctx context.Context, from string) (*types.TransactionOption, error) {
	blockNumber, err := c.c.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}

	block, err := c.c.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, err
	}

	gasFeeCap := block.BaseFee()

	gasTipCap, err := c.c.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, err
	}

	nonce, err := c.PendingNonceAt(ctx, from)
	if err != nil {
		return nil, err
	}

	transactionOption := &types.TransactionOption{
		Nonce:     nonce,
		GasTipCap: gasTipCap.Mul(gasTipCap, big.NewInt(2)),
		GasFeeCap: gasFeeCap.Mul(gasFeeCap, big.NewInt(2)),
	}

	// gasFeeCap < gasTipCap 이면 아래 에러 발생
	// "max priority fee per gas higher than max fee per gas"
	if transactionOption.GasFeeCap.Cmp(transactionOption.GasTipCap) == -1 {
		transactionOption.GasFeeCap = transactionOption.GasFeeCap.Add(transactionOption.GasTipCap, big.NewInt(1))
	}

	return transactionOption, nil
}

func (c *client) GetTransactionData(request *types.RequestTransaction) (*types.Transaction, error) {
	toAddress := ethercommon.HexToAddress(request.To)
	txData := types.NewTx(c.chain, &ethertypes.DynamicFeeTx{
		Nonce:     request.Nonce,
		GasFeeCap: request.GasFeeCap,
		GasTipCap: request.GasTipCap,
		Gas:       request.GasLimit,
		To:        &toAddress,
		Value:     request.Value,
		Data:      request.Data,
	})
	return txData, nil
}
