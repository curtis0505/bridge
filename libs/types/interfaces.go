package types

import (
	arb "github.com/curtis0505/arbitrum"
	arbcommon "github.com/curtis0505/arbitrum/common"
	base "github.com/curtis0505/base"
	basecommon "github.com/curtis0505/base/common"
	"github.com/ethereum/go-ethereum"
	ethercommon "github.com/ethereum/go-ethereum/common"
	"github.com/klaytn/klaytn"
	klaycommon "github.com/klaytn/klaytn/common"
	"math/big"
)

type CallMsgUnmarshaler interface {
	Unmarshal([]any)
}

type CallMsg struct {
	From     string
	To       string
	Gas      uint64
	GasPrice *big.Int
	Value    *big.Int
	Data     []byte
}

func (c CallMsg) Klaytn() klaytn.CallMsg {
	to := klaycommon.HexToAddress(c.To)
	return klaytn.CallMsg{
		From:     klaycommon.HexToAddress(c.From),
		To:       &to,
		Gas:      c.Gas,
		GasPrice: c.GasPrice,
		Value:    c.Value,
		Data:     c.Data,
	}
}

func (c CallMsg) Ethereum() ethereum.CallMsg {
	to := ethercommon.HexToAddress(c.To)
	return ethereum.CallMsg{
		From:     ethercommon.HexToAddress(c.From),
		To:       &to,
		Gas:      c.Gas,
		GasPrice: c.GasPrice,
		Value:    c.Value,
		Data:     c.Data,
	}
}

func (c CallMsg) Arbitrum() arb.CallMsg {
	to := arbcommon.HexToAddress(c.To)
	return arb.CallMsg{
		From:     arbcommon.HexToAddress(c.From),
		To:       &to,
		Gas:      c.Gas,
		GasPrice: c.GasPrice,
		Value:    c.Value,
		Data:     c.Data,
	}
}

func (c CallMsg) Base() base.CallMsg {
	to := basecommon.HexToAddress(c.To)
	return base.CallMsg{
		From:     basecommon.HexToAddress(c.From),
		To:       &to,
		Gas:      c.Gas,
		GasPrice: c.GasPrice,
		Value:    c.Value,
		Data:     c.Data,
	}
}

type CallWasm interface {
	WasmKey() string
}
