package stable

import (
	"github.com/curtis0505/bridge/libs/common"
	"math/big"
)

type InputDeposit struct {
	AssetAddr  common.Address `abi:"_assetAddr"`
	Assets     *big.Int       `abi:"_assets"`
	DaiPercent uint8          `abi:"_daiPercent"`
	MinDAI     *big.Int       `abi:"_minDAI"`
	MinUSDe    *big.Int       `abi:"_minUSDe"`
}

type InputWithdraw struct {
	SDAIAmount  *big.Int `abi:"_sDAIAmount"`
	SUSDeAmount *big.Int `abi:"_sUSDeAmount"`
	MinBoostDAI *big.Int `abi:"_minBoostDAI"`
}

type InputClaimForBoost struct {
	MinDAI *big.Int `abi:"_minDAI"`
}
