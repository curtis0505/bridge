// StrategySDAI.sol

package strategy

const (
	StrategyRWA = "strategy-RWA"

	ProtocolINameMorpho = "Morpho"

	ProtocolCurve2  = "Curve2"
	ProtocolMorpho1 = "Morpho1"
	ProtocolSky1    = "Sky1"

	ProviderCurve  = "Curve"
	ProviderMorpho = "Morpho"
)

const (
	Maintenance = "rwa"

	StrategyRWAID      = "strategy-rwa" //contract id
	UsdcID             = "usd-coin"
	UsualBoostedUsdcID = "usual-boosted-usdc"

	RWABoostVault    = "rwa-boost-vault" //item name
	StrategyRWAIName = "rwa-boost-vault"
	RWABoost         = "RWA-BOOST"
)

const (
	MethodDeposit = "deposit"
	MethodRedeem  = "redeem"

	MethodConvertToAssets = "convertToAssets"
	MethodConvertToShares = "convertToShares"

	MethodGetPositions     = "getPositions"
	MethodComponents       = "components"
	MethodComponentDataMap = "componentDataMap"

	MethodFee         = "fee"
	MethodMaxWithdraw = "maxWithdraw"
)

const (
	EventNameDeposit = "DepositEvent"
	EventNameRedeem  = "RedeemEvent"
)

const (
	OndoAPYAppKey       = "ondo_usdy_apy"
	MorphoVaultAppKey   = "morpho1_vault_address"
	CurveUsdyUsdcAppKey = "curve2_usdy_usdc_pool_address"
)
