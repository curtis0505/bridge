package staking

import (
	"github.com/curtis0505/bridge/libs/common"
	"math/big"
)

var (
	_ InputCommonStaking = &InputEnterLeaveStaking{}
	_ InputCommonStaking = &InputDepositWithdrawStaking{}
)

type InputCommonStaking interface {
	SafePid() *big.Int
	SafeAmount() *big.Int
	SafeProof() common.Bytes
}

type InputEnterLeaveStaking struct {
	Amount *big.Int
	Proof  common.Bytes
}

func (input InputEnterLeaveStaking) SafePid() *big.Int {
	return big.NewInt(0)
}

func (input InputEnterLeaveStaking) SafeAmount() *big.Int {
	if input.Amount != nil {
		return input.Amount
	}
	return big.NewInt(0)
}

func (input InputEnterLeaveStaking) SafeProof() common.Bytes {
	return common.Bytes{}
}

type InputDepositWithdrawStaking struct {
	Pid    *big.Int
	Amount *big.Int
	Proof  common.Bytes
}

func (input InputDepositWithdrawStaking) SafePid() *big.Int {
	if input.Pid != nil {
		return input.Pid
	}
	return nil
}

func (input InputDepositWithdrawStaking) SafeAmount() *big.Int {
	if input.Amount != nil {
		return input.Amount
	}
	return big.NewInt(0)
}

func (input InputDepositWithdrawStaking) SafeProof() common.Bytes {
	if input.Proof != nil {
		return input.Proof
	}
	return common.Bytes{}
}

type InputClaimReward struct {
	Spid  *big.Int
	Bpid  *big.Int
	Proof common.Bytes
}

type InputClaimBonus struct {
	Spid  *big.Int
	Bpid  *big.Int
	User  common.Address
	Proof common.Bytes
}

type InputConfirmRequest struct {
	Id         *big.Int
	FunctionId uint8
	FirstArg   [32]uint8
	SecondArg  [32]uint8
	ThirdArg   [32]uint8
}

type InputWithdawApprovedStaking struct {
	Id *big.Int
}

type InputSubmitApproveStakingWithdrawal struct {
	To    common.Address
	Value *big.Int
}

type InputRequestWithdraw struct {
	// NpKlayAmount
	NpKlayAmount *big.Int
	NpEthAmount  *big.Int
	// AmountNP : npETH
	AmountNp *big.Int
	Amount   *big.Int
	Proof    common.Bytes
}

func (input InputRequestWithdraw) SafeAmount() *big.Int {
	if input.NpKlayAmount != nil {
		return input.NpKlayAmount
	}

	if input.NpEthAmount != nil {
		return input.NpEthAmount
	}

	if input.AmountNp != nil {
		return input.AmountNp
	}

	if input.Amount != nil {
		return input.Amount
	}

	return big.NewInt(0)
}

type InputRequestClaim struct {
	AmountEth *big.Int
	Proof     common.Bytes
}

type InputStake struct {
	Proof common.Bytes
}

type InputHandleExitedValidator struct {
	AccBeaconValidatorsCount *big.Int
	BeaconBalance            *big.Int
	ValidatorPubKey          []byte
	ExitedAmount             *big.Int
}

type InputGovernanceVote struct {
	Choice  *big.Int
	Id      string
	Address common.Address
}

type InputClaimMembershipReward struct {
	Address common.Address
}

type InputBuyVoucher struct {
	Amount          *big.Int
	MinSharesToMint *big.Int
}

type InputSellVoucherNew struct {
	ClaimAmount         *big.Int
	MaximumSharesToBurn *big.Int
}

type InputUnstakeClaimTokensNew struct {
	UnbondNonce *big.Int
}
