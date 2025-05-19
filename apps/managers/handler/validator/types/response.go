package types

import (
	"github.com/curtis0505/bridge/apps/managers/types"
	mongoentity "github.com/curtis0505/bridge/libs/entity/mongo"
)

type TxResponse struct {
	*types.ResponseHeader

	Status         string `json:"status"`
	From           string `json:"from"`
	FromChain      string `json:"fromChain"`
	FromTxHash     string `json:"fromTxHash"`
	FromCurrencyId string `json:"fromCurrencyId"`
	To             string `json:"to"`
	ToChain        string `json:"toChain"`
	ToTxHash       string `json:"toTxHash"`
	ToCurrencyId   string `json:"toCurrencyId"`
	GroupId        string `json:"groupId"`
	Amount         string `json:"amount"`
	Tax            string `json:"tax"`
	Confirmation   int    `json:"confirmation"`
	Execute        bool   `json:"execute"`

	BridgeHistoryList   []*BridgeTxInfo                      `json:"bridgeHistoryList,omitempty"`
	TransferHistoryList []*mongoentity.BridgeTransferHistory `json:"transferHistoryList,omitempty"`
}

func (tx *TxResponse) AddBridgeHistory(summary *ValidatorSummary, history *mongoentity.BridgeTxHistory) {
	tx.BridgeHistoryList = append(tx.BridgeHistoryList, &BridgeTxInfo{
		ValidatorInfo: summary,
		TxHistory:     history,
	})
}

func (tx *TxResponse) AddBridgeTransferHistory(history *mongoentity.BridgeTransferHistory) {
	tx.TransferHistoryList = append(tx.TransferHistoryList, history)
}

type BridgeTxInfo struct {
	ValidatorInfo *ValidatorSummary            `json:"validatorInfo,omitempty"`
	TxHistory     *mongoentity.BridgeTxHistory `json:"tx"`
}

type ValidatorStatisticsResponse struct {
	*types.ResponseHeader
}
