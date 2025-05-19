package farm

const (
	AddLiquidity          = "addLiquidity"
	AddLiquidityETH       = "addLiquidityETH"
	AddLiquiditySingle    = "addLiquiditySingle"
	AddLiquidityETHSingle = "addLiquidityETHSingle"

	RemoveLiquidity                                 = "removeLiquidity"
	RemoveLiquidityETH                              = "removeLiquidityETH"
	RemoveLiquidityETHSupportingFeeOnTransferTokens = "removeLiquidityETHSupportingFeeOnTransferTokens"

	MethodUserInfo      = "userInfo"
	MethodGetUserAmount = "getUserAmount"
	MethodPendingReward = "pendingReward"
	MethodPendingBonus  = "pendingBonus"

	MethodPoolInfo    = "poolInfo"
	MethodToken0      = "token0"
	MethodToken1      = "token1"
	MethodGetReserves = "getReserves"
	MethodTotalSupply = "totalSupply"

	EasyRouterId = "easyrouter"
	MasterChefId = "masterchef"
	BonusChefId  = "bonuschef"
	FactoryId    = "uniswapv2factory"
	IV2PairId    = "iuniswapv2pair"

	MasterChef = "MasterChef"
	EasyRouter = "EasyRouter"
	BonusChef  = "BonusChef"

	YieldFarm = "CONTRACT-YDF"
)

const (
	EventNameDeposit          = "Deposit"
	EventNameWithdraw         = "Withdraw"
	EventNameMint             = "Mint"
	EventNameBurn             = "Burn"
	EventNameClaimReward      = "ClaimReward"
	EventNameClaimBonus       = "ClaimBonus"
	EventNameAttachMasterChef = "AttachMasterChef"
	EventNameDetachMasterChef = "DetachMasterChef"
	EventNameUpdateEndBlock   = "UpdateEndBlock"
)

const (
	Maintenance = "farm"
)
