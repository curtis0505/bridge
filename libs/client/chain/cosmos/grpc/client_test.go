package cosmos

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/go-bip39"
	conf2 "github.com/curtis0505/bridge/libs/client/chain/conf"
	cosmostypes "github.com/curtis0505/bridge/libs/client/chain/cosmos/types"
	clienttypes "github.com/curtis0505/bridge/libs/client/chain/types"
	cosmoscommon "github.com/curtis0505/bridge/libs/common/cosmos"
	"github.com/curtis0505/bridge/libs/testutil"
	"github.com/curtis0505/bridge/libs/types"
	tokentypes "github.com/curtis0505/bridge/libs/types/token"
	"github.com/curtis0505/grpc-idl/finschia/collection"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"math"
	"testing"
	"time"
)

const (
	CoinTypeFinschia = 428
	CoinTypeCosmos   = 118
)

var (
	DevFinschiaConfig = conf2.ClientConfig{
		Chain:     types.ChainFNSA,
		ChainName: types.ChainFNSA,
		Url:       "node.dev.neopin.io:32300",
	}

	EbonyConfig = conf2.ClientConfig{
		Chain:     types.ChainTFNSA,
		ChainName: types.ChainTFNSA,
		Url:       "hapi.neoply.io:19090",
	}

	DevCosmosConfig = conf2.ClientConfig{
		Chain:     types.ChainATOM,
		ChainName: types.ChainATOM,
		Url:       "node.dev.neopin.io:49090",
	}

	DQCosmosConfig = conf2.ClientConfig{
		Chain:     types.ChainATOM,
		ChainName: types.ChainATOM,
		Url:       "hapi.neoply.io:49090",
	}
)

func GetTestClient(t *testing.T, env testutil.TestENV, chain string) clienttypes.CosmosClient {
	switch chain {
	case types.ChainFNSA:
		c, err := NewClient(DevFinschiaConfig)
		assert.NoError(t, err)
		return c
	case types.ChainTFNSA:
		c, err := NewClient(EbonyConfig)
		assert.NoError(t, err)
		return c
	case types.ChainATOM:
		if env == testutil.DEV {
			c, err := NewClient(DevCosmosConfig)
			assert.NoError(t, err)
			return c
		} else {
			c, err := NewClient(DQCosmosConfig)
			assert.NoError(t, err)
			return c
		}
	}
	return nil
}

func GetTestAccount(t *testing.T, privateKey string) *types.Account {
	account, err := types.NewAccountFromPK(privateKey)
	assert.NoError(t, err)
	return account
}

func GetTestPrivKey(t *testing.T, mnemonic string, coinType int, index int) *secp256k1.PrivKey {
	seed := bip39.NewSeed(mnemonic, "")
	master, ch := hd.ComputeMastersFromSeed(seed)
	priv, err := hd.DerivePrivateKeyForPath(master, ch, fmt.Sprintf("m/44'/%d'/0'/0/%d", coinType, index))

	privKey := &secp256k1.PrivKey{Key: priv}
	assert.NoError(t, err)
	return privKey
}

type TxTestSuite struct {
	suite.Suite
	ctx context.Context
	//redisDB *model.RedisDB
	//redis   bool
	env testutil.TestENV

	client           clienttypes.CosmosClient
	chain            string
	validatorAddress string
	toAddress        string

	testCoins    cosmossdk.Coins
	testMaxCoins cosmossdk.Coins

	signer *types.Account

	// sent tx hash
	txHash string
}

func TestRunTxTestSuite(t *testing.T) {
	suite.Run(t, new(TxTestSuite))
}

const Chain = types.ChainATOM

func (suite *TxTestSuite) SetupSuite() {
	suite.chain = Chain
	suite.env = testutil.DEV
	suite.client = GetTestClient(suite.T(), suite.env, suite.chain)
	suite.testCoins = cosmossdk.NewCoins(cosmossdk.NewInt64Coin(tokentypes.DenomByChain(suite.chain), 1e6))
	suite.testMaxCoins = cosmossdk.NewCoins(cosmossdk.NewInt64Coin(tokentypes.DenomByChain(suite.chain), math.MaxInt64))
	suite.ctx = context.Background()

	switch suite.chain {
	case types.ChainFNSA:
		suite.signer = GetTestAccount(suite.T(), "83c13018679900c69be911758784782f693215f754c995fc17e304c9fbf90afb")
		suite.validatorAddress = "linkvaloper1c20t099jq2j3a9pg3hlld8um8jfa3g6yjyeztk"
		suite.toAddress = "link1887ty9wmywzvspl4d9qkvjhzar0ezhclvq00hj"
	case types.ChainTFNSA:
		suite.signer = GetTestAccount(suite.T(), "8932705e9a280ff7f36401ba26479f56d4e2400a2a93563dd885fe7f99bbb144")
		suite.validatorAddress = "tlinkvaloper1q2tr2qtq8wn2ln2c5n82prc4dswa5cna7xuv8j"
		suite.toAddress = ""
	case types.ChainATOM:
		suite.signer = GetTestAccount(suite.T(), "83c13018679900c69be911758784782f693215f754c995fc17e304c9fbf90afb")
		suite.validatorAddress = "cosmosvaloper1srt0vu2qcjkq4dvcrgssvg6ug53j53kfyyjdlt"
		suite.toAddress = "cosmos1887ty9wmywzvspl4d9qkvjhzar0ezhclew0dv6"
	}

	suite.T().Log("chain:", suite.chain)
	suite.T().Log("signer:", suite.SignerAddress())
}

