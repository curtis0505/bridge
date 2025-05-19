package cw20

type EventTransfer struct {
	Amount string `json:"amount"`
	From   string `json:"from"`
	To     string `json:"to"`
}

func (e EventTransfer) EventName() string {
	return "transfer"
}

type EventIncreaseAllowance struct {
	Amount  string `json:"amount"`
	Owner   string `json:"owner"`
	Spender string `json:"spender"`
}

func (e EventIncreaseAllowance) EventName() string {
	return "increase_allowance"
}

type EventDecreaseAllowance struct {
	Amount  string `json:"amount"`
	Owner   string `json:"owner"`
	Spender string `json:"spender"`
}

func (e EventDecreaseAllowance) EventName() string {
	return "decrease_allowance"
}

type EventMint struct {
	Amount string `json:"amount"`
	To     string `json:"to"`
}

func (e EventMint) EventName() string {
	return "mint"
}

type EventBurn struct {
	Amount string `json:"amount"`
	From   string `json:"to"`
}

func (e EventBurn) EventName() string {
	return "burn"
}

type EventSend struct {
	Amount string `json:"amount"`
	From   string `json:"from"`
	To     string `json:"to"`
}

func (e EventSend) EventName() string {
	return "send"
}
