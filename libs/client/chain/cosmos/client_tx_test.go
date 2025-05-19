package cosmos

import (
	"context"
	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	kmultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	signing2 "github.com/cosmos/cosmos-sdk/x/auth/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	types3 "github.com/curtis0505/bridge/libs/client/chain/cosmos/types"
	clienttypes "github.com/curtis0505/bridge/libs/client/chain/types"
	cosmoscommon "github.com/curtis0505/bridge/libs/common/cosmos"
	"github.com/curtis0505/bridge/libs/logger/v2"
	"github.com/curtis0505/bridge/libs/testutil"
	"github.com/curtis0505/bridge/libs/types"
	tokentypes "github.com/curtis0505/bridge/libs/types/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TxTestSuite struct {
	suite.Suite

	env    testutil.TestENV
	ctx    context.Context
	client clienttypes.CosmosClient

	chain      string
	feeGranter *types.Account
	signer     *types.Account
	//validatorSigner  []*types.Account
	multiSigPriv []*secp256k1.PrivKey
}

func TestTxTestSuite(t *testing.T) {
	logger.InitLog(logger.Config{UseTerminal: true, VerbosityTerminal: 5})
	suite.Run(t, new(TxTestSuite))
}

func (suite *TxTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.env = testutil.DEV
	suite.chain = types.ChainFNSA
	suite.client, _ = NewClient(envConfig[suite.chain][suite.env].restConfig)
	suite.client.SetAccountPrefix()
	suite.feeGranter = GetTestAccount(suite.T(), "58d7ad6ece2c88ce01648aaf0799e27eb760e37f11d30b937f36529bdee926db")
	suite.signer = GetTestAccount(suite.T(), "83c13018679900c69be911758784782f693215f754c995fc17e304c9fbf90afb")

	mnemonic := "notice oak worry limit wrap speak medal online prefer cluster roof addict wrist behave treat actual wasp year salad speed social layer crew genius"
	suite.multiSigPriv = make([]*secp256k1.PrivKey, 3)
	suite.multiSigPriv[0] = GetTestPrivKey(suite.T(), mnemonic, CoinTypeFinschia, 0)
	suite.multiSigPriv[1] = GetTestPrivKey(suite.T(), mnemonic, CoinTypeFinschia, 1)
	suite.multiSigPriv[2] = GetTestPrivKey(suite.T(), mnemonic, CoinTypeFinschia, 2)
}

func (suite *TxTestSuite) feeGranterAddress() string {
	return cosmoscommon.FromPublicKeyUnSafe(suite.chain, suite.feeGranter.Secp256k1().PubKey().Bytes()).String()
}

func (suite *TxTestSuite) signerAddress() string {
	return cosmoscommon.FromPublicKeyUnSafe(suite.chain, suite.signer.Secp256k1().PubKey().Bytes()).String()
}

func (suite *TxTestSuite) multiSigAddress() string {
	pubKeys := []cryptotypes.PubKey{
		suite.multiSigPriv[0].PubKey(),
		suite.multiSigPriv[1].PubKey(),
		suite.multiSigPriv[2].PubKey(),
	}

	multiSigPubKey := kmultisig.NewLegacyAminoPubKey(2, pubKeys)
	multiSigAddress := cosmoscommon.FromAddress(suite.chain, multiSigPubKey.Address().Bytes())
	return multiSigAddress.String()
}

func (suite *TxTestSuite) TestAllowFeeGrant() {
	allowance := &feegrant.BasicAllowance{
		SpendLimit: cosmossdk.NewCoins(cosmossdk.NewInt64Coin("cony", 1e6)),
	}
	allowanceAny, err := types2.NewAnyWithValue(allowance)
	suite.NoError(err)

	msg := &feegrant.MsgGrantAllowance{
		Granter:   suite.feeGranterAddress(),
		Grantee:   suite.signerAddress(),
		Allowance: allowanceAny,
	}

	resp, err := suite.client.SendTransaction(suite.ctx, suite.feeGranter, types3.WithMsgs(msg), types3.WithFeeAmount(cosmossdk.NewInt64Coin("cony", 1e5)))
	suite.NoError(err)
	_ = resp
}

func (suite *TxTestSuite) TestFeeGranter() {
	feePayerAddress := cosmoscommon.FromPublicKeyUnSafe(suite.chain, suite.feeGranter.Secp256k1().PubKey().Bytes()).String()
	signerAddress := cosmoscommon.FromPublicKeyUnSafe(suite.chain, suite.signer.Secp256k1().PubKey().Bytes()).String()

	msg := &banktypes.MsgSend{
		FromAddress: signerAddress,
		ToAddress:   feePayerAddress,
		Amount:      cosmossdk.NewCoins(cosmossdk.NewInt64Coin("cony", 1e5)),
	}

	resp, err := suite.client.SendTransaction(suite.ctx, suite.signer,
		types3.WithMsgs(msg),
		types3.WithFeeAmount(cosmossdk.NewInt64Coin("cony", 1e5)),
		types3.WithFeeGranter(feePayerAddress),
	)
	suite.NoError(err)
	if err == nil {
		suite.T().Log(resp.Hash)
	}
}

