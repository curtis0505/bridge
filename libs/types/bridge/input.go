package bridge

import (
	"github.com/curtis0505/bridge/libs/common"
	"math/big"
)

type InputDeposit struct {
	TokenAddr   common.Address
	ToChainName string
	To          common.Bytes
	Amount      *big.Int
	Proof       common.Bytes
}

type InputBurn struct {
	WTokenAddr  common.Address
	ToChainName string
	To          common.Bytes
	Amount      *big.Int
	Proof       common.Bytes
}

type InputDepositBurn struct {
	TokenAddr   common.Address
	WTokenAddr  common.Address
	ToChainName string
	To          common.Bytes
	Amount      *big.Int
	Proof       common.Bytes
}

type InputDepositFxERC20RootTunnel struct {
	RootToken common.Address
	User      common.Address
	Amount    *big.Int
	Data      common.Bytes
	Proof     common.Bytes
}

type InputWithdrawFxERC20ChildTunnel struct {
	RootToken  common.Address
	ChildToken common.Address
	Amount     *big.Int
	Proof      common.Bytes
}
