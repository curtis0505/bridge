package cosmos

import (
	"bytes"
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	kmultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/go-bip39"
	cosmostypes "github.com/curtis0505/bridge/libs/client/chain/cosmos/types"
	clienttypes "github.com/curtis0505/bridge/libs/client/chain/types"
	cosmoscommon "github.com/curtis0505/bridge/libs/common/cosmos"
	"github.com/curtis0505/bridge/libs/logger/v2"
	"github.com/curtis0505/bridge/libs/testutil"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/types/cosmos/bridge"
	"github.com/curtis0505/bridge/libs/types/cosmos/cw20"
	"github.com/curtis0505/bridge/libs/types/cosmos/cwutil"
	tokentypes "github.com/curtis0505/bridge/libs/types/token"
	"github.com/curtis0505/bridge/libs/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"math/big"
	"sort"
	"testing"
)

type BridgeTestSuite struct {
	suite.Suite

	env    testutil.TestENV
	ctx    context.Context
	client clienttypes.CosmosClient

	chain   string
	toChain string
	signer  *types.Account
	//validatorSigner  []*types.Account
	multiSigPriv []*secp256k1.PrivKey
	vaultAddress string
	mintAddress  string

	evmTokenAddress  string //vault 기준 child, minter 기준 parent
	fnsaTokenAddress string //vault 기준 parent, minter 기준 child

	evmChainName  string
	fnsaChainName string
}

func TestBridgeSuite(t *testing.T) {
	logger.InitLog(logger.Config{UseTerminal: true, VerbosityTerminal: 5})
	suite.Run(t, new(BridgeTestSuite))
}

const (
	CoinTypeFinschia = 438
	CoinTypeCosmos   = 118
)

var (
	contractConfig = map[string]map[testutil.TestENV]struct {
		vaultAddress     string
		mintAddress      string
		chainName        string
		fnsaTokenAddress string
	}{
		types.ChainFNSA: {
			testutil.DEV: {
				vaultAddress:     "link1sqctf7u55yhu4s00d4f33nj7k2a09tjx4x0ev74u83zvpqkhvfqsvqqnpz",
				mintAddress:      "link1un8cmeumu2fw77xzc36uz88n67zgzf8zlyt66hkp9nayua7dkj8qtwyr4a",
				chainName:        getChainName(types.ChainFNSA, testutil.DEV),
				fnsaTokenAddress: "0x0000000000000000000000000000000000000001",
			},
			testutil.DQ: {
				vaultAddress:     "tlink1cej4htumfduzh38eh8nw3vj00733dyruqjycsaqhvj7qfvxcnrvshenclc",
				mintAddress:      "tlink12kcmqa46dl0xkun9dlxekgtjsnt0zvday2zfypfmc0ntzzw38r5s0h8paw",
				chainName:        getChainName(types.ChainFNSA, testutil.DQ),
				fnsaTokenAddress: "0x0000000000000000000000000000000000000001",
			},
		},
		types.ChainKLAY: {
			testutil.DEV: {
				vaultAddress:     "",
				mintAddress:      "",
				chainName:        types.ChainKLAY,
				fnsaTokenAddress: "0x93de1ec8a5dab7cdf31e013837391cc2dff356b3",
			},
			testutil.DQ: {
				vaultAddress:     "",
				mintAddress:      "",
				chainName:        types.ChainKLAY,
				fnsaTokenAddress: "0xb0c0a345b3bf609e8f5f33d5e0d2a908e03157f0",
			},
		},
	}
)

func getChainName(chain string, env testutil.TestENV) string {
	if chain == types.ChainFNSA {
		if env == testutil.DQ {
			return types.ChainTFNSA
		}
	}
	return chain
}

func getPrivateKey(env testutil.TestENV) string {
	switch env {
	case testutil.DEV:
		return "7129133f30389e8aebd1231e2f3a5240127534ac1a9d8d29e1083b83eaa6932e"
	case testutil.DQ:
		return "eb81c11c21a9b07edcd83c88b20689b7a372805ba18f3831ae4802d0c9653e8d"
	}
	return ""
}

