package strategy

import "github.com/curtis0505/bridge/libs/types/base"

type LiquidityInfo struct {
	TotalSupply     base.OutputBigInt //LP(srUSD totalSupply)
	CurveLiquidity  base.OutputBigInt //Curve USDC Liquidity
	MorphoLiquidity base.OutputBigInt //Morpho USDC Liquidity
}
