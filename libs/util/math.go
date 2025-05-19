package util

import (
	"fmt"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sort"
)

func ComputeMedian(values []interface{}) (decimal.Decimal, error) {

	if values == nil || len(values) == 0 {
		return decimal.Zero, fmt.Errorf("empty")
	}

	var decimals []decimal.Decimal

	for _, v := range values {
		switch v := v.(type) {
		case float64:
			decimals = append(decimals, decimal.NewFromFloat(v))
		case primitive.Decimal128:
			dec, _ := decimal.NewFromString(v.String())
			decimals = append(decimals, dec)
		default:
			return decimal.Zero, fmt.Errorf("unsupported type: %T", v)
		}
	}

	sortedDecimals := SortValues(decimals)

	mid := len(sortedDecimals) / 2
	if len(sortedDecimals)%2 == 0 {
		return sortedDecimals[mid-1].Add(sortedDecimals[mid]).Div(decimal.NewFromFloat(2)), nil
	}
	return sortedDecimals[mid], nil
}

func SortValues(values []decimal.Decimal) []decimal.Decimal {

	sort.Slice(values, func(i, j int) bool {
		return values[i].LessThan(values[j])
	})

	return values
}