func GetTestPrivKey(t *testing.T, mnemonic string, coinType int, index int) *secp256k1.PrivKey {
	seed := bip39.NewSeed(mnemonic, "")
	master, ch := hd.ComputeMastersFromSeed(seed)
	priv, err := hd.DerivePrivateKeyForPath(master, ch, fmt.Sprintf("m/44'/%d'/0'/0/%d", coinType, index))

	privKey := &secp256k1.PrivKey{Key: priv}
	assert.NoError(t, err)
	return privKey

}

func (suite *BridgeTestSuite) SetupSuite() {
	suite.ctx = context.Background()

	mnemonic := "notice oak worry limit wrap speak medal online prefer cluster roof addict wrist behave treat actual wasp year salad speed social layer crew genius"
	suite.multiSigPriv = make([]*secp256k1.PrivKey, 3)
	suite.multiSigPriv[0] = GetTestPrivKey(suite.T(), mnemonic, CoinTypeFinschia, 0)
	suite.multiSigPriv[1] = GetTestPrivKey(suite.T(), mnemonic, CoinTypeFinschia, 1)
	suite.multiSigPriv[2] = GetTestPrivKey(suite.T(), mnemonic, CoinTypeFinschia, 2)

	suite.env = testutil.DQ //DQ
	suite.chain = types.ChainFNSA
	suite.toChain = types.ChainKLAY
	suite.signer = GetTestAccount(suite.T(), getPrivateKey(suite.env))

	suite.client, _ = NewClient(envConfig[suite.chain][suite.env].restConfig)
	suite.client.SetAccountPrefix()

	suite.vaultAddress = contractConfig[types.ChainFNSA][suite.env].vaultAddress
	suite.mintAddress = contractConfig[types.ChainFNSA][suite.env].mintAddress

	suite.fnsaTokenAddress = contractConfig[suite.chain][suite.env].fnsaTokenAddress  //fnsa
	suite.evmTokenAddress = contractConfig[suite.toChain][suite.env].fnsaTokenAddress //wfnsa

	suite.fnsaChainName = contractConfig[suite.chain][suite.env].chainName
	suite.evmChainName = contractConfig[suite.toChain][suite.env].chainName

	suite.T().Log(suite.SignerAddress())

}

//func (suite *BridgeTestSuite) ValidatorSignerAddress() []string {
//	return cosmoscommon.FromPublicKeyUnSafe(suite.chain, suite.validatorSigner.Secp256k1().PubKey().Bytes()).String()
//}

func (suite *BridgeTestSuite) SignerAddress() string {
	return cosmoscommon.FromPublicKeyUnSafe(suite.fnsaChainName, suite.signer.Secp256k1().PubKey().Bytes()).String()
}

/**************************** vault ****************************/
func (suite *BridgeTestSuite) TestQueryMsgChildChain() {
	response := bridge.QueryChildChainResponse{}
	err := suite.client.CallWasm(suite.ctx, suite.vaultAddress, bridge.QueryChildChainRequest{
		ChildChainName: suite.evmChainName, //vault 기준 ETH는 child chain
	}, &response)
	assert.NoError(suite.T(), err)
	logger.Info("QueryChildChainRequest", logger.BuildLogInput().WithData("response", response))

}

func (suite *BridgeTestSuite) TestQueryMsgChildToken() { // Vault 토큰 정보 등록 X
	response := bridge.QueryChildTokenResponse{}
	err := suite.client.CallWasm(suite.ctx, suite.vaultAddress, bridge.QueryChildTokenRequest{
		ChildChainName: suite.evmChainName, //vault 기준 ETH는 child chain
		TokenAddr:      suite.fnsaTokenAddress,
	}, &response)
	assert.NoError(suite.T(), err)
	logger.Info("QueryChildTokenRequest", logger.BuildLogInput().WithData("response", response))
}

