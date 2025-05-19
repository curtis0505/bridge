package pair

import (
	"github.com/curtis0505/bridge/libs/common"
	"math/big"
)

type CreatePairByUserInput struct {
	TokenA         common.Address `abi:"tokenA"`
	TokenB         common.Address `abi:"tokenB"`
	AmountADesired *big.Int       `abi:"amountADesired"`
	AmountBDesired *big.Int       `abi:"amountBDesired"`
	AmountAMin     *big.Int       `abi:"amountAMin"`
	AmountBMin     *big.Int       `abi:"amountBMin"`
	To             common.Address `abi:"to"`
	Deadline       *big.Int       `abi:"deadline"`
}

type CreatePairETHByUserInput struct {
	Token              common.Address `abi:"token"`
	AmountTokenDesired *big.Int       `abi:"amountTokenDesired"`
	AmountTokenMin     *big.Int       `abi:"amountTokenMin"`
	AmountETHMin       *big.Int       `abi:"amountETHMin"`
	To                 common.Address `abi:"to"`
	Deadline           *big.Int       `abi:"deadline"`
}

type CreatePairETHByUserInputCommon struct {
	TokenA             common.Address
	TokenB             common.Address
	Token              common.Address
	AmountADesired     *big.Int
	AmountBDesired     *big.Int
	AmountTokenDesired *big.Int
	AmountAMin         *big.Int
	AmountBMin         *big.Int
	AmountTokenMin     *big.Int
	AmountETHMin       *big.Int
	To                 common.Address
	Deadline           *big.Int
}

func (input CreatePairETHByUserInputCommon) CreateABInfo(method string, value *big.Int) *CreatePairByUserTokenInfo {
	//tokenA, tokenB common.Address, lpAmount, amountA, amountB *big.Int, isCoin bool) {

	createPairByUserTokenInfo := &CreatePairByUserTokenInfo{}

	switch method {
	case MethodCreatePairByUser:
		createPairByUserTokenInfo.TokenA = input.TokenA
		createPairByUserTokenInfo.TokenB = input.TokenB
		createPairByUserTokenInfo.AmountA = safeAmount(input.AmountADesired)
		createPairByUserTokenInfo.AmountB = safeAmount(input.AmountBDesired)
		createPairByUserTokenInfo.HaveCoin = false
	case MethodCreatePairETHByUser:
		createPairByUserTokenInfo.TokenB = input.Token
		createPairByUserTokenInfo.AmountA = safeAmount(value)
		createPairByUserTokenInfo.AmountB = safeAmount(input.AmountTokenDesired)
		createPairByUserTokenInfo.HaveCoin = true
	}

	return createPairByUserTokenInfo
}

func safeAmount(amount *big.Int) *big.Int {
	if amount == nil {
		return big.NewInt(0)
	}
	return amount
}
