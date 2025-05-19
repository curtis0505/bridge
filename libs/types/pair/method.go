package pair

const (
	AbiCreatePairByUser    = "createPairByUser(address,address,uint256,uint256,uint256,uint256,address,uint256)"
	AbiCreatePairETHByUser = "createPairETHByUser(address,uint256,uint256,uint256,address,uint256)"
	AbiAddLiquidity        = "addLiquidity(address,address,uint256,uint256,uint256,uint256,address,uint256)"
	AbiAddLiquidityETH     = "addLiquidityETH(address,uint256,uint256,uint256,address,uint256)"
	AbiRemoveLiquidity     = "removeLiquidity(address,address,uint256,uint256,uint256,address,uint256)"
	AbiRemoveLiquidityETH  = "removeLiquidityETH(address,uint256,uint256,uint256,address,uint256)"

	AbiCreatePairByUserProof    = "createPairByUser(address,address,uint256,uint256,uint256,uint256,address,uint256,bytes)"
	AbiCreatePairETHByUserProof = "createPairETHByUser(address,uint256,uint256,uint256,address,uint256,bytes)"
	AbiAddLiquidityProof        = "addLiquidity(address,address,uint256,uint256,uint256,uint256,address,uint256,bytes)"
	AbiAddLiquidityETHProof     = "addLiquidityETH(address,uint256,uint256,uint256,address,uint256,bytes)"
	AbiRemoveLiquidityProof     = "removeLiquidity(address,address,uint256,uint256,uint256,address,uint256,bytes)"
	AbiRemoveLiquidityETHProof  = "removeLiquidityETH(address,uint256,uint256,uint256,address,uint256,bytes)"
)
