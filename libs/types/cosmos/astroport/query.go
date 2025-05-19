package astroport

import (
	"github.com/curtis0505/bridge/libs/types"
)

var (
	_ types.CallWasm = &QueryPairRequest{}
	_ types.CallWasm = &QueryPoolRequest{}
	_ types.CallWasm = &QuerySimulationRequest{}
	_ types.CallWasm = &QueryReverseSimulationRequest{}
	_ types.CallWasm = &QueryCumulativePricesRequest{}
	_ types.CallWasm = &QueryPairsRequest{}
	_ types.CallWasm = &QueryFeeInfoRequest{}
	_ types.CallWasm = &QuerySimulateSwapOperationsRequest{}
	_ types.CallWasm = &QueryPoolLengthRequest{}
	//_ types.CallWasm = &Query
)

type QueryConfig struct{}

func (QueryConfig) WasmKey() string { return "config" }

// QueryPairRequest Pair, Factory Contract
type QueryPairRequest struct {
	// AssetInfos require call Factory Contract
	AssetInfos []*Asset `json:"asset_infos,omitempty"`
}

func (query QueryPairRequest) WasmKey() string { return "pair" }

type QueryPairResponse struct {
	Assets          []*Asset `json:"asset_infos"`
	ContractAddress string   `json:"contract_addr"`
	LiquidityToken  string   `json:"liquidity_token"`
	PairType        PairType `json:"pair_type"`
}

// QueryPoolRequest Pair Contract
type QueryPoolRequest struct{}

func (query QueryPoolRequest) WasmKey() string { return "pool" }

type QueryPoolResponse struct {
	Assets     []AssetInfo `json:"assets"`
	TotalShare string      `json:"total_share"`
}

// QuerySimulationRequest Pair Contract
type QuerySimulationRequest struct {
	OfferAsset *AssetInfo `json:"offer_asset"`
}

func (query QuerySimulationRequest) WasmKey() string { return "simulation" }

type QuerySimulationResponse struct {
	CommissionAmount string `json:"commission_amount"`
	ReturnAmount     string `json:"return_amount"`
	SpreadAmount     string `json:"spread_amount"`
}

// QueryReverseSimulationRequest Pair Contract
type QueryReverseSimulationRequest struct {
	AskAsset *AssetInfo `json:"ask_asset"`
}

func (query QueryReverseSimulationRequest) WasmKey() string { return "reverse_simulation" }

type QueryReverseSimulationResponse struct {
	CommissionAmount string `json:"commission_amount"`
	ReturnAmount     string `json:"return_amount"`
	SpreadAmount     string `json:"spread_amount"`
}

// QueryCumulativePricesRequest Pair Contract
type QueryCumulativePricesRequest struct{}

func (query QueryCumulativePricesRequest) WasmKey() string { return "cumulative_prices" }

type QueryCumulativePricesResponse struct {
	Assets           []AssetInfo       `json:"assets"`
	TotalShare       string            `json:"total_share"`
	CumulativePrices []CumulativePrice `json:"cumulative_prices"`
}

// QueryPairsRequest Factory Contract
type QueryPairsRequest struct {
	StartAfter []*Asset `json:"start_after"`
	Limit      int64    `json:"limit"`
}

func (query QueryPairsRequest) WasmKey() string { return "pairs" }

type QueryPairsResponse struct {
	Pairs []QueryPairResponse `json:"pairs"`
}

// QueryFeeInfoRequest Factory Contract
type QueryFeeInfoRequest struct {
	PairType PairType `json:"pair_type"`
}

func (query QueryFeeInfoRequest) WasmKey() string { return "fee_info" }

type QueryFeeInfoResponse struct {
	FeeAddress  string `json:"fee_address"`
	MakerFeeBPS int64  `json:"maker_fee_bps"`
	TotalFeeBPS int64  `json:"total_fee_bps"`
}

// QuerySimulateSwapOperationsRequest Router Contract
type QuerySimulateSwapOperationsRequest struct {
	OfferAmount    string           `json:"offer_amount"`
	SwapOperations []*SwapOperation `json:"operations"`
}

func (query QuerySimulateSwapOperationsRequest) WasmKey() string { return "simulate_swap_operations" }

type QueryMsgSimulateSwapOperationsResponse struct {
	Amount string `json:"amount"`
}

// QueryPoolLengthRequest Generator Contract
type QueryPoolLengthRequest struct{}

func (query QueryPoolLengthRequest) WasmKey() string { return "pool_length" }

type QueryPoolLengthResponse int64

// QueryPendingTokenRequest Generator Contract
type QueryPendingTokenRequest struct {
	LpToken string `json:"lp_token"`
	User    string `json:"user"`
}

