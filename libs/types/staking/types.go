package staking

import "fmt"

type eventName string

func (e eventName) String() string { return string(e) }

const (
	NptStaking   = "NPT-staking"
	FnsaStaking  = "FNSA-staking"
	TfnsaStaking = "TFNSA-staking"
	// Deprecated: KlayStaking
	KlayStaking = "KLY-staking"
	KaiaStaking = "KAIA-staking"
	// Deprecated: KlayLiquidStaking
	KlayLiquidStaking = "KLAY-liquid-staking"
	KaiaLiquidStaking = "KAIA-liquid-staking"
	TronStakingV2     = "TRX-staking-v2"
	EthStaking        = "ETH-staking"
	EthLiquidStaking  = "ETH-liquid-staking"
	MaticStaking      = "MATIC-staking"
	CosmosStaking     = "%s-staking"

	Stk         = "STK"
	ContractStk = "CONTRACT-STK"
	StkWithdraw = "STK-WITHDRAW"
	UnStk       = "UNSTK"
	UnStkr      = "UNSTKR"
	ClaimReward = "CLAIM-REWARD"
	ClaimBonus  = "CLAIM-BONUS"
	Claim       = "CLAIM"

	InitialLockupStaking = "initialLockupStaking"
	StStakeChef          = "StStakeChef"
	StBonusChef          = "StBonusChef"

	StStakeChefId = "ststakechef"
	StBonusChefId = "stbonuschef"

	MethodPendingReward = "pendingReward"
	MethodUserInfo      = "userInfo"
	MethodPoolInfo      = "poolInfo"

	StakingEnterStaking    = "enterStaking"
	StakingDeposit         = "deposit"
	StakingLeaveStaking    = "leaveStaking"
	StakingWithdraw        = "withdraw"
	StakingUpdateBonusChef = "updateBonusChef"

	EventStakeDeposit        eventName = "Deposit"
	EventStakeWithdraw       eventName = "Withdraw"
	EventStakeClaimReward    eventName = "ClaimReward"
	EventStakeClaimBonus     eventName = "ClaimBonus"
	EventStakeUpdateEndBlock eventName = "UpdateEndBlock"
	EventAttachStakeChef     eventName = "AttachStakeChef"
	EventDetachStakeChef     eventName = "DetachStakeChef"

	// Lcokup staking
	MethodDepositToStaking          = "depositToStaking"
	MethodWithdrawFromStaking       = "withdrawFromStaking"
	MethodWithdrawWhenFailedFunding = "withdrawWhenFailedFunding"

	MethodGetTotalStake    = "getTotalStake"
	MethodGetLiquidRewards = "getLiquidRewards"
	MethodCurrentEpoch     = "currentEpoch"

	ValidatorShareId = "validator-share"
	StakingInfoId    = "staking-info"
	EventsHubId      = "events-hub"
	StakeManagerId   = "stake-manager"
)

const (
	LiquidStk    = "LIQUID-STK"
	LiquidUnStk  = "LIQUID-UNSTK"
	LiquidUnStkr = "LIQUID-UNSTKR"

	NpEth = "NpEth"

	LockupStk = "LOCKUP-STK"

	// method names
	Submit       = "stakeKlay"
	LiquidSubmit = "submit"

	Stake      = "stake"
	UnStake    = "unstake"
	ClaimStake = "claimStake"

	RequestWithdraw         = "submitApproveStakingWithdrawal"
	LiquidRequestWithdraw   = "requestWithdraw"
	LiquidDistributeRewards = "distributeRewards"
	LiquidRequestClaim      = "requestClaim"

	ConfirmRequest          = "confirmRequest"
	WithdrawApprovedStaking = "withdrawApprovedStaking"

	HandleDepositValidators   = "handleDepositValidators"
	HandleExitValidator       = "handleExitValidator"
	HandleExitedValidator     = "handleExitedValidator"
	HandleElRewards           = "handleElRewards"
	HandleWithdrawReq         = "handleWithdrawReq"
	HandleCompensateBufferEth = "handleCompensateBufferEth"

	// event names
	StakeKlay                    = "StakeKlay"
	SubmitRequest                = "SubmitRequest"
	SubmitEvent                  = "SubmitEvent"
	RequestWithdrawEvent         = "RequestWithdrawEvent"
	ApproveStakingWithdrawal     = "ApproveStakingWithdrawal"
	WithdrawApprovedStakingEvent = "WithdrawApprovedStaking"
	DistributeRewardsEvent       = "DistributeRewardsEvent"
	HandleElRewardsEvent         = "HandleElRewardsEvent"

	StakeEvent   = "StakeEvent"
	UnstakeEvent = "UnstakeEvent"

	HandleDepositValidatorsEvent = "HandleDepositEvent"

	BuyVoucher            = "buyVoucher"
	SellVoucherNew        = "sellVoucher_new"
	UnstakeClaimTokensNew = "unstakeClaimTokens_new"
	WithdrawRewards       = "withdrawRewards"
)

const (
	Maintenance = "staking"
)

type StakeType uint8

const (
	StakeTypeLiquid = StakeType(0) + iota
	StakeTypeStake
)

const (
	ClaimTypeUnstake = "UNSTK"
	ClaimTypeReward  = "REWARD"
)

type StakingOrderKeyType string

func GetStakingOrderKey(chain string, itemName string) StakingOrderKeyType {
	return StakingOrderKeyType(fmt.Sprintf("%s:%s", chain, itemName))
}
