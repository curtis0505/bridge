package farm

import (
	"github.com/curtis0505/bridge/libs/common"
	"github.com/curtis0505/bridge/libs/types"
	"math/big"
)

var (
	_ types.CallMsgUnmarshaler = &OutputPendingReward{}
	_ types.CallMsgUnmarshaler = &OutputPoolInfo{}
	_ types.CallMsgUnmarshaler = &OutputBonusPoolInfo{}
	_ types.CallMsgUnmarshaler = &OutputToken{}
	_ types.CallMsgUnmarshaler = &OutputTotalSupply{}
	_ types.CallMsgUnmarshaler = &OutputUserAmount{}
	_ types.CallMsgUnmarshaler = &OutputUserInfo{}
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

// OutputPoolInfo
// MasterChef - poolInfo
type OutputPoolInfo struct {
	StakeToken         string
	RewardToken        string
	LastRewardBlock    *big.Int
	RewardPerBlock     *big.Int
	AccRewardPerShare  *big.Int
	NextRewardPerBlock *big.Int
	NextBlockNumber    *big.Int
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
	output.BonusChef = v[7].(common.Address).String()
	output.Bpid = v[8].(*big.Int)
}

// OutputBonusPoolInfo
// BonusChef - poolInfo
type OutputBonusPoolInfo struct {
	RewardToken        string
	LastRewardBlock    *big.Int
	RewardPerBlock     *big.Int
	AccRewardPerShare  *big.Int
	StartBlock         *big.Int
	EndBlock           *big.Int
	NextRewardPerBlock *big.Int
	NextBlockNumber    *big.Int
	Spid               *big.Int
	StakeSupply        *big.Int
	IsAttached         bool
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

type OutputUserInfo struct {
	Amount        *big.Int
	RewardDebt    *big.Int
	ClaimedReward *big.Int
}

func (output *OutputUserInfo) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
	output.RewardDebt = v[1].(*big.Int)
	output.ClaimedReward = v[2].(*big.Int)
}

type OutputUserAmount struct {
	Amount *big.Int
}

func (output *OutputUserAmount) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
}

type OutputToken struct {
	Address common.Address
}

func (output *OutputToken) Unmarshal(v []interface{}) {
	output.Address = v[0].(common.Address)
}

type OutputGetReserves struct {
	Token0Amount *big.Int
	Token1Amount *big.Int
}

func (output *OutputGetReserves) Unmarshal(v []interface{}) {
	output.Token0Amount = v[0].(*big.Int)
	output.Token1Amount = v[1].(*big.Int)
}

type OutputTotalSupply struct {
	Amount *big.Int
}

func (output *OutputTotalSupply) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
}

type OutputSlot0 struct {
	SqrtPriceX96               *big.Int
	Tick                       uint
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}

func (output *OutputSlot0) Unmarshal(v []interface{}) {
	output.SqrtPriceX96 = v[0].(*big.Int)
}
