package base

import (
	"context"
	"github.com/curtis0505/bridge/libs/client/chain/conf"
	clienttypes "github.com/curtis0505/bridge/libs/client/chain/types"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var cfg = conf.ClientConfig{
	Url:       "https://sepolia.base.org",
	ChainName: "Base",
	Chain:     "BASE",
	Proxy:     false,
}

func NewTestClient(t *testing.T) clienttypes.Client {
	c, err := NewClient(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, c)
	return c
}

func TestClient_BalanceAt(t *testing.T) {
	c := NewTestClient(t)
	b, err := c.BalanceAt(context.Background(), "0xb8882ab9B22Eed4d3A464655b379036987072a58", nil)
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
	b, err := NewTestClient(t).GetTransaction(context.Background(), "0x548afbda6cdee9e7a68b788584dd9dc09a1d1c02557217f817a02643f8b01161")
	assert.NoError(t, err)
	assert.NotNil(t, b)
	t.Log(b.From())
}

func TestClient_GetTransactionReceipt(t *testing.T) {
	b, err := NewTestClient(t).GetTransactionReceipt(context.Background(), "0x548afbda6cdee9e7a68b788584dd9dc09a1d1c02557217f817a02643f8b01161")
	assert.NoError(t, err)
	assert.NotNil(t, b)
	t.Log(b.Status())
}

func TestClient_EstimateGas(t *testing.T) {
	c := NewTestClient(t)
	from := "0x77777779f3Ef12bc8C3Ae9E67dC058eb84FCeb79"
	amount := big.NewInt(0)
	to := "0x01d2E89CD35260b2B453Ca1063d0354c0d25AD29" // dev2 uuid=84a65bd8c94a39d37178a39961504a45

	msg := types.CallMsg{
		From:  from,
		To:    to,
		Gas:   100000000,
		Value: amount,
		Data:  []byte{},
	}

	b, err := c.EstimateGas(context.Background(), msg)
	assert.NoError(t, err)
	assert.NotNil(t, b)
	t.Log(b)
}
