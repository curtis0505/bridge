package farm

import (
	"github.com/curtis0505/bridge/libs/common"
	"math/big"
)

type EventDeposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
}

type EventMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
}

type EventBurn struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
}

type EventWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
}

type EventLogAttachDetach struct {
	Spid *big.Int
	Bpid *big.Int
}

type EventUpdateEndBlock struct {
	Bpid     *big.Int
	EndBlock *big.Int
}

type EventClaimReward struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
}

type EventClaimBonus struct {
	User   common.Address
	Spid   *big.Int
	Bpid   *big.Int
	Amount *big.Int
}
