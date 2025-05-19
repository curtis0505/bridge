package farm

const (
	AbiAddLiquidity          = "addLiquidity(address,address,uint256,uint256,uint256,uint256,address,uint256,uint256)"
	AbiAddLiquiditySingle    = "addLiquiditySingle(address,address,uint256,uint256,uint256)"
	AbiAddLiquidityETH       = "addLiquidityETH(address,uint256,uint256,uint256,address,uint256,uint256)"
	AbiAddLiquidityETHSingle = "addLiquidityETHSingle(address,uint256,uint256,uint256)"
	AbiRemoveLiquidity       = "removeLiquidity(address,address,uint256,uint256,uint256,address,uint256,uint256)"
	AbiRemoveLiquidityETH    = "removeLiquidityETH(address,uint256,uint256,uint256,address,uint256,uint256)"
	AbiDeposit               = "deposit(uint256,uint256)"
	AbiClaimReward           = "claimReward(uint256,uint256)"

	AbiAddLiquidityProof          = "addLiquidity(address,address,uint256,uint256,uint256,uint256,address,uint256,uint256,bytes)"
	AbiAddLiquiditySingleProof    = "addLiquiditySingle(address,address,uint256,uint256,uint256,bytes)"
	AbiAddLiquidityETHProof       = "addLiquidityETH(address,uint256,uint256,uint256,address,uint256,uint256,bytes)"
	AbiAddLiquidityETHSingleProof = "addLiquidityETHSingle(address,uint256,uint256,uint256,bytes)"
	AbiRemoveLiquidityProof       = "removeLiquidity(address,address,uint256,uint256,uint256,address,uint256,uint256,bytes)"
	AbiRemoveLiquidityETHProof    = "removeLiquidityETH(address,uint256,uint256,uint256,address,uint256,uint256,bytes)"
	AbiDepositProof               = "deposit(uint256,uint256,bytes)"
	AbiClaimRewardProof           = "claimReward(uint256,uint256,bytes)"
)
