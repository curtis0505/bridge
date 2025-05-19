package staking

import (
	"encoding/json"
	"github.com/curtis0505/bridge/libs/common"
	"github.com/curtis0505/bridge/libs/types"
	"math/big"
)

var (
	_ types.CallMsgUnmarshaler = &OutputPendingReward{}
	_ types.CallMsgUnmarshaler = &OutputUserInfo{}
	_ types.CallMsgUnmarshaler = &OutputPoolInfo{}
	_ types.CallMsgUnmarshaler = &OutputBonusUserInfo{}
	_ types.CallMsgUnmarshaler = &OutputBonusPoolInfo{}
	_ types.CallMsgUnmarshaler = &OutputGetUserAmount{}
	_ types.CallMsgUnmarshaler = &OutputGetRequestIds{}
	_ types.CallMsgUnmarshaler = &OutputGetRequestInfo{}
	_ types.CallMsgUnmarshaler = &OutputGetApprovedStakingWithdrawalInfo{}

	_ types.CallMsgUnmarshaler = &OutputConvertNpKlayToKlay{}
	_ types.CallMsgUnmarshaler = &OutputConvertKlayToNpKlay{}
	_ types.CallMsgUnmarshaler = &OutputConvertNpEthToEth{}
	_ types.CallMsgUnmarshaler = &OutputConvertEthToNpEth{}
	_ types.CallMsgUnmarshaler = &OutputBeaconBalance{}
	_ types.CallMsgUnmarshaler = &OutPutTotalSupply{}
	_ types.CallMsgUnmarshaler = &OutputBalanceOf{}
	_ types.CallMsgUnmarshaler = &OutputTotalWithdrawableAmount{}
	_ types.CallMsgUnmarshaler = &OutputTotalWithdrawReqAmount{}
	_ types.CallMsgUnmarshaler = &OutputUsers{}
	_ types.CallMsgUnmarshaler = &BufferedEthResp{}
	_ types.CallMsgUnmarshaler = &OutputWithdrawReqLogNextStartIndex{}
	_ types.CallMsgUnmarshaler = &OutputWithdrawReqLogLength{}
)

// OutputPendingReward
// StStakeChef - pendingReward
// BonusChef - pendingReward
type OutputPendingReward struct {
	Amount *big.Int
}

func (output *OutputPendingReward) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
}

type OutputUserInfo struct {
	Amount            *big.Int
	RewardDebt        *big.Int
	ClaimReward       *big.Int
	UnlockBlockNumber *big.Int
}

func (output *OutputUserInfo) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
	output.RewardDebt = v[1].(*big.Int)
	output.ClaimReward = v[2].(*big.Int)
	output.UnlockBlockNumber = v[3].(*big.Int)
}

type OutputPoolInfo struct {
	StakeToken         string
	RewardToken        string
	LastRewardBlock    *big.Int
	RewardPerBlock     *big.Int
	AccRewardPerShare  *big.Int
	NextRewardPerBlock *big.Int
	NextBlockNumber    *big.Int
	LockupDuration     *big.Int
	BonusChef          string
	Bpid               *big.Int
}

func (output *OutputPoolInfo) Unmarshal(v []interface{}) {
	output.StakeToken = v[0].(common.Address).String()
	output.RewardToken = v[1].(common.Address).String()
	output.LastRewardBlock = v[2].(*big.Int)
	output.RewardPerBlock = v[3].(*big.Int)
	output.AccRewardPerShare = v[4].(*big.Int)
	output.NextRewardPerBlock = v[5].(*big.Int)
	output.NextBlockNumber = v[6].(*big.Int)
	output.LockupDuration = v[7].(*big.Int)
	output.BonusChef = v[8].(common.Address).String()
	output.Bpid = v[9].(*big.Int)
}

type OutputBonusUserInfo struct {
	Amount      *big.Int
	RewardDebt  *big.Int
	ClaimReward *big.Int
}

func (output *OutputBonusUserInfo) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
	output.RewardDebt = v[1].(*big.Int)
	output.ClaimReward = v[2].(*big.Int)
}

