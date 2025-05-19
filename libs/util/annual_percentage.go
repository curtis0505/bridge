package util

import (
	"github.com/shopspring/decimal"
	"math/big"
)

// EstimateLiquidStakingAnnualPercentageYieldAndRate
// reward, totalPooled, feeBP를 기반으로 apr, apy를 계산
func EstimateLiquidStakingAnnualPercentageYieldAndRate(
	reward *big.Int,
	totalPooled *big.Int,
	feeBP *big.Int,
) (
	apr float64,
	apy float64,
) {
	rate := ToDecimal(reward, 0).
		Mul(decimal.NewFromInt(365)).
		Div(ToDecimal(totalPooled, 0))

	// (1-FeeRate) = 1 - 0.08 = 0.92
	userRewardRate := decimal.NewFromInt(1).
		Sub(ToDecimal(feeBP, 4))

	// (1 + (APR * (1-FeeRate)/365) ^ 365) - 1
	yield := decimal.NewFromInt(1).
		Add(
			rate.Mul(userRewardRate).
				Div(decimal.NewFromInt(365)),
		).
		Pow(decimal.NewFromInt(365)).
		Sub(decimal.NewFromInt(1))

	apr, _ = rate.Mul(decimal.NewFromInt(100)).Round(2).Float64()
	//apy = EstimateLiquidStakingAnnualPercentageYield(apr, feeBP)
	apy, _ = yield.Mul(decimal.NewFromInt(100)).Round(2).Float64()

	return
}

// EstimateLiquidStakingAnnualPercentageYield
// apr% 과 feeBP를 기반으로 apy% 를 계산 오차 있을 수 있음.
func EstimateLiquidStakingAnnualPercentageYield(
	apr float64,
	feeBP *big.Int,
) (
	apy float64,
) {
	userRewardRate := decimal.NewFromInt(1).
		Sub(ToDecimal(feeBP, 4))

	yield := decimal.NewFromInt(1).
		Add(
			decimal.NewFromFloat(apr).
				Div(decimal.NewFromInt(100)).
				Mul(userRewardRate).
				Div(decimal.NewFromInt(365)),
		).
		Pow(decimal.NewFromInt(365)).
		Sub(decimal.NewFromInt(1))

	apy, _ = yield.Mul(decimal.NewFromInt(100)).Round(2).Float64()
	return
}

// EstimateLiquidStakingAnnualPercentageRate
// reward, totalPooled, feeBP를 기반으로 apr 을 계산
func EstimateLiquidStakingAnnualPercentageRate(
	reward *big.Int,
	totalPooled *big.Int,
) (
	apr float64,
) {
	rate := ToDecimal(reward, 0).
		Mul(decimal.NewFromInt(365)).
		Div(ToDecimal(totalPooled, 0))

	apr, _ = rate.Mul(decimal.NewFromInt(100)).Round(2).Float64()
	return
}

func EstimateCosmosStakingAnnualPercentageRate(
	totalSupply, totalBonded, inflationRate,
	commissionRate, communityTax, foundationTax decimal.Decimal,
) (
	apr float64,
) {
	rate := inflationRate.Mul(
		decimal.NewFromInt(1).Sub(foundationTax),
	).Mul(
		decimal.NewFromInt(1).Sub(communityTax),
	).Mul(
		decimal.NewFromInt(1).Sub(commissionRate),
	).Div(
		totalBonded.Div(totalSupply),
	)

	apr, _ = rate.Mul(decimal.NewFromInt(100)).Round(2).Float64()
	return apr
}
