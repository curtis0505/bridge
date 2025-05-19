package cosmos

import (
	"context"
	conf2 "github.com/curtis0505/bridge/libs/client/chain/conf"
	cosmostypes "github.com/curtis0505/bridge/libs/client/chain/cosmos/types"
	clienttypes "github.com/curtis0505/bridge/libs/client/chain/types"
	"github.com/curtis0505/bridge/libs/testutil"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"math/big"
	"testing"
)

const (
	REST = "rest"
	GRPC = "grpc"
)

var (

	// chain/env/protocol
	envConfig = map[string]map[testutil.TestENV]struct {
		restConfig       conf2.ClientConfig
		grpcConfig       conf2.ClientConfig
		Address          string
		ValidatorAddress string
	}{
		types.ChainATOM: {
			testutil.LIVE: {
				restConfig: conf2.ClientConfig{
					Chain:     types.ChainATOM,
					ChainName: types.ChainATOM,
					Url:       "https://cosmos-api.neopin.io",
				},
				grpcConfig: conf2.ClientConfig{},
			},
		},

		types.ChainFNSA: {
			testutil.DEV: {
				Address:          "link1887ty9wmywzvspl4d9qkvjhzar0ezhclgh7lwh",
				ValidatorAddress: "linkvaloper1q2tr2qtq8wn2ln2c5n82prc4dswa5cna7xuv8j",
				restConfig: conf2.ClientConfig{
					Chain:     types.ChainFNSA,
					ChainName: types.ChainFNSA,
					Url:       "https://finschia-rest.dev.neopin.io",
				},
				grpcConfig: conf2.ClientConfig{
					Chain:     types.ChainFNSA,
					ChainName: types.ChainFNSA,
					Url:       "node.dev.neopin.io:32300",
				},
			},
			testutil.DQ: {
				Address:          "tlink1887ty9wmywzvspl4d9qkvjhzar0ezhclgh7lwh",
				ValidatorAddress: "tlinkvaloper1q2tr2qtq8wn2ln2c5n82prc4dswa5cna7xuv8j",
				restConfig: conf2.ClientConfig{
					Chain:     types.ChainTFNSA,
					ChainName: types.ChainTFNSA,
					Url:       "https://finschia-rest.hapi.neoply.io",
				},
				grpcConfig: conf2.ClientConfig{
					Chain:     types.ChainTFNSA,
					ChainName: types.ChainTFNSA,
					Url:       "hapi.neoply.io:19090",
				},
			},
			testutil.LIVE: {
				Address:          "link1nddwnkc47p9v9apruhp9xtq4rquf2y49xcw67r",
				ValidatorAddress: "linkvaloper1nddwnkc47p9v9apruhp9xtq4rquf2y495vv8ss",
				restConfig: conf2.ClientConfig{
					Chain:     types.ChainFNSA,
					ChainName: types.ChainFNSA,
					Url:       "https://dsvt-finschia.line-apps.com",
				},
				grpcConfig: conf2.ClientConfig{
					Chain:     types.ChainFNSA,
					ChainName: types.ChainFNSA,
					Url:       "finschia-grpc.neopin.io:39090",
				},
			},
		},
	}
)

func GetTestAccount(t *testing.T, privateKey string) *types.Account {
	account, err := types.NewAccountFromPK(privateKey)
	assert.NoError(t, err)
	return account
}

func TestRunClientSuite(t *testing.T) {
	suite.Run(t, new(TestClientSuite))
}

type TestClientSuite struct {
	suite.Suite

	chain            string
	env              testutil.TestENV
	address          string
	validatorAddress string
	ctx              context.Context
	grpcClient       clienttypes.CosmosClient
	restClient       clienttypes.CosmosClient
}

func (suite *TestClientSuite) SetupSuite() {
	suite.chain = types.ChainFNSA
	suite.env = testutil.DEV

	var err error
	suite.grpcClient, err = NewClient(envConfig[suite.chain][suite.env].grpcConfig)
	assert.NoError(suite.T(), err)

	suite.restClient, err = NewClient(envConfig[suite.chain][suite.env].restConfig)
	assert.NoError(suite.T(), err)

	suite.address = envConfig[suite.chain][suite.env].Address
	suite.validatorAddress = envConfig[suite.chain][suite.env].ValidatorAddress
	suite.ctx = context.Background()

	assert.Equal(suite.T(), suite.grpcClient.ChainId(), suite.restClient.ChainId())
}

