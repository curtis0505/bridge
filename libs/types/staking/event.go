package staking

import (
	"github.com/curtis0505/bridge/libs/common"
	"math/big"
)

type EventDeposit struct {
	User   common.Address `json:"user"`
	Amount *big.Int       `json:"amount"`
	Pid    *big.Int       `json:"pid"`
}

type EventWithdraw struct {
	User   common.Address `json:"user"`
	Amount *big.Int       `json:"amount"`
	Pid    *big.Int       `json:"pid"`
}

type EventClaimReward struct {
	User   common.Address `json:"user"`
	Amount *big.Int       `json:"amount"`
	Pid    *big.Int       `json:"pid"`
}

type EventClaimBonus struct {
	User   common.Address `json:"user"`
	Spid   *big.Int       `json:"spid"`
	Bpid   *big.Int       `json:"bpid"`
	Amount *big.Int       `json:"amount"`
}

type EventLogAttachDetach struct {
	Spid *big.Int
	Bpid *big.Int
}

type EventSubmitEvent struct {
	From         common.Address
	Amount       *big.Int
	AmountToMint *big.Int
}

type EventRequestWithdraw struct {
	From             common.Address
	AmountToBurn     *big.Int
	AmountToWithdraw *big.Int
}

type EventNpEthRequestWithdraw struct {
	From                common.Address
	AmountNpToBurn      *big.Int
	AmountEth           *big.Int
	TotalWithdrawReqEth *big.Int
	WithdrawReqIndex    *big.Int
}

type EventSubmitRequest struct {
	Id         *big.Int
	From       common.Address
	FunctionId uint8
	FirstArg   [32]uint8
	SecondArg  [32]uint8
	ThirdArg   [32]uint8
}

type EventApproveStakingWithdrawal struct {
	ApprovedWithdrawalId *big.Int
	To                   common.Address
	Value                *big.Int
	WithdrawableFrom     *big.Int
}

type EventWithdrawApprovedStaking struct {
	ApprovedWithdrawalId *big.Int
	To                   common.Address
	Value                *big.Int
}

type EventRequestClaim struct {
	From            common.Address
	AmountEth       *big.Int
	WithdrawableEth *big.Int
	WithdrawReqEth  *big.Int
}

type EventStakeEvent struct {
	From         common.Address
	Amount       *big.Int
	AmountToMint *big.Int
}

type EventUnstakeEvent struct {
	From                common.Address
	AmountNpToBurn      *big.Int
	AmountEth           *big.Int
	TotalWithdrawReqEth *big.Int
	WithdrawReqIndex    *big.Int
}

type EventHandleDepositEvent struct {
	PubKey common.Bytes
	Amount *big.Int
}

type EventDistributeRewardsEvent struct {
	Amount       *big.Int
	AmountToMint *big.Int
}

type EventHandleElRewardsEvent struct {
	Amount       *big.Int
	AmountToMint *big.Int
}

type EventUpdateEndBlock struct {
	Bpid     *big.Int
	EndBlock *big.Int
}

// Lockup Staking
type EventLockupDeposit struct {
	From   common.Address `json:"_from"`
	Pid    *big.Int       `json:"_pid"`
	Amount *big.Int       `json:"_amount"`
}

type EventLockupWithdraw EventLockupDeposit

type EventShareMinted struct {
	ValidatorId *big.Int
	User        common.Address
	Amount      *big.Int
	Tokens      *big.Int
}

type EventShareBurnedWithId struct {
	ValidatorId *big.Int
	User        common.Address
	Amount      *big.Int
	Tokens      *big.Int
	Nonce       *big.Int
}

type EventDelegatorUnstakeWithId struct {
	ValidatorId *big.Int
	User        common.Address
	Amount      *big.Int
	Nonce       *big.Int
}
type EventDelegatorClaimedRewards struct {
	ValidatorId *big.Int
	User        common.Address
	Rewards     *big.Int
}
