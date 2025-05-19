package rest

import (
	"context"
	"fmt"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cometbft/cometbft/crypto"
	kmultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	cosmostypes "github.com/curtis0505/bridge/libs/client/chain/cosmos/types"
	clienttypes "github.com/curtis0505/bridge/libs/client/chain/types"
	cosmoscommon "github.com/curtis0505/bridge/libs/common/cosmos"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/types/bridge"
	tokentypes "github.com/curtis0505/bridge/libs/types/token"
	"github.com/curtis0505/bridge/libs/util"
	"github.com/shopspring/decimal"
)

const (
	gasLimit = 500_000
)

func (c *client) NewMinter(address string, abi []map[string]interface{}) (bridge.Minter, error) {
	return newMinter(c, address), nil
}

func (c *client) NewVault(address string, abi []map[string]interface{}) (bridge.Vault, error) {
	return newVault(c, address), nil
}

func (c *client) NewMultiSigWallet(address string) (bridge.MultiSigWallet, error) {
	return newMultiSigWallet(c, address), nil
}

func (c *client) SignMultiSig(ctx context.Context, account *types.Account, multiSig string, msg *wasmtypes.MsgExecuteContract) (signing.SignatureV2, error) {
	minGasPrice, err := c.GetMinimumGasPrice(ctx)
	if err != nil {
		return signing.SignatureV2{}, err
	}

	gasPrice := util.ToDecimal(minGasPrice, 0).Mul(decimal.NewFromInt(gasLimit)).Round(0).IntPart()
	txOption := &cosmostypes.TxOption{
		FeeAmount: cosmossdk.NewCoins(cosmossdk.NewInt64Coin(tokentypes.DenomByChain(c.Chain()), gasPrice)),
		GasLimit:  gasLimit,
	}

	txBuilder := c.TxConfig().NewTxBuilder()
	txBuilder.SetFeeAmount(txOption.FeeAmount)
	txBuilder.SetGasLimit(txOption.GasLimit)
	err = txBuilder.SetMsgs(msg)
	if err != nil {
		return signing.SignatureV2{}, types.WrapError("SetMsgs", err)
	}

	num, seq, err := c.GetAccountNumberAndSequence(ctx, multiSig)
	if err != nil {
		return signing.SignatureV2{}, types.WrapError("GetAccountNumberAndSequence", err)
	}

	if account.Type == types.KMS {
		pubKeyUncompressed, err := account.KMS.PublicKey(context.Background())
		if err != nil {
			return signing.SignatureV2{}, types.WrapError("KMS:PublicKey", err)
		}
		pubKey, err := cosmoscommon.NewPubKey(pubKeyUncompressed)
		if err != nil {
			return signing.SignatureV2{}, types.WrapError("KMS:FromPublicKey", err)
		}

		signMode := signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON
		err = txBuilder.SetSignatures(
			signing.SignatureV2{
				PubKey: pubKey,
				Data: &signing.SingleSignatureData{
					SignMode:  signMode,
					Signature: nil,
				},
				Sequence: seq,
			},
		)
		if err != nil {
			return signing.SignatureV2{}, types.WrapError("KMS:SetSignatures", err)
		}
		bytesToSign, err := c.TxConfig().SignModeHandler().GetSignBytes(
			signMode,
			xauthsigning.SignerData{
				ChainID:       c.ChainId(),
				AccountNumber: num,
				Sequence:      seq,
				PubKey:        pubKey,
				Address:       multiSig,
			},
			txBuilder.GetTx(),
		)
		if err != nil {
			return signing.SignatureV2{}, types.WrapError("KMS:bytesToSign", err)
		}

		signature, err := account.KMS.Sign(context.Background(), crypto.Sha256(bytesToSign))
		if err != nil {
			return signing.SignatureV2{}, types.WrapError("KMS:Sign", err)
		}

		return signing.SignatureV2{
			PubKey: pubKey,
			Data: &signing.SingleSignatureData{
				SignMode:  signMode,
				Signature: signature,
			},
			Sequence: seq,
		}, nil
	} else {
		priv := account.Secp256k1()
		err = txBuilder.SetSignatures(
			signing.SignatureV2{
				PubKey: priv.PubKey(),
				Data: &signing.SingleSignatureData{
					SignMode:  signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
					Signature: nil,
				},
				Sequence: seq,
			})
		if err != nil {
			return signing.SignatureV2{}, types.WrapError("SetSignatures", err)
		}

		bytesToSign, err := c.TxConfig().SignModeHandler().GetSignBytes(
			signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
			xauthsigning.SignerData{
				ChainID:       c.ChainId(),
				AccountNumber: num,
				Sequence:      seq,
				PubKey:        priv.PubKey(),
				Address:       multiSig,
			},
			txBuilder.GetTx(),
		)
		if err != nil {
			return signing.SignatureV2{}, types.WrapError("bytesToSign", err)
		}

		signature, err := priv.Sign(bytesToSign)
		if err != nil {
			return signing.SignatureV2{}, types.WrapError("Sign", err)
		}

		return signing.SignatureV2{
			PubKey: priv.PubKey(),
			Data: &signing.SingleSignatureData{
				SignMode:  signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
				Signature: signature,
			},
			Sequence: seq,
		}, nil
	}
}

func (c *client) SendMultiSigTransaction(ctx context.Context, multiSigPubKey *kmultisig.LegacyAminoPubKey, msg *wasmtypes.MsgExecuteContract, singedTxs ...signing.SignatureV2) (*clienttypes.SendTxAsyncResult, error) {
	minGasPrice, err := c.GetMinimumGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	gasPrice := util.ToDecimal(minGasPrice, 0).Mul(decimal.NewFromInt(gasLimit)).Round(0).IntPart()
	txOption := &cosmostypes.TxOption{
		FeeAmount: cosmossdk.NewCoins(cosmossdk.NewInt64Coin(tokentypes.DenomByChain(c.Chain()), gasPrice)),
		GasLimit:  gasLimit,
	}

	multiSigTxBuilder := c.TxConfig().NewTxBuilder()
	multiSigTxBuilder.SetFeeAmount(txOption.FeeAmount)
	multiSigTxBuilder.SetGasLimit(txOption.GasLimit)
	err = multiSigTxBuilder.SetMsgs(msg)
	if err != nil {
		return nil, types.WrapError("SetMsgs", err)
	}

	multiSigSign := multisig.NewMultisig(len(multiSigPubKey.GetPubKeys()))

	sequence := uint64(0)
	for _, signedTx := range singedTxs {
		if sequence == uint64(0) {
			sequence = signedTx.Sequence
		}

		if sequence != signedTx.Sequence {
			return nil, types.WrapError("Sequence", fmt.Errorf(""))
		}

		multisig.AddSignatureV2(multiSigSign, signedTx, multiSigPubKey.GetPubKeys())
	}

	multiSigTxBuilder.SetSignatures(signing.SignatureV2{
		PubKey:   multiSigPubKey,
		Data:     multiSigSign,
		Sequence: sequence,
	})

	rawTx, err := c.TxConfig().TxEncoder()(multiSigTxBuilder.GetTx())
	if err != nil {
		return nil, types.WrapError("txBytes", err)
	}

	return c.BroadcastRawTx(ctx, rawTx, txtypes.BroadcastMode_BROADCAST_MODE_SYNC)
}
