package base

import (
	"encoding/json"
	arbcommon "github.com/curtis0505/arbitrum/common"
	"github.com/curtis0505/bridge/libs/common"
	"github.com/curtis0505/bridge/libs/types"
	ethercommon "github.com/ethereum/go-ethereum/common"
	klaycommon "github.com/klaytn/klaytn/common"
	"math/big"
)

var (
	_ types.CallMsgUnmarshaler = &OutputAddress{}
	_ types.CallMsgUnmarshaler = &OutputHash{}
	_ types.CallMsgUnmarshaler = &OutputBigInt{}
	_ types.CallMsgUnmarshaler = &OutputBool{}
	_ types.CallMsgUnmarshaler = &OutputString{}
	_ types.CallMsgUnmarshaler = &OutputUInt8{}
	_ types.CallMsgUnmarshaler = &OutputAddresses{}
)

type OutputAny struct {
	Value any
}

type OutputAddress struct {
	Value common.Address
}

func (output *OutputAddress) Unmarshal(v []any) {
	output.Value = v[0].(common.Address)
}

func (output *OutputAddress) UnmarshalJSON(b []byte) error {
	o := OutputAny{}
	err := json.Unmarshal(b, &o)
	if err != nil {
		return err
	}

	output.Value = ethercommon.HexToAddress(o.Value.(string))
	return nil
}

type OutputHash struct {
	Value common.Hash
}

func (output *OutputHash) Unmarshal(v []any) {
	output.Value = v[0].(common.Hash)
}

type OutputUInt8 struct {
	Value uint8
}

func (output *OutputUInt8) Unmarshal(v []any) {
	output.Value = v[0].(uint8)
}

type OutputUInt16 struct {
	Value uint16
}

func (output *OutputUInt16) Unmarshal(v []any) {
	output.Value = v[0].(uint16)
}

type OutputBigInt struct {
	Value *big.Int
}

func (output *OutputBigInt) Unmarshal(v []any) {
	output.Value = v[0].(*big.Int)
}

type OutputBool struct {
	Value bool
}

func (output *OutputBool) Unmarshal(v []any) {
	output.Value = v[0].(bool)
}

type OutputString struct {
	Value string
}

func (output *OutputString) Unmarshal(v []any) {
	output.Value = v[0].(string)
}

type OutputAddresses struct {
	Value []common.Address
}

func (output *OutputAddresses) Unmarshal(v []any) {
	switch addresses := v[0].(type) {
	case []ethercommon.Address:
		for _, address := range addresses {
			output.Value = append(output.Value, address)
		}
	case []klaycommon.Address:
		for _, address := range addresses {
			output.Value = append(output.Value, address)
		}
	case []arbcommon.Address:
		for _, address := range addresses {
			output.Value = append(output.Value, address)
		}
	}
}
