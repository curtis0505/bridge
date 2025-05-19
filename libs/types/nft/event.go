package nft

import (
	"github.com/curtis0505/bridge/libs/common"
	"math/big"
)

type EventTransfer struct {
	// Common
	From common.Address
	To   common.Address

	// ERC-721
	TokenId *big.Int

	// ERC-1155
	Operator common.Address // Signer
	Id       *big.Int       // TokenId
	Value    *big.Int       // Amount
}

func (event EventTransfer) Amount() *big.Int {
	if event.Value != nil {
		return event.Value
	}

	return big.NewInt(0)
}

func (event EventTransfer) TokenID() *big.Int {
	if event.TokenId != nil {
		return event.TokenId
	}

	if event.Id != nil {
		return event.Id
	}

	return big.NewInt(0)
}

type EventTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
}
