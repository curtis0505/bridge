package chain

import (
	"context"
	"fmt"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	retry "github.com/avast/retry-go/v4"
	"github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	"github.com/curtis0505/bridge/libs/cache"
	"github.com/curtis0505/bridge/libs/client/chain/arb"
	"github.com/curtis0505/bridge/libs/client/chain/base"
	"github.com/curtis0505/bridge/libs/client/chain/conf"
	"github.com/curtis0505/bridge/libs/client/chain/cosmos"
	"github.com/curtis0505/bridge/libs/client/chain/ether"
	"github.com/curtis0505/bridge/libs/client/chain/klay"
	clienttypes "github.com/curtis0505/bridge/libs/client/chain/types"
	"github.com/curtis0505/bridge/libs/database"
	mongoServiceDB "github.com/curtis0505/bridge/libs/database/mongo/service_db"
	"github.com/curtis0505/bridge/libs/logger/v2"
	"github.com/curtis0505/bridge/libs/multicall"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/types/bridge"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
	"math/big"
	"reflect"
	"sync"
	"time"
)

type Client struct {
	config  conf.Config
	clients map[string]clienttypes.Client
	nodes   map[string][]Node
	chains  []string
}

type Node struct {
	clienttypes.Client
	*sync.RWMutex

	url         string
	blockNumber int64
}

func NewNode(client clienttypes.Client, url string) Node {
	return Node{
		RWMutex: &sync.RWMutex{},
		Client:  client,

		url:         url,
		blockNumber: 0,
	}
}

var (
	NotFoundProxy = fmt.Errorf("invalid chain")
	NotSupport    = fmt.Errorf("this chain is not supported")
)

const (
	defaultDelay    = time.Millisecond * 300
	defaultAttempts = 5
)

func InitClient(clientConf conf.Chains) *Client {
	nodeInfo, err := database.GetRegistry().ServiceDB().FindNodes(context.Background(), bson.M{"chain": bson.M{"$in": clientConf.GetClientChains()}})
	if err != nil {
		logger.Error("InitClient", logger.BuildLogInput().WithError(err))
		return nil
	}

	if len(nodeInfo) == 0 {
		logger.Error("InitClient", logger.BuildLogInput().WithError(fmt.Errorf("node info is empty")))
		return nil
	}

	return NewClient(nodeInfo)
}

func NewClient(nodes []*mongoServiceDB.Nodes) *Client {
	client := &Client{
		nodes:   make(map[string][]Node),
		clients: make(map[string]clienttypes.Client),
		chains:  make([]string, 0),
	}

	lo.ForEach[*mongoServiceDB.Nodes](nodes, func(node *mongoServiceDB.Nodes, _ int) {
		for _, url := range node.ServerNodeURL {
			if err := client.AddClient(conf.ClientConfig{
				Chain: node.Chain,
				Url:   url,
			}); err != nil {
				logger.Error("NewClient",
					logger.BuildLogInput().
						WithChain(node.Chain).
						WithError(err).
						WithData("url", url),
				)
				panic(err)
			}
		}
	})

	client.setActiveClient()
	go client.iterate()

	return client
}

/*
NewClientByConfig bridge-validator 에서만 사용중
*/
func NewClientByConfig(config conf.Config) *Client {
	client := &Client{
		nodes:   make(map[string][]Node),
		clients: make(map[string]clienttypes.Client),
		chains:  make([]string, 0),
	}

	for _, clientConfig := range config {
		if err := client.AddClient(clientConfig); err != nil {
			//logger.Error("AddClient", "chain", clientConfig.Chain, "url", clientConfig.Url, "err", err)
			panic(err)
		}
	}

	client.setActiveClient()
	go client.iterate()

	return client
}

func (p *Client) AddClient(config conf.ClientConfig) error {
	var client clienttypes.Client
	var err error

	switch config.Chain {
	// TODO: 체인 추가시 체크 필요
	case types.ChainKLAY:
		client, err = klay.NewClient(config)
	case types.ChainARB:
		client, err = arb.NewClient(config)
	case types.ChainBASE:
		client, err = base.NewClient(config)
	case types.ChainETH, types.ChainMATIC:
		client, err = ether.NewClient(config)
	case types.ChainATOM, types.ChainFNSA, types.ChainTFNSA, types.ChainOSMO, types.ChainKAVA:
		client, err = cosmos.NewClient(config)
	default:
		return fmt.Errorf("invalid chain %s", config.Chain)
	}

	if err != nil {
		return err
	}

	p.nodes[config.Chain] = append(p.nodes[config.Chain], NewNode(client, config.Url))
	p.chains = append(p.chains, config.Chain)

	return nil
}

