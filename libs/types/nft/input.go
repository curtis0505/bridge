package nft

import (
	"github.com/curtis0505/bridge/libs/common"
	"math/big"
)

// InputTransferFrom
// ERC 721
type InputTransferFrom struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
}

// InputTransferSingle
// ERC 1155
type InputTransferSingle struct {
	From   common.Address
	To     common.Address
	Id     *big.Int
	Amount *big.Int
	Data   []byte
}

// InputTransferBatch
// ERC 1155
type InputTransferBatch struct {
	From    common.Address
	To      common.Address
	Ids     []*big.Int
	Amounts []*big.Int
	Data    []byte
}
