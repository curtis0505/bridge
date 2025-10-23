package migrate

import (
	"github.com/kaiachain/kaia/common"
	"math/big"
)

// MigrateInput migrate(uint256 _amountFNSA, address to);
type MigrateInput struct {
	AmountFNSA *big.Int       `abi:"_amountFNSA"`
	To         common.Address `abi:"to"`
}
