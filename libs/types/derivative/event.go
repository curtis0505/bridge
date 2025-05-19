package derivative

import (
	"github.com/curtis0505/bridge/libs/common"
	"math/big"
)

type EventDepositETHEvent struct {
	From        common.Address
	EthAmount   *big.Int `abi:"_ethAmount"`
	StEthAmount *big.Int `abi:"_stEthAmount"`
	ShareAmount *big.Int `abi:"_shareAmount"`
}

type EventWithdrawEvent struct {
	From         common.Address
	ShareAmount  *big.Int `abi:"_shareAmount"`
	Assets       *big.Int `abi:"_assets"`
	StEthAmount  *big.Int `abi:"_stEthAmount"`
	WstEthAmount *big.Int `abi:"_wstEthAmount"`
	WethAmount   *big.Int `abi:"_wethAmount"`
}

type EventClaimEvent struct {
	From      common.Address
	RequestId *big.Int `abi:"_requestId"`
	EthAmount *big.Int `abi:"_ethAmount"`
}
