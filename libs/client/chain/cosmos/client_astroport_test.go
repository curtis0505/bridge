package cosmos

import (
	"context"
	"fmt"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	cosmostypes "github.com/curtis0505/bridge/libs/client/chain/cosmos/types"
	clienttypes "github.com/curtis0505/bridge/libs/client/chain/types"
	cosmoscommon "github.com/curtis0505/bridge/libs/common/cosmos"
	"github.com/curtis0505/bridge/libs/logger/v2"
	"github.com/curtis0505/bridge/libs/testutil"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/types/cosmos/astroport"
	"github.com/curtis0505/bridge/libs/types/cosmos/cw20"
	"github.com/curtis0505/bridge/libs/types/cosmos/cwutil"
	"github.com/curtis0505/bridge/libs/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"math/big"
	"testing"
)

type AstroPortTestSuite struct {
	suite.Suite

	env    testutil.TestENV
	ctx    context.Context
	client clienttypes.CosmosClient

	chain  string
	signer *types.Account

	factoryAddress   string
	tokenAddress     string
	xtokenAddress    string
	poolAddress      string
	lpAddress        string
	routerAddress    string
	vestingAddress   string
	generatorAddress string
	stakingAddress   string
	makerAddress     string
}

func TestAstroportSuite(t *testing.T) {
	logger.InitLog(logger.Config{UseTerminal: true, VerbosityTerminal: 5})
	suite.Run(t, new(AstroPortTestSuite))
}

func (suite *AstroPortTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.env = testutil.DEV
	suite.chain = types.ChainFNSA
	suite.client, _ = NewClient(envConfig[suite.chain][suite.env].restConfig)
	suite.client.SetAccountPrefix()

	suite.signer = GetTestAccount(suite.T(), "58d7ad6ece2c88ce01648aaf0799e27eb760e37f11d30b937f36529bdee926db")

	suite.T().Log(suite.SignerAddress())
	suite.factoryAddress = "link1jnnggf8enr4dx35qzpkjuth0ujzluz9jwdfee07rgcj082rxnunsvlms4a"   // factory
	suite.poolAddress = "link1lqum4f08cf9cs7jc7482cu4wr96u2tmleqd4lvkhsg2m9vrsc2yq42gmgs"      // pool
	suite.routerAddress = "link10fnppdrugtsxmr4nndjdun9jr8ahfv3j5aw7ea04wgnlg53m87tqus2sap"    // swap router
	suite.generatorAddress = "link1x3x472ydsvvz3609pdx0xpvm6xfxgjrmdmawhp2f264pjqua22wqjpj2d2" // reward
	suite.lpAddress = "link1h004lf2cx9hkshj4wz7gyrwkwf0whme2xvnehvwuhqyml3gulgjqglddj4"
	suite.tokenAddress = "link16tfcsjcg4trlnj2t4cpdq20q4d7nzusew3ev50scz30mmfzdx8ssw3wuhh"  // FSG
	suite.xtokenAddress = "link152dkhakhe9g8vuj6drvy88zcne4furtqy8pkggu66k877wjtecwsq82cfv" // xFSG
	suite.stakingAddress = "link1fuxfrnxcec6tyy37rrme04utqtw29wmpt0qz8266crdd5s0r9tvscg63za"
	suite.makerAddress = "link1nufa563efn866lwvmrlg8k7cfj5wj2u2d45wwqhsr0m0hc86zxqsfmpma6"
	suite.vestingAddress = "link1rqm73d4s5c2c7ccxt2dndhgduc9xqz3huyqwl2mpr7gptj5sx9tq9dw36e"
}

func (suite *AstroPortTestSuite) SignerAddress() string {
	return cosmoscommon.FromPublicKeyUnSafe(suite.chain, suite.signer.Secp256k1().PubKey().Bytes()).String()
}

func (suite *AstroPortTestSuite) TestQueryRouter() {
	var result1 astroport.QueryMsgSimulateSwapOperationsResponse
	err := suite.client.CallWasm(suite.ctx, suite.routerAddress, astroport.QuerySimulateSwapOperationsRequest{
		OfferAmount: util.ToString(big.NewInt(123456)),
		SwapOperations: []*astroport.SwapOperation{
			astroport.NewSwapOperation(
				astroport.NewAsset(astroport.AssetTypeCoin, "cony"),
				astroport.NewAsset(astroport.AssetTypeToken, suite.tokenAddress),
			),
		},
	}, &result1)
	assert.NoError(suite.T(), err)
	logger.Info("QuerySimulateSwapOperationsRequest", logger.BuildLogInput().WithData("response", result1))
}