type OutputBonusPoolInfo struct {
	RewardToken        string   `json:"rewardToken"`
	LastRewardBlock    *big.Int `json:"lastRewardBlock"`
	RewardPerBlock     *big.Int `json:"rewardPerBlock"`
	AccRewardPerShare  *big.Int `json:"accRewardPerShare"`
	StartBlock         *big.Int `json:"startBlock"`
	EndBlock           *big.Int `json:"endBlock"`
	NextRewardPerBlock *big.Int `json:"nextRewardPerBlock"`
	NextBlockNumber    *big.Int `json:"nextBlockNumber"`
	Spid               *big.Int `json:"spid"`
	StakeSupply        *big.Int `json:"stakeSupply"`
	IsAttached         bool     `json:"isAttached"`
}

func (output *OutputBonusPoolInfo) Unmarshal(v []interface{}) {
	output.RewardToken = v[0].(common.Address).String()
	output.LastRewardBlock = v[1].(*big.Int)
	output.RewardPerBlock = v[2].(*big.Int)
	output.AccRewardPerShare = v[3].(*big.Int)
	output.StartBlock = v[4].(*big.Int)
	output.EndBlock = v[5].(*big.Int)
	output.NextRewardPerBlock = v[6].(*big.Int)
	output.NextBlockNumber = v[7].(*big.Int)
	output.Spid = v[8].(*big.Int)
	output.StakeSupply = v[9].(*big.Int)
	output.IsAttached = v[10].(bool)

}

type OutputGetUserAmount struct {
	Amount *big.Int
}

func (output *OutputGetUserAmount) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
}

type OutputGetRequestIds struct {
	Ids []*big.Int
}

func (output *OutputGetRequestIds) Unmarshal(v []interface{}) {
	output.Ids = v[0].([]*big.Int)
}

type OutputGetRequestInfo struct {
	FunctionId uint8
	FirstArg   [32]uint8
	SecondArg  [32]uint8
	ThirdArg   [32]uint8
	Proposer   common.Address
	Confirmers interface{}
	State      uint8
}

func (output *OutputGetRequestInfo) Unmarshal(v []interface{}) {
	output.FunctionId = v[0].(uint8)
	output.FirstArg = v[1].([32]uint8)
	output.SecondArg = v[2].([32]uint8)
	output.ThirdArg = v[3].([32]uint8)
	output.Proposer = v[4].(common.Address)
	output.Confirmers = v[5]
	output.State = v[6].(uint8)
}

type OutputGetApprovedStakingWithdrawalInfo struct {
	To               common.Address
	Value            *big.Int
	WithdrawableFrom *big.Int
	State            uint8
}

func (output *OutputGetApprovedStakingWithdrawalInfo) Unmarshal(v []interface{}) {
	output.To = v[0].(common.Address)
	output.Value = v[1].(*big.Int)
	output.WithdrawableFrom = v[2].(*big.Int)
	output.State = v[3].(uint8)
}

type OutputConvertKlayToNpKlay struct {
	Amount *big.Int
}

type OutputConvertNpEthToEth struct {
	Amount *big.Int
}

func (output *OutputConvertKlayToNpKlay) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
}

type OutputConvertNpKlayToKlay struct {
	Amount *big.Int
}

func (output *OutputConvertNpKlayToKlay) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
}

type OutputGetTotalPooledKlay struct {
	Amount *big.Int
}

func (output *OutputGetTotalPooledKlay) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
}

type OutputGetTotalPooledEth struct {
	Amount *big.Int
}

func (output *OutputGetTotalPooledEth) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
}

type OutputGetProtocolFeeBP struct {
	Amount *big.Int
}

func (output *OutputGetProtocolFeeBP) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
}

type OutputGetTotalStake struct {
	Amount *big.Int
}

func (output *OutputGetTotalStake) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
}

func (output *OutputConvertNpEthToEth) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
}

type OutputConvertEthToNpEth struct {
	Amount *big.Int
}

