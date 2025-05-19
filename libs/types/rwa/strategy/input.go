package strategy

import (
	"github.com/curtis0505/bridge/libs/common"
	"math/big"
)

type InputDeposit struct {
	Assets    *big.Int       `abi:"assets"`
	MinShares *big.Int       `abi:"minShares"`
	Receiver  common.Address `abi:"receiver"`
}

type InputRedeem struct {
	Shares   *big.Int       `abi:"shares"`
	MinAsset *big.Int       `abi:"minAsset"`
	Receiver common.Address `abi:"receiver"`
	Owner    common.Address `abi:"owner"`
}
