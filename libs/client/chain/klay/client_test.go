package klay

import (
	"context"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/curtis0505/bridge/libs/client/chain/conf"
	clienttypes "github.com/curtis0505/bridge/libs/client/chain/types"
	"github.com/curtis0505/bridge/libs/types"
)

func NewTestClient(t *testing.T) clienttypes.KlayClient {
	c, err := NewClient(conf.ClientConfig{
		//Url:       "https://klay.dev.neopin.io",
		Url:   "http://3.36.36.184:13100",
		Chain: types.ChainKLAY,
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
	b, err := NewTestClient(t).GetTransaction(context.Background(), "0x10fdea4bd8c09c2d29d33894408db97a0e09cfe4255b83701c084be2db9c3e81")
	assert.NoError(t, err)
	assert.NotNil(t, b)
	t.Log(b.From())
}

func TestClient_GetTransactionReceipt(t *testing.T) {
	b, err := NewTestClient(t).GetTransactionReceipt(context.Background(), "0x10fdea4bd8c09c2d29d33894408db97a0e09cfe4255b83701c084be2db9c3e81")
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
