package util

import (
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/exp/constraints"
	"math/big"
)

func Contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func SafeBigIntToDecimal128(value *big.Int) primitive.Decimal128 {
	decimalZero, _ := primitive.ParseDecimal128("0")
	if value == nil {
		return decimalZero
	}

	valueD, err := primitive.ParseDecimal128(value.String())
	if err != nil {
		return decimalZero
	}

	return valueD
}

func ToEther(val interface{}) decimal.Decimal {
	return ToEtherWithDecimal(val, 18)
}

type Number interface {
	constraints.Integer | constraints.Float
}

func ToEtherWithDecimal[T Number](val interface{}, d T) decimal.Decimal {
	value := new(big.Int)
	switch v := val.(type) {
	case string:
		value.SetString(v, 10)
	case int64:
		value = big.NewInt(v)
	case primitive.Decimal128:
		value, _, _ = v.BigInt()
	case *big.Int:
		value = v
	case decimal.Decimal:
		value = v.BigInt()
	case *decimal.Decimal:
		value = v.BigInt()
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(d)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}
