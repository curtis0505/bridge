package cw20

import (
	"github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/types/cosmos/cwutil"
)

var (
	_ types.CallWasm = &QueryBalanceRequest{}
	_ types.CallWasm = &QueryTokenInfoRequest{}
	_ types.CallWasm = &QueryAllowanceRequest{}
)

// QueryBalanceRequest CW20 Contract
type QueryBalanceRequest struct {
	Address string `json:"address"`
}

func (query QueryBalanceRequest) WasmKey() string { return "balance" }

type QueryBalanceResponse struct {
	Balance string `json:"balance"`
}

// QueryTokenInfoRequest CW20 Contract
type QueryTokenInfoRequest struct{}

func (query QueryTokenInfoRequest) WasmKey() string { return "token_info" }

type QueryTokenInfoResponse struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Decimal     int64  `json:"decimals"`
	TotalSupply string `json:"total_supply"`
}

// QueryAllowanceRequest CW20 Contract
type QueryAllowanceRequest struct {
	Owner   string `json:"owner"`
	Spender string `json:"spender"`
}

func (query QueryAllowanceRequest) WasmKey() string { return "allowance" }

type QueryAllowanceResponse struct {
	Allowance  string            `json:"allowance"`
	Expiration cwutil.Expiration `json:"expires"`
}

type QueryMarketingInfoRequest struct{}

func (QueryMarketingInfoRequest) WasmKey() string { return "marketing_info" }

type QueryMarketingInfoResponse struct {
	Project     string `json:"project"`
	Description string `json:"description"`
	Logo        struct {
		URL string `json:"url"`
	} `json:"logo"`
	Marketing string `json:"marketing"`
}