func (suite *AstroPortTestSuite) TestQueryFactory() {
	var result1 astroport.QueryPairsResponse
	err := suite.client.CallWasm(suite.ctx, suite.factoryAddress, astroport.QueryPairsRequest{
		Limit: 10,
	}, &result1)
	assert.NoError(suite.T(), err)
	logger.Info("QueryPairsRequest", logger.BuildLogInput().WithData("response", result1))

	var result2 astroport.QueryPairResponse
	err = suite.client.CallWasm(suite.ctx, suite.factoryAddress, astroport.QueryPairRequest{
		AssetInfos: []*astroport.Asset{
			astroport.NewAsset(astroport.AssetTypeCoin, "cony"),
			astroport.NewAsset(astroport.AssetTypeToken, suite.tokenAddress),
		},
	}, &result2)
	assert.NoError(suite.T(), err)
	logger.Info("QueryPairRequest", logger.BuildLogInput().WithData("response", result2))

	var result3 astroport.QueryFeeInfoResponse
	err = suite.client.CallWasm(suite.ctx, suite.factoryAddress, astroport.QueryFeeInfoRequest{
		PairType: astroport.PairType{
			Xyk: &struct{}{},
		},
	}, &result3)
	assert.NoError(suite.T(), err)
	logger.Info("QueryFeeInfoRequest", logger.BuildLogInput().WithData("response", result3))
}

func (suite *AstroPortTestSuite) TestQueryPool() {
	var result1 astroport.QueryPoolResponse
	err := suite.client.CallWasm(suite.ctx, suite.poolAddress, astroport.QueryPoolRequest{}, &result1)
	assert.NoError(suite.T(), err)
	logger.Info("QueryPoolRequest", logger.BuildLogInput().WithData("response", result1))

	var result2 astroport.QueryPairResponse
	err = suite.client.CallWasm(suite.ctx, suite.poolAddress, astroport.QueryPairRequest{}, &result2)
	assert.NoError(suite.T(), err)
	logger.Info("QueryPairRequest", logger.BuildLogInput().WithData("response", result2))

	var result3 astroport.QuerySimulationResponse
	err = suite.client.CallWasm(suite.ctx, suite.poolAddress, astroport.QuerySimulationRequest{
		OfferAsset: &astroport.AssetInfo{
			Info:   astroport.NewAsset(astroport.AssetTypeToken, suite.tokenAddress),
			Amount: "100000",
		},
	}, &result3)
	assert.NoError(suite.T(), err)
	logger.Info("QuerySimulationRequest", logger.BuildLogInput().WithData("response", result3))

	var result4 astroport.QueryReverseSimulationResponse
	err = suite.client.CallWasm(suite.ctx, suite.poolAddress, astroport.QueryReverseSimulationRequest{
		AskAsset: &astroport.AssetInfo{
			Info:   astroport.NewAsset(astroport.AssetTypeToken, suite.tokenAddress),
			Amount: "100000",
		},
	}, &result4)
	assert.NoError(suite.T(), err)
	logger.Info("QueryReverseSimulationRequest", logger.BuildLogInput().WithData("response", result4))

	var result5 astroport.QueryCumulativePricesResponse
	err = suite.client.CallWasm(suite.ctx, suite.poolAddress, astroport.QueryCumulativePricesRequest{}, &result5)
	assert.NoError(suite.T(), err)
	logger.Info("QueryCumulativePricesResponse", logger.BuildLogInput().WithData("response", result5))
}

