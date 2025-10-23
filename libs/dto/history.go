package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ScBridgeTransferHistoryData struct {
	Chain      string               `bson:"chain" json:"chain"`
	From       string               `bson:"from" json:"from"`
	To         string               `bson:"to" json:"to"`
	TxHash     string               `bson:"tx_hash" json:"txHash"`
	Cate       string               `bson:"cate" json:"cate"`
	Symbol     string               `bson:"symbol" json:"symbol"`
	CurrencyId string               `bson:"currency_id" json:"currencyId"`
	GroupId    string               `bson:"group_id" json:"groupId"`
	Amount     primitive.Decimal128 `bson:"amount" json:"amount"`
	FromChain  string               `bson:"from_chain" json:"fromChain,omitempty"`
	FromTxHash string               `bson:"from_tx_hash" json:"fromTxHash,omitempty"`
	ToChain    string               `bson:"to_chain" json:"toChain,omitempty"`
	ToTxHash   string               `bson:"to_tx_hash" json:"toTxHash,omitempty"`
	CreateAt   time.Time            `bson:"create_at" json:"createAt"`
}