func (suite *TxTestSuite) SetupTest() {

}

func (suite *TxTestSuite) BeforeTest(name, test string) {
	suite.T().Log("run test:", test)
}

//func (suite *TxTestSuite) AfterTest(name, test string) {
//	time.Sleep(5 * time.Second)
//	suite.T().Log("end test:", test, "txhash", suite.txHash)
//
//	if suite.redis {
//		tx, err := suite.client.GetTransaction(suite.ctx, suite.txHash)
//		if err != nil {
//			return
//		}
//		assert.NoError(suite.T(), err)
//		castedTx := tx.Inner().(*cosmostxtypes.Tx)
//
//		protoTxBytes, err := castedTx.Marshal()
//		assert.NoError(suite.T(), err)
//
//		receipt, err := suite.client.GetTransactionReceipt(suite.ctx, suite.txHash)
//		assert.NoError(suite.T(), err)
//
//		responseBytes, err := receipt.MarshalBinary()
//		assert.NoError(suite.T(), err)
//
//		redisData := &dto.CosmosRedisData{
//			Chain:         suite.chain,
//			TxHash:        suite.txHash,
//			RawTx:         protoTxBytes,
//			RawTxResponse: responseBytes,
//		}
//
//		suite.redisDB.LPushCosmosContract(redisData)
//	}
//}

func (suite *TxTestSuite) SignerAddress() string {
	return cosmoscommon.FromPublicKeyUnSafe(suite.chain, suite.signer.Secp256k1().PubKey().Bytes()).String()
}

func (suite *TxTestSuite) TestMsgSend() {
	msg := &banktypes.MsgSend{
		FromAddress: suite.SignerAddress(),
		ToAddress:   suite.toAddress,
		Amount:      suite.testCoins,
	}

	resp, err := suite.client.SendTransaction(suite.ctx, suite.signer, cosmostypes.WithMsgs(msg))
	assert.NoError(suite.T(), err)
	suite.txHash = resp.Hash
}

func (suite *TxTestSuite) TestFailMsgSend() {
	msg := &banktypes.MsgSend{
		FromAddress: suite.SignerAddress(),
		ToAddress:   suite.toAddress,
		Amount:      suite.testMaxCoins,
	}

	resp, err := suite.client.SendTransaction(suite.ctx, suite.signer, cosmostypes.WithMsgs(msg))
	assert.NoError(suite.T(), err)
	suite.txHash = resp.Hash
}

func (suite *TxTestSuite) TestMsgDelegate() {
	ok, coin := suite.testCoins.Find(tokentypes.DenomByChain(suite.chain))
	assert.Equal(suite.T(), ok, true)

	msg := &stakingtypes.MsgDelegate{
		DelegatorAddress: suite.SignerAddress(),
		ValidatorAddress: suite.validatorAddress,
		Amount:           coin,
	}
	resp, err := suite.client.SendTransaction(suite.ctx, suite.signer, cosmostypes.WithMsgs(msg))
	assert.NoError(suite.T(), err)
	suite.txHash = resp.Hash
}

func (suite *TxTestSuite) TestMsgSubmitProposal() {
	msg := &govtypesv1.MsgSubmitProposal{}
	resp, err := suite.client.SendTransaction(suite.ctx, suite.signer, cosmostypes.WithMsgs(msg))
	assert.NoError(suite.T(), err)
	suite.txHash = resp.Hash
}

func (suite *TxTestSuite) TestFailMsgDelegate() {
	ok, coin := suite.testMaxCoins.Find(tokentypes.DenomByChain(suite.chain))
	assert.Equal(suite.T(), ok, true)

	msg := &stakingtypes.MsgDelegate{
		DelegatorAddress: suite.SignerAddress(),
		ValidatorAddress: suite.validatorAddress,
		Amount:           coin,
	}
	resp, err := suite.client.SendTransaction(suite.ctx, suite.signer, cosmostypes.WithMsgs(msg))
	assert.NoError(suite.T(), err)
	suite.txHash = resp.Hash
}

func (suite *TxTestSuite) TestMsgUndelegate() {
	ok, coin := suite.testCoins.Find(tokentypes.DenomByChain(suite.chain))
	assert.Equal(suite.T(), ok, true)

	msg := &stakingtypes.MsgUndelegate{
		DelegatorAddress: suite.SignerAddress(),
		ValidatorAddress: suite.validatorAddress,
		Amount:           coin,
	}

	resp, err := suite.client.SendTransaction(suite.ctx, suite.signer, cosmostypes.WithMsgs(msg))
	assert.NoError(suite.T(), err)
	suite.txHash = resp.Hash
}