func (suite *AstroPortTestSuite) TestQueryGenerator() {
	var result1 astroport.QueryPoolLengthResponse
	err := suite.client.CallWasm(suite.ctx, suite.generatorAddress, astroport.QueryPoolLengthRequest{}, &result1)
	assert.NoError(suite.T(), err)
	logger.Info("QueryPoolLengthRequest", logger.BuildLogInput().WithData("response", result1))

	var result2 astroport.QueryPendingTokenResponse
	err = suite.client.CallWasm(suite.ctx, suite.generatorAddress, astroport.QueryPendingTokenRequest{
		LpToken: suite.lpAddress,
		User:    suite.SignerAddress(),
	}, &result2)
	assert.NoError(suite.T(), err)
	logger.Info("QueryPendingTokenRequest", logger.BuildLogInput().WithData("response", result2))

	var result3 astroport.QueryRewardInfoResponse
	err = suite.client.CallWasm(suite.ctx, suite.generatorAddress, astroport.QueryRewardInfoRequest{
		LpToken: suite.lpAddress,
	}, &result3)
	assert.NoError(suite.T(), err)
	logger.Info("QueryRewardInfoRequest", logger.BuildLogInput().WithData("response", result3))

	var result4 astroport.QueryPoolInfoResponse
	err = suite.client.CallWasm(suite.ctx, suite.generatorAddress, astroport.QueryPoolInfoRequest{
		LpToken: suite.lpAddress,
	}, &result4)
	assert.NoError(suite.T(), err)
	logger.Info("QueryPoolInfoRequest", logger.BuildLogInput().WithData("response", result4))

	var result5 astroport.QueryGeneratorConfigResponse
	err = suite.client.CallWasm(suite.ctx, suite.generatorAddress, astroport.QueryConfig{}, &result5)
	assert.NoError(suite.T(), err)
	logger.Info("QueryConfig", logger.BuildLogInput().WithData("response", result5))
}

func (suite *AstroPortTestSuite) TestQueryStaking() {
	var result1 astroport.QueryStakingConfigResponse
	err := suite.client.CallWasm(suite.ctx, suite.stakingAddress, astroport.QueryConfig{}, &result1)
	assert.NoError(suite.T(), err)
	logger.Info("QueryStakingConfigResponse", logger.BuildLogInput().WithData("response", result1))

	var result2 astroport.QueryTotalShareResponse
	err = suite.client.CallWasm(suite.ctx, suite.stakingAddress, astroport.QueryTotalShareRequest{}, &result2)
	assert.NoError(suite.T(), err)
	logger.Info("QueryTotalShareResponse", logger.BuildLogInput().WithData("response", result2))

	var result3 astroport.QueryTotalDepositResponse
	err = suite.client.CallWasm(suite.ctx, suite.stakingAddress, astroport.QueryTotalDepositRequest{}, &result3)
	assert.NoError(suite.T(), err)
	logger.Info("QueryTotalDepositResponse", logger.BuildLogInput().WithData("response", result3))
}

func (suite *AstroPortTestSuite) TestQueryMaker() {
	var result1 interface{}
	err := suite.client.CallWasm(suite.ctx, suite.makerAddress, astroport.QueryConfig{}, &result1)
	assert.NoError(suite.T(), err)
	logger.Info("QueryMakerConfigResponse", logger.BuildLogInput().WithData("response", result1))
}

