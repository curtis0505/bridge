package nft

import (
	"github.com/curtis0505/bridge/libs/common"
	"github.com/curtis0505/bridge/libs/types"
	"math/big"
)

var (
	_ types.CallMsgUnmarshaler = &OutputBalanceOf{}
	_ types.CallMsgUnmarshaler = &OutputTokenURI{}
)

type OutputBalanceOf struct {
	Amount *big.Int
}

func (output *OutputBalanceOf) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
}

type OutputTokenURI struct {
	URI string
}

func (output *OutputTokenURI) Unmarshal(v []interface{}) {
	output.URI = v[0].(string)
}

type OutputOwnerOf struct {
	Address common.Address
}

func (output *OutputOwnerOf) Unmarshal(v []interface{}) {
	output.Address = v[0].(common.Address)
}
