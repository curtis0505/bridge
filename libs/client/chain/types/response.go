package types

import (
	"errors"
	"github.com/curtis0505/bridge/libs/common"
	"math/big"
	"time"
)

var (
	NotFound = errors.New("not found")
)

type SendTxAsyncResult struct {
	Result SendTxResultType `json:"result"`
	Error  string           `json:"error"`
	TxHash common.Hash      `json:"txHash"`
	Hash   string           `json:"hash"`
}

func NewSendTxAsyncResult(result SendTxResultType, errMsg string, txHash common.Hash) *SendTxAsyncResult {
	return &SendTxAsyncResult{Result: result, Error: errMsg, TxHash: txHash}
}

type TronStaking struct {
	BandwidthFrozenAmountV1 *big.Int
	EnergyFrozenAmountV1    *big.Int
	BandwidthFrozenAmountV2 *big.Int
	EnergyFrozenAmountV2    *big.Int
}

type Staking struct {
	Chain        string
	Amount       *big.Int // 스테이킹 수량
	AmountLegacy *big.Int // Delegated 스테이킹 수량 (trx v1)

	Reward     Reward       // Reward 보상 수량
	Claimable  []Claimable  // Claimable 출금 가능
	Withdrawal []Withdrawal // Withdrawal 출금 대기

	TronStaking *TronStaking // 트론 스테이킹 Frozen 상세정보
}

func (s *Staking) GetRewardAmount() *big.Int { return s.Reward.Amount }
func (s *Staking) GetStakedAmount() *big.Int { return s.Amount }
func (s *Staking) GetClaimableAmount() *big.Int {
	amount := big.NewInt(0)

	for _, claimable := range s.Claimable {
		amount = new(big.Int).Add(amount, claimable.Amount)
	}
	return amount
}
func (s *Staking) GetWithdrawalAmount() *big.Int {
	if len(s.Withdrawal) == 0 {
		return big.NewInt(0)
	}
	amount := big.NewInt(0)
	for _, withdrawal := range s.Withdrawal {
		amount = new(big.Int).Add(amount, withdrawal.Amount)
	}

	return amount
}

type Reward struct {
	Chain         string
	Amount        *big.Int
	ClaimableTime time.Time
}

type Claimable struct {
	Chain  string
	Amount *big.Int
}

type Withdrawal struct {
	Chain          string
	Amount         *big.Int
	CompletionTime time.Time
}
