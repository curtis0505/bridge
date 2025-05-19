package swap

import (
	"github.com/curtis0505/bridge/libs/common"
	"github.com/curtis0505/bridge/libs/util"
	"math/big"
)

type EventSwap struct {
	// Common
	Sender common.Address

	// V2
	To         common.Address
	Amount0In  *big.Int
	Amount0Out *big.Int
	Amount1In  *big.Int
	Amount1Out *big.Int

	// V3
	Recipient    common.Address
	Amount0      *big.Int
	Amount1      *big.Int
	SqrtPriceX96 *big.Int
	Liquidity    *big.Int
	Tick         *big.Int
}

func (event EventSwap) GetSender() string { return event.Sender.String() }
func (event EventSwap) GetRecipient() string {
	if !util.IsNil(event.To) {
		return event.To.String()
	}

	if !util.IsNil(event.Recipient) {
		return event.Recipient.String()
	}

	return ""
}

func (event EventSwap) AmountIn() *big.Int {
	// V2
	if !util.IsZero(event.Amount0In) {
		return event.Amount1In
	}

	if !util.IsZero(event.Amount1In) {
		return event.Amount0In
	}

	// V3
	if event.Amount0.Cmp(big.NewInt(0)) <= 0 {
		return event.Amount0
	}

	if event.Amount1.Cmp(big.NewInt(0)) <= 0 {
		return event.Amount1
	}

	return big.NewInt(0)
}

func (event EventSwap) AmountOut() *big.Int {
	// V2
	if !util.IsZero(event.Amount0Out) {
		return event.Amount1Out
	}

	if !util.IsZero(event.Amount1Out) {
		return event.Amount0Out
	}

	// V3
	if event.Amount0.Cmp(big.NewInt(0)) > 0 {
		return event.Amount0
	}

	if event.Amount1.Cmp(big.NewInt(0)) > 0 {
		return event.Amount1
	}

	return big.NewInt(0)
}
