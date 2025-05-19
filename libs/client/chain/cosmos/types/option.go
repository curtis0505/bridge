package types

import (
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
)

type TxOption struct {
	Msgs       []cosmossdk.Msg
	FeeAmount  cosmossdk.Coins
	FeeGranter string
	GasLimit   uint64
	Memo       string

	BroadcastMode txtypes.BroadcastMode
}

func NewTxOption() *TxOption {
	return &TxOption{
		Msgs:          make([]cosmossdk.Msg, 0),
		FeeAmount:     cosmossdk.NewCoins(),
		GasLimit:      1_000_000,
		BroadcastMode: txtypes.BroadcastMode_BROADCAST_MODE_ASYNC,
	}
}

func (opt *TxOption) Apply(options ...TxOptionFunc) *TxOption {
	for _, option := range options {
		option(opt)
	}
	return opt
}

type TxOptionFunc func(option *TxOption)

func WithGasLimit(gasLimit uint64) TxOptionFunc {
	return func(option *TxOption) {
		option.GasLimit = gasLimit
	}
}

func WithMemo(memo string) TxOptionFunc {
	return func(option *TxOption) {
		option.Memo = memo
	}
}

func WithFeeAmount(coin cosmossdk.Coin) TxOptionFunc {
	return func(option *TxOption) {
		option.FeeAmount = append(option.FeeAmount, coin)
	}
}

func WithMsgs(msgs ...cosmossdk.Msg) TxOptionFunc {
	return func(option *TxOption) {
		for _, msg := range msgs {
			option.Msgs = append(option.Msgs, msg)
		}
	}
}

func WithFeeGranter(feeGranter string) TxOptionFunc {
	return func(option *TxOption) {
		option.FeeGranter = feeGranter
	}
}

func WithBroadcastMode(mode txtypes.BroadcastMode) TxOptionFunc {
	return func(option *TxOption) {
		option.BroadcastMode = mode
	}
}
