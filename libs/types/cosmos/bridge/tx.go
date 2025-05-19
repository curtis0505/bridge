package bridge

import "github.com/curtis0505/bridge/libs/types"

var (
	_ types.CallWasm = &MsgDeposit{}
	_ types.CallWasm = &MsgWithdraw{}
	_ types.CallWasm = &MsgMint{}
	_ types.CallWasm = &MsgBurn{}
)

type MsgDeposit struct {
	ChildChainName string `json:"child_chain_name"`
	ChildTokenAddr string `json:"child_token_addr"`
	ToAddr         string `json:"to_addr"`
	Amount         string `json:"amount"`
}

func (MsgDeposit) WasmKey() string { return "deposit" }

type MsgDepositCoin struct {
	ChildChainName string `json:"child_chain_name"`
	ChildTokenAddr string `json:"child_token_addr"`
	ToAddr         string `json:"to_addr"`
}

func (MsgDepositCoin) WasmKey() string { return "deposit_coin" }

type MsgWithdraw struct {
	ChildChainName string `json:"child_chain_name"`
	ChildTokenAddr string `json:"child_token_addr"`
	ChildTx        string `json:"child_tx"`
	FromAddr       string `json:"from_addr"`
	ToAddr         string `json:"to_addr"`
	Amount         string `json:"amount"`
}

func (MsgWithdraw) WasmKey() string { return "withdraw" }

type MsgWithdrawCoin struct {
	ChildChainName string `json:"child_chain_name"`
	ChildTokenAddr string `json:"child_token_addr"`
	ChildTx        string `json:"child_tx"`
	FromAddr       string `json:"from_addr"`
	ToAddr         string `json:"to_addr"`
	Amount         string `json:"amount"`
}

func (MsgWithdrawCoin) WasmKey() string { return "withdraw_coin" }

type MsgMint struct {
	ParentChainName string `json:"parent_chain_name"`
	ParentTokenAddr string `json:"parent_token_addr"`
	ParentTx        string `json:"parent_tx"`
	FromAddr        string `json:"from_addr"`
	ToAddr          string `json:"to_addr"`
	Amount          string `json:"amount"`
}

func (MsgMint) WasmKey() string { return "mint" }

type MsgBurn struct {
	ParentChainName string `json:"parent_chain_name"`
	ParentTokenAddr string `json:"parent_token_addr"`
	ToAddr          string `json:"to_addr"`
	Amount          string `json:"amount"`
}

func (MsgBurn) WasmKey() string { return "burn" }
