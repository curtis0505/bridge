package util

import (
	"fmt"
	"github.com/curtis0505/bridge/libs/logger"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/exp/constraints"
	"math/big"
	"strconv"
	"strings"
)

type Number interface {
	constraints.Integer | constraints.Float
}

var ZeroDecimal128, _ = primitive.ParseDecimal128("0")

// ToDecimal128 null safe
// convert value to primitive.Decimal128
func ToDecimal128(v any) primitive.Decimal128 {
	decimalZero, _ := primitive.ParseDecimal128("0")
	if v == nil {
		return decimalZero
	}
	valueD, err := primitive.ParseDecimal128(ToString(v))
	if err != nil {
		return decimalZero
	}
	return valueD
}

func ToDecimal128V2(v any) primitive.Decimal128 {
	decimalZero, _ := primitive.ParseDecimal128("0")
	if v == nil {
		return decimalZero
	}

	valueS := ToString(v)
	if valueS == "0" {
		return decimalZero
	}

	valueD, err := primitive.ParseDecimal128(ToString(v))
	if err != nil {
		return decimalZero
	}
	return valueD
}

// ToBigInt null safe
// convert value to *big.Int
func ToBigInt(val any) *big.Int {
	bigZero := big.NewInt(0)
	switch v := val.(type) {
	case *big.Int:
		if v == nil {
			return bigZero
		}
		return v
	case *primitive.Decimal128:
		if v.IsZero() {
			return bigZero
		}
		value, ok := new(big.Int).SetString(strings.Split(v.String(), ".")[0], 0)
		if !ok {
			return bigZero
		}
		return value

	case primitive.Decimal128:
		if v.IsZero() {
			return bigZero
		}
		value, ok := new(big.Int).SetString(strings.Split(v.String(), ".")[0], 0)
		if !ok {
			return bigZero
		}
		return value
	case decimal.Decimal:
		return v.BigInt()
	case *decimal.Decimal:
		return v.BigInt()
	case int64:
		return big.NewInt(v)
	case int32:
		return big.NewInt(int64(v))
	case int8:
		return big.NewInt(int64(v))
	case uint64:
		return big.NewInt(int64(v))
	case uint32:
		return big.NewInt(int64(v))
	case uint8:
		return big.NewInt(int64(v))
	case int:
		return big.NewInt(int64(v))
	case float64:
		return big.NewInt(int64(v))
	case string:
		value, ok := new(big.Int).SetString(strings.Split(v, ".")[0], 0)
		if !ok {
			return bigZero
		}
		return value
	default:
		return bigZero
	}
}

func IsZero(val any) bool {
	switch v := val.(type) {
	case float64:
		return v == 0
	}
	return ToBigInt(val).Cmp(big.NewInt(0)) == 0
}

// ToString null safe
// convert value to string
func ToString(val any) string {
	switch v := val.(type) {
	case primitive.Decimal128:
		if v.IsZero() {
			return ToString(0)
		}
		return v.String()
	case fmt.Stringer:
		return v.String()
	case float64:
		return fmt.Sprintf("%f", v)
	case float32:
		return fmt.Sprintf("%f", v)
	case int64:
		return fmt.Sprintf("%d", v)
	case int32:
		return fmt.Sprintf("%d", v)
	case uint64:
		return fmt.Sprintf("%d", v)
	case int:
		return fmt.Sprintf("%d", v)
	case string:
		return v
	case nil:
		return ""
	default:
		return fmt.Sprintf("%v", v)
	}
}

func ToEther(val any) decimal.Decimal {
	return ToEtherWithDecimal(val, 18)
}

func ToEtherWithDecimal[T Number](val any, d T) decimal.Decimal {
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(d)))
	num, _ := decimal.NewFromString(ToBigInt(val).String())
	result := num.Div(mul)

	return result
}

// ToDecimal wei to decimals
func ToDecimal[T Number](val any, d T) decimal.Decimal {
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(d)))
	num, _ := decimal.NewFromString(ToDecimal128(val).String())
	result := num.Div(mul)
	return result
}

func ToWeiWithDecimal[T Number](val any, d T) *big.Int {
	num, _ := decimal.NewFromString(ToString(val))
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(d)))
	return ToBigInt(num.Mul(mul))
}

// ToWei ether to wei
func ToWei(v any) *big.Int {
	return ToWeiWithDecimal(v, 18)
}

func Round[T Number](v any, d T) float64 {
	num, _ := ToDecimal(v, 0).Round(int32(d)).Float64()
	return num
}

// IsVisibleAmount 네오핀에서 6자리까지 표시하고 6자리 미만은 체크하지 않아 사용하는 함수
// unit: 가장 작은 단위(ex. wei, satoshi, uatom), tokenDecimal: 해당 토큰의 decimal
func IsVisibleAmount(unit any, tokenDecimal int) bool {
	amount := ToDecimal(unit, tokenDecimal-types.NeopinVisibleDecimal)
	amount = amount.RoundFloor(0)

	return amount.IsPositive()
}

// Deprecated: SafeBigIntToDecimal128
// use ToDecimal128() instead
func SafeBigIntToDecimal128(v any) primitive.Decimal128 {
	logger.Warn("Deprecated", "method", "SafeBigIntToDecimal128", "alert", "deprecated")
	return ToDecimal128(v)
}

// Deprecated: SafeIntToDecimal128
// use ToDecimal128() instead
func SafeIntToDecimal128(v any) primitive.Decimal128 {
	logger.Warn("Deprecated", "method", "SafeIntToDecimal128", "alert", "deprecated")
	return ToDecimal128(v)
}

// Deprecated: SafeValueToBigInt
// use ToBigInt() instead
func SafeValueToBigInt(v any) *big.Int {
	logger.Warn("Deprecated", "method", "SafeValueToBigInt", "alert", "deprecated")
	return ToBigInt(v)
}

// Deprecated: SafeDecimalToBigInt
// use ToBigInt() instead
func SafeDecimalToBigInt(v any) (*big.Int, error) {
	logger.Warn("Deprecated", "method", "SafeDecimalToBigInt", "alert", "deprecated")
	return ToBigInt(v), nil
}

// Deprecated: BalanceToDecimal
func BalanceToDecimal(balance *big.Int, price float32, quantity int) decimal.Decimal {
	logger.Warn("Deprecated", "method", "BalanceToDecimal", "alert", "deprecated")

	// div = 10 ** quantity
	div := decimal.New(int64(1), int32(quantity))

	priceDecimal := decimal.NewFromFloat32(price)
	balanceDecimal := decimal.NewFromBigInt(balance, 0)
	mulValue := priceDecimal.Mul(balanceDecimal)

	return mulValue.Div(div)
}

func StringToInt64(sVal string) int64 {
	n, err := strconv.ParseInt(sVal, 10, 64)
	if err != nil {
		return -1
	}
	return n
}

// IntToBool 0 이면 false 아니면 true
func IntToBool(input int) bool {
	if input == 0 {
		return false
	} else {
		return true
	}
}