func (p *Client) iterate() {
	nodeTicker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-nodeTicker.C:
			p.setActiveClient()
		}
	}
}

func (p *Client) setActiveClient() {
	for _, chain := range p.chains {
		var wg sync.WaitGroup
		for i := range p.nodes[chain] {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()

				node := &p.nodes[chain][i]
				block, err := node.Client.BlockNumber(context.Background())
				if err != nil {
					logger.Error("setActiveClient",
						logger.BuildLogInput().
							WithChain(chain).
							WithError(err).
							WithData("url", node.url),
					)
					return
				}

				node.RWMutex.Lock()
				node.blockNumber = block.Int64()
				node.RWMutex.Unlock()
			}(i)
		}
		wg.Wait()

		// 노드들의 블록 번호를 비교하여 가장 큰 노드를 active client 로 설정
		for i, node := range p.nodes[chain] {
			if p.nodes[chain] == nil || p.nodes[chain][i].blockNumber < node.blockNumber {
				p.clients[chain] = node.Client
				logger.Info("change best node",
					logger.BuildLogInput().
						WithChain(chain).
						WithData("url", node.url),
				)
			}
		}

		if p.clients[chain] == nil {
			p.clients[chain] = p.nodes[chain][0].Client
		}
	}
}

func (p *Client) RetryDo(ctx context.Context, doFunc retry.RetryableFunc, logFunc retry.OnRetryFunc) error {
	return retry.Do(
		doFunc,
		retry.Context(ctx),
		retry.Delay(defaultDelay),
		retry.Attempts(defaultAttempts),
		retry.DelayType(retry.FixedDelay),
		retry.LastErrorOnly(true),
		retry.OnRetry(logFunc),
	)
}

func (p *Client) RetryLog(msg string, params ...interface{}) retry.OnRetryFunc {
	return func(n uint, err error) {
		logs := append([]interface{}{
			"retry", n + 1,
			"err", err,
		}, params...)
		logger.Warn(msg, logger.BuildLogInput().WithData(logs...))
	}
}

func (p *Client) NetworkId(ctx context.Context, chain string) (*big.Int, error) {
	if _, ok := p.clients[chain]; !ok {
		return nil, NotFoundProxy
	}

	return p.clients[chain].NetworkId(ctx)
}

func (p *Client) GetChainID(ctx context.Context, chain string) (*big.Int, error) {
	if _, ok := p.clients[chain]; !ok {
		return nil, NotFoundProxy
	}

	return p.clients[chain].GetChainID(ctx)
}

func (p *Client) CallContract(ctx context.Context, chain string, msg types.CallMsg, blockNumber *big.Int) ([]byte, error) {
	if _, ok := p.clients[chain]; !ok {
		return nil, NotFoundProxy
	}

	client, ok := p.clients[chain].(clienttypes.EVMClient)
	if !ok {
		return nil, NotSupport
	}

	var (
		bz  []byte
		err error
	)

	err = p.RetryDo(
		ctx,
		func() error {
			bz, err = client.CallContract(ctx, msg, blockNumber)
			if err != nil {
				return err
			}
			return nil
		},
		p.RetryLog("CallContract", "chain", chain, "address", msg.To),
	)

	return bz, err
}

func (p *Client) PendingNonceAt(ctx context.Context, chain, address string) (uint64, error) {
	if _, ok := p.clients[chain]; !ok {
		return 0, NotFoundProxy
	}

	var (
		nonce = uint64(0)
		err   error
	)

	err = p.RetryDo(
		ctx,
		func() error {
			nonce, err = p.clients[chain].PendingNonceAt(ctx, address)
			if err != nil {
				return err
			}
			return nil
		},
		p.RetryLog("PendingNonceAt", "chain", chain, "address", address),
	)

	return nonce, err
}