func (suite *BridgeTestSuite) TestQueryMsgToken() { // Vault 토큰 정보 등록 X
	response := bridge.QueryCoinResponse{}
	err := suite.client.CallWasm(suite.ctx, suite.vaultAddress, bridge.QueryCoinRequest{
		ChildChainName: suite.evmChainName, //vault 기준 ETH는 child chain
		ChildTokenAddr: suite.evmTokenAddress,
	}, &response)
	assert.NoError(suite.T(), err)
	logger.Info("QueryTokenRequest", logger.BuildLogInput().WithData("response", response))
}

func (suite *BridgeTestSuite) TestExecuteDepositCoin() {
	res, err := suite.client.SendTransaction(suite.ctx, suite.signer,
		cosmostypes.WithMsgs(
			cosmostypes.NewMsgExecuteContract(
				suite.SignerAddress(),
				suite.vaultAddress,
				bridge.MsgDepositCoin{
					ChildChainName: suite.evmChainName,
					ChildTokenAddr: suite.evmTokenAddress,
					ToAddr:         "0x01d2E89CD35260b2B453Ca1063d0354c0d25AD29",
				},
				cosmossdk.NewCoins(cosmossdk.NewInt64Coin(tokentypes.DenomByChain(suite.fnsaChainName), 2000000)),
			),
		),
		cosmostypes.WithFeeAmount(cosmossdk.NewInt64Coin(tokentypes.DenomByChain(suite.fnsaChainName), 300000)),
		cosmostypes.WithBroadcastMode(txtypes.BroadcastMode_BROADCAST_MODE_SYNC),
	)

	assert.NoError(suite.T(), err)
	suite.T().Log(res.Hash)
}

/**************************** minter ****************************/
func (suite *BridgeTestSuite) TestQueryMsgConfig() {
	response := bridge.QueryConfigResponse{}
	err := suite.client.CallWasm(suite.ctx, suite.mintAddress, bridge.QueryConfigRequest{}, &response)
	assert.NoError(suite.T(), err)
	logger.Info("QueryConfigRequest", logger.BuildLogInput().WithData("response", response))

	t, _ := new(big.Float).SetString(response.TaxRate)
	taxRateBp, _ := new(big.Int).SetString(t.Mul(t, big.NewFloat(10000)).String(), 10)
	logger.Info("QueryConfigRequest", logger.BuildLogInput().WithData("tax", taxRateBp))
}

func (suite *BridgeTestSuite) TestQueryMsgParentChain() {
	response := bridge.QueryParentChainResponse{}
	err := suite.client.CallWasm(suite.ctx, suite.mintAddress, bridge.QueryParentChainRequest{
		ParentChainName: suite.evmChainName, //minter 기준 parent는 ETH
	}, &response)
	assert.NoError(suite.T(), err)
	logger.Info("QueryParentChainRequest", logger.BuildLogInput().WithData("response", response))
	fee, _ := new(big.Int).SetString(response.Fee, 10)
	//fee랑 tax 반대로!

	logger.Info("QueryChildChainRequest", logger.BuildLogInput().WithData("fee", fee))
}

func (suite *BridgeTestSuite) TestQueryMsgParentToken() {
	response := bridge.QueryParentTokenResponse{}
	err := suite.client.CallWasm(suite.ctx, suite.mintAddress, bridge.QueryParentTokenRequest{
		ParentChainName: suite.evmChainName, //minter 기준 parent는 ETH
		TokenAddr:       suite.fnsaTokenAddress,
	}, &response)
	assert.NoError(suite.T(), err)
	logger.Info("QueryParentTokenRequest", logger.BuildLogInput().WithData("response", response))
}

func (suite *BridgeTestSuite) TestQueryMsgTokenV2() {
	response := bridge.QueryTokenResponseV2{}
	err := suite.client.CallWasm(suite.ctx, suite.mintAddress, bridge.QueryTokenRequestV2{
		ParentChainName: suite.evmChainName, //minter 기준 parent는 ETH
		ParentTokenAddr: suite.evmTokenAddress,
	}, &response)
	assert.NoError(suite.T(), err)
	logger.Info("QueryTokenRequestV2", logger.BuildLogInput().WithData("response", response))
}

