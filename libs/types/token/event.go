package token

import (
	"encoding/json"
	"github.com/curtis0505/bridge/libs/common"
	ethercommon "github.com/ethereum/go-ethereum/common"
	"math/big"
)

type EventTransfer struct {
	// ERC 20
	From  common.Address `json:"from"`
	To    common.Address `json:"to"`
	Value *big.Int       `json:"value"`

	// Orbit Bridge
	Amount *big.Int `json:"amount"`

	// dLLT
	Tokens *big.Int `json:"tokens"`

	// ERC 677
	// https://github.com/ethereum/EIPs/issues/677
	Data common.Bytes `json:"data"`

	// Neopin
	Spender common.Address `json:"spender"`

	// Wrapped ETH
	Src common.Address
	Dst common.Address
	Wad *big.Int
}

func (event *EventTransfer) UnmarshalJSON(b []byte) error {
	eventTransfer := struct {
		From  string   `json:"from"`
		To    string   `json:"to"`
		Value *big.Int `json:"value"`

		Src string   `json:"src"`
		Dst string   `json:"dst"`
		Wad *big.Int `json:"wad"`
	}{}

	err := json.Unmarshal(b, &eventTransfer)
	if err != nil {
		return err
	}

	event.From = ethercommon.HexToAddress(eventTransfer.From)
	event.Src = ethercommon.HexToAddress(eventTransfer.Src)

	event.To = ethercommon.HexToAddress(eventTransfer.To)
	event.Dst = ethercommon.HexToAddress(eventTransfer.Dst)

	event.Value = eventTransfer.Value
	event.Wad = eventTransfer.Wad

	return nil
}

func (event EventTransfer) SafeFrom() string {
	if !common.EmptyAddress(event.From) {
		return event.From.String()
	}

	if !common.EmptyAddress(event.Src) {
		return event.Src.String()
	}

	return common.EmptyAddressString
}

func (event EventTransfer) SafeTo() string {
	if !common.EmptyAddress(event.To) {
		return event.To.String()
	}

	if !common.EmptyAddress(event.Dst) {
		return event.Dst.String()
	}

	return common.EmptyAddressString
}

func (event EventTransfer) SafeValue() *big.Int {
	if event.Value != nil {
		return event.Value
	}

	if event.Amount != nil {
		return event.Amount
	}

	if event.Wad != nil {
		return event.Wad
	}

	if event.Tokens != nil {
		return event.Tokens
	}

	return big.NewInt(0)
}

type EventApproval struct {
	Owner    common.Address `json:"owner"`
	Spender  common.Address `json:"spender"`
	OldValue *big.Int       `json:"old_value"`
	Value    *big.Int       `json:"value"`

	// Orbit Bridge
	Holder common.Address `json:"holder"`
	Amount *big.Int       `json:"amount"`

	// Wrapped ETH
	Src common.Address
	Guy common.Address
	Wad *big.Int
}

func (event EventApproval) SafeValue() *big.Int {
	if event.Value != nil {
		return event.Value
	}

	if event.Amount != nil {
		return event.Amount
	}

	if event.Wad != nil {
		return event.Wad
	}

	return big.NewInt(0)
}

func (event EventApproval) SafeOwner() string {
	if !common.EmptyAddress(event.Owner) {
		return event.Owner.String()
	}

	if !common.EmptyAddress(event.Holder) {
		return event.Holder.String()
	}

	if !common.EmptyAddress(event.Src) {
		return event.Src.String()
	}

	return common.EmptyAddressString
}

func (event EventApproval) SafeSpender() string {
	if !common.EmptyAddress(event.Spender) {
		return event.Spender.String()
	}

	if !common.EmptyAddress(event.Guy) {
		return event.Guy.String()
	}

	return common.EmptyAddressString
}
