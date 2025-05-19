package stable

import (
	"encoding/json"
	"math/big"
)

// OutputEstimate
// - MethodEstimateDAI
// = MethodEstimateUSDC
// - MethodEstimateUSDT
type OutputEstimate struct {
	DAIAmount  *big.Int
	USDeAmount *big.Int
}

func (output *OutputEstimate) Unmarshal(v []any) {
	output.DAIAmount = v[0].(*big.Int)
	output.USDeAmount = v[1].(*big.Int)
}

type OutputConvertToDAI struct {
	DAIForSDAI  *big.Int
	DAIForSUSDe *big.Int
	DAIForUSDe  *big.Int
}

func (output *OutputConvertToDAI) Unmarshal(v []any) {
	output.DAIForSDAI = v[0].(*big.Int)
	output.DAIForSUSDe = v[1].(*big.Int)
	output.DAIForUSDe = v[2].(*big.Int)
}

// OutputCoolDowns
// - MethodGetCoolDownForBoost
// - MethodCoolDowns
type OutputCoolDowns struct {
	CoolDownEnd      *big.Int `json:"cooldownEnd"`
	UnderlyingAmount *big.Int `json:"underlyingAmount"`
}

func (output *OutputCoolDowns) Unmarshal(v []any) {
	bz, err := json.Marshal(v[0])
	if err != nil {
		return
	}
	json.Unmarshal(bz, output)
}

type OutputCoolDownDuration struct {
	CoolDownDuration *big.Int
}

func (output *OutputCoolDownDuration) Unmarshal(v []any) {
	output.CoolDownDuration = v[0].(*big.Int)
}

// OutputUsers 유저가 실제로 스테이킹한 수량
// - MethodUsers
type OutputUsers struct {
	SDAIAmount        *big.Int
	SUSDeAmount       *big.Int
	PendingUSDeAmount *big.Int
	LastWithdrawTime  *big.Int
}

func (output *OutputUsers) Unmarshal(v []any) {
	output.SDAIAmount = v[0].(*big.Int)
	output.SUSDeAmount = v[1].(*big.Int)
	output.PendingUSDeAmount = v[2].(*big.Int)
	output.LastWithdrawTime = v[3].(*big.Int)
}

// OutputConvertToAssets
// - MethodConvertToAssets
type OutputConvertToAssets struct {
	DAIAmount      *big.Int
	USDeAmount     *big.Int
	DAIBoostAmount *big.Int
}

func (output *OutputConvertToAssets) Unmarshal(v []any) {
	output.DAIAmount = v[0].(*big.Int)
	output.USDeAmount = v[1].(*big.Int)
	output.DAIBoostAmount = v[2].(*big.Int)
}

// OutputConvertToShares
// - MethodConvertToShares
type OutputConvertToShares struct {
	SDAIAmount  *big.Int
	SUSDeAmount *big.Int
}

func (output *OutputConvertToShares) Unmarshal(v []any) {
	output.SDAIAmount = v[0].(*big.Int)
	output.SUSDeAmount = v[1].(*big.Int)
}