func (output *OutputConvertEthToNpEth) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
}

type OutputStakes struct {
	Amount *big.Int
}

func (output *OutputStakes) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
}

type OutputUsers struct {
	WithdrawReqAmount  *big.Int
	WithdrawableAmount *big.Int
}

func (output *OutputUsers) Unmarshal(v []interface{}) {
	output.WithdrawReqAmount = v[0].(*big.Int)
	output.WithdrawableAmount = v[1].(*big.Int)
}

type BufferedEthResp struct {
	BufferedEth *big.Int
}

func (b *BufferedEthResp) Unmarshal(v []interface{}) {
	b.BufferedEth = v[0].(*big.Int)
}

type OutputTotalWithdrawReqAmount struct {
	TotalWithdrawReqAmount *big.Int
}

func (output *OutputTotalWithdrawReqAmount) Unmarshal(v []interface{}) {
	output.TotalWithdrawReqAmount = v[0].(*big.Int)
}

type OutputTotalWithdrawableAmount struct {
	TotalWithdrawableAmount *big.Int
}

func (output *OutputTotalWithdrawableAmount) Unmarshal(v []interface{}) {
	output.TotalWithdrawableAmount = v[0].(*big.Int)
}

type OutputWithdrawReqLogLength struct {
	WithdrawReqLogLength *big.Int
}

func (output *OutputWithdrawReqLogLength) Unmarshal(v []interface{}) {
	output.WithdrawReqLogLength = v[0].(*big.Int)
}

type OutputWithdrawReqLogNextStartIndex struct {
	OutputWithdrawReqLogNextStartIndex *big.Int
}

func (output *OutputWithdrawReqLogNextStartIndex) Unmarshal(v []interface{}) {
	output.OutputWithdrawReqLogNextStartIndex = v[0].(*big.Int)
}

type OutputWithdrawReqLog struct {
	Amount *big.Int
	User   common.Address
	Stype  uint8
	Status uint8
}

func (output *OutputWithdrawReqLog) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
	output.User = v[1].(common.Address)
	output.Stype = v[2].(uint8)
	output.Status = v[3].(uint8)
}

type OutPutTotalSupply struct {
	TotalSupply *big.Int
}

func (output *OutPutTotalSupply) Unmarshal(v []interface{}) {
	output.TotalSupply = v[0].(*big.Int)
}

type OutputBeaconBalance struct {
	BeaconBalance *big.Int
}

func (output *OutputBeaconBalance) Unmarshal(v []interface{}) {
	output.BeaconBalance = v[0].(*big.Int)
}

type OutputBalanceOf struct {
	BalanceOf *big.Int
}

func (output *OutputBalanceOf) Unmarshal(v []interface{}) {
	output.BalanceOf = v[0].(*big.Int)
}

type OutputUsersLockupStaking struct {
	FundingAmount   *big.Int
	WithdrawnAmount *big.Int
	IsWithdrawn     bool
}

func (output *OutputUsersLockupStaking) Unmarshal(v []interface{}) {
	output.FundingAmount = v[0].(*big.Int)
	output.WithdrawnAmount = v[1].(*big.Int)
	output.IsWithdrawn = v[2].(bool)
}

type TimeLockParams struct {
	StartTs         *big.Int
	EndTs           *big.Int
	WithdrawTs      *big.Int // endTs + termDays * 1 days
	TermDays        uint32
	CollectGapDays  uint32
	WithdrawGapDays uint32
}

type AnnRateBPS struct {
	User     uint16 // user annual interest rate basis points
	Operator uint16 // operator annual interest rate basis points
}

type OutputProductsLockupStaking struct {
	TimeLockParams       TimeLockParams
	AnnRateBPS           AnnRateBPS
	TargetFundingAmount  *big.Int       // target funding amount
	CurrentFundingAmount *big.Int       // current funding amount
	CollectedAmount      *big.Int       // collected amount to delegatee
	ReturnedAmount       *big.Int       // returned amount from delegatee = collectedAmount + interestAmount
	OperatorFeeAmount    *big.Int       // operator fee amount
	WithdrawnAmount      *big.Int       // withdrawn amount
	FundingTokenAddr     common.Address // funding token address
	DelegateeAddr        common.Address // delegatee address
	LiquidAddr           common.Address // liquid address
	IsSuccessFunding     bool           // is successed
}

