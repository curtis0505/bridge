package cosmos

import (
	"context"
	kmultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	signing2 "github.com/cosmos/cosmos-sdk/x/auth/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
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

type MultiSigSuite struct {
	suite.Suite

	env    testutil.TestENV
	ctx    context.Context
	client clienttypes.CosmosClient

	chain  string
	signer *types.Account
	//validatorSigner  []*types.Account
	multiSigPriv []*secp256k1.PrivKey
}

func TestMultiSigSuite(t *testing.T) {
	logger.InitLog(logger.Config{UseTerminal: true, VerbosityTerminal: 5})
	suite.Run(t, new(MultiSigSuite))
}

func (suite *MultiSigSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.env = testutil.DEV
	suite.chain = types.ChainFNSA
	suite.client, _ = NewClient(envConfig[suite.chain][suite.env].restConfig)
	suite.client.SetAccountPrefix()

	mnemonic := "notice oak worry limit wrap speak medal online prefer cluster roof addict wrist behave treat actual wasp year salad speed social layer crew genius"
	suite.multiSigPriv = make([]*secp256k1.PrivKey, 3)
	suite.multiSigPriv[0] = GetTestPrivKey(suite.T(), mnemonic, CoinTypeFinschia, 0)
	suite.multiSigPriv[1] = GetTestPrivKey(suite.T(), mnemonic, CoinTypeFinschia, 1)
	suite.multiSigPriv[2] = GetTestPrivKey(suite.T(), mnemonic, CoinTypeFinschia, 2)
}

func (suite *MultiSigSuite) TestSign() {
	pubKeys := []cryptotypes.PubKey{
		suite.multiSigPriv[0].PubKey(),
		suite.multiSigPriv[1].PubKey(),
		suite.multiSigPriv[2].PubKey(),
	}

	multiSigPubKey := kmultisig.NewLegacyAminoPubKey(2, pubKeys)
	multiSigAddress := cosmoscommon.FromAddress(suite.chain, multiSigPubKey.Address().Bytes())
	testMsg := banktypes.MsgSend{
		FromAddress: "link1dex3p248ez2eznlsaxr7eey99vrw6rra09qm5z",
		ToAddress:   "link1dex3p248ez2eznlsaxr7eey99vrw6rra09qm5z",
		Amount:      cosmossdk.NewCoins(cosmossdk.NewInt64Coin("cony", 1000)),
	}

	accountNum, accountSeq, err := suite.client.GetAccountNumberAndSequence(suite.ctx, multiSigAddress.String())
	assert.NoError(suite.T(), err)

	signMode := signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON

	suite.T().Log(accountNum, accountSeq)

	multiSigTxBuilder := suite.client.TxConfig().NewTxBuilder()
	multiSigTxBuilder.SetMsgs(&testMsg)
	multiSigTxBuilder.SetFeeAmount(cosmossdk.NewCoins(cosmossdk.NewInt64Coin(tokentypes.DenomByChain(suite.chain), 5_000)))
	multiSigTxBuilder.SetGasLimit(2_000_000)

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
		assert.NoError(suite.T(), err)

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

		suite.T().Log(string(bytesToSign))
		//multisig.AddSignatureV2(multisigSig, sigV2, multiSigPubKey.GetPubKeys())
		bz, _ := suite.client.TxConfig().MarshalSignatureJSON([]signing.SignatureV2{sigV2})
		suite.T().Log(string(bz))
	}

	sigV2 := signing.SignatureV2{
		PubKey:   multiSigPubKey,
		Data:     multisigSig,
		Sequence: accountSeq,
	}

	multiSigTxBuilder.SetSignatures(sigV2)
	//	txBytes, err := suite.client.TxConfig().TxEncoder()(multiSigTxBuilder.GetTx())
	//
	//	suite.T().Log(suite.client.RawTxAsync(context.Background(), txBytes, nil))
}
