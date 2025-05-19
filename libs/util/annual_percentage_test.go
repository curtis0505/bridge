package util

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestAnnualPercentage(t *testing.T) {
	totalPooledList := []float64{
		160,
		160.32,
		162,
		163,
		160,
		163.4,
		164.45,
		166.31,
		170.234,
		172.34,
		174.43,
		177.324,
		179.682,
		182.174,
		184.666,
		187.158,
		189.65,
		192.142,
		194.634,
		197.126,
		199.618,
		202.11,
		204.602,
		207.094,
		209.586,
		212.078,
		214.57,
		217.062,
		219.554,
		222.046,
	}
	rewardList := []float64{
		0,
		0.02,
		0.021,
		0.0213,
		0.024,
		0.02,
		0.021,
		0.0213,
		0.023,
		0.024,
		0.023,
		0.0234,
		0.0238,
		0.0242,
		0.0246,
		0.025,
		0.0254,
		0.0258,
		0.0262,
		0.0266,
		0.027,
		0.0274,
		0.0278,
		0.0282,
		0.0286,
		0.029,
		0.0294,
		0.0298,
		0.0302,
		0.0306,
	}

	expectedApr := []float64{
		0.00,
		4.55,
		4.73,
		4.77,
		5.48,
		4.47,
		4.66,
		4.67,
		4.93,
		5.08,
		4.81,
		4.82,
		4.83,
		4.85,
		4.86,
		4.88,
		4.89,
		4.90,
		4.91,
		4.93,
		4.94,
		4.95,
		4.96,
		4.97,
		4.98,
		4.99,
		5.00,
		5.01,
		5.02,
		5.03,
	}

	expectedApy := []float64{
		0.00,
		4.28,
		4.45,
		4.49,
		5.17,
		4.20,
		4.38,
		4.39,
		4.64,
		4.79,
		4.53,
		4.53,
		4.55,
		4.56,
		4.57,
		4.59,
		4.60,
		4.61,
		4.62,
		4.64,
		4.65,
		4.66,
		4.67,
		4.68,
		4.69,
		4.70,
		4.71,
		4.72,
		4.73,
		4.74,
	}

	feeBP := big.NewInt(800)

	for i, _ := range totalPooledList {
		totalPooled := ToWei(totalPooledList[i])
		reward := ToWei(rewardList[i])

		apr, apy := EstimateLiquidStakingAnnualPercentageYieldAndRate(reward, totalPooled, feeBP)
		apyByApr := EstimateLiquidStakingAnnualPercentageYield(apr, feeBP)
		assert.Equal(t, expectedApr[i], apr)
		assert.Equal(t, expectedApy[i], apy)
		assert.Equal(t, expectedApy[i], apyByApr)
	}
}

func TestFinschiaAnnualPercentageRate(t *testing.T) {
	var totalSupply, totalBonded, inflation, commission, communityTax, foundationTax decimal.Decimal
	totalSupply = ToDecimal(7094854, 0)
	totalBonded = ToDecimal(4152428, 0)
	inflation = ToDecimal(0.15, 0)
	commission = ToDecimal(0.05, 0)
	communityTax = ToDecimal(0.375, 0)
	foundationTax = ToDecimal(0.2, 0)

	apr := EstimateCosmosStakingAnnualPercentageRate(totalSupply, totalBonded, inflation, commission, communityTax, foundationTax)
	assert.Equal(t, 12.17, apr)
	t.Log(apr)
}