func (suite *TxTestSuite) TestFailMsgUndelegate() {
	ok, coin := suite.testMaxCoins.Find(tokentypes.DenomByChain(suite.chain))
	assert.Equal(suite.T(), ok, true)

	msg := &stakingtypes.MsgUndelegate{
		DelegatorAddress: suite.SignerAddress(),
		ValidatorAddress: suite.validatorAddress,
		Amount:           coin,
	}
	resp, err := suite.client.SendTransaction(suite.ctx, suite.signer, cosmostypes.WithMsgs(msg))
	assert.NoError(suite.T(), err)
	suite.txHash = resp.Hash
}

func (suite *TxTestSuite) TestMsgWithdrawDelegatorReward() {
	msg := &distributiontypes.MsgWithdrawDelegatorReward{
		DelegatorAddress: suite.SignerAddress(),
		ValidatorAddress: suite.validatorAddress,
	}

	resp, err := suite.client.SendTransaction(suite.ctx, suite.signer, cosmostypes.WithMsgs(msg))
	assert.NoError(suite.T(), err)
	suite.txHash = resp.Hash
}

func (suite *TxTestSuite) TestFailMsgWithdrawDelegatorReward() {
	msg := &distributiontypes.MsgWithdrawDelegatorReward{
		DelegatorAddress: suite.SignerAddress(),
		ValidatorAddress: "",
	}

	resp, err := suite.client.SendTransaction(suite.ctx, suite.signer, cosmostypes.WithMsgs(msg))
	assert.NoError(suite.T(), err)
	suite.txHash = resp.Hash
}

func (suite *TxTestSuite) TestCollectionMsgMsgCreateContract() {
	if suite.chain == types.ChainFNSA || suite.chain == types.ChainTFNSA {
		msg := &collection.MsgCreateContract{
			Owner: suite.SignerAddress(),
		}

		resp, err := suite.client.SendTransaction(suite.ctx, suite.signer, cosmostypes.WithMsgs(msg))
		assert.NoError(suite.T(), err)
		suite.txHash = resp.Hash
	}
}

type StressTestSuite struct {
	suite.Suite
	ctx context.Context

	mnemonic string
	index    int
	msgCount int
	coinType int

	env testutil.TestENV

	client clienttypes.CosmosClient
	chain  string

	testCoins    cosmossdk.Coins
	testMaxCoins cosmossdk.Coins

	privKeyList []*secp256k1.PrivKey

	// sent tx hash
	txHash string
}

func TestRunStressTestSuite(t *testing.T) {
	suite.Run(t, new(StressTestSuite))
}

func (suite *StressTestSuite) SetupSuite() {
	suite.chain = Chain
	suite.env = testutil.DEV
	suite.client = GetTestClient(suite.T(), suite.env, suite.chain)
	suite.testCoins = cosmossdk.NewCoins(cosmossdk.NewInt64Coin(tokentypes.DenomByChain(suite.chain), 1e6))
	suite.testMaxCoins = cosmossdk.NewCoins(cosmossdk.NewInt64Coin(tokentypes.DenomByChain(suite.chain), math.MaxInt64))
	suite.ctx = context.Background()
	suite.mnemonic = "lock load dad patient economy receive ridge language confirm file spoil chunk oak humble need brick swallow play song soccer sudden mansion shine lecture"
	suite.index = 100
	suite.msgCount = 2

	switch suite.chain {
	case types.ChainFNSA:
		suite.coinType = CoinTypeFinschia
		var err error
		suite.client, err = NewClient(DevFinschiaConfig)
		assert.NoError(suite.T(), err)
	}

	for i := 0; i < suite.index; i++ {
		suite.privKeyList = append(suite.privKeyList, GetTestPrivKey(suite.T(), suite.mnemonic, suite.coinType, i))
	}
}

func (suite *StressTestSuite) AfterTest(name, test string) {
	time.Sleep(5 * time.Second)
}

func (suite *StressTestSuite) TestMsgSendSelf() {
	for _, privKey := range suite.privKeyList {
		go func(privKey *secp256k1.PrivKey) {
			var msgs []cosmossdk.Msg
			msg := &banktypes.MsgSend{
				FromAddress: cosmoscommon.FromPublicKeyUnSafe(suite.chain, privKey.PubKey().Bytes()).String(),
				ToAddress:   cosmoscommon.FromPublicKeyUnSafe(suite.chain, privKey.PubKey().Bytes()).String(),
				Amount:      suite.testCoins,
			}

			for i := 0; i < suite.msgCount; i++ {
				msgs = append(msgs, msg)
			}

			resp, err := suite.client.SendTransactionPrivKey(suite.ctx, privKey, cosmostypes.WithMsgs(msgs...))
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), resp.Result, clienttypes.SendTxResultType_Success)
		}(privKey)
	}
}
