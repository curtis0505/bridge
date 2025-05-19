package types

import (
	"github.com/curtis0505/bridge/libs/common"
	"math/big"
)

type BridgeBalanceHistoryData struct {
	Chain       string   `bson:"chain" json:"chain"`
	Name        string   `json:"name" bson:"name"`
	Address     string   `bson:"address" json:"address"`
	CurrencyId  string   `bson:"currency_id" json:"currencyId"`
	GroupId     string   `bson:"group_id" json:"groupId"`
	BlockNumber int64    `bson:"block_number" json:"blockNumber"`
	Event       string   `bson:"event" json:"event"`
	Balance     *big.Int `bson:"balance" json:"balance,omitempty"`
}

type MultiSigTransaction struct {
	Address  common.Address `json:"address"`
	Value    *big.Int       `json:"value"`
	Data     []byte         `json:"data"`
	Proof    [32]byte       `json:"proof"`
	Required *big.Int       `json:"required"`
	Executed bool           `json:"executed"`
}

func NewMultiSigTransaction(tx []interface{}) *MultiSigTransaction {
	if len(tx) == 2 {
		return &MultiSigTransaction{
			Proof:    tx[0].([32]byte),
			Executed: tx[1].(bool),
		}
	}

	return &MultiSigTransaction{
		Address:  tx[0].(common.Address),
		Value:    tx[1].(*big.Int),
		Data:     tx[2].([]byte),
		Proof:    tx[3].([32]byte),
		Required: tx[4].(*big.Int),
		Executed: tx[5].(bool),
	}
}
