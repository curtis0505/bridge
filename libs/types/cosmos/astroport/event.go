package astroport

import cosmossdk "github.com/cosmos/cosmos-sdk/types"

type EventSend struct {
	Amount string `json:"amount"`
	From   string `json:"from"`
	To     string `json:"to"`
}

func (EventSend) EventName() string { return "send" }

type EventSwap struct {
	AskAsset         string `json:"ask_asset"`
	CommissionAmount string `json:"commission_amount"`
	MakerFeeAmount   string `json:"maker_fee_amount"`
	OfferAmount      string `json:"offer_amount"`
	OfferAsset       string `json:"offer_asset"`
	Receiver         string `json:"receiver"`
	ReturnAmount     string `json:"return_amount"`
	Sender           string `json:"sender"`
	SpreadAmount     string `json:"spread_amount"`
}

func (EventSwap) EventName() string { return "swap" }

type EventProvideLiquidity struct {
	Assets   string `json:"assets"`
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
	Share    string `json:"share"`
}

func (event *EventProvideLiquidity) GetAssets() cosmossdk.Coins {
	coins, err := cosmossdk.ParseCoinsNormalized(event.Assets)
	if err != nil {
		return coins
	}
	return coins
}

func (EventProvideLiquidity) EventName() string { return "provide_liquidity" }

type EventCreatePair struct {
	Pair string `json:"pair"`
}

func (EventCreatePair) EventName() string { return "create_pair" }

type EventWithdrawLiquidity struct {
	RefundAssets   string `json:"refund_assets"`
	Sender         string `json:"sender"`
	WithdrawnShare string `json:"withdrawn_share"`
}

func (EventWithdrawLiquidity) EventName() string { return "withdraw_liquidity" }

func (event *EventWithdrawLiquidity) GetRefundAssets() cosmossdk.Coins {
	coins, err := cosmossdk.ParseCoinsNormalized(event.RefundAssets)
	if err != nil {
		return coins
	}
	return coins
}

type EventEnter struct {
	NeopinAmount  string `json:"neopin_amount"`
	Recipient     string `json:"recipient"`
	XNeopinAmount string `json:"xneopin_amount"`
}

func (EventEnter) EventName() string { return "enter" }

type EventLeave struct {
	NeopinAmount  string `json:"neopin_amount"`
	Recipient     string `json:"recipient"`
	XNeopinAmount string `json:"xneopin_amount"`
}

func (EventLeave) EventName() string { return "leave" }

type EventClaim struct {
	Address         string `json:"address"`
	AvailableAmount string `json:"available_amount"`
	ClaimedAmount   string `json:"claimed_amount"`
}

func (EventClaim) EventName() string { return "claim" }

type EventDistributeNeopin struct {
	NeopinDistribution           string `json:"neopin_distribution"`
	PreUpgradeNeopinDistribution string `json:"preupgrade_neopin_distribution"`
}

func (EventDistributeNeopin) EventName() string { return "distribute_neopin" }

type EventLiquidityProviderTokenMinted struct {
	Amount string `json:"amount"`
	To     string `json:"to"`
}

func (EventLiquidityProviderTokenMinted) EventName() string { return "mint" }

type EventStakingTokenMinted struct {
	Amount string `json:"amount"`
	To     string `json:"to"`
}

func (EventStakingTokenMinted) EventName() string { return "mint" }