func (suite *BridgeTestSuite) TestQueryCW20() {
	var result1 cw20.QueryBalanceResponse
	err := suite.client.CallWasm(suite.ctx, suite.fnsaTokenAddress, cw20.QueryBalanceRequest{
		Address: suite.SignerAddress(),
	}, &result1)
	assert.NoError(suite.T(), err)
	logger.Info("QueryBalanceRequest", logger.BuildLogInput().WithData("response", result1))
}

func (suite *BridgeTestSuite) TestExecuteMint() {
	amount := "10000"
	res, err := suite.client.SendTransaction(suite.ctx, suite.signer,
		cosmostypes.WithMsgs(cosmostypes.NewMsgExecuteContract(
			suite.SignerAddress(),
			suite.fnsaTokenAddress,
			cw20.MsgIncreaseAllowance{
				Spender: suite.mintAddress,
				Amount:  amount,
				Expires: cwutil.NewExpiration(),
			},
			cosmossdk.NewCoins(cosmossdk.NewInt64Coin("cony", 1e6)),
		)),
	)
	assert.NoError(suite.T(), err)
	suite.T().Log(res.Hash)
}

func (suite *BridgeTestSuite) TestExecuteBurn() {
	//amount := "200000" //usdt
	amount := "10000000000000000" //npt
	res, err := suite.client.SendTransaction(suite.ctx, suite.signer,
		cosmostypes.WithMsgs(
			cosmostypes.NewMsgExecuteContract(
				suite.SignerAddress(),
				suite.fnsaTokenAddress,
				cw20.MsgIncreaseAllowance{
					Spender: suite.mintAddress,
					Amount:  amount,
					Expires: cwutil.NewExpiration(),
				},
				cosmossdk.NewCoins(cosmossdk.NewInt64Coin("cony", 0)),
			),
			cosmostypes.NewMsgExecuteContract(
				suite.SignerAddress(),
				suite.mintAddress,
				bridge.MsgBurn{
					ParentChainName: suite.evmChainName,
					ParentTokenAddr: suite.evmTokenAddress,
					ToAddr:          "0xbD1649F9a1599301639d6525Bb0C87e5081e28d3",
					Amount:          amount,
				},
				cosmossdk.NewCoins(cosmossdk.NewInt64Coin("cony", 10000)),
			),
		),
	)
	assert.NoError(suite.T(), err)
	suite.T().Log(res.Hash)
}

func (suite *BridgeTestSuite) TestTransferToOasis() {
	resp, err := suite.client.SendTransaction(suite.ctx, suite.signer,
		cosmostypes.WithMsgs(
			cosmostypes.NewMsgExecuteContract(
				suite.SignerAddress(),
				suite.fnsaTokenAddress,
				cw20.MsgTransfer{
					Recipient: "link1887ty9wmywzvspl4d9qkvjhzar0ezhclvq00hj",
					Amount:    util.ToString(big.NewInt(1000000000)),
				},
				cosmossdk.NewCoins(),
			),
		))
	assert.NoError(suite.T(), err)
	if err == nil {
		PrintWasmLogs(suite.T(), suite.client, resp.Hash)
	}
}

