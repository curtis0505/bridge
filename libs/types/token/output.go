package token

import (
	"github.com/curtis0505/bridge/libs/types"
	"math/big"
)

var (
	_ types.CallMsgUnmarshaler = &OutputTotalSupply{}
	_ types.CallMsgUnmarshaler = &OutputBalanceOf{}
	_ types.CallMsgUnmarshaler = &OutputAllowance{}
)

type OutputBalanceOf struct {
	Amount *big.Int
}

func (output *OutputBalanceOf) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
}

type OutputAllowance struct {
	Amount *big.Int
}

func (output *OutputAllowance) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
}

type OutputTotalSupply struct {
	Amount *big.Int
}

func (output *OutputTotalSupply) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
}

type OutputIERC20Decimals struct {
	Value uint8
}

func (output *OutputIERC20Decimals) Unmarshal(v []interface{}) {
	switch vv := v[0].(type) {
	case uint8:
		output.Value = vv
	case *big.Int:
		output.Value = uint8(vv.Int64())
	}
}
