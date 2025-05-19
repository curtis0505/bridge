package util

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/big"
	"testing"
)

var (
	emptyD128   = primitive.Decimal128{}                           // empty
	zeroD128, _ = primitive.ParseDecimal128("0")                   // 0
	testD128    = primitive.NewDecimal128(3476215962376601600, 11) // 1.1
	oneD128, _  = primitive.ParseDecimal128("1")

	emptyD = decimal.Decimal{}         // empty
	zeroD  = decimal.NewFromInt(0)     // 0
	testD  = decimal.NewFromFloat(1.1) // 1.1

	emptyString = ""    // empty
	zeroString  = "0"   // 0
	testString  = "1.1" // 1.1

	emptyBigInt = &big.Int{}    // empty
	zeroBigInt  = big.NewInt(0) // 0
	testBigInt  = big.NewInt(1) // 1.1
)

func TestDecimal128(t *testing.T) {
	// note min decimal 128 Exp is -6178
	assert.Equal(t, -6176, primitive.MinDecimal128Exp)

	// go.mongodb.org/mongo-driver/bson/primitive/decimal.go
	//
	// ...
	//
	// if exp < 0 {
	//	 return string(repr[last+posSign:]) + "E" + strconv.Itoa(exp)
	// }
	//
	// empty decimal128 : "0" + "E" + "primitive.MinDecimal128Exp" = "0E-6176"
	const EmptyDecimal128 = "0E-6176"
	assert.Equal(t, EmptyDecimal128, emptyD128.String())

	// "0" != emptyD128
	assert.NotEqual(t, "0", emptyD128.String())

	// "0" == ToString(emptyD128)
	assert.Equal(t, "0", ToString(emptyD128))

	// emptyD128 != zeroD128
	assert.NotEqual(t, emptyD128.String(), zeroD128.String())

	// emptyD128 != ToString(emptyD128)
	assert.NotEqual(t, emptyD128.String(), ToString(emptyD128))

	// in to deep
	// compare uint64 bytes
	emptyHigh, emptyLow := emptyD128.GetBytes()
	zeroHigh, zeroLow := zeroD128.GetBytes()

	// compare high, not equal
	assert.NotEqual(t, emptyHigh, zeroHigh)

	// compare low, equal
	assert.Equal(t, emptyLow, zeroLow)

	// "0" == zeroD128
	assert.Equal(t, "0", zeroD128.String())

	// "0" == ToString(zeroD128)
	assert.Equal(t, "0", ToString(zeroD128))

	// zeroD128 == ToString(zeroD128)
	assert.Equal(t, zeroD128.String(), ToString(zeroD128))

	// "1.1" == testD128
	assert.Equal(t, "1.1", testD128.String())

	// "1.1" == ToString(testD128)
	assert.Equal(t, "1.1", ToString(testD128))

	// testD128 = ToString(testD128)
	assert.Equal(t, testD128.String(), ToString(testD128))
}

func TestToString(t *testing.T) {
	assert.Equal(t, "0", ToString(emptyD))
	assert.Equal(t, "0", ToString(emptyD128))
	assert.Equal(t, "0", ToString(emptyBigInt))
	assert.Equal(t, "", ToString(emptyString))

	assert.Equal(t, "0", ToString(zeroD))
	assert.Equal(t, "0", ToString(zeroD128))
	assert.Equal(t, "0", ToString(zeroBigInt))
	assert.Equal(t, "0", ToString(zeroString))

	assert.Equal(t, "1.1", ToString(testD))
	assert.Equal(t, "1.1", ToString(testD128))
	assert.Equal(t, "1", ToString(testBigInt))
	assert.Equal(t, "1.1", ToString(testString))
}

func TestToBigInt(t *testing.T) {
	bigZero := big.NewInt(0)
	bigOne := big.NewInt(1)

	assert.Equal(t, bigZero, ToBigInt(emptyD))
	assert.Equal(t, bigZero, ToBigInt(emptyD128))
	assert.Equal(t, bigZero, ToBigInt(emptyBigInt))
	assert.Equal(t, bigZero, ToBigInt(emptyString))

	assert.Equal(t, bigZero, ToBigInt(zeroD))
	assert.Equal(t, bigZero, ToBigInt(zeroD128))
	assert.Equal(t, bigZero, ToBigInt(zeroBigInt))
	assert.Equal(t, bigZero, ToBigInt(zeroString))

	assert.Equal(t, bigOne, ToBigInt(testD))
	assert.Equal(t, bigOne, ToBigInt(testD128))
	assert.Equal(t, bigOne, ToBigInt(testBigInt))
	assert.Equal(t, bigOne, ToBigInt(testString))
}

func TestToDecimal128(t *testing.T) {
	assert.Equal(t, zeroD128, ToDecimal128(emptyD))
	assert.Equal(t, zeroD128, ToDecimal128(emptyD128))
	assert.Equal(t, zeroD128, ToDecimal128(emptyBigInt))
	assert.Equal(t, zeroD128, ToDecimal128(emptyString))

	assert.Equal(t, zeroD128, ToDecimal128(zeroD))
	assert.Equal(t, zeroD128, ToDecimal128(zeroD128))
	assert.Equal(t, zeroD128, ToDecimal128(zeroBigInt))
	assert.Equal(t, zeroD128, ToDecimal128(zeroString))

	assert.Equal(t, testD128, ToDecimal128(testD))
	assert.Equal(t, testD128, ToDecimal128(testD128))
	assert.Equal(t, oneD128, ToDecimal128(testBigInt))
	assert.Equal(t, testD128, ToDecimal128(testString))
}

func TestIsVisibleAmount(t *testing.T) {
	assert.Equal(t, true, IsVisibleAmount(1_123_000_000_000, 18))
	assert.Equal(t, true, IsVisibleAmount(1_000_000_000_000, 18))
	assert.Equal(t, false, IsVisibleAmount(999_000_000_000, 18))
	assert.Equal(t, false, IsVisibleAmount(-1_999_000_000_000, 18))
	assert.Equal(t, true, IsVisibleAmount(1, 6))
	assert.Equal(t, false, IsVisibleAmount(10, 8))
	assert.Equal(t, false, IsVisibleAmount(0, 8))
	assert.Equal(t, false, IsVisibleAmount(0, 18))
	assert.Equal(t, false, IsVisibleAmount(0, 1))
	assert.Equal(t, true, IsVisibleAmount(1, 0))
	assert.Equal(t, false, IsVisibleAmount(0, 0))
	assert.Equal(t, false, IsVisibleAmount(-0, 0))
}
