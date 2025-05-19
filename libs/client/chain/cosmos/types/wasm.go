package types

import (
	"encoding/json"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/curtis0505/bridge/libs/types"
)

func NewMsgExecuteContract(sender, contract string, msg types.CallWasm, funds cosmossdk.Coins) *wasmtypes.MsgExecuteContract {
	bz, _ := MarshalCallWasm(msg)
	return &wasmtypes.MsgExecuteContract{
		Sender:   sender,
		Contract: contract,
		Msg:      bz,
		Funds:    funds,
	}
}

func MarshalCallWasm(wasm types.CallWasm) ([]byte, error) {
	call := map[string]interface{}{
		wasm.WasmKey(): wasm,
	}
	return json.Marshal(call)
}

func UnmarshalCallWasm(bz []byte) (string, interface{}, error) {
	var raw map[string]interface{}

	if err := json.Unmarshal(bz, &raw); err != nil {
		return "", nil, err
	}

	var wasmKey string
	var wasmValue interface{}
	for key, value := range raw {
		wasmKey = key
		wasmValue = value
		break
	}
	return wasmKey, wasmValue, nil
}

// MsgCallWasm is base msg
type MsgCallWasm struct {
	key   string
	Value interface{}
}

func (msg MsgCallWasm) WasmKey() string { return msg.key }

func (msg *MsgCallWasm) UnmarshalJSON(bz []byte) error {
	key, value, err := UnmarshalCallWasm(bz)
	if err != nil {
		return err
	}
	msg.key = key
	msg.Value = value
	return nil
}

func (msg *MsgCallWasm) Unmarshal(v interface{}) error {
	bz, _ := json.Marshal(msg.Value)
	return json.Unmarshal(bz, &v)
}
