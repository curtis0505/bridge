package types

import (
	"github.com/curtis0505/bridge/apps/managers/types"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
)

type PendingTxRequest struct {
	Chain  string `json:"chain"`
	From   string `json:"from"`
	TxHash string `json:"txHash"`
}

type SubmitTransactionResponse struct {
	*types.ResponseHeader
}

type CosmosMultiSigTx struct {
	Chain  string
	TxHash string

	EventName string
	EventLog  bridgetypes.CommonEventLog
}

func NewCosmosMultiSigTx(chain, txHash, eventName string, eventLog bridgetypes.CommonEventLog) CosmosMultiSigTx {
	return CosmosMultiSigTx{
		Chain:     chain,
		TxHash:    txHash,
		EventName: eventName,
		EventLog:  eventLog,
	}
}