func (query QueryPendingTokenRequest) WasmKey() string { return "pending_token" }

type QueryPendingTokenResponse struct {
	Pending        string `json:"pending"`
	PendingOnProxy string `json:"pending_on_proxy"`
}

// QueryRewardInfoRequest Generator Contract
type QueryRewardInfoRequest struct {
	LpToken string `json:"lp_token"`
}

func (query QueryRewardInfoRequest) WasmKey() string { return "reward_info" }

type QueryRewardInfoResponse struct {
	BaseRewardToken  *Asset `json:"base_reward_token"`
	ProxyRewardToken *Asset `json:"proxy_reward_token"`
}

// QueryPoolInfoRequest Generator Contract
type QueryPoolInfoRequest struct {
	LpToken string `json:"lp_token"`
}

func (query QueryPoolInfoRequest) WasmKey() string { return "pool_info" }

type QueryPoolInfoResponse struct {
	AllocPoint                      string        `json:"alloc_point"`
	NeopinTokensPerBlock            string        `json:"neopin_tokens_per_block"`
	LastRewardBlock                 int           `json:"last_reward_block"`
	CurrentBlock                    int           `json:"current_block"`
	GlobalRewardIndex               string        `json:"global_reward_index"`
	PendingNeopinRewards            string        `json:"pending_neopin_rewards"`
	RewardProxy                     interface{}   `json:"reward_proxy"`
	PendingProxyRewards             interface{}   `json:"pending_proxy_rewards"`
	AccumulatedProxyRewardsPerShare []interface{} `json:"accumulated_proxy_rewards_per_share"`
	ProxyRewardBalanceBeforeUpdate  string        `json:"proxy_reward_balance_before_update"`
	OrphanProxyRewards              []interface{} `json:"orphan_proxy_rewards"`
	LpSupply                        string        `json:"lp_supply"`
}

type QueryGeneratorConfigResponse struct {
	Owner                    string        `json:"owner"`
	Factory                  string        `json:"factory"`
	GeneratorController      interface{}   `json:"generator_controller"`
	VotingEscrow             interface{}   `json:"voting_escrow"`
	VotingEscrowDelegation   interface{}   `json:"voting_escrow_delegation"`
	NeopinToken              *Asset        `json:"neopin_token"`
	TokensPerBlock           string        `json:"tokens_per_block"`
	TotalAllocPoint          string        `json:"total_alloc_point"`
	StartBlock               string        `json:"start_block"`
	VestingContract          string        `json:"vesting_contract"`
	ActivePools              []ActivePool  `json:"active_pools"`
	BlockedTokensList        []interface{} `json:"blocked_tokens_list"`
	Guardian                 interface{}   `json:"guardian"`
	CheckpointGeneratorLimit interface{}   `json:"checkpoint_generator_limit"`
}

type QueryStakingConfigResponse struct {
	DepositTokenAddress string `json:"deposit_token_addr"`
	ShareTokenAddress   string `json:"share_token_addr"`
}

type QueryTotalShareRequest struct{}

func (QueryTotalShareRequest) WasmKey() string { return "total_shares" }

type QueryTotalShareResponse string

type QueryTotalDepositRequest struct{}

func (QueryTotalDepositRequest) WasmKey() string { return "total_deposit" }

type QueryTotalDepositResponse string

type QueryBalancesRequest struct {
	Assets []*Asset `json:"assets"`
}

func (QueryBalancesRequest) WasmKey() string { return "balances" }

type QueryBalancesResponse struct {
	Balances []*AssetInfo `json:"balances"`
}

type QueryBridgesRequest struct{}

func (QueryBridgesRequest) WasmKey() string { return "bridges" }

type QueryBridgeResponse [][]string

type QueryVestingAccountRequest struct {
	Address string `json:"address"`
}

func (QueryVestingAccountRequest) WasmKey() string {
	return "vesting_account"
}

type QueryVestingAccountResponse struct {
	Address     string      `json:"address"`
	VestingInfo VestingInfo `json:"info"`
}

type QueryVestingAccountsRequest struct {
	StartAfter string   `json:"start_after,omitempty"`
	Limit      int64    `json:"limit,omitempty"`
	OrderBy    *OrderBy `json:"order_by,omitempty"`
}

func (QueryVestingAccountsRequest) WasmKey() string {
	return "vesting_accounts"
}

type QueryVestingAccountsResponse struct {
	VestingAccounts []QueryVestingAccountResponse `json:"vesting_accounts"`
}

type QueryAvailableAmountRequest struct {
	Address string `json:"address"`
}

func (QueryAvailableAmountRequest) WasmKey() string {
	return "available_amount"
}

type QueryAvailableAmountResponse string