func (p *Client) NonceAt(ctx context.Context, chain, address string) (uint64, error) {
	if _, ok := p.clients[chain]; !ok {
		return 0, NotFoundProxy
	}

	var (
		nonce = uint64(0)
		err   error
	)

	err = p.RetryDo(
		ctx,
		func() error {
			nonce, err = p.clients[chain].NonceAt(ctx, address)
			if err != nil {
				return err
			}
			return nil
		},
		p.RetryLog("NonceAt", "chain", chain, "address", address),
	)

	return nonce, err
}

func (p *Client) GasPrice(ctx context.Context, chain string) (*big.Int, error) {
	if _, ok := p.clients[chain]; !ok {
		return nil, NotFoundProxy
	}

	var (
		gasPrice = big.NewInt(0)
		err      error
	)

	err = p.RetryDo(
		ctx,
		func() error {
			gasPrice, err = p.clients[chain].GasPrice(ctx)
			if err != nil {
				return err
			}
			return nil
		},
		p.RetryLog("GasPrice", "chain", chain),
	)

	return gasPrice, err
}

func (p *Client) RawSendTxAsync(ctx context.Context, chain string, rawTx []byte, rawProxyRequest []byte) (*clienttypes.SendTxAsyncResult, error) {
	if _, ok := p.clients[chain]; !ok {
		return nil, NotFoundProxy
	}

	return p.clients[chain].RawTxAsync(ctx, rawTx, rawProxyRequest)
}

func (p *Client) RawSendTxAsyncByTx(ctx context.Context, chain string, tx *types.Transaction) (*clienttypes.SendTxAsyncResult, error) {
	if _, ok := p.clients[chain]; !ok {
		return nil, NotFoundProxy
	}

	return p.clients[chain].RawTxAsyncByTx(ctx, tx)
}

// Deprecated: SendTxAsync
func (p *Client) SendTxAsync(ctx context.Context, chain string, rlpTx string, proxyRequest ...clienttypes.ProxyRequest) (*clienttypes.SendTxAsyncResult, error) {
	if _, ok := p.clients[chain]; !ok {
		return nil, NotFoundProxy
	}

	if len(proxyRequest) == 0 {
		return p.clients[chain].TxAsync(ctx, rlpTx, clienttypes.ProxyRequest{})
	}

	return p.clients[chain].TxAsync(ctx, rlpTx, proxyRequest[0])
}

func (p *Client) SendTx(ctx context.Context, chain string, rlpTx string) (*clienttypes.SendTxAsyncResult, error) {
	if _, ok := p.clients[chain]; !ok {
		return nil, NotFoundProxy
	}

	return p.clients[chain].TxAsync(ctx, rlpTx, clienttypes.ProxyRequest{})
}

func (p *Client) BalanceAt(ctx context.Context, chain, address string, blockNumber *big.Int) (*big.Int, error) {
	if _, ok := p.clients[chain]; !ok {
		return nil, NotFoundProxy
	}

	if len(address) == 0 {
		return big.NewInt(0), nil
	}

	var (
		balance = big.NewInt(0)
		err     error
	)

	err = p.RetryDo(
		ctx,
		func() error {
			balance, err = p.clients[chain].BalanceAt(ctx, address, blockNumber)
			if err != nil {
				return err
			}
			return nil
		},
		p.RetryLog("BalanceAt", "chain", chain, "address", address),
	)

	return balance, err
}

func (p *Client) BlockNumber(ctx context.Context, chain string) (*big.Int, error) {
	if _, ok := p.clients[chain]; !ok {
		return nil, NotFoundProxy
	}

	var (
		blockNumber = big.NewInt(0)
		err         error
	)

	err = p.RetryDo(
		ctx,
		func() error {
			blockNumber, err = p.clients[chain].BlockNumber(ctx)
			if err != nil {
				return err
			}
			return nil
		},
		p.RetryLog("BlockNumber", "chain", chain, "blockNumber", blockNumber),
	)

	return blockNumber, err
}

func (p *Client) GetTransactionReceipt(ctx context.Context, chain, txHash string) (*types.Receipt, error) {
	if _, ok := p.clients[chain]; !ok {
		return nil, NotFoundProxy
	}

	var (
		receipt *types.Receipt
		err     error
	)

	err = p.RetryDo(
		ctx,
		func() error {
			receipt, err = p.clients[chain].GetTransactionReceipt(ctx, txHash)
			if err != nil {
				return err
			}
			return nil
		},
		p.RetryLog("GetTransactionReceipt", "chain", chain, "txHash", txHash),
	)

	return receipt, err
}

