package farm

import (
	"github.com/curtis0505/bridge/libs/common"
	"github.com/curtis0505/bridge/libs/util"
	ethercommon "github.com/ethereum/go-ethereum/common"
	"math/big"
)

type InputAddLiquidity struct {
	TokenA         common.Address
	TokenB         common.Address
	AmountADesired *big.Int
	AmountBDesired *big.Int
	AmountAMin     *big.Int
	AmountBMin     *big.Int
	To             common.Address
	Deadline       *big.Int
	Pid            *big.Int
	Proof          common.Bytes
}

type InputAddLiquidityETH struct {
	Token              common.Address
	AmountTokenDesired *big.Int
	AmountTokenMin     *big.Int
	AmountETHMin       *big.Int
	To                 common.Address
	Deadline           *big.Int
	Pid                *big.Int
	Proof              common.Bytes
}

type InputAddLiquiditySingle struct {
	TokenIn          common.Address
	TokenOut         common.Address
	AmountIn         *big.Int
	AmountSwapOutMin *big.Int
	Pid              *big.Int
	Proof            common.Bytes
}

type InputAddLiquidityETHSingle struct {
	Token            common.Address
	AmountIn         *big.Int
	AmountSwapOutMin *big.Int
	Pid              *big.Int
	Proof            common.Bytes
}

type InputCommonAddLiquidity struct {
	Token              common.Address
	TokenIn            common.Address
	TokenOut           common.Address
	TokenA             common.Address
	TokenB             common.Address
	AmountADesired     *big.Int
	AmountBDesired     *big.Int
	AmountAMin         *big.Int
	AmountBMin         *big.Int
	AmountTokenDesired *big.Int
	AmountTokenMin     *big.Int
	AmountETHMin       *big.Int
	AmountIn           *big.Int
	AmountSwapOutMin   *big.Int
	Deadline           *big.Int
	Pid                *big.Int
	To                 common.Address
	Proof              common.Bytes
	Liquidity          *big.Int
}

func (input InputCommonAddLiquidity) LiquidityABInfo(method string, value *big.Int) (
	tokenA, tokenB common.Address, lpAmount, amountA, amountB, pid *big.Int) {
	pid = input.Pid
	switch method {

	case AddLiquidity:
		tokenA = input.TokenA
		tokenB = input.TokenB
		amountA = input.AmountADesired
		amountB = input.AmountBDesired
	case AddLiquidityETH:
		tokenB = input.Token
		amountA = value
		amountB = input.AmountTokenDesired
	case AddLiquiditySingle:
		tokenA = input.TokenIn
		tokenB = input.TokenOut
		amountA = input.AmountIn
	case AddLiquidityETHSingle:
		if value.Cmp(big.NewInt(0)) == 0 {
			tokenA = input.Token
			amountA = input.AmountIn
		} else {
			tokenB = input.Token
			amountA = value
		}
	case RemoveLiquidity:
		tokenA = input.TokenA
		tokenB = input.TokenB
		lpAmount = input.Liquidity
		amountA = input.AmountAMin
		amountB = input.AmountBMin
	case RemoveLiquidityETH:
		tokenB = input.Token
		lpAmount = input.Liquidity
		amountA = input.AmountETHMin
		amountB = input.AmountTokenMin
	}

	lpAmount, amountA, amountB = util.ToBigInt(lpAmount), util.ToBigInt(amountA), util.ToBigInt(amountB)

	return
}

func (input InputCommonAddLiquidity) SafeTokenA() common.Address {
	if input.TokenIn != nil {
		return input.TokenIn
	}

	if input.TokenA != nil {
		return input.TokenA
	}

	return ethercommon.Address{}
}

func (input InputCommonAddLiquidity) SafeTokenB() common.Address {
	if input.Token != nil {
		return input.Token
	}

	if input.TokenOut != nil {
		return input.TokenOut
	}

	if input.TokenB != nil {
		return input.TokenB
	}

	return ethercommon.Address{}
}

func (input InputCommonAddLiquidity) SafeAmountA(value *big.Int) *big.Int {
	if value.Cmp(big.NewInt(0)) == 1 {
		return value
	}

	if input.AmountADesired != nil {
		return input.AmountADesired
	}

	if input.AmountIn != nil {
		return input.AmountIn
	}

	return big.NewInt(0)
}

func (input InputCommonAddLiquidity) SafeAmountB() *big.Int {
	if input.AmountTokenDesired != nil {
		return input.AmountTokenDesired
	}

	if input.AmountBDesired != nil {
		return input.AmountADesired
	}

	if input.AmountSwapOutMin != nil {
		return input.AmountSwapOutMin
	}

	return input.Pid
}

func (input InputCommonAddLiquidity) SafePid() *big.Int {

	return input.Pid
}

type InputRemoveLiquidity struct {
	TokenA     common.Address
	TokenB     common.Address
	Liquidity  *big.Int
	AmountAMin *big.Int
	AmountBMin *big.Int
	To         common.Address
	Deadline   *big.Int
	Pid        *big.Int
	Proof      common.Bytes
}

type InputRemoveLiquidityETH struct {
	Token          common.Address
	Liquidity      *big.Int
	AmountTokenMin *big.Int
	AmountETHMin   *big.Int
	To             common.Address
	Deadline       *big.Int
	Pid            *big.Int
	Proof          common.Bytes
}

type InputFarmDeposit struct {
	Pid    *big.Int
	Amount *big.Int
	Proof  common.Bytes
}

type InputClaimReward struct {
	Spid  *big.Int
	Bpid  *big.Int
	Proof common.Bytes
}
