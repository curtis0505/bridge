package swap

import (
	"github.com/curtis0505/bridge/libs/common"
	"github.com/curtis0505/bridge/libs/types"
	"math/big"
)

var (
	_ types.CallMsgUnmarshaler = &OutputToken{}
	_ types.CallMsgUnmarshaler = &OutputAllPairLength{}
	_ types.CallMsgUnmarshaler = &OutputGetAmountsOut{}
)

// OutputToken
// UniswapV2Router02 - token0 or token1
type OutputToken struct {
	Token common.Address
}

func (output *OutputToken) Unmarshal(v []interface{}) {
	output.Token = v[0].(common.Address)
}

type OutputAllPairLength struct {
	Length int64
}

func (output *OutputAllPairLength) Unmarshal(v []interface{}) {
	output.Length = v[0].(*big.Int).Int64()
}

type OutputGetAmountsOut struct {
	AmountIn  *big.Int
	AmountOut []*big.Int
}

func (output *OutputGetAmountsOut) Unmarshal(v []interface{}) {
	out := v[0].([]*big.Int)
	output.AmountIn = out[0]
	if len(out) > 1 {
		output.AmountOut = out[1:]
	}
}
