package tron

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/curtis0505/bridge/libs/client/chain/conf"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
	"time"
	"unsafe"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"math/big"

	clienttypes "github.com/curtis0505/bridge/libs/client/chain/types"
	"github.com/curtis0505/bridge/libs/common"
	troncommon "github.com/curtis0505/bridge/libs/common/tron"
	"github.com/curtis0505/bridge/libs/types"
	tronapi "github.com/curtis0505/grpc-idl/tron/api"
	troncore "github.com/curtis0505/grpc-idl/tron/core"
	etherabi "github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	_ clienttypes.Client        = &client{}
	_ clienttypes.StakingClient = &client{}
	_ clienttypes.TronClient    = &client{}
)

type client struct {
	cc *grpc.ClientConn

	c tronapi.WalletClient

	networkId *big.Int
	chain     string
}

func NewClient(config conf.ClientConfig) (clienttypes.TronClient, error) {
	c := &client{
		chain: config.Chain,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var err error
	c.cc, err = grpc.DialContext(ctx, config.Url, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, types.WrapError("DialContext", err)
	}

	c.c = tronapi.NewWalletClient(c.cc)

	return c, nil
}

func ProxyClient(proxy clienttypes.Proxy) (clienttypes.TronClient, error) {
	c := proxy.ProxyClient(types.ChainTRX)
	if c == nil {
		return nil, fmt.Errorf("not found proxy")
	}

	client, ok := c.(clienttypes.TronClient)
	if !ok {
		return nil, fmt.Errorf("failed to casting client")
	}

	return client, nil
}

func (c *client) Chain() string              { return c.chain }
func (c *client) ChainType() types.ChainType { return types.GetChainType(c.chain) }
func (c *client) GetChainID(ctx context.Context) (*big.Int, error) {
	return nil, clienttypes.NotImplemented
}
func (c *client) NetworkId(ctx context.Context) (*big.Int, error) {
	return nil, clienttypes.NotImplemented
}

func (c *client) CallContract(ctx context.Context, msg types.CallMsg, blockNumber *big.Int) ([]byte, error) {
	contractAddress, err := troncommon.FromBase58(msg.From)
	if err != nil {
		return nil, err
	}

	in := troncore.TriggerSmartContract{
		ContractAddress: contractAddress,
		Data:            msg.Data,
	}

	resp, err := c.c.TriggerConstantContract(ctx, &in)
	if err != nil {
		return nil, err
	}

	var b []byte
	for _, result := range resp.GetConstantResult() {
		b = append(b, result...)
	}

	return b, nil
}

func (c *client) PendingNonceAt(ctx context.Context, address string) (uint64, error) {
	return 0, clienttypes.NotImplemented
}

func (c *client) NonceAt(ctx context.Context, address string) (uint64, error) {
	return c.PendingNonceAt(ctx, address)
}

func (c *client) GasPrice(ctx context.Context) (*big.Int, error) {
	in := tronapi.EmptyMessage{}
	resp, err := c.c.GetChainParameters(ctx, &in)
	if err != nil {
		return nil, err
	}

	for _, param := range resp.GetChainParameter() {
		return big.NewInt(param.Value), nil
	}

	return nil, clienttypes.NotImplemented
}

func (c *client) RawTxAsync(ctx context.Context, rawTx []byte, rawProxyRequest []byte) (*clienttypes.SendTxAsyncResult, error) {
	var tx troncore.Transaction
	if err := proto.Unmarshal(rawTx, &tx); err != nil {
		return nil, err
	}

	r, err := c.c.BroadcastTransaction(ctx, &tx)
	if err != nil {
		return nil, err
	}

	sha256hash := sha256.New()
	raw, err := proto.Marshal(tx.GetRawData())
	if err != nil {
		return nil, err
	}
	sha256hash.Write(raw)
	txHash := sha256hash.Sum(nil)
	var h common.Hash
	copy(h[:], txHash)

	var result clienttypes.SendTxAsyncResult
	if r.Result {
		result.Result = clienttypes.SendTxResultType_Success
	} else {
		result.Result = clienttypes.SendTxResultType_SendTxError
		result.Error = string(r.Message)
	}
	result.TxHash = h

	return &result, nil
}

func (c *client) RawTxAsyncByTx(ctx context.Context, tx *types.Transaction) (*clienttypes.SendTxAsyncResult, error) {
	_tx, err := tx.TronTransaction()
	if err != nil {
		return nil, err
	}

	r, err := c.c.BroadcastTransaction(ctx, _tx)
	if err != nil {
		return nil, err
	}

	sha256hash := sha256.New()
	raw, err := proto.Marshal(_tx.GetRawData())
	if err != nil {
		return nil, err
	}
	sha256hash.Write(raw)
	txHash := sha256hash.Sum(nil)
	var h common.Hash
	copy(h[:], txHash)

	var result clienttypes.SendTxAsyncResult
	if r.Result {
		result.Result = clienttypes.SendTxResultType_Success
	} else {
		result.Result = clienttypes.SendTxResultType_SendTxError
		result.Error = string(r.Message)
	}
	result.TxHash = h

	return &result, nil
}

func (c *client) TxAsync(ctx context.Context, rlpTx string, proxyRequest clienttypes.ProxyRequest) (*clienttypes.SendTxAsyncResult, error) {
	b, err := hex.DecodeString(rlpTx)
	if err != nil {
		return nil, err
	}

	return c.RawTxAsync(ctx, b, nil)
}

func (c *client) BalanceAt(ctx context.Context, address string, blockNumber *big.Int) (*big.Int, error) {
	addr, err := troncommon.FromBase58(address)
	if err != nil {
		return nil, err
	}

	in := troncore.Account{Address: addr}
	resp, err := c.c.GetAccount(ctx, &in)
	if err != nil {
		return nil, err
	}
	return big.NewInt(resp.Balance), nil
}

func (c *client) BlockNumber(ctx context.Context) (*big.Int, error) {
	block, err := c.c.GetNowBlock2(ctx, &tronapi.EmptyMessage{})
	if err != nil {
		return nil, err
	}

	return big.NewInt(block.GetBlockHeader().GetRawData().Number), nil
}

func (c *client) GetTransactionReceipt(ctx context.Context, txHash string) (*types.Receipt, error) {
	b, err := hex.DecodeString(txHash)
	if err != nil {
		return nil, err
	}
	in := tronapi.BytesMessage{
		Value: b,
	}

	info, err := c.c.GetTransactionInfoById(ctx, &in)
	if err != nil {
		return nil, err
	}

	return types.NewReceipt(info, types.ChainTRX), nil
}

// Tron은 진짜 좀더 고민!
func (c *client) GetHeaderByHash(ctx context.Context, txHash string) (*types.Header, error) {
	b, err := hex.DecodeString(txHash)
	if err != nil {
		return nil, err
	}

	block, err := c.c.GetBlockById(ctx, &tronapi.BytesMessage{Value: b})
	if err != nil {
		return nil, err
	}

	return types.NewHeader(block.BlockHeader), nil
}

func (c *client) GetTransaction(ctx context.Context, txHash string) (*types.Transaction, error) {
	b, err := hex.DecodeString(txHash)
	if err != nil {
		return nil, err
	}
	in := tronapi.BytesMessage{
		Value: b,
	}

	tx, err := c.c.GetTransactionById(ctx, &in)
	if err != nil {
		return nil, err
	}

	return types.NewTransaction(tx, types.ChainTRX), nil
}

func (c *client) GetTransactionWithReceipt(ctx context.Context, txHash string) (*types.Transaction, bool, error) {
	return nil, false, clienttypes.NotImplemented
}

func (c *client) CallMsg(ctx context.Context, from, to, methodName string, abi []map[string]interface{}, args ...interface{}) ([]interface{}, error) {
	if from == "" {
		from = to
	}

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

// ----------
// TRX implementation
// ----------

func (c *client) GetAccount(ctx context.Context, address string) (*troncore.Account, error) {
	addr, err := troncommon.FromBase58(address)
	if err != nil {
		return nil, err
	}

	in := troncore.Account{Address: addr}
	resp, err := c.c.GetAccount(ctx, &in)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *client) GetReward(ctx context.Context, address string) (int64, error) {
	addr, err := troncommon.FromBase58(address)
	if err != nil {
		return 0, err
	}

	in := tronapi.BytesMessage{
		Value: addr,
	}
	resp, err := c.c.GetRewardInfo(ctx, &in)
	if err != nil {
		return 0, err
	}

	return resp.Num, nil
}

func (c *client) GetStaking(ctx context.Context, address string) (*clienttypes.Staking, error) {
	return nil, clienttypes.NotImplemented
}

func (c *client) GetTronStaking(ctx context.Context, address string, validatorAddress string) (*clienttypes.Staking, error) {
	account, err := c.GetAccount(ctx, address)
	if err != nil {
		return nil, err
	}
	reward, _ := c.GetReward(ctx, address)

	staking := &clienttypes.Staking{
		Amount:       big.NewInt(0),
		AmountLegacy: big.NewInt(0),
		Reward: clienttypes.Reward{
			Chain:         c.chain,
			Amount:        big.NewInt(reward),
			ClaimableTime: time.Unix(account.GetLatestWithdrawTime()/Epoch+86400, 0),
		},
	}

	for _, vote := range account.GetVotes() {
		if troncommon.FromBytes(vote.GetVoteAddress()).String() == validatorAddress {
			staking.Amount = new(big.Int).Mul(big.NewInt(vote.GetVoteCount()), big.NewInt(1000000))
		}
	}

	var bandwidthFrozenAmountV1 int64
	var energyFrozenAmountV1 int64
	var bandwidthFrozenAmountV2 int64
	var energyFrozenAmountV2 int64

	// V1 Bandwidth
	for _, frozen := range account.GetFrozen() {
		bandwidthFrozenAmountV1 += frozen.GetFrozenBalance()
	}

	// V1 Energy
	if account.GetAccountResource().GetFrozenBalanceForEnergy() != nil {
		energyFrozenAmountV1 = account.GetAccountResource().GetFrozenBalanceForEnergy().GetFrozenBalance()
	}

	// V2 Bandwidth, Energy
	for _, frozen := range account.GetFrozenV2() {
		switch frozen.GetType() {
		case troncore.ResourceCode_ENERGY:
			energyFrozenAmountV2 += frozen.GetAmount()
		case troncore.ResourceCode_BANDWIDTH:
			bandwidthFrozenAmountV2 += frozen.GetAmount()
		case troncore.ResourceCode_TRON_POWER:
			// TODO: Check Tron Power
		}
	}
	staking.TronStaking = &clienttypes.TronStaking{
		BandwidthFrozenAmountV1: big.NewInt(bandwidthFrozenAmountV1),
		EnergyFrozenAmountV1:    big.NewInt(energyFrozenAmountV1),
		BandwidthFrozenAmountV2: big.NewInt(bandwidthFrozenAmountV2),
		EnergyFrozenAmountV2:    big.NewInt(energyFrozenAmountV2),
	}

	currentTime := time.Now()
	for _, unfrozen := range account.GetUnfrozenV2() {
		if currentTime.Unix() > unfrozen.GetUnfreezeExpireTime()/Epoch {
			// 언스테이킹 후 14일 경과
			staking.Claimable = append(staking.Claimable, clienttypes.Claimable{
				Chain:  c.chain,
				Amount: big.NewInt(unfrozen.GetUnfreezeAmount()),
			})
		} else {
			// 언스테이킹 후 14일 대기
			staking.Withdrawal = append(staking.Withdrawal, clienttypes.Withdrawal{
				Chain:          c.chain,
				Amount:         big.NewInt(unfrozen.GetUnfreezeAmount()),
				CompletionTime: time.Unix(unfrozen.GetUnfreezeExpireTime()/Epoch, 0),
			})
		}
	}

	return staking, nil
}

func (c *client) CreateTransferTransaction(ctx context.Context, from, to string, amount *big.Int) (*types.Transaction, error) {
	fromAddress, err := troncommon.FromBase58(from)
	if err != nil {
		return nil, err
	}
	toAddress, err := troncommon.FromBase58(to)
	if err != nil {
		return nil, err
	}

	contract := &troncore.TransferContract{
		OwnerAddress: fromAddress,
		ToAddress:    toAddress,
		Amount:       amount.Int64(),
	}

	tx, err := c.c.CreateTransaction2(ctx, contract)
	if err != nil {
		return nil, err
	}

	return types.NewTransaction(tx.Transaction, types.ChainTRX), nil
}

func (c *client) GetBlock(ctx context.Context) (*tronapi.BlockExtention, error) {
	return c.c.GetNowBlock2(ctx, &tronapi.EmptyMessage{})
}

func (c *client) CreateTransaction(ctx context.Context, contract proto.Message) (*types.Transaction, error) {
	tx, err := c.createTransaction(ctx, contract)
	if err != nil {
		return nil, err
	}
	return types.NewTransaction(tx, types.ChainTRX), nil
}

func (c *client) createTransaction(ctx context.Context, contract proto.Message) (*troncore.Transaction, error) {
	block, err := c.c.GetNowBlock2(ctx, &tronapi.EmptyMessage{})
	if err != nil {
		return nil, err
	}

	sha256hash := sha256.New()
	rawHeader, err := proto.Marshal(block.GetBlockHeader().GetRawData())
	if err != nil {
		return nil, err
	}
	sha256hash.Write(rawHeader)
	blockHash := sha256hash.Sum(nil)

	IntToByteArray := func(num int64) []byte {
		size := int(unsafe.Sizeof(num))
		arr := make([]byte, size)
		for i := 0; i < size; i++ {
			byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(i)))
			arr[size-1-i] = byt
		}
		return arr
	}

	blockBytes := IntToByteArray(block.GetBlockHeader().GetRawData().GetNumber())

	protoAny, err := anypb.New(contract)
	if err != nil {
		return nil, err
	}
	protoName := protoAny.MessageName().Name()

	tx := &troncore.Transaction{
		RawData: &troncore.TransactionRaw{
			RefBlockHash:  blockHash[8:16],
			RefBlockBytes: blockBytes[6:8],
			RefBlockNum:   block.GetBlockHeader().GetRawData().GetNumber(),
			Timestamp:     time.Now().Unix() * 1000,
			Expiration:    block.GetBlockHeader().GetRawData().GetTimestamp() + 10*60*60*1000,
			Contract: []*troncore.Transaction_Contract{
				{
					Parameter: protoAny,
					Type:      troncore.Transaction_Contract_ContractType(troncore.Transaction_Contract_ContractType_value[string(protoName)]),
				},
			},
		},
	}

	return tx, nil
}

func (c *client) EstimateGas(ctx context.Context, callMsg types.CallMsg) (*big.Int, error) {
	panic("implement me")
}

func (c *client) HeaderByNumber(ctx context.Context, blockNumber *big.Int) (*types.Header, error) {
	panic("implement me")
}

func (c *client) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	panic("implement me")
}

func (c *client) GetTransactionOption(ctx context.Context, from string) (*types.TransactionOption, error) {
	return nil, clienttypes.NotImplemented
}

func (c *client) GetTransactionData(data *types.RequestTransaction) (*types.Transaction, error) {
	return nil, clienttypes.NotImplemented
}
