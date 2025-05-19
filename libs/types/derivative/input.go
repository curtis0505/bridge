package derivative

import "math/big"

type InputWithdraw struct {
	ShareAmount *big.Int
}

type InputClaim struct {
	RequestId *big.Int
}
