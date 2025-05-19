package migrate

import (
	"github.com/klaytn/klaytn/common"
	"math/big"
)

// EventMigrate Migrate(address indexed src, uint256 amountFNSA, uint256 amountKaia);
type EventMigrate struct {
	Src        common.Address
	AmountFNSA *big.Int `abi:"amountFNSA"`
	AmountKaia *big.Int `abi:"amountKaia"`
}
