package cw20

import (
	"encoding/base64"
	"encoding/json"
	cosmostypes "github.com/curtis0505/bridge/libs/client/chain/cosmos/types"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/types/cosmos/cwutil"
)

var (
	_ types.CallWasm = &MsgTransfer{}
	_ types.CallWasm = &MsgIncreaseAllowance{}
	_ types.CallWasm = &MsgDecreaseAllowance{}
	_ types.CallWasm = &MsgSend{}
)

type MsgTransfer struct {
	Recipient string `json:"recipient"`
	Amount    string `json:"amount"`
}

func (msg MsgTransfer) WasmKey() string { return "transfer" }

type MsgIncreaseAllowance struct {
	Spender string            `json:"spender"`
	Amount  string            `json:"amount"`
	Expires cwutil.Expiration `json:"expires"`
}

func (msg MsgIncreaseAllowance) WasmKey() string { return "increase_allowance" }

type MsgDecreaseAllowance struct {
	Spender string            `json:"spender"`
	Amount  string            `json:"amount"`
	Expires cwutil.Expiration `json:"expires"`
}

func (msg MsgDecreaseAllowance) WasmKey() string { return "decrease_allowance" }

type msgSend struct {
	Contract string `json:"contract"`
	Amount   string `json:"amount"`
	Msg      string `json:"msg"`
}

type MsgSend struct {
	Contract string         `json:"contract"`
	Amount   string         `json:"amount"`
	Msg      types.CallWasm `json:"msg"`
}

func (msg MsgSend) WasmKey() string { return "send" }

func (msg MsgSend) MarshalJSON() ([]byte, error) {
	bz, err := cosmostypes.MarshalCallWasm(msg.Msg)
	if err != nil {
		return nil, err
	}

	return json.Marshal(msgSend{
		Contract: msg.Contract,
		Amount:   msg.Amount,
		Msg:      base64.StdEncoding.EncodeToString(bz),
	})
}

func (msg *MsgSend) UnmarshalJSON(bz []byte) error {
	callWasm := cosmostypes.MsgCallWasm{}
	if err := callWasm.UnmarshalJSON(bz); err != nil {
		return err
	}

	msgSendRaw := msgSend{}
	if err := callWasm.Unmarshal(&msgSendRaw); err != nil {
		return err
	}

	msg.Contract = msgSendRaw.Contract
	msg.Amount = msgSendRaw.Amount

	encodedMsg, err := base64.StdEncoding.DecodeString(msgSendRaw.Msg)
	if err != nil {
		return err
	}

	callWasmMsg := cosmostypes.MsgCallWasm{}
	if err := callWasmMsg.UnmarshalJSON(encodedMsg); err != nil {
		return err
	}

	msg.Msg = callWasmMsg
	return nil
}
