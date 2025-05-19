package token

import (
	"github.com/curtis0505/bridge/libs/common"
	"math/big"
)

type InputApproval struct {
	// ERC 20
	Spender common.Address
	Value   *big.Int
	Amount  *big.Int

	// KIP 17
	AddedValue      *big.Int // increaseAllowance
	SubtractedValue *big.Int // decreaseAllowance

	//Wrapped ETh
	Usr       common.Address
	Wad       *big.Int
	RawAmount *big.Int
}

func (input InputApproval) SafeSpender() string {
	if input.Spender != nil {
		return input.Spender.String()
	}

	if input.Usr != nil {
		return input.Usr.String()
	}

	return common.EmptyAddressString
}

func (input InputApproval) SafeValue() *big.Int {
	if input.Value != nil {
		return input.Value
	}

	if input.Amount != nil {
		return input.Amount
	}

	if input.Wad != nil {
		return input.Wad
	}

	if input.RawAmount != nil {
		return input.RawAmount
	}

	return big.NewInt(0)
}

type InputTransfer struct {
	Spender   common.Address
	Recipient common.Address
	Amount    *big.Int

	From  common.Address
	To    common.Address
	Value *big.Int

	// Wrapped ETh
	Src common.Address
	Dst common.Address
	Wad *big.Int

	// Uniswap
	RawAmount *big.Int

	//dLLT(lillus)
	Tokens *big.Int
}

func (input InputTransfer) SafeFrom() string {
	if input.Spender != nil {
		return input.Spender.String()
	}

	if input.From != nil {
		return input.From.String()
	}

	if input.Src != nil {
		return input.Src.String()
	}

	return common.EmptyAddressString
}

func (input InputTransfer) SafeTo() string {
	if input.Recipient != nil {
		return input.Recipient.String()
	}

	if input.To != nil {
		return input.To.String()
	}

	if input.Dst != nil {
		return input.Dst.String()
	}

	return common.EmptyAddressString
}

func (input InputTransfer) SafeValue() *big.Int {
	if input.Value != nil {
		return input.Value
	}

	if input.Amount != nil {
		return input.Amount
	}

	if input.Wad != nil {
		return input.Wad
	}

	if input.RawAmount != nil {
		return input.RawAmount
	}

	if input.Tokens != nil {
		return input.Tokens
	}

	return big.NewInt(0)
}
