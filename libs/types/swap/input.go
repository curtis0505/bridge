package swap

import (
	"github.com/curtis0505/bridge/libs/common"
	"math/big"
)

type InputSwapTokens struct {
	AmountIn     *big.Int
	AmountInMax  *big.Int
	AmountOut    *big.Int
	AmountOutMin *big.Int
	Path         []common.Address
	To           common.Address
	Deadline     *big.Int
	Proof        common.Bytes
}

func (input InputSwapTokens) SafeAmountOut() *big.Int {
	if input.AmountOut != nil {
		return input.AmountOut
	}

	if input.AmountOutMin != nil {
		return input.AmountOutMin
	}

	return big.NewInt(0)
}

func (input InputSwapTokens) SafeAmountIn(value *big.Int) *big.Int {
	if value.Cmp(big.NewInt(0)) == 1 {
		return value
	}

	if input.AmountIn != nil {
		return input.AmountIn
	}

	if input.AmountInMax != nil {
		return input.AmountInMax
	}

	return big.NewInt(0)
}
