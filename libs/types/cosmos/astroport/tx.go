package astroport

import "github.com/curtis0505/bridge/libs/types"

var (
	_ types.CallWasm = &MsgExecuteSwapOperation{}
	_ types.CallWasm = &MsgExecuteSwapOperations{}
	_ types.CallWasm = &MsgProvideLiquidity{}
	_ types.CallWasm = &MsgWithdrawLiquidity{}
)

type MsgExecuteSwapOperation struct {
	SwapOperation *SwapOperation `json:"operation"`
	To            string         `json:"to"`
	MaxSpread     string         `json:"max_spread"`
	Single        bool           `json:"single"`
}

func (msg MsgExecuteSwapOperation) WasmKey() string { return "execute_swap_operation" }

type MsgExecuteSwapOperations struct {
	SwapOperations []*SwapOperation `json:"operations"`
	MinimumReceive string           `json:"minimum_receive"`
	To             string           `json:"to"`
	MaxSpread      string           `json:"max_spread"`
}

func (msg MsgExecuteSwapOperations) WasmKey() string { return "execute_swap_operations" }

type MsgProvideLiquidity struct {
	Assets            []AssetInfo `json:"assets"`
	SlippageTolerance *string     `json:"slippage_tolerance,omitempty"`
	AutoStake         bool        `json:"auto_stake"`
	Receiver          string      `json:"receiver"`
}

func (msg MsgProvideLiquidity) WasmKey() string { return "provide_liquidity" }

type MsgWithdrawLiquidity struct{}

func (msg MsgWithdrawLiquidity) WasmKey() string { return "withdraw_liquidity" }

type MsgExecuteSwap struct {
	OfferAsset  AssetInfo `json:"offer_asset"`
	BeliefPrice string    `json:"belief_price"`
	MaxSpread   string    `json:"max_spread"`
	To          string    `json:"to"`
}

func (msg MsgExecuteSwap) WasmKey() string { return "swap" }

type MsgEnter struct{}

func (msg MsgEnter) WasmKey() string { return "enter" }

type MsgLeave struct{}

func (msg MsgLeave) WasmKey() string { return "leave" }

type MsgCollect struct {
	Assets []*AssetWithLimit `json:"assets"`
}

func (MsgCollect) WasmKey() string { return "collect" }

type MsgClaim struct {
	Recipient string `json:"recipient"`
	Amount    string `json:"amount"`
}

func (MsgClaim) WasmKey() string { return "claim" }
