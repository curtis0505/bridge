package tron

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/curtis0505/bridge/libs/client/chain/conf"
	"github.com/curtis0505/bridge/libs/common"
	troncommon "github.com/curtis0505/bridge/libs/common/tron"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/types/bridge/abi"
	"github.com/curtis0505/bridge/libs/types/bridge/abi/ether"
	"github.com/curtis0505/bridge/libs/types/token"
	troncore "github.com/curtis0505/grpc-idl/tron/core"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"testing"
	"time"
)

const (
	NodeGRPC       = "10.30.172.141:24000"
	TestPrivateKey = "e1838b5796ee7f19cce25fd8f6ca55717257f373e1738951e9692baddeffc7d1"
)

var (
	NileConfig = conf.ClientConfig{
		Chain:     types.ChainTRX,
		ChainName: types.ChainTRX,
		Url:       "grpc.nile.trongrid.io:50051",
	}

	MainnetConfig = conf.ClientConfig{
		Chain:     types.ChainTRX,
		ChainName: types.ChainTRX,
		Url:       "grpc.trongrid.io:50051",
	}
)

func Test_BalanceAt(t *testing.T) {
	c, err := NewClient(NileConfig)
	ctx := context.Background()
	assert.NoError(t, err)

	address := "TF3tMewZn3YM9c2Ujohst45Z1PMVJLWLUr"

	b, err := c.BalanceAt(ctx, address, nil)
	assert.NoError(t, err)
	t.Log("balance", b)
}

func Test_TransactionInfo(t *testing.T) {
	c, err := NewClient(NileConfig)
	ctx := context.Background()
	assert.NoError(t, err)

	txHash := "e055acbdfe759d9803f9779aeee6e6e5ed34b491ad68a4f5878b35f7ed9eaef4"
	r, err := c.GetTransactionReceipt(ctx, txHash)
	assert.NoError(t, err)

	info := r.Inner().(*troncore.TransactionInfo)

	t.Log(info)
}

func Test_Parameter(t *testing.T) {
	//c, err := NewClient(NodeGRPC, types.ChainTRX, types.ChainTRX)
	//ctx := context.Background()
	//assert.NoError(t, err)
	//
	//in := tronapi.EmptyMessage{}
	//list, err := c.c.GetChainParameters(ctx, &in)
	//
	//for _, witness := range list.GetChainParameter() {
	//	t.Log(witness.GetKey(), witness.GetValue())
	//}
}

func Test_ListWitnesses(t *testing.T) {
	//c, err := NewClient(NodeGRPC, types.ChainTRX, types.ChainTRX)
	//ctx := context.Background()
	//assert.NoError(t, err)
	//
	//in := tronapi.EmptyMessage{}
	//list, err := c.c.ListWitnesses(ctx, &in)
	//
	//for _, witness := range list.GetWitnesses() {
	//	t.Log(troncommon.FromBytes(witness.GetAddress()).String(), witness.Url, witness.GetVoteCount())
	//}
}

func Test_Transfer(t *testing.T) {
	c, err := NewClient(NileConfig)
	ctx := context.Background()
	assert.NoError(t, err)

	for {
		from := "TF3tMewZn3YM9c2Ujohst45Z1PMVJLWLUr"
		to := "TJP51ccL2W2YNk6DEdPziNQBPRBU5Akuox"
		amount := int64(1000000)
		transfer := &troncore.TransferContract{
			OwnerAddress: troncommon.FromBase58Unsafe(from),
			ToAddress:    troncommon.FromBase58Unsafe(to),
			Amount:       amount,
		}

		tx, err := c.CreateTransaction(ctx, transfer)
		assert.NoError(t, err)

		account, err := types.NewAccountFromPK(TestPrivateKey)
		assert.NoError(t, err)

		tx, err = account.Sign(tx, nil)
		assert.NoError(t, err)

		raw, err := proto.Marshal(tx.Inner().(*troncore.Transaction))
		assert.NoError(t, err)

		r, err := c.RawTxAsync(ctx, raw, nil)
		assert.NoError(t, err)

		t.Log("result:", r)
		t.Log("txHash:", tx.TxHash())
		time.Sleep(time.Second)
	}

}

func Test_WithdrawBalance(t *testing.T) {
	c, err := NewClient(NileConfig)
	ctx := context.Background()
	assert.NoError(t, err)

	from := "TF3tMewZn3YM9c2Ujohst45Z1PMVJLWLUr"
	transfer := &troncore.WithdrawBalanceContract{
		OwnerAddress: troncommon.FromBase58Unsafe(from),
	}

	tx, err := c.CreateTransaction(ctx, transfer)
	assert.NoError(t, err)

	account, err := types.NewAccountFromPK(TestPrivateKey)
	assert.NoError(t, err)

	tx, err = account.Sign(tx, nil)
	assert.NoError(t, err)

	raw, err := proto.Marshal(tx.Inner().(*troncore.Transaction))
	assert.NoError(t, err)

	r, err := c.RawTxAsync(ctx, raw, nil)
	assert.NoError(t, err)

	t.Log("result:", r.Error)
	t.Log("txHash:", tx.TxHash())
}

