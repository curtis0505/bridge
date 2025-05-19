package astroport

import (
	"encoding/json"
	"time"
)

type AssetType string
type PoolType string

const (
	AssetTypeCoin  AssetType = "coin"
	AssetTypeToken AssetType = "token"

	PoolTypePersistent  PoolType = "xyk"
	PoolTypeSubstantial PoolType = "stable"
)

type RoutePath struct {
	From            *Asset
	To              *Asset
	ContractAddress string
	PoolFeeRate     float64
	PoolType        string
}

func NewSwapOperation(offer, ask *Asset) *SwapOperation {
	if offer.Type() == AssetTypeCoin && ask.Type() == AssetTypeCoin {
		return &SwapOperation{
			NativeSwap: &NativeSwap{
				OfferDenom: offer.NativeToken.Denom,
				AskDenom:   ask.NativeToken.Denom,
			},
		}
	} else {
		return &SwapOperation{
			AstroSwap: &NeopinSwap{
				OfferAssetInfo: offer,
				AskAssetInfo:   ask,
			},
		}
	}
}

type SwapOperation struct {
	NativeSwap *NativeSwap `json:"native_swap,omitempty"`
	AstroSwap  *NeopinSwap `json:"neopin_swap,omitempty"`
}

type NativeSwap struct {
	OfferDenom string `json:"offer_denom"`
	AskDenom   string `json:"ask_denom"`
}

// NeopinSwap renamed from AstroSwap
type NeopinSwap struct {
	OfferAssetInfo *Asset `json:"offer_asset_info"`
	AskAssetInfo   *Asset `json:"ask_asset_info"`
}

type PairType struct {
	Xyk    *struct{} `json:"xyk,omitempty"`
	Stable *struct{} `json:"stable,omitempty"`
}

func (pair PairType) Type() PoolType {
	if pair.Xyk != nil {
		return PoolTypePersistent
	} else {
		return PoolTypeSubstantial
	}
}

func (pair PairType) String() string {
	return string(pair.Type())
}

type CumulativePrice struct {
	AssetA *Asset
	AssetB *Asset
	Price  string
}

func (c *CumulativePrice) UnmarshalJSON(bz []byte) error {
	v := []interface{}{
		&Asset{}, &Asset{}, "",
	}
	if err := json.Unmarshal(bz, &v); err != nil {
		return err
	}

	c.AssetA = v[0].(*Asset)
	c.AssetB = v[1].(*Asset)
	c.Price = v[2].(string)
	return nil
}

type AssetWithLimit struct {
	Info  *Asset `json:"info"`
	Limit string `json:"limit,omitempty"`
}

type AssetInfo struct {
	Info   *Asset `json:"info"`
	Amount string `json:"amount"`
}

type Asset struct {
	// NativeToken AssetTypeCoin
	NativeToken *NativeToken `json:"native_token,omitempty"`

	// Token AssetTypeToken
	Token *Token `json:"token,omitempty"`
}

func NewAsset(coinType AssetType, data string) *Asset {
	if coinType == AssetTypeCoin {
		return &Asset{
			NativeToken: &NativeToken{
				Denom: data,
			},
		}
	} else {
		return &Asset{
			Token: &Token{
				ContractAddress: data,
			},
		}
	}
}

func (asset *Asset) Type() AssetType {
	if asset.NativeToken != nil {
		return AssetTypeCoin
	} else {
		return AssetTypeToken
	}
}

func (asset *Asset) String() string {
	if asset.NativeToken != nil {
		return asset.NativeToken.Denom
	} else if asset.Token != nil {
		return asset.Token.ContractAddress
	} else {
		return ""
	}
}

type Token struct {
	ContractAddress string `json:"contract_addr"`
}

type NativeToken struct {
	Denom string `json:"denom"`
}

type ActivePool struct {
	LpTokenAddress string
	AllocPoint     string
}

func (pool *ActivePool) UnmarshalJSON(bz []byte) error {
	var raw []string
	if err := json.Unmarshal(bz, &raw); err != nil {
		return err
	}
	pool.LpTokenAddress = raw[0]
	pool.AllocPoint = raw[1]

	return nil
}

type VestingInfo struct {
	Schedules      []VestingSchedule `json:"schedules"`
	ReleasedAmount string            `json:"released_amount"`
}

type VestingSchedule struct {
	StartPoint VestingPoint `json:"start_point"`
	EndPoint   VestingPoint `json:"end_point"`
}

type VestingPoint struct {
	Time   int64  `json:"time"`
	Amount string `json:"amount"`
}

func (v VestingPoint) GetTime() time.Time {
	return time.Unix(v.Time, 0)
}

type OrderByType string

const (
	OrderByDesc = OrderByType("desc")
	OrderByAsc  = OrderByType("asc")
)

type OrderBy struct {
	DESC *struct{} `json:"desc,omitempty"`
	ASC  *struct{} `json:"asc,omitempty"`
}

func NewOrderBy(orderBy OrderByType) *OrderBy {
	if orderBy == OrderByDesc {
		return &OrderBy{
			DESC: &struct{}{},
		}
	} else {
		return &OrderBy{
			ASC: &struct{}{},
		}
	}
}
