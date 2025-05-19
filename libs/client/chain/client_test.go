package chain

import (
	"context"
	"fmt"
	arbitrumcommon "github.com/curtis0505/arbitrum/common"
	arbitrumclient "github.com/curtis0505/arbitrum/ethclient"
	"github.com/curtis0505/bridge/libs/client/chain/cosmos"
	"github.com/curtis0505/bridge/libs/client/chain/tron"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/util"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestNewProxy(t *testing.T) {
	NewTestProxy(t)
}

func TestProxy_BlockNumber(t *testing.T) {
	proxy := NewTestProxy(t)

	for chain, _ := range testProxy {
		blockNumber, err := proxy.BlockNumber(context.Background(), chain)
		assert.NoError(t, err)
		t.Log(chain, blockNumber)
	}
}

func TestProxy_BalanceAt(t *testing.T) {
	proxy := NewTestProxy(t)

	for chain, address := range testAddress {
		balance, err := proxy.BalanceAt(context.Background(), chain, address, nil)
		if err != nil {
			continue
		}
		assert.NoError(t, err)
		t.Log(chain, balance)
	}
}

func TestProxy_GetTransaction(t *testing.T) {
	proxy := NewTestProxy(t)

	for chain, txHash := range testTxHash {
		tx, err := proxy.GetTransaction(context.Background(), chain, txHash)
		if err != nil {
			continue
		}
		assert.NoError(t, err)
		t.Log(tx.From())
	}
}

func TestProxy_GetTransactionReceipt(t *testing.T) {
	proxy := NewTestProxy(t)

	for chain, txHash := range testTxHash {
		receipt, err := proxy.GetTransactionReceipt(context.Background(), chain, txHash)
		if err != nil {
			continue
		}
		assert.NoError(t, err)
		t.Log(chain, receipt.Status(), receipt.GasUsed(), receipt.BlockNumber())
	}
}

func TestProxy_Cosmos(t *testing.T) {
	proxy := NewTestProxy(t)
	client, err := cosmos.ProxyClient(proxy, types.ChainATOM)
	assert.NoError(t, err)

	staking, err := client.GetStaking(context.Background(), testAddress[types.ChainATOM])
	t.Log(
		"staked:", util.ToEtherWithDecimal(staking.GetStakedAmount(), 6),
		"unstaked:", util.ToEtherWithDecimal(staking.GetWithdrawalAmount(), 6),
		"reward:", util.ToEtherWithDecimal(staking.GetRewardAmount(), 6),
	)
	for _, unbond := range staking.Withdrawal {
		t.Log(unbond)
	}

	t.Log(client.GetValidatorApr(context.Background(), "cosmosvaloper1c4k24jzduc365kywrsvf5ujz4ya6mwympnc4en"))
}

func TestProxy_Tron(t *testing.T) {
	proxy := NewTestProxy(t)
	client, err := tron.ProxyClient(proxy)
	assert.NoError(t, err)
	t.Log(client.ChainName())
}

func TestArbitrum(t *testing.T) {
	t.Log(arbitrumcommon.HexToAddress("0x00"))

	url := "https://arb-sepolia.g.alchemy.com/v2/w_6Kf_v-FaWG4BSE6WFkgttsLmz-xlPV"
	dial, err := arbitrumclient.Dial(url)
	if err != nil {
		return
	}

	ctx := context.Background()
	number, err := dial.BlockNumber(ctx)
	if err != nil {
		return
	}

	byNumber, err := dial.BlockByNumber(ctx, big.NewInt(int64(number)))
	if err != nil {
		return
	}
	fmt.Printf("%+v\n", byNumber)
	t.Log("byNumber", byNumber)

	for _, tx := range byNumber.Transactions() {
		json, err := tx.MarshalJSON()
		if err != nil {
			return
		}

		t.Log("byNumber", string(json))
		t.Log("byNumber", tx.Hash().String())
	}

}