func Test_FreezeBalanceV2Contract(t *testing.T) {
	c, err := NewClient(NileConfig)
	ctx := context.Background()
	assert.NoError(t, err)

	from := "TF3tMewZn3YM9c2Ujohst45Z1PMVJLWLUr"
	amount := int64(1000000)
	transfer := &troncore.FreezeBalanceV2Contract{
		OwnerAddress:  troncommon.FromBase58Unsafe(from),
		FrozenBalance: amount,
	}

	tx, err := c.CreateTransaction(ctx, transfer)
	assert.NoError(t, err)

	account, err := types.NewAccountFromPK(TestPrivateKey)
	assert.NoError(t, err)

	tx, err = account.Sign(tx, nil)
	assert.NoError(t, err)

	raw, err := proto.Marshal(tx.Inner().(*troncore.Transaction))
	assert.NoError(t, err)

	r, err := c.RawTxAsync(ctx, raw, nil)
	assert.NoError(t, err)

	t.Log("result:", r)
	t.Log("txHash:", tx.TxHash())
}

func Test_VoteWitnessContract(t *testing.T) {
	c, err := NewClient(NileConfig)
	ctx := context.Background()
	assert.NoError(t, err)

	from := "TF3tMewZn3YM9c2Ujohst45Z1PMVJLWLUr"
	transfer := &troncore.VoteWitnessContract{
		OwnerAddress: troncommon.FromBase58Unsafe(from),
		Votes: []*troncore.VoteWitnessContract_Vote{
			{
				VoteAddress: troncommon.FromBase58Unsafe("TF3tMewZn3YM9c2Ujohst45Z1PMVJLWLUr"),
				VoteCount:   10,
			},
		},
	}

	tx, err := c.CreateTransaction(ctx, transfer)
	assert.NoError(t, err)

	account, err := types.NewAccountFromPK(TestPrivateKey)
	assert.NoError(t, err)

	tx, err = account.Sign(tx, nil)
	assert.NoError(t, err)

	raw, err := proto.Marshal(tx.Inner().(*troncore.Transaction))
	assert.NoError(t, err)

	r, err := c.RawTxAsync(ctx, raw, nil)
	assert.NoError(t, err)

	t.Log("result:", r)
	t.Log("txHash:", tx.TxHash())
}

func Test_TriggerSmartContract(t *testing.T) {
	c, err := NewClient(NileConfig)
	ctx := context.Background()
	assert.NoError(t, err)

	USDTAddress := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	HolderAddress := "TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9"

	var erc20Abi []map[string]interface{}
	json.Unmarshal([]byte(abi.ERC20Abi), &erc20Abi)

	// CallContract
	balanceOfResp, err := c.CallMsg(ctx, "", USDTAddress, "balanceOf", erc20Abi, common.HexToAddress(types.ChainTRX, HolderAddress))
	assert.NoError(t, err)

	t.Log("Balance:", balanceOfResp)

	receipt, err := c.GetTransactionReceipt(ctx, "713bc5a01f14dc58cdbe997633d4b6215b56e16053a8e24a6ac8be4532befb85")
	assert.NoError(t, err)

	var eventTransfer token.EventTransfer
	log := receipt.Logs()[0]

	// Set ABI to log
	log, err = log.Abi(ether.ERC20)
	assert.NoError(t, err)

	// Unmarshal to EventTransfer
	err = log.Unmarshal(&eventTransfer)
	assert.NoError(t, err)

	t.Log(eventTransfer)

	// Convert ETH based address To Base58 TRX address
	t.Log("From:", troncommon.FromEthAddress(eventTransfer.From.Bytes()))
	t.Log("To:", troncommon.FromEthAddress(eventTransfer.To.Bytes()))
	t.Log("Value:", eventTransfer.Value)
}

