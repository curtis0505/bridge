// StrategySDAI.sol

package stable

const (
	StrategySDAIID = "strategysdai"
	StrategySDAI   = "StrategySDAI"

	MakerDAOPotID = "mdopot"

	RWAStable = "3pool-derivative"

	ProtocolSpark1  = "Spark1"
	ProtocolEthena1 = "Ethena1"
)

const (
	Maintenance = "rwa"
)

const (
	MethodDeposit       = "deposit"
	MethodWithdraw      = "withdraw"
	MethodClaimForBoost = "claimForBoost"

	MethodConvertToAssets = "convertToAssets"
	MethodConvertToShares = "convertToShares"
	MethodConvertToDAI    = "convertToDAI"

	MethodEstimateDAI  = "estimateDai"
	MethodEstimateUSDC = "estimateUSDC"
	MethodEstimateUSDT = "estimateUSDT"

	MethodGetCoolDownForBoost = "getCooldownForBoost"
	MethodCoolDowns           = "cooldowns"

	MethodWithdrawDuration = "withdrawDuration"
	MethodCoolDownDuration = "cooldownDuration"

	MethodUsers = "users"

	MethodTotalPendingSUSDeAmount = "totalPendingSUSDeAmount"
	MethodTotalPendingUSDeAmount  = "totalPendingUSDeAmount"
	MethodTotalSDAIAmount         = "totalSDAIAmount"
	MethodTotalSUSDeAmount        = "totalSUSDeAmount"

	MethodWithdrawBatchForBoost = "withdrawBatchForBoost"
)

const (
	LpTokenTypeRWA   = "RWA"
	LpTokenTypeBoost = "BOOST"
)

const (
	EventNameDeposit       = "DepositEvent"
	EventNameWithdraw      = "WithdrawEvent"
	EventNameClaimForBoost = "ClaimForBoostEvent"
)
