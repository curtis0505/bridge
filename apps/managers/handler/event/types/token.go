package types

import "math/big"

type Token struct {
	Chain   string
	Address string
	Amount  *big.Int
	Decimal uint8
}