func Test_GetAccount(t *testing.T) {
	c, err := NewClient(NileConfig)
	ctx := context.Background()
	assert.NoError(t, err)

	address := "TSNSh61A6Sq7xhZXZ9GhHa7xh4f7QqLDnr"

	account, err := c.GetAccount(ctx, address)
	if err != nil {
		t.Error(err)
		return
	}

	var bandwidthStakedAmountV1 int64
	var bandwidthDeletedStakedAmountV1 int64
	var energyStakedAmountV1 int64
	var energyDeletedStakedAmountV1 int64
	var stakedAmountV1 int64
	// bandwidth
	for _, v1 := range account.GetFrozen() {
		bandwidthStakedAmountV1 += v1.GetFrozenBalance()
	}
	// energy
	energyStakedAmountV1 += account.GetAccountResource().GetFrozenBalanceForEnergy().GetFrozenBalance()
	// delegated bandwidth
	bandwidthDeletedStakedAmountV1 += account.GetDelegatedFrozenBalanceForBandwidth()
	// delegated energy
	energyDeletedStakedAmountV1 += account.GetAccountResource().GetDelegatedFrozenBalanceForEnergy()
	stakedAmountV1 = bandwidthStakedAmountV1 + bandwidthDeletedStakedAmountV1 + energyStakedAmountV1 + energyDeletedStakedAmountV1

	var bandwidthStakedAmount int64
	var bandwidthDeletedStakedAmount int64
	var energyStakedAmount int64
	var energyDeletedStakedAmount int64
	var stakedAmount int64

	for _, v2 := range account.GetFrozenV2() {
		switch v2.GetType() {
		case troncore.ResourceCode_BANDWIDTH:
			bandwidthStakedAmount += v2.GetAmount()
		case troncore.ResourceCode_ENERGY:
			energyStakedAmount += v2.GetAmount()
		default:
			t.Log(v2.String())
		}
		stakedAmount += v2.GetAmount()
	}
	bandwidthDeletedStakedAmount = account.GetDelegatedFrozenV2BalanceForBandwidth()
	energyDeletedStakedAmount = account.GetAccountResource().GetDelegatedFrozenV2BalanceForEnergy()
	stakedAmount = bandwidthStakedAmount + bandwidthDeletedStakedAmount + energyStakedAmount + energyDeletedStakedAmount

	//
	var unstakingAmount int64
	var withdrawableAmount int64
	for _, v2 := range account.GetUnfrozenV2() {
		if time.Now().UnixMilli() > v2.GetUnfreezeExpireTime() {
			withdrawableAmount += v2.GetUnfreezeAmount()
		} else {
			unstakingAmount += v2.GetUnfreezeAmount()
		}
	}

	//

	fmt.Println("staking 1.0")
	fmt.Printf("staked amount: %d\n", stakedAmountV1/1000000)
	fmt.Printf("ㄴenergy: %d, (my energy: %d, delegated energy: %d)\n", (energyStakedAmountV1+energyDeletedStakedAmountV1)/1000000, energyStakedAmountV1/1000000, energyDeletedStakedAmountV1/1000000)
	fmt.Printf("ㄴbandwitdh: %d, (my bandwitdh: %d, delegated bandwitdh: %d)\n", (bandwidthStakedAmountV1+bandwidthDeletedStakedAmountV1)/1000000, bandwidthStakedAmountV1/1000000, bandwidthDeletedStakedAmountV1/1000000)
	fmt.Println()
	fmt.Println()
	fmt.Println("staking 2.0")
	fmt.Printf("staked amount: %d\n", stakedAmount/1000000)
	fmt.Printf("ㄴenergy: %d, (my energy: %d, delegated energy: %d)\n", (energyStakedAmount+energyDeletedStakedAmount)/1000000, energyStakedAmount/1000000, energyDeletedStakedAmount/1000000)
	fmt.Printf("ㄴbandwitdh: %d, (my bandwitdh: %d, delegated bandwitdh: %d)\n", (bandwidthStakedAmount+bandwidthDeletedStakedAmount)/1000000, bandwidthStakedAmount/1000000, bandwidthDeletedStakedAmount/1000000)
	fmt.Println()
	fmt.Println("unstkaing")
	fmt.Printf("wait unstaking amount: %d\n", unstakingAmount/1000000)
	fmt.Printf("withdrawable unstaking amount: %d\n", withdrawableAmount/1000000)
	fmt.Println()
	fmt.Println()
	fmt.Println("vote")
	for _, vote := range account.GetVotes() {
		fmt.Printf("address: %s, vote: %d\n", troncommon.FromBytes(vote.GetVoteAddress()), vote.GetVoteCount())
	}

	fmt.Println()
	fmt.Println("특징")
	fmt.Println("- delegate 하면 claim(delegated 한거 회수) 후 unstaking(unfroze) 할 수 있다.")
	fmt.Println("- delegate 한 금액도 vote 가능하다.")
	fmt.Println()

	fmt.Println("account.GetTronPower()", account.GetTronPower())
	fmt.Println("account.GetOldTronPower()", account.GetOldTronPower())

}
