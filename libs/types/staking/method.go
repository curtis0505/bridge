package staking

const (
	// none proof[contract(token) staking]
	AbiDeposit      = "deposit(uint256,uint256)"
	AbiEnterStaking = "enterStaking(uint256)"
	AbiWithdraw     = "withdraw(uint256,uint256)"
	AbiLeaveStaking = "leaveStaking(uint256)"
	AbiClaimReward  = "claimReward(uint256,uint256)"

	// proof[contract(token) staking]
	AbiDepositProof      = "deposit(uint256,uint256,bytes)"
	AbiEnterStakingProof = "enterStaking(uint256,bytes)"
	AbiWithdrawProof     = "withdraw(uint256,uint256,bytes)"
	AbiLeaveStakingProof = "leaveStaking(uint256,bytes)"
	AbiClaimRewardProof  = "claimReward(uint256,uint256,bytes)"

	// lockup staking
	AbiDepositToStaking          = "depositToStaking(uint256,uint256)"
	AbiWithdrawFromStaking       = "withdrawFromStaking(uint256)"
	AbiWithdrawWhenFailedFunding = "withdrawWhenFailedFunding(uint256)"

	AbiSubmit          = "submit(bytes)"
	AbiRequestWithdraw = "requestWithdraw(uint256,bytes)"
	AbiRequestClaim    = "requestClaim(uint256,bytes)"

	AbiStake      = "stake(bytes)"
	AbiUnstake    = "unstake(uint256,bytes)"
	AbiClaimStake = "claimStake(uint256,bytes)"

	//polygon staking
	AbiBuyVoucher            = "buyVoucher(uint256,uint256)"
	AbiSellVoucherNew        = "sellVoucher_new(uint256,uint256)"
	AbiUnstakeClaimTokensNew = "unstakeClaimTokens_new(uint256)"
	AbiWithdrawRewards       = "withdrawRewards()"
)
