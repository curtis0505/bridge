package derivative

import (
	"github.com/curtis0505/bridge/libs/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var (
	_ types.CallMsgUnmarshaler = &OutputUser{}
	_ types.CallMsgUnmarshaler = &OutputWithdrawalRequestStatus{}
)

type OutputUser struct {
	ShareAmount   *big.Int
	ClaimedAmount *big.Int
}

func (output *OutputUser) Unmarshal(v []interface{}) {
	output.ShareAmount = v[0].(*big.Int)
	output.ClaimedAmount = v[1].(*big.Int)
}

type OutputWithdrawalRequestStatus struct {
	WithdrawalRequests []WithdrawalRequestStatus
}

type WithdrawalRequestStatus struct {
	AmountOfStETH  *big.Int
	AmountOfShares *big.Int
	Owner          common.Address
	Timestamp      *big.Int
	IsFinalized    bool
	IsClaimed      bool
}

func (output *OutputWithdrawalRequestStatus) Unmarshal(v []interface{}) {
	output.WithdrawalRequests = make([]WithdrawalRequestStatus, 0)
	out0 := *abi.ConvertType(v[0], new([]WithdrawalRequestStatus)).(*[]WithdrawalRequestStatus)
	output.WithdrawalRequests = out0
}

type OutputRequestIds struct {
	RequestIds []*big.Int
}

func (output *OutputRequestIds) Unmarshal(v []interface{}) {
	output.RequestIds = v[0].([]*big.Int)
}

type OutputTotalShareAmount struct {
	Amount *big.Int
}

func (output *OutputTotalShareAmount) Unmarshal(v []interface{}) {
	output.Amount = v[0].(*big.Int)
}

type OutputConvertToAssets struct {
	Assets *big.Int
}

func (output *OutputConvertToAssets) Unmarshal(v []interface{}) {
	output.Assets = v[0].(*big.Int)
}

type OutputProtocolFeeBP struct {
	ProtocolFeeBp uint16
}

func (output *OutputProtocolFeeBP) Unmarshal(v []interface{}) {
	output.ProtocolFeeBp = v[0].(uint16)
}

type OutputProtocolFeeTo struct {
	Address common.Address
}

func (output *OutputProtocolFeeTo) Unmarshal(v []interface{}) {
	output.Address = v[0].(common.Address)
}
