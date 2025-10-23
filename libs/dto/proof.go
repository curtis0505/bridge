package dto

import (
	"math/big"

	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
)

type ProofBridgeReq struct {
	SenderVault     []byte              `json:"senderVault"`
	SenderMinter    []byte              `json:"senderMinter"`
	SenderHash      []byte              `json:"senderHash"`
	SenderChainName string              `json:"senderChainName"`
	MsgReceiver     string              `json:"msgReceiver"`
	SendTokenAddr   string              `json:"sendTokenAddr"`
	RcvAmount       *big.Int            `json:"rcvAmount"`
	RcvChainName    string              `json:"rcvChainName"`
	Version         bridgetypes.Version `json:"version"`
}