func (output *OutputProductsLockupStaking) Unmarshal(v []interface{}) {
	timeLockParams := TimeLockParams{}
	annRateBPS := AnnRateBPS{}

	timeLockParams_, _ := json.Marshal(v[0].(interface{}))
	json.Unmarshal(timeLockParams_, &timeLockParams)

	annRateBPS_, _ := json.Marshal(v[1].(interface{}))
	json.Unmarshal(annRateBPS_, &annRateBPS)

	output.TimeLockParams = timeLockParams
	output.AnnRateBPS = annRateBPS
	output.TargetFundingAmount = v[2].(*big.Int)
	output.CurrentFundingAmount = v[3].(*big.Int)
	output.CollectedAmount = v[4].(*big.Int)
	output.ReturnedAmount = v[5].(*big.Int)
	output.OperatorFeeAmount = v[6].(*big.Int)
	output.WithdrawnAmount = v[7].(*big.Int)
	output.FundingTokenAddr = v[8].(common.Address)
	output.DelegateeAddr = v[9].(common.Address)
	output.LiquidAddr = v[10].(common.Address)
	output.IsSuccessFunding = v[11].(bool)
}

type OutputUnbondNonces struct {
	Nonce *big.Int
}

func (output *OutputUnbondNonces) Unmarshal(v []interface{}) {
	output.Nonce = v[0].(*big.Int)
}

type OutputGetLiquidRewards struct {
	Amount *big.Int
}

func (output *OutputGetLiquidRewards) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
}

type OutputGetTotalStaked struct { //function name duplicate 이슈로 + d 추가
	Amount *big.Int
	Rate   *big.Int
}

func (output *OutputGetTotalStaked) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
	output.Rate = v[1].(*big.Int)
}

type OutputValidatorState struct {
	Amount      *big.Int // Total Staked on chain
	StakerCount *big.Int // Total Validator Count
}

func (output *OutputValidatorState) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
	output.StakerCount = v[1].(*big.Int)
}

type OutputValidators struct {
	Amount                *big.Int // Validator Staked
	Reward                *big.Int
	ActivationEpoch       *big.Int
	DeactivationEpoch     *big.Int
	JailTime              *big.Int
	Signer                common.Address
	ContractAddress       common.Address
	Status                uint8
	CommissionRate        *big.Int
	LastCommissionUpdate  *big.Int
	DelegatorsReward      *big.Int
	DelegatedAmount       *big.Int // Delegator Staked
	InitialRewardPerStake *big.Int
}

func (output *OutputValidators) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
	output.Reward = v[1].(*big.Int)
	output.ActivationEpoch = v[2].(*big.Int)
	output.DeactivationEpoch = v[3].(*big.Int)
	output.JailTime = v[4].(*big.Int)
	output.Signer = v[5].(common.Address)
	output.ContractAddress = v[6].(common.Address)
	output.Status = v[7].(uint8)
	output.CommissionRate = v[8].(*big.Int)
	output.LastCommissionUpdate = v[9].(*big.Int)
	output.DelegatorsReward = v[10].(*big.Int)
	output.DelegatedAmount = v[11].(*big.Int)
	output.InitialRewardPerStake = v[12].(*big.Int)
}

type OutputUnBondsNew struct {
	Shares        *big.Int
	WithdrawEpoch *big.Int
}

func (output *OutputUnBondsNew) Unmarshal(v []interface{}) {
	output.Shares = v[0].(*big.Int)
	output.WithdrawEpoch = v[1].(*big.Int)
}

type OutputCurrentEpoch struct {
	CurrentEpoch *big.Int
}

func (output *OutputCurrentEpoch) Unmarshal(v []interface{}) {
	output.CurrentEpoch = v[0].(*big.Int)
}
