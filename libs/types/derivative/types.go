package derivative

const (
	AutoBoost = "AUTO-BOOST"

	ProtocolINameSomelier = "Sommelier"
	EthAutoBoost          = "ETH-Auto-Boost"
	StrategyTurboStETH    = "StrategyTurboStETH"
	StrategyTurboStETHId  = "strategy-turbo-steth"
	ProtocolSommelier1    = "Sommelier1"
)

const (
	Maintenance = "derivative"
)

const (
	AbiDepositEth         = "depositETH()"
	AbiWithdraw           = "withdraw(uint256)"
	AbiClaim              = "claim(uint256)"
	MethodUsers           = "users"
	MethodRequestsByOwner = "requestsByOwner"

	MethodDepositETH           = "depositETH"
	MethodWithdraw             = "withdraw"
	MethodClaim                = "claim"
	MethodGetRequestIdsByOwner = "getRequestIdsByOwner"
	MethodGetWithdrawalStatus  = "getWithdrawalStatus"
	MethodConvertToAssets      = "convertToAssets"
	MethodConvertToShares      = "convertToShares"
	MethodTotalShareAmount     = "totalShareAmount"
	MethodProtocolFeeBP        = "protocolFeeBP"
	MethodProtocolFeeTo        = "protocolFeeTo"

	InitializeEvent       = "InitializeEvent"
	EmergencyEvent        = "EmergencyEvent"
	SetProtocolFeeBPEvent = "SetProtocolFeeBPEvent"
	SetProtocolFeeToEvent = "SetProtocolFeeToEvent"
	DepositETHEvent       = "DepositETHEvent"
	WithdrawEvent         = "WithdrawEvent"
	ClaimEvent            = "ClaimEvent"
)
