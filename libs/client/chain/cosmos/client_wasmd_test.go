package cosmos

import (
	"context"
	"encoding/json"
	"fmt"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	cosmostypes "github.com/curtis0505/bridge/libs/client/chain/cosmos/types"
	clienttypes "github.com/curtis0505/bridge/libs/client/chain/types"
	cosmoscommon "github.com/curtis0505/bridge/libs/common/cosmos"
	"github.com/curtis0505/bridge/libs/logger/v2"
	"github.com/curtis0505/bridge/libs/testutil"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/types/cosmos/cw20"
	"github.com/curtis0505/bridge/libs/types/cosmos/cwutil"
	"github.com/curtis0505/bridge/libs/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"math/big"
	"testing"
	"time"
)

type CW20TestSuite struct {
	suite.Suite

	env    testutil.TestENV
	ctx    context.Context
	client clienttypes.CosmosClient

	chain  string
	signer *types.Account

	tokenAddress string
	poolAddress  string
}

func TestCW20TestSuite(t *testing.T) {
	logger.InitLog(logger.Config{UseTerminal: true, VerbosityTerminal: 5})
	suite.Run(t, new(CW20TestSuite))
}

func (suite *CW20TestSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.env = testutil.DEV
	suite.chain = types.ChainFNSA

	suite.client, _ = NewClient(envConfig[suite.chain][suite.env].restConfig)
	suite.client.SetAccountPrefix()

	suite.signer = GetTestAccount(suite.T(), "58d7ad6ece2c88ce01648aaf0799e27eb760e37f11d30b937f36529bdee926db")
	suite.tokenAddress = "link1dny07jvnkemv5sndvawn0xurxy57p7whn9gpxmfshq85rlctszgshxw99x"
	suite.poolAddress = "link1um0shuef3euaykf4yjt3mlwkr32r3h2cmxff63dpshrypljwd0eqww4vha"

	suite.T().Log(suite.SignerAddress())
}

func (suite *CW20TestSuite) SignerAddress() string {
	return cosmoscommon.FromPublicKeyUnSafe(suite.chain, suite.signer.Secp256k1().PubKey().Bytes()).String()
}

func (suite *CW20TestSuite) TestQueryCW20() {
	var result1 cw20.QueryBalanceResponse
	err := suite.client.CallWasm(suite.ctx, suite.tokenAddress, cw20.QueryBalanceRequest{
		Address: suite.SignerAddress(),
	}, &result1)
	assert.NoError(suite.T(), err)
	logger.Info("QueryBalanceRequest", logger.BuildLogInput().WithData("response", result1))

	var result2 cw20.QueryTokenInfoResponse
	err = suite.client.CallWasm(suite.ctx, suite.tokenAddress, cw20.QueryTokenInfoRequest{}, &result2)
	assert.NoError(suite.T(), err)
	logger.Info("QueryBalanceRequest", logger.BuildLogInput().WithData("response", result2))

	var result3 cw20.QueryAllowanceResponse
	err = suite.client.CallWasm(suite.ctx, suite.tokenAddress, cw20.QueryAllowanceRequest{
		Owner:   suite.SignerAddress(),
		Spender: suite.poolAddress,
	}, &result3)
	assert.NoError(suite.T(), err)
	logger.Info("QueryBalanceRequest", logger.BuildLogInput().WithData("response", result3))

	var result4 cw20.QueryMarketingInfoResponse
	err = suite.client.CallWasm(suite.ctx, suite.tokenAddress, cw20.QueryMarketingInfoRequest{}, &result4)
	assert.NoError(suite.T(), err)
	logger.Info("QueryBalanceRequest", logger.BuildLogInput().WithData("response", result4))
}

func (suite *CW20TestSuite) TestExecuteCW20() {
	blockNumber, err := suite.client.BlockNumber(suite.ctx)
	assert.NoError(suite.T(), err)

	resp, err := suite.client.SendTransaction(suite.ctx, suite.signer,
		cosmostypes.WithMsgs(
			cosmostypes.NewMsgExecuteContract(
				suite.SignerAddress(),
				suite.tokenAddress,
				cw20.MsgTransfer{
					Recipient: suite.SignerAddress(),
					Amount:    util.ToString(big.NewInt(1e6)),
				},
				cosmossdk.NewCoins(),
			)),
	)
	assert.NoError(suite.T(), err)
	if err == nil {
		PrintWasmLogs(suite.T(), suite.client, resp.Hash)
	}
	expire := blockNumber.Int64() + 100

	resp, err = suite.client.SendTransaction(suite.ctx, suite.signer,
		cosmostypes.WithMsgs(
			cosmostypes.NewMsgExecuteContract(
				suite.SignerAddress(),
				suite.tokenAddress,
				cw20.MsgIncreaseAllowance{
					Spender: suite.poolAddress,
					Amount:  util.ToString(big.NewInt(1e6)),
					Expires: cwutil.NewExpirationAtHeight(expire),
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
				suite.tokenAddress,
				cw20.MsgDecreaseAllowance{
					Spender: suite.poolAddress,
					Amount:  util.ToString(big.NewInt(1e6)),
					Expires: cwutil.NewExpirationAtHeight(expire),
				},
				cosmossdk.NewCoins(),
			)),
	)
	assert.NoError(suite.T(), err)
	if err == nil {
		PrintWasmLogs(suite.T(), suite.client, resp.Hash)
	}
}

func PrintWasmLogs(t *testing.T, client clienttypes.CosmosClient, txHash string) cosmostypes.ReceiptLogs {
	t.Log("TxHash", txHash)
	time.Sleep(time.Second * 5)

	receipt, err := client.GetTransactionReceipt(context.Background(), txHash)
	if err != nil {
		t.Log(err)
		return cosmostypes.ReceiptLogs{}
	}
	inner := receipt.Inner().(*cosmossdk.TxResponse)
	if inner.Code != 0 {
		t.Log(inner.RawLog)
		return cosmostypes.ReceiptLogs{}
	}

	logs := cosmostypes.ParseLogs(inner)
	wasmLogs := logs.GetLogsByType(wasmtypes.ModuleName)
	for _, log := range wasmLogs {
		t.Log("\n Event:", log.EventName, "\n Address:", log.ContractAddress, "\n Logs:", log.Logs)
		b, _ := json.Marshal(log.Logs)
		fmt.Println(string(b))
	}

	return wasmLogs
}
