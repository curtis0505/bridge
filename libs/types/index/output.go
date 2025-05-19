package index

import (
	"encoding/json"
	"github.com/curtis0505/bridge/libs/common"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/types/base"
	"math/big"
)

type OutputGetComponents struct {
	base.OutputAddresses
}

func (OutputGetComponents) MethodName() string { return "getComponents" }

type OutputGetDefaultPositionRealUnit struct {
	base.OutputBigInt
}

func (OutputGetDefaultPositionRealUnit) MethodName() string { return "getDefaultPositionRealUnit" }

type OutputTotalSupply struct {
	base.OutputBigInt
}

func (OutputTotalSupply) MethodName() string { return "totalSupply" }

/*
OutputGetEstimatedIssueSetAmount

	ISetToken _setToken,
	IERC20 _inputToken,
	uint256 _amountInput,
	DEXAdapter.SwapData memory _ethSwap,           // inputToken -> eth
	DEXAdapter.SwapData[] memory _componentsSwap   // (eth -> component)[]
*/
type OutputGetEstimatedIssueSetAmount struct {
	base.OutputBigInt
}

func (OutputGetEstimatedIssueSetAmount) MethodName() string { return "getEstimatedIssueSetAmount" }

/*
OutputGetAmountInToIssueExactSet

	ISetToken _setToken,
	IERC20 _inputToken,
	uint256 _amountSetToken,
	DEXAdapter.SwapData memory _ethSwap,          // inputToken -> eth
	DEXAdapter.SwapData[] memory _componentsSwap  // (eth -> component)[]
*/
type OutputGetAmountInToIssueExactSet struct {
	base.OutputBigInt
}

func (OutputGetAmountInToIssueExactSet) MethodName() string { return "getAmountInToIssueExactSet" }

/*
OutputGetAmountOutOnRedeemSet

	ISetToken _setToken,
	address _outputToken,
	uint256 _amountSetToken,
	DEXAdapter.SwapData memory _ethSwap,          // eth -> outputToken
	DEXAdapter.SwapData[] memory _componentsSwap  // (component -> eth)[]
*/
type OutputGetAmountOutOnRedeemSet struct {
	base.OutputBigInt
}

func (OutputGetAmountOutOnRedeemSet) MethodName() string { return "getAmountOutOnRedeemSet" }

type OutputFeeState struct {
	FeeRecipient              common.Address `json:"feeRecipient"`
	MaxStreamingFeePercentage *big.Int       `json:"maxStreamingFeePercentage"`
	StreamingFeePercentage    *big.Int       `json:"streamingFeePercentage"`
	LastStreamingFeeTimeStamp *big.Int       `json:"lastStreamingFeeTimeStamp"`
}

func (OutputFeeState) MethodName() string { return "feeStates" }

func (o *OutputFeeState) UnmarshalJSON(bz []byte) error {
	feeState := struct {
		FeeRecipient              string   `json:"feeRecipient"`
		MaxStreamingFeePercentage *big.Int `json:"maxStreamingFeePercentage"`
		StreamingFeePercentage    *big.Int `json:"streamingFeePercentage"`
		LastStreamingFeeTimeStamp *big.Int `json:"lastStreamingFeeTimeStamp"`
	}{}

	err := json.Unmarshal(bz, &feeState)
	if err != nil {
		return err
	}

	o.FeeRecipient = common.HexToAddress(types.ChainETH, feeState.FeeRecipient)
	o.MaxStreamingFeePercentage = feeState.MaxStreamingFeePercentage
	o.StreamingFeePercentage = feeState.StreamingFeePercentage
	o.LastStreamingFeeTimeStamp = feeState.LastStreamingFeeTimeStamp

	return nil
}

func (o *OutputFeeState) Unmarshal(v []any) {
	o.FeeRecipient = v[0].(common.Address)
	o.MaxStreamingFeePercentage = v[1].(*big.Int)
	o.StreamingFeePercentage = v[2].(*big.Int)
	o.LastStreamingFeeTimeStamp = v[3].(*big.Int)
}