func (suite *AstroPortTestSuite) TestExecuteRouter() {
	// Uniswap: Swap ETH for tokens
	resp, err := suite.client.SendTransaction(suite.ctx, suite.signer,
		cosmostypes.WithGasLimit(3_000_000),
		cosmostypes.WithFeeAmount(cosmossdk.NewInt64Coin("cony", 5000)),
		cosmostypes.WithMsgs(
			cosmostypes.NewMsgExecuteContract(
				suite.SignerAddress(),
				suite.routerAddress,
				astroport.MsgExecuteSwapOperations{
					SwapOperations: []*astroport.SwapOperation{
						astroport.NewSwapOperation(
							astroport.NewAsset(astroport.AssetTypeCoin, "cony"),
							astroport.NewAsset(astroport.AssetTypeToken, suite.tokenAddress),
						),
					},
					To:             suite.SignerAddress(),
					MaxSpread:      util.ToString(0.15),
					MinimumReceive: util.ToString(0),
				},
				cosmossdk.NewCoins(cosmossdk.NewInt64Coin("cony", 1e6)),
			),
		),
	)
	assert.NoError(suite.T(), err)
	if err == nil {
		logs := PrintWasmLogs(suite.T(), suite.client, resp.Hash)
		eventSwap := astroport.EventSwap{}
		logs.GetLogByEvent("swap").Unmarshal(&eventSwap)
		fmt.Println(eventSwap)
	}

	// Uniswap: Swap tokens for tokens
	resp, err = suite.client.SendTransaction(suite.ctx, suite.signer,
		// swap
		cosmostypes.WithGasLimit(3_000_000),
		cosmostypes.WithFeeAmount(cosmossdk.NewInt64Coin("cony", 5000)),
		cosmostypes.WithMsgs(
			cosmostypes.NewMsgExecuteContract(
				suite.SignerAddress(),
				suite.tokenAddress,
				cw20.MsgSend{
					Contract: suite.routerAddress,
					Amount:   util.ToString(100000),
					Msg: astroport.MsgExecuteSwapOperations{
						SwapOperations: []*astroport.SwapOperation{
							astroport.NewSwapOperation(
								astroport.NewAsset(astroport.AssetTypeToken, suite.tokenAddress),
								astroport.NewAsset(astroport.AssetTypeCoin, "cony"),
							),
							astroport.NewSwapOperation(
								astroport.NewAsset(astroport.AssetTypeCoin, "cony"),
								astroport.NewAsset(astroport.AssetTypeToken, suite.tokenAddress),
							),
							astroport.NewSwapOperation(
								astroport.NewAsset(astroport.AssetTypeToken, suite.tokenAddress),
								astroport.NewAsset(astroport.AssetTypeCoin, "cony"),
							),
							astroport.NewSwapOperation(
								astroport.NewAsset(astroport.AssetTypeCoin, "cony"),
								astroport.NewAsset(astroport.AssetTypeToken, suite.tokenAddress),
							),
						},
						To:             suite.SignerAddress(),
						MaxSpread:      util.ToString(0.15),
						MinimumReceive: util.ToString(0),
					},
				},
				cosmossdk.NewCoins(),
			),
		),
	)

	assert.NoError(suite.T(), err)
	if err == nil {
		logs := PrintWasmLogs(suite.T(), suite.client, resp.Hash)
		eventSwap := astroport.EventSwap{}
		logs.GetLogByEvent("swap").Unmarshal(&eventSwap)
		fmt.Println(eventSwap)
	}
}

func (suite *AstroPortTestSuite) TestExecutePool() {
	var baseAmount = "100000"
	// Simulate
	var simulateResponse astroport.QuerySimulationResponse
	err := suite.client.CallWasm(suite.ctx, suite.poolAddress, astroport.QuerySimulationRequest{
		OfferAsset: &astroport.AssetInfo{
			Info:   astroport.NewAsset(astroport.AssetTypeCoin, "cony"),
			Amount: baseAmount,
		},
	}, &simulateResponse)

	slippage := util.ToString(0.49)
	resp, err := suite.client.SendTransaction(suite.ctx, suite.signer,
		// approve simulated amount
		cosmostypes.WithMsgs(
			cosmostypes.NewMsgExecuteContract(
				suite.SignerAddress(),
				suite.tokenAddress,
				cw20.MsgIncreaseAllowance{
					Spender: suite.poolAddress,
					Amount:  simulateResponse.ReturnAmount,
					Expires: cwutil.NewExpiration(),
				},
				cosmossdk.NewCoins(),
			),
			// add liquidity base value and simulated amount
			cosmostypes.NewMsgExecuteContract(
				suite.SignerAddress(),
				suite.poolAddress,
				astroport.MsgProvideLiquidity{
					Assets: []astroport.AssetInfo{
						{
							Info:   astroport.NewAsset(astroport.AssetTypeToken, suite.tokenAddress),
							Amount: simulateResponse.ReturnAmount,
						},
						{
							Info:   astroport.NewAsset(astroport.AssetTypeCoin, "cony"),
							Amount: baseAmount,
						},
					},
					SlippageTolerance: &slippage,
					AutoStake:         false,
					Receiver:          suite.SignerAddress(),
				},
				cosmossdk.NewCoins(cosmossdk.NewInt64Coin("cony", 100000)),
			),
		),
	)
	assert.NoError(suite.T(), err)
	if err == nil {
		PrintWasmLogs(suite.T(), suite.client, resp.Hash)
	}

	resp, err = suite.client.SendTransaction(suite.ctx, suite.signer,
		cosmostypes.WithMsgs(
			cosmostypes.NewMsgExecuteContract(
				suite.SignerAddress(),
				suite.lpAddress,
				cw20.MsgSend{
					Contract: suite.poolAddress,
					Amount:   util.ToString(1000),
					Msg:      astroport.MsgWithdrawLiquidity{},
				},
				cosmossdk.NewCoins(),
			)),
	)
	assert.NoError(suite.T(), err)
	if err == nil {
		PrintWasmLogs(suite.T(), suite.client, resp.Hash)
	}
}

