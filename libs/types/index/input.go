package index

import (
	"github.com/curtis0505/bridge/libs/common"
	"math/big"
)

// InputIssueSetForExactToken issueSetForExactToken
type InputIssueSetForExactToken struct {
	SetToken      common.Address `abi:"_setToken"`
	InputToken    common.Address `abi:"_inputToken"`
	AmountInput   *big.Int       `abi:"_amountInput"`
	MinSetReceive *big.Int       `abi:"_minSetReceive"`
	ToETHSwap     SwapData       `abi:"_toEthSwap"`
	//ToComponentsSwap []SwapData     `abi:"_toComponentsSwap"`
}

func (InputIssueSetForExactToken) MethodName() string { return "issueSetForExactToken" }

// InputIssueSetForExactETH issueSetForExactETH
type InputIssueSetForExactETH struct {
	SetToken      common.Address `abi:"_setToken"`
	MinSetReceive *big.Int       `abi:"_minSetReceive"`
	//ToComponentsSwap []SwapData     `abi:"_toComponentsSwap"`
}

func (InputIssueSetForExactETH) MethodName() string { return "issueSetForExactETH" }

// InputIssueExactSetFromToken issueExactSetFromToken
type InputIssueExactSetFromToken struct {
	SetToken            common.Address `abi:"_setToken"`
	InputToken          common.Address `abi:"_inputToken"`
	AmountSetToken      *big.Int       `abi:"_amountSetToken"`
	MaxAmountInputToken *big.Int       `abi:"_maxAmountInputToken"`
	ToETHSwap           SwapData       `abi:"_toEthSwap"`
	//ToComponentsSwap    []SwapData     `abi:"_toComponentsSwap"`
}

func (InputIssueExactSetFromToken) MethodName() string { return "issueExactSetFromToken" }

// InputIssueExactSetFromETH issueExactSetFromETH
type InputIssueExactSetFromETH struct {
	SetToken       common.Address `abi:"_setToken"`
	AmountSetToken *big.Int       `abi:"_amountSetToken"`
	//ToComponentsSwap []SwapData     `abi:"_toComponentsSwap"`
}

func (InputIssueExactSetFromETH) MethodName() string { return "issueExactSetFromETH" }

// InputRedeemExactSet redeemExactSet
type InputRedeemExactSet struct {
	SetToken       common.Address `abi:"_setToken"`
	AmountSetToken *big.Int       `abi:"_amountSetToken"`
}

func (InputRedeemExactSet) MethodName() string { return "redeemExactSet" }

// InputRedeemExactSetForToken redeemExactSetForToken
type InputRedeemExactSetForToken struct {
	SetToken         common.Address `abi:"_setToken"`
	OutputToken      common.Address `abi:"_outputToken"`
	AmountSetToken   *big.Int       `abi:"_amountSetToken"`
	MinOutputReceive *big.Int       `abi:"_minOutputReceive"`
	//FromComponentsSwap []SwapData     `abi:"_fromComponentsSwap"` // (component -> eth)[]
}

func (InputRedeemExactSetForToken) MethodName() string { return "redeemExactSetForToken" }

// InputRedeemExactSetForETH redeemExactSetForETH
type InputRedeemExactSetForETH struct {
	SetToken       common.Address `abi:"_setToken"`
	AmountSetToken *big.Int       `abi:"_amountSetToken"`
	MinETHOut      *big.Int       `abi:"_minEthOut"`
	//FromComponentsSwap []SwapData     `abi:"_fromComponentsSwap"` // (component -> eth)[]
}

func (InputRedeemExactSetForETH) MethodName() string { return "redeemExactSetForETH" }
