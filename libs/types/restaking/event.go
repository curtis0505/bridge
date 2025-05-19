package restaking

import (
	"github.com/curtis0505/bridge/libs/common"
	"math/big"
)

type EventDepositRestakeToken struct {
	ToChainName        string         `abi:"toChainName"`
	FromAddr           common.Address `abi:"fromAddr"`
	To                 common.Bytes   `abi:"to"`
	RestakeTokenAddr   common.Address `abi:"restakeTokenAddr"`
	Decimal            uint8          `abi:"decimal"`
	RestakeTokenAmount *big.Int       `abi:"restakeTokenAmount"`
	DepositNonce       *big.Int       `abi:"depositNonce"`
}

type EventWithdrawRestakeToken struct {
	FromChainName    string       `abi:"fromChainName"`
	From             common.Bytes `abi:"from"`
	To               common.Bytes `abi:"to"`
	RestakeTokenAddr common.Bytes `abi:"restakeTokenAddr"`
	Bytes32s         common.Hash  `abi:"bytes32s"`
	Uints            []*big.Int   `abi:"uints"`
	IsDirect         bool         `abi:"isDirect"`
	WithdrawAmount   *big.Int     `abi:"withdrawAmount"`
}