func (suite *TxTestSuite) TestMultiSigAllowFeeGrant() {
	allowance := &feegrant.BasicAllowance{
		SpendLimit: cosmossdk.NewCoins(cosmossdk.NewInt64Coin("cony", 1e6)),
	}
	allowanceAny, err := types2.NewAnyWithValue(allowance)
	suite.NoError(err)

	msg := &feegrant.MsgGrantAllowance{
		Granter:   suite.feeGranterAddress(),
		Grantee:   suite.multiSigAddress(),
		Allowance: allowanceAny,
	}

	resp, err := suite.client.SendTransaction(suite.ctx, suite.feeGranter, types3.WithMsgs(msg), types3.WithFeeAmount(cosmossdk.NewInt64Coin("cony", 1e5)))
	suite.NoError(err)
	_ = resp
}

func (suite *TxTestSuite) TestMultiSigFeeGranter() {
	pubKeys := []cryptotypes.PubKey{
		suite.multiSigPriv[0].PubKey(),
		suite.multiSigPriv[1].PubKey(),
		suite.multiSigPriv[2].PubKey(),
	}

	multiSigPubKey := kmultisig.NewLegacyAminoPubKey(2, pubKeys)
	multiSigAddress := cosmoscommon.FromAddress(suite.chain, multiSigPubKey.Address().Bytes())
	testMsg := banktypes.MsgSend{
		FromAddress: suite.multiSigAddress(),
		ToAddress:   suite.feeGranterAddress(),
		Amount:      cosmossdk.NewCoins(cosmossdk.NewInt64Coin("cony", 1000)),
	}

	suite.T().Log(multiSigAddress.String())
	accountNum, accountSeq, err := suite.client.GetAccountNumberAndSequence(suite.ctx, multiSigAddress.String())
	assert.NoError(suite.T(), err)

	signMode := signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON

	suite.T().Log(accountNum, accountSeq)

	multiSigTxBuilder := suite.client.TxConfig().NewTxBuilder()
	multiSigTxBuilder.SetMsgs(&testMsg)
	multiSigTxBuilder.SetFeeAmount(cosmossdk.NewCoins(cosmossdk.NewInt64Coin(tokentypes.DenomByChain(suite.chain), 5_000)))
	multiSigTxBuilder.SetGasLimit(2_000_000)

	feeGranter, _ := cosmossdk.AccAddressFromBech32(suite.feeGranterAddress())
	multiSigTxBuilder.SetFeeGranter(feeGranter)

	multisigSig := multisig.NewMultisig(len(multiSigPubKey.PubKeys))
	for i := 0; i < len(multiSigPubKey.PubKeys); i++ {
		txBuilder := suite.client.TxConfig().NewTxBuilder()
		txBuilder.SetFeeAmount(cosmossdk.NewCoins(cosmossdk.NewInt64Coin(tokentypes.DenomByChain(suite.chain), 5_000)))
		txBuilder.SetGasLimit(2_000_000)
		txBuilder.SetMsgs(&testMsg)

		signerData := signing2.SignerData{
			ChainID:       suite.client.ChainId(),
			AccountNumber: accountNum,
			Sequence:      accountSeq,
			PubKey:        pubKeys[i],
			Address:       multiSigAddress.String(),
		}

		sigData := signing.SingleSignatureData{
			SignMode:  signMode,
			Signature: nil,
		}

		sigV2 := signing.SignatureV2{
			PubKey:   pubKeys[i],
			Data:     &sigData,
			Sequence: accountSeq,
		}

		// Round1
		txBuilder.SetSignatures(sigV2)
		bytesToSign, err := suite.client.TxConfig().SignModeHandler().GetSignBytes(
			signMode,
			signerData,
			txBuilder.GetTx(),
		)
		if err != nil {
			suite.T().Log(err)
		}

		signature, err := suite.multiSigPriv[i].Sign(bytesToSign)
		sigData = signing.SingleSignatureData{
			SignMode:  signMode,
			Signature: signature,
		}

		sigV2 = signing.SignatureV2{
			PubKey:   pubKeys[i],
			Data:     &sigData,
			Sequence: accountSeq,
		}

		//suite.T().Log(string(bytesToSign))
		err = multisig.AddSignatureV2(multisigSig, sigV2, multiSigPubKey.GetPubKeys())
		if err != nil {
			suite.T().Log(err)
		}
		//bz, _ := suite.client.TxConfig().MarshalSignatureJSON([]signing.SignatureV2{sigV2})
		//suite.T().Log(string(bz))
	}

	suite.T().Log(multisigSig)

	sigV2 := signing.SignatureV2{
		PubKey:   multiSigPubKey,
		Data:     multisigSig,
		Sequence: accountSeq,
	}

	err = multiSigTxBuilder.SetSignatures(sigV2)
	if err != nil {
		suite.T().Log(err)
	}
	txBytes, err := suite.client.TxConfig().TxEncoder()(multiSigTxBuilder.GetTx())

	suite.T().Log(suite.client.BroadcastRawTx(context.Background(), txBytes, txtypes.BroadcastMode_BROADCAST_MODE_SYNC))
}