func (p *Client) GetHeaderByHash(ctx context.Context, chain, txHash string) (*types.Header, error) {
	if _, ok := p.clients[chain]; !ok {
		return nil, NotFoundProxy
	}

	var (
		header *types.Header
		err    error
	)

	err = p.RetryDo(
		ctx,
		func() error {
			header, err = p.clients[chain].GetHeaderByHash(ctx, txHash)
			if err != nil {
				return err
			}
			return nil
		},
		p.RetryLog("GetHeaderByHash", "chain", chain, "txHash", txHash),
	)

	return header, err
}

func (p *Client) GetTransaction(ctx context.Context, chain, txHash string) (*types.Transaction, error) {
	if _, ok := p.clients[chain]; !ok {
		return nil, NotFoundProxy
	}

	var (
		tx  *types.Transaction
		err error
	)

	err = p.RetryDo(
		ctx,
		func() error {
			tx, err = p.clients[chain].GetTransaction(ctx, txHash)
			if err != nil {
				return err
			}
			return nil
		},
		p.RetryLog("GetTransaction", "chain", chain, "txHash", txHash),
	)

	return tx, err
}

func (p *Client) GetTransactionWithReceipt(ctx context.Context, chain, txHash string) (*types.Transaction, bool, error) {
	if _, ok := p.clients[chain]; !ok {
		return nil, false, NotFoundProxy
	}

	var (
		tx        *types.Transaction
		isPending bool
		err       error
	)

	err = p.RetryDo(
		ctx,
		func() error {
			tx, isPending, err = p.clients[chain].GetTransactionWithReceipt(ctx, txHash)
			if err != nil {
				return err
			}
			return nil
		},
		p.RetryLog("GetTransactionWithReceipt", "chain", chain, "txHash", txHash),
	)

	return tx, isPending, err
}

func (p *Client) CallWasm(ctx context.Context, chain, address string, wasm types.CallWasm, v interface{}) error {
	if _, ok := p.clients[chain]; !ok {
		return NotFoundProxy
	}
	client, ok := p.clients[chain].(clienttypes.WasmClient)
	if !ok {
		return NotSupport
	}

	if err := client.CallWasm(ctx, address, wasm, &v); err != nil {
		return err
	}

	return nil
}

func (p *Client) CallMsg(ctx context.Context, chain, from, to, methodName string, abi []map[string]interface{}, args ...interface{}) ([]interface{}, error) {
	if _, ok := p.clients[chain]; !ok {
		return nil, NotFoundProxy
	}
	client, ok := p.clients[chain].(clienttypes.EVMClient)
	if !ok {
		return nil, NotSupport
	}

	var (
		result []interface{}
		err    error
	)

	err = p.RetryDo(
		ctx,
		func() error {
			result, err = client.CallMsg(ctx, from, to, methodName, abi, args...)
			if err != nil {
				return err
			}
			return nil
		},
		p.RetryLog("CallMsg", "chain", chain, "to", to),
	)

	return result, err
}