func (suite *AstroPortTestSuite) TestExecuteStaking() {
	var baseAmount = "100000"
	resp, err := suite.client.SendTransaction(suite.ctx, suite.signer,
		cosmostypes.WithMsgs(
			cosmostypes.NewMsgExecuteContract(
				suite.SignerAddress(),
				suite.tokenAddress,
				cw20.MsgSend{
					Contract: suite.stakingAddress,
					Amount:   baseAmount,
					Msg:      astroport.MsgEnter{},
				},
				cosmossdk.NewCoins(),
			)),
	)
	assert.NoError(suite.T(), err)
	if err == nil {
		PrintWasmLogs(suite.T(), suite.client, resp.Hash)
	}

	resp, err = suite.client.SendTransaction(suite.ctx, suite.signer,
		cosmostypes.WithMsgs(
			cosmostypes.NewMsgExecuteContract(
				suite.SignerAddress(),
				suite.xtokenAddress,
				cw20.MsgSend{
					Contract: suite.stakingAddress,
					Amount:   baseAmount,
					Msg:      astroport.MsgLeave{},
				},
				cosmossdk.NewCoins(),
			)),
	)
	assert.NoError(suite.T(), err)
	if err == nil {
		PrintWasmLogs(suite.T(), suite.client, resp.Hash)
	}
}

func (suite *AstroPortTestSuite) TestQueryVesting() {
	var result1 interface{}
	err := suite.client.CallWasm(suite.ctx, suite.vestingAddress, astroport.QueryConfig{}, &result1)
	assert.NoError(suite.T(), err)
	logger.Info("QueryConfigRequest", logger.BuildLogInput().WithData("response", result1))

	var result2 astroport.QueryVestingAccountResponse
	err = suite.client.CallWasm(suite.ctx, suite.vestingAddress, astroport.QueryVestingAccountRequest{
		Address: "link1887ty9wmywzvspl4d9qkvjhzar0ezhclvq00hj",
	}, &result2)
	assert.NoError(suite.T(), err)
	logger.Info("QueryConfigRequest", logger.BuildLogInput().WithData("response", result2))

	var result3 astroport.QueryAvailableAmountResponse
	err = suite.client.CallWasm(suite.ctx, suite.vestingAddress, astroport.QueryAvailableAmountRequest{
		Address: "link1887ty9wmywzvspl4d9qkvjhzar0ezhclvq00hj",
	}, &result3)
	assert.NoError(suite.T(), err)
	logger.Info("QueryConfigRequest", logger.BuildLogInput().WithData("response", result3))

	var result4 astroport.QueryVestingAccountsResponse
	err = suite.client.CallWasm(suite.ctx, suite.vestingAddress, astroport.QueryVestingAccountsRequest{
		StartAfter: "",
		Limit:      20,
	}, &result4)
	assert.NoError(suite.T(), err)
	for _, vestingAccount := range result4.VestingAccounts {
		var availableAmountResponse astroport.QueryAvailableAmountResponse
		err = suite.client.CallWasm(suite.ctx, suite.vestingAddress, astroport.QueryAvailableAmountRequest{
			Address: vestingAccount.Address,
		}, &availableAmountResponse)

		for _, schedule := range vestingAccount.VestingInfo.Schedules {
			logger.Info("QueryVestingAccountsRequest",
				logger.BuildLogInput().WithData(
					"start", schedule.StartPoint.GetTime(), "end", schedule.EndPoint.GetTime(),
					"startAmount", schedule.StartPoint.Amount, "endAmount", schedule.EndPoint.Amount,
				))
		}

		logger.Info("QueryVestingAccountsRequest",
			logger.BuildLogInput().WithData(
				"address", vestingAccount.Address,
				"released", util.ToDecimal(vestingAccount.VestingInfo.ReleasedAmount, 6),
				"available", util.ToDecimal(availableAmountResponse, 6)),
		)
	}
}
