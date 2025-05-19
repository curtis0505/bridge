package ether

import (
	"context"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/curtis0505/bridge/libs/client/chain/conf"
	clienttypes "github.com/curtis0505/bridge/libs/client/chain/types"
	"github.com/curtis0505/bridge/libs/types"
)

func NewTestClient(t *testing.T) clienttypes.EtherClient {
	c, err := NewClient(conf.ClientConfig{
		Url:       "http://node1.dq.neopin.pmang.cloud:12100",
		Chain:     types.ChainMATIC,
		ChainName: "a-polygon",
	})
	assert.NoError(t, err)
	assert.NotNil(t, c)
	return c
}

func TestNewClient(t *testing.T) {
	NewTestClient(t)
}

func TestClient_BalanceAt(t *testing.T) {
	c := NewTestClient(t)
	b, err := c.BalanceAt(context.Background(), "0x01d2E89CD35260b2B453Ca1063d0354c0d25AD29", nil)
	assert.NoError(t, err)
	assert.NotNil(t, b)
	t.Log(b)
}

func TestClient_BlockNumber(t *testing.T) {
	c := NewTestClient(t)
	b, err := c.BlockNumber(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, b)
	t.Log(b)
}

func TestClient_GasPrice(t *testing.T) {
	c := NewTestClient(t)
	b, err := c.GasPrice(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, b)
	t.Log(b)
}

func TestClient_GetTransaction(t *testing.T) {
	b, err := NewTestClient(t).GetTransaction(context.Background(), "0xa585baf86a61b01f405fd2ab820b270debcf32e3804eb0c90e0570b2818884d1")
	assert.NoError(t, err)
	assert.NotNil(t, b)
	t.Log(b.From())
}

func TestClient_GetTransactionReceipt(t *testing.T) {
	b, err := NewTestClient(t).GetTransactionReceipt(context.Background(), "0xa585baf86a61b01f405fd2ab820b270debcf32e3804eb0c90e0570b2818884d1")
	assert.NoError(t, err)
	assert.NotNil(t, b)
	t.Log(b.Status())
}

func TestClient_EstimateGas(t *testing.T) {
	c := NewTestClient(t)
	from := "0x01d2E89CD35260b2B453Ca1063d0354c0d25AD29"
	amount := big.NewInt(1e9)
	to := "0x41b020615C975155CBB3E36C8c174dcD56Aa6ac3" // dev2 uuid=84a65bd8c94a39d37178a39961504a45

	msg := types.CallMsg{
		From:  from,
		To:    to,
		Gas:   100000,
		Value: amount,
		Data:  []byte{},
	}

	b, err := c.EstimateGas(context.Background(), msg)
	assert.NoError(t, err)
	assert.NotNil(t, b)
	t.Log(b)
}