func (suite *BridgeTestSuite) TestSignatureV2Conversions() {
	pubKeys := []cryptotypes.PubKey{
		suite.multiSigPriv[0].PubKey(),
		suite.multiSigPriv[1].PubKey(),
		suite.multiSigPriv[2].PubKey(),
	}

	//https://github.com/cosmos/cosmos-sdk/blob/v0.42.2/client/keys/add.go (line 172 참고)
	sort.Slice(pubKeys, func(i, j int) bool {
		return bytes.Compare(pubKeys[i].Address(), pubKeys[j].Address()) < 0
	})

	multisigPk := kmultisig.NewLegacyAminoPubKey(2, pubKeys)
	multisigTxBuilder := suite.client.TxConfig().NewTxBuilder()

	//1차 중은님 address와 다름.
	multisigAddr := cosmoscommon.FromAddress(suite.chain, multisigPk.Address().Bytes())
	assert.Equal(suite.T(), multisigAddr.String(), "link1ut5srsfjcvuh0zwzq2wwyq32dedv5kd3m7ew8w")

	suite.T().Log(multisigPk)
	suite.T().Log(multisigAddr.String())
	msg := cosmostypes.NewMsgExecuteContract(
		multisigAddr.String(),
		suite.mintAddress,
		bridge.MsgMint{
			ParentChainName: suite.evmChainName,
			ParentTokenAddr: suite.evmTokenAddress,
			ParentTx:        "0x0002",
			FromAddr:        "0x0000000",
			ToAddr:          "link1p8hkv3f3669fl9dj6z46urmvk3fp0mdplr259v",
			Amount:          "10000000000000000000",
		},
		cosmossdk.NewCoins(),
	)
	err := multisigTxBuilder.SetMsgs(msg)
	multisigTxBuilder.SetFeeAmount(cosmossdk.NewCoins(cosmossdk.NewInt64Coin(tokentypes.DenomByChain(suite.chain), 5_000)))
	multisigTxBuilder.SetGasLimit(3_000_000)

	assert.NoError(suite.T(), err)
	assert.Error(suite.T(), multisigTxBuilder.GetTx().ValidateBasic())

	num, seq, err := suite.client.GetAccountNumberAndSequence(
		context.Background(), multisigAddr.String())

	msigData := multisig.NewMultisig(2)
	for _, pub := range pubKeys {
		txBuilder := suite.client.TxConfig().NewTxBuilder()
		txBuilder.SetFeeAmount(cosmossdk.NewCoins(cosmossdk.NewInt64Coin(tokentypes.DenomByChain(suite.chain), 5_000)))
		txBuilder.SetGasLimit(2_000_000)
		txBuilder.SetMsgs(msg)

		//1
		sigData := signingtypes.SingleSignatureData{
			SignMode:  signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
			Signature: nil,
		}

		sigV2 := signingtypes.SignatureV2{
			PubKey:   pub,
			Data:     &sigData,
			Sequence: seq,
		}
		txBuilder.SetSignatures(sigV2)

		signerData := signing.SignerData{
			Address:       multisigAddr.String(),
			ChainID:       suite.client.ChainId(),
			AccountNumber: num,
			Sequence:      seq,
			PubKey:        pub,
		}
		mSignBytes, _ := suite.client.TxConfig().SignModeHandler().GetSignBytes(
			signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON, signerData, txBuilder.GetTx())

		for _, priv := range suite.multiSigPriv {
			if !priv.PubKey().Equals(pub) {
				continue
			}
			sign, _ := priv.Sign(mSignBytes)
			sigData = signingtypes.SingleSignatureData{
				SignMode:  signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
				Signature: sign,
			}

			sigV2 = signingtypes.SignatureV2{
				PubKey:   pub,
				Data:     &sigData,
				Sequence: seq,
			}
			multisig.AddSignatureV2(msigData, sigV2, multisigPk.GetPubKeys())
		}
	}

	msig := signingtypes.SignatureV2{PubKey: multisigPk, Data: msigData, Sequence: seq}

	err = multisigTxBuilder.SetSignatures(msig)
	assert.NoError(suite.T(), err)
	txBytes, err := suite.client.TxConfig().TxEncoder()(multisigTxBuilder.GetTx())
	assert.NoError(suite.T(), err)

	result, _ := suite.client.RawTxAsync(suite.ctx, txBytes, nil)
	suite.T().Log(result)

	suite.client.RawTxAsync(context.Background(), txBytes, nil)

	assert.NoError(suite.T(), err)
	//sigTx = multisigTxBuilder.GetTx()
	//sigsV2, err := sigTx.GetSignaturesV2()

	//fmt.Println(len(sigsV2))
}
