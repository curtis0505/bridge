package bridge

import (
	"github.com/curtis0505/bridge/libs/types"
)

var (
	_ types.CallWasm = &QueryChildChainRequest{}
	_ types.CallWasm = &QueryChildTokenRequest{}
	_ types.CallWasm = &QueryTokenRequestV1{}

	_ types.CallWasm = &QueryParentChainRequest{}
	_ types.CallWasm = &QueryParentTokenRequest{}
	_ types.CallWasm = &QueryTokenRequestV2{}
)

type QueryChildChainRequest struct {
	ChildChainName string `json:"child_chain_name"`
}

func (QueryChildChainRequest) WasmKey() string { return "child_chain" }

type QueryChildChainResponse struct {
	Fee     string `json:"fee"`
	Support bool   `json:"support"`
}

type QueryChildTokenRequest struct {
	ChildChainName string `json:"child_chain_name"`
	TokenAddr      string `json:"token_addr"`
}

func (QueryChildTokenRequest) WasmKey() string { return "child_token" }

type QueryChildTokenResponse struct {
	TokenAddr string `json:"token_addr"`
}

type QueryTokenRequestV1 struct {
	ChildChainName string `json:"child_chain_name"`
	ChildTokenAddr string `json:"child_token_addr"`
}

func (QueryTokenRequestV1) WasmKey() string { return "token" }

type QueryTokenResponseV1 struct {
	TokenAddr string `json:"token_addr"`
}

type QueryConfigRequest struct {
}

func (QueryConfigRequest) WasmKey() string { return "config" }

type QueryConfigResponse struct {
	Owner          string `json:"owner"`
	ChainName      string `json:"chain_name"`
	TaxRate        string `json:"tax_rate"`
	TaxToAddr      string `json:"tax_to_addr"`
	ChainFeeToAddr string `json:"chain_fee_to_addr"`
	Validator      string `json:"validator"`
}

type QueryParentChainRequest struct {
	ParentChainName string `json:"parent_chain_name"`
}

func (QueryParentChainRequest) WasmKey() string { return "parent_chain" }

type QueryParentChainResponse struct {
	Fee     string `json:"fee"`
	Support bool   `json:"support"`
}

type QueryParentTokenRequest struct {
	ParentChainName string `json:"parent_chain_name"`
	TokenAddr       string `json:"token_addr"`
}

func (QueryParentTokenRequest) WasmKey() string { return "parent_token" }

type QueryParentTokenResponse struct {
	TokenAddr string `json:"token_addr"`
}

type QueryTokenRequestV2 struct {
	ParentChainName string `json:"parent_chain_name"`
	ParentTokenAddr string `json:"parent_token_addr"`
}

func (QueryTokenRequestV2) WasmKey() string { return "token" }

type QueryTokenResponseV2 struct {
	TokenAddr string `json:"token_addr"`
}

type QueryTokenRequest struct {
	// ChildTokenAddr
	ChildChainName string `json:"child_chain_name,omitempty"`
	ChildTokenAddr string `json:"child_token_addr,omitempty"`

	// ParentTokenAddr
	ParentChainName string `json:"parent_chain_name,omitempty"`
	ParentTokenAddr string `json:"parent_token_addr,omitempty"`
}

func (QueryTokenRequest) WasmKey() string { return "token" }

type QueryTokenResponse struct {
	TokenAddr string `json:"token_addr"`
}

type QueryCoinRequest struct {
	// ChildTokenAddr
	ChildChainName string `json:"child_chain_name,omitempty"`
	ChildTokenAddr string `json:"child_token_addr,omitempty"`
}

func (QueryCoinRequest) WasmKey() string { return "coin" }

type QueryCoinResponse struct {
	NativeDenom string `json:"native_denom"`
}
