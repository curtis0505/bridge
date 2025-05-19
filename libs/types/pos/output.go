package pos

import (
	"github.com/curtis0505/bridge/libs/types"
	"math/big"
)

var (
	_ types.CallMsgUnmarshaler = &OutputLastChildBlock{}
	_ types.CallMsgUnmarshaler = &OutputLastChildBlock{}
)

type OutputRootBlockInfo struct {
	HeaderBlockNumber *big.Int
	Start             *big.Int
	End               *big.Int
}

func (output *OutputRootBlockInfo) Unmarshal(v []interface{}) {
	output.Start = v[1].(*big.Int)
	output.End = v[2].(*big.Int)
}

type OutputLastChildBlock struct {
	LastChildBlock *big.Int
}

func (output *OutputLastChildBlock) Unmarshal(v []interface{}) {
	output.LastChildBlock = v[0].(*big.Int)
}
