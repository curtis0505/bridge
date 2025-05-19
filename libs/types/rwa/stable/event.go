package stable

import (
	"github.com/curtis0505/bridge/libs/common"
	"math/big"
)

type EventDeposit struct {
	From        common.Address
	TokenAddr   common.Address
	DAIAmount   *big.Int `abi:"_daiAmount"`
	DAIPercent  uint8    `abi:"_daiPercent"`
	USDeAmount  *big.Int `abi:"_usdeAmount"`
	SDAIAmount  *big.Int `abi:"_sDaiAmount"`
	SUSDeAmount *big.Int `abi:"_sUSDeAmount"`
}

type EventWithdraw struct {
	From           common.Address
	SDAIAmount     *big.Int `abi:"_sDaiAmount"`
	SUSDeAmount    *big.Int `abi:"_sUSDeAmount"`
	DAIAmount      *big.Int `abi:"_daiAmount"`
	DAIAmountBoost *big.Int `abi:"_daiAmountBoost"`
	USDeAmount     *big.Int `abi:"_usdeAmount"`
}

type EventClaimForBoost struct {
	From       common.Address
	USDeAmount *big.Int `abi:"_usdeAmount"`
	DAIAmount  *big.Int `abi:"_daiAmount"`
}

type EventWithdrawBatchForBoost struct {
	From        common.Address
	SDAIAmount  *big.Int `abi:"_sDaiAmount"`
	SUSDeAmount *big.Int `abi:"_sUSDeAmount"`
}
