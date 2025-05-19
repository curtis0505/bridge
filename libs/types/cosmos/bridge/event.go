package bridge

type EventDeposit struct {
	ChildChainName string `json:"child_chain_name"`
	ChildTokenAddr string `json:"child_token_addr"`
	TokenAddr      string `json:"token_addr"`
	FromAddr       string `json:"from_addr"`
	ToAddr         string `json:"to_addr"`
	Tax            string `json:"tax"`
	ChainFee       string `json:"chain_fee"`
	Decimals       string `json:"decimals"`
	Amount         string `json:"amount"`
}

func (EventDeposit) EventName() string { return "deposit" }

type EventDepositCoin struct {
	ChildChainName string `json:"child_chain_name"`
	ChildTokenAddr string `json:"child_token_addr"`
	NativeDenom    string `json:"native_denom"`
	FromAddr       string `json:"from_addr"`
	ToAddr         string `json:"to_addr"`
	Tax            string `json:"tax"`
	ChainFee       string `json:"chain_fee"`
	Decimals       string `json:"decimals"`
	Amount         string `json:"amount"`
}

func (EventDepositCoin) EventName() string { return "deposit_coin" }

type EventWithdraw struct {
	ChildChainName string `json:"child_chain_name"`
	ChildTokenAddr string `json:"child_token_addr"`
	ChildTx        string `json:"child_tx"`
	TokenAddr      string `json:"token_addr"`
	FromAddr       string `json:"from_addr"`
	ToAddr         string `json:"to_addr"`
	Amount         string `json:"amount"`
}

func (EventWithdraw) EventName() string { return "withdraw" }

type EventWithdrawCoin struct {
	ChildChainName string `json:"child_chain_name"`
	ChildTokenAddr string `json:"child_token_addr"`
	ChildTx        string `json:"child_tx"`
	NativeDenom    string `json:"native_denom"`
	FromAddr       string `json:"from_addr"`
	ToAddr         string `json:"to_addr"`
	Amount         string `json:"amount"`
}

func (EventWithdrawCoin) EventName() string { return "withdraw_coin" }

type EventMint struct {
	ParentChainName string `json:"parent_chain_name"`
	ParentTokenAddr string `json:"parent_token_addr"`
	ParentTx        string `json:"parent_tx"`
	TokenAddr       string `json:"token_addr"`
	FromAddr        string `json:"from_addr"`
	ToAddr          string `json:"to_addr"`
	Amount          string `json:"amount"`
}

func (EventMint) EventName() string { return "mint" }

type EventBurn struct {
	ParentChainName string `json:"parent_chain_name"`
	ParentTokenAddr string `json:"parent_token_addr"`
	FromAddr        string `json:"from_addr"`
	ToAddr          string `json:"to_addr"`
	Tax             string `json:"tax"`
	ChainFee        string `json:"chain_fee"`
	Decimals        string `json:"decimals"`
	Amount          string `json:"amount"`
}

func (EventBurn) EventName() string { return "burn" }
