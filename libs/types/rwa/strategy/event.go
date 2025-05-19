package strategy

import (
	"github.com/curtis0505/bridge/libs/common"
	"math/big"
)

type EventDeposit struct {
	From     common.Address //`abi:"from"`
	Receiver common.Address //`abi:"receiver"`
	Assets   *big.Int       //`abi:"assets"`
	Shares   *big.Int       //`abi:"shares"`
}

type EventRedeem struct {
	From     common.Address //`abi:"from"`
	Receiver common.Address //`abi:"receiver"`
	Shares   *big.Int       //`abi:"shares"`
	Assets   *big.Int       //`abi:"assets"`
}
