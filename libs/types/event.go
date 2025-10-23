package types

import (
	arbabi "github.com/curtis0505/arbitrum/accounts/abi"
	baseabi "github.com/curtis0505/base/accounts/abi"
	etherabi "github.com/ethereum/go-ethereum/accounts/abi"
	klayabi "github.com/kaiachain/kaia/accounts/abi"
)

const (
	GradeCalculationByPosition = "POSITION"
	GradeCalculationByPoint    = "POINT"
)

type Event struct {
	chain string
	inner any
}

func NewEvent(chain string, inner any) *Event {
	return &Event{
		chain: chain,
		inner: inner,
	}
}

func (e *Event) RawName() string {
	switch v := e.inner.(type) {
	case *klayabi.Event:
		return v.RawName
	case *etherabi.Event:
		return v.RawName
	case *baseabi.Event:
		return v.RawName
	case *arbabi.Event:
		return v.RawName
	default:
		return ""
	}
}

func (e *Event) Name() string {
	switch v := e.inner.(type) {
	case *klayabi.Event:
		return v.Name
	case *etherabi.Event:
		return v.Name
	case *baseabi.Event:
		return v.Name
	case *arbabi.Event:
		return v.Name
	default:
		return ""
	}
}
