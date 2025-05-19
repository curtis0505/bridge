package index

import (
	"github.com/curtis0505/bridge/libs/common"
	"math/big"
)

type EventExchangeIssue struct {
	Recipient        common.Address
	SetToken         common.Address
	InputToken       common.Address
	AmountInputToken *big.Int
	AmountSetIssued  *big.Int
}

func (EventExchangeIssue) EventName() string { return "ExchangeIssue" }

type EventExchangeRedeem struct {
	Recipient         common.Address
	SetToken          common.Address
	OutputToken       common.Address
	AmountSetRedeemed *big.Int
	AmountOutputToken *big.Int
}

func (EventExchangeRedeem) EventName() string { return "ExchangeRedeem" }

type EventRefund struct {
	Recipient    common.Address `abi:"_recipient"`
	RefundAmount *big.Int       `abi:"_refundAmount"`
}

func (EventRefund) EventName() string { return "Refund" }