func (suite *TestClientSuite) TestValidatorDelegations() {
	total1, total2 := big.NewInt(0), big.NewInt(0)

	delegations1, err := suite.restClient.GetValidatorDelegations(suite.ctx, suite.validatorAddress)
	assert.NoError(suite.T(), err)

	for _, delegation := range delegations1 {
		total1.Add(total1, delegation.Balance.Amount.BigInt())
	}
	suite.T().Log(util.ToDecimal(total1, 6))

	delegations2, err := suite.grpcClient.GetValidatorDelegations(suite.ctx, suite.validatorAddress)
	assert.NoError(suite.T(), err)

	for _, delegation := range delegations2 {
		total2.Add(total2, delegation.Balance.Amount.BigInt())
	}

	suite.T().Log(util.ToDecimal(total2, 6))

	assert.Equal(suite.T(), total1, total2)
	assert.Equal(suite.T(), len(delegations1), len(delegations2))
}

func (suite *TestClientSuite) TestBlockNumber() {
	num1, err := suite.grpcClient.BlockNumber(suite.ctx)
	assert.NoError(suite.T(), err)

	num2, err := suite.restClient.BlockNumber(suite.ctx)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), num1, num2)
}

func (suite *TestClientSuite) TestGetBlockByNumber() {
	num1, err := suite.grpcClient.GetBlockByNumber(suite.ctx, big.NewInt(0))
	assert.NoError(suite.T(), err)

	num2, err := suite.restClient.GetBlockByNumber(suite.ctx, big.NewInt(0))
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), num1.GetHeader().ChainID, num2.GetHeader().ChainID)
}

func (suite *TestClientSuite) TestValidatorApr() {
	num1, err := suite.grpcClient.GetValidatorApr(suite.ctx, suite.validatorAddress)
	assert.NoError(suite.T(), err)

	num2, err := suite.restClient.GetValidatorApr(suite.ctx, suite.validatorAddress)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), num1, num2)
}

func (suite *TestClientSuite) TestBalanceAt() {
	num1, err := suite.grpcClient.BalanceAt(suite.ctx, suite.address, nil)
	assert.NoError(suite.T(), err)

	num2, err := suite.restClient.BalanceAt(suite.ctx, suite.address, nil)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), num1, num2)
}

func (suite *TestClientSuite) TestPendingNonceAt() {
	num1, err := suite.grpcClient.PendingNonceAt(suite.ctx, suite.address)
	assert.NoError(suite.T(), err)

	num2, err := suite.restClient.PendingNonceAt(suite.ctx, suite.address)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), num1, num2)
}

func (suite *TestClientSuite) TestStaking() {
	resp1, err := suite.grpcClient.GetStaking(suite.ctx, suite.address)
	assert.NoError(suite.T(), err)

	resp2, err := suite.restClient.GetStaking(suite.ctx, suite.address)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), resp1, resp2)
}

func (suite *TestClientSuite) TestGetTxs() {
	resp1, err := suite.grpcClient.GetTxs(suite.ctx, cosmostypes.NewQueryTxEvent("wasm", "_contract_address", "link16tfcsjcg4trlnj2t4cpdq20q4d7nzusew3ev50scz30mmfzdx8ssw3wuhh"))
	assert.NoError(suite.T(), err)

	resp2, err := suite.restClient.GetTxs(suite.ctx, cosmostypes.NewQueryTxEvent("wasm", "_contract_address", "link16tfcsjcg4trlnj2t4cpdq20q4d7nzusew3ev50scz30mmfzdx8ssw3wuhh"))
	assert.NoError(suite.T(), err)

	for _, txResponse := range resp1.GetTxResponses() {
		logs := cosmostypes.ParseLogs(txResponse)
		for _, log := range logs.GetLogsByEvent("swap") {
			suite.T().Log("GRPC", log)
		}
	}

	for _, txResponse := range resp2.GetTxResponses() {
		logs := cosmostypes.ParseLogs(txResponse)
		for _, log := range logs.GetLogsByEvent("swap") {
			suite.T().Log("REST", log)
		}
	}
}