func (p *Client) CallMsgMultiCall(ctx context.Context, chain string, multiCall *multicall.MultiCall, opts ...multicall.CallOptionFunc) error {
	option := multicall.NewCallOption()
	for _, fn := range opts {
		fn(option)
	}

	c, err := cache.ContractCache().GetContractByContractID(ctx, chain, multicall.ContractID)
	if err != nil {
		return err
	}
	for chunkNum := 0; chunkNum < multiCall.ChunkSize(option.GetChunkSize()); chunkNum++ {
		chunk := multiCall.Chunk(chunkNum, option.GetChunkSize())
		err := p.CallMsgUnmarshal(ctx, chain, "", c.Address, multicall.Aggregate, multicall.Abi, chunk, chunk)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Client) CallMsgMultiCall3(ctx context.Context, chain string, multiCall *multicall.MultiCall3, opts ...multicall.CallOptionFunc) error {
	option := multicall.NewCallOption()
	for _, fn := range opts {
		fn(option)
	}

	c, err := cache.ContractCache().GetContractByContractID(ctx, chain, multicall.ContractID)
	if err != nil {
		return err
	}
	for chunkNum := 0; chunkNum < multiCall.ChunkSize(option.GetChunkSize()); chunkNum++ {
		chunk := multiCall.Chunk(chunkNum, option.GetChunkSize())
		err = p.CallMsgUnmarshal(ctx, chain, "", c.Address, multicall.Aggregate3, multicall.Abi, chunk, chunk)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Client) CallMsgContract(ctx context.Context, contract mongoServiceDB.Contracts, methodName string, args ...interface{}) ([]interface{}, error) {
	return p.CallMsg(ctx, contract.Chain, "", contract.Address, methodName, contract.ABI, args...)
}

func (p *Client) CallMsgUnmarshal(ctx context.Context, chain, from, to, methodName string, abi []map[string]interface{}, v types.CallMsgUnmarshaler, args ...interface{}) error {
	if reflect.ValueOf(v).Kind() != reflect.Pointer {
		return fmt.Errorf("not pointer: %v", v)
	}

	if call, ok := v.(*multicall.MultiCall); ok {
		if call.Len() > multicall.DefaultBatchSize {
			logger.Warn("CallMsgUnmarshal", logger.BuildLogInput().WithError(fmt.Errorf("maximum batch size exceeded. len(%d) Please refrain from using CallMsgUnmarshal for batches larger than the specified limit. Consider switching to CallMsgMultiCall for efficient handling of larger batches", call.Len())))
		}
	}

	for i, arg := range args {
		v := reflect.ValueOf(arg)
		if v.Kind() == reflect.Ptr && v.IsNil() {
			return fmt.Errorf("call to %s, method: %s arg index[%d] type(%v) is nil pointer", to, methodName, i, reflect.TypeOf(arg))
		}
	}

	resp, err := p.CallMsg(ctx, chain, from, to, methodName, abi, args...)
	if err != nil {
		return err
	}

	v.Unmarshal(resp)
	return nil
}

func (p *Client) CallMsgUnmarshalContract(ctx context.Context, contract mongoServiceDB.Contracts, methodName string, v types.CallMsgUnmarshaler, args ...interface{}) error {
	return p.CallMsgUnmarshal(ctx, contract.Chain, "", contract.Address, methodName, contract.ABI, v, args...)
}

func (p *Client) CallMsgUnmarshalContract2(ctx context.Context, contract *mongoServiceDB.Contracts, methodName string, v types.CallMsgUnmarshaler, args ...interface{}) error {
	return p.CallMsgUnmarshal(ctx, contract.Chain, "", contract.Address, methodName, contract.ABI, v, args...)
}

func (p *Client) NewMinter(chain, address string, abi []map[string]interface{}) (bridge.Minter, error) {
	if _, ok := p.clients[chain]; !ok {
		return nil, NotFoundProxy
	}

	client, ok := p.clients[chain].(clienttypes.BridgeClient)
	if !ok {
		return nil, NotSupport
	}

	return client.NewMinter(address, abi)
}

func (p *Client) NewVault(chain, address string, abi []map[string]interface{}) (bridge.Vault, error) {
	if _, ok := p.clients[chain]; !ok {
		return nil, NotFoundProxy
	}

	client, ok := p.clients[chain].(clienttypes.BridgeClient)
	if !ok {
		return nil, NotSupport
	}

	return client.NewVault(address, abi)
}

func (p *Client) NewFxERC20RootTunnel(address string, abi []map[string]interface{}) (bridge.FxERC20RootTunnel, error) {
	if _, ok := p.clients[types.ChainETH]; !ok {
		return nil, NotFoundProxy
	}

	client, ok := p.clients[types.ChainETH].(clienttypes.FxPortalClient)
	if !ok {
		return nil, NotSupport
	}

	return client.NewFxERC20RootTunnel(address, abi)
}

func (p *Client) NewFxERC20ChildTunnel(address string, abi []map[string]interface{}) (bridge.FxERC20ChildTunnel, error) {
	if _, ok := p.clients[types.ChainMATIC]; !ok {
		return nil, NotFoundProxy
	}

	client, ok := p.clients[types.ChainMATIC].(clienttypes.FxPortalClient)
	if !ok {
		return nil, NotSupport
	}

	return client.NewFxERC20ChildTunnel(address, abi)
}

func (p *Client) DenomTrace(chain, hash string) (*ibctransfertypes.DenomTrace, error) {
	if _, ok := p.clients[chain]; !ok {
		return nil, NotFoundProxy
	}

	client, ok := p.clients[chain].(clienttypes.IBCClient)
	if !ok {
		return nil, NotSupport
	}

	return client.IBCDenomTrace(context.Background(), hash)
}

func (p *Client) ProxyClient(chain string) clienttypes.Client {
	return p.clients[chain]
}

func (p *Client) CosmosClients() []clienttypes.CosmosClient {
	clients := make([]clienttypes.CosmosClient, 0)
	for chain, client := range p.clients {
		if types.GetChainType(chain) == types.ChainTypeCOSMOS {
			cosmosClient, ok := client.(clienttypes.CosmosClient)
			if ok {
				clients = append(clients, cosmosClient)
			}
		}
	}
	return clients
}

// GetChains return chain
// ETH, MATIC, KLAY, TRX ...
func (p *Client) GetChains() []string { return p.chains }

// GetEVMChains return chain
// ETH, MATIC, KLAY, TRX ...
func (p *Client) GetEVMChains() []string {
	var evmChains []string

	for _, chain := range p.GetChains() {
		if types.GetChainType(chain) == types.ChainTypeEVM {
			evmChains = append(evmChains, chain)
		}
	}
	return evmChains
}

func (p *Client) EstimateGas(ctx context.Context, chain string, msg types.CallMsg) (*big.Int, error) {
	if _, ok := p.clients[chain]; !ok {
		return nil, NotFoundProxy
	}

	return p.clients[chain].EstimateGas(ctx, msg)
}

func (p *Client) HeaderByNumber(ctx context.Context, chain string, blockNumber *big.Int) (*types.Header, error) {
	client, ok := p.clients[chain].(clienttypes.EVMClient)
	if !ok {
		return nil, NotSupport
	}

	return client.HeaderByNumber(ctx, blockNumber)
}

func (p *Client) SuggestGasTipCap(ctx context.Context, chain string) (*big.Int, error) {
	client, ok := p.clients[chain].(clienttypes.EVMClient)
	if !ok {
		return nil, NotSupport
	}

	return client.SuggestGasTipCap(ctx)
}

func (p *Client) Subscribe(ctx context.Context, chain string, cb func(eventLog types.Log), addresses ...string) error {
	if _, ok := p.clients[chain]; !ok {
		return NotFoundProxy
	}

	return p.clients[chain].Subscribe(ctx, cb, addresses...)
}

func (p *Client) SignMultiSig(ctx context.Context, chain string, account *types.Account, multiSig string, msg *wasmtypes.MsgExecuteContract) (signing.SignatureV2, error) {
	switch chain {
	case types.ChainFNSA, types.ChainTFNSA:
		for _, client := range p.CosmosClients() {
			return client.SignMultiSig(ctx, account, multiSig, msg)
		}
	}
	return signing.SignatureV2{}, fmt.Errorf("invalid chain")
}

func (p *Client) SendMultiSigSinged(ctx context.Context, chain string, multiSigPubKey *multisig.LegacyAminoPubKey, msg *wasmtypes.MsgExecuteContract, singedTxs ...signing.SignatureV2) (*clienttypes.SendTxAsyncResult, error) {
	switch chain {
	case types.ChainFNSA, types.ChainTFNSA:
		for _, client := range p.CosmosClients() {
			return client.SendMultiSigTransaction(ctx, multiSigPubKey, msg, singedTxs...)
		}
	}
	return nil, fmt.Errorf("invalid chain")
}

func (p *Client) GetTransactionOption(ctx context.Context, chain string, from string) (*types.TransactionOption, error) {
	if _, ok := p.clients[chain]; !ok {
		return nil, NotFoundProxy
	}

	return p.clients[chain].GetTransactionOption(ctx, from)
}

func (p *Client) GetTransactionData(chain string, data *types.RequestTransaction) (*types.Transaction, error) {
	if _, ok := p.clients[chain]; !ok {
		return nil, NotFoundProxy
	}

	return p.clients[chain].GetTransactionData(data)
}

func (p *Client) BroadcastRawTx(ctx context.Context, chain string, rawTx []byte, mode txtypes.BroadcastMode) (*clienttypes.SendTxAsyncResult, error) {
	client, ok := p.clients[chain].(clienttypes.CosmosClient)
	if !ok {
		return nil, NotSupport
	}

	return client.BroadcastRawTx(ctx, rawTx, mode)
}
