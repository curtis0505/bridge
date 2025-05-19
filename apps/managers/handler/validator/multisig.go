package validator

import (
	"bytes"
	"context"
	"errors"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/avast/retry-go/v4"
	kmultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	eventtypes "github.com/curtis0505/bridge/apps/managers/handler/event/types"
	validatortypes "github.com/curtis0505/bridge/apps/managers/handler/validator/types"
	"github.com/curtis0505/bridge/libs/cache"
	cosmostypes "github.com/curtis0505/bridge/libs/client/chain/cosmos/types"
	clienttypes "github.com/curtis0505/bridge/libs/client/chain/types"
	cosmoscommon "github.com/curtis0505/bridge/libs/common/cosmos"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	"github.com/curtis0505/bridge/libs/types/cosmos/bridge"
	"github.com/curtis0505/bridge/libs/types/token"
	"github.com/curtis0505/bridge/libs/util"
	"github.com/shopspring/decimal"
	"sort"
	"strings"
	"time"
)

type MultiSigAddressResponse struct {
	Address string `json:"address"`
}

func (p *ValidatorHandler) MultiSigPubKey(chain string, threshold int) *kmultisig.LegacyAminoPubKey {
	var pubKeys []cryptotypes.PubKey

	for _, validator := range p.validatorList {
		pubKey, err := validator.GetPubKey(chain)
		if err != nil {
			p.logger.Warn("event", "MultiSigPubKey", "validator", validator.Name, "err", err)
			return nil
		}
		pubKeys = append(pubKeys, pubKey)
	}

	sort.Slice(pubKeys, func(i, j int) bool {
		return bytes.Compare(pubKeys[i].Address(), pubKeys[j].Address()) < 0
	})

	return kmultisig.NewLegacyAminoPubKey(threshold, pubKeys)
}

const redisKeyLockMultiSigLock = "bridge:manager:multisig:lock"

func (p *ValidatorHandler) AddCosmosMultiSigTx(tx validatortypes.CosmosMultiSigTx) {
	p.logger.Info("event", "AddCosmosMultiSigTx", "chain", tx.Chain, "txHash", tx.TxHash, "eventName", tx.EventName)
	p.multiSigTxChan <- tx
}

func (p *ValidatorHandler) CosmosMultiSigTx() {
	if len(p.multiSigTxChan) == 0 {
		return
	}

	if err := p.multiSigTxLock.TryLock(); err != nil {
		p.logger.Warn("event", "CosmosMultiSigTx", "lock", err, "pendingTx", len(p.multiSigTxChan))
		return
	}
	defer func() {
		if ok, err := p.multiSigTxLock.Unlock(); !ok || err != nil {
			p.logger.Warn("event", "CosmosMultiSigTx", "unlock", err)
		}
	}()

	select {
	case multiSigTx := <-p.multiSigTxChan:
		if multiSigTx.TxHash == "" || multiSigTx.Chain == "" {
			p.logger.Warn("event", "CosmosMultiSigTx", "channel", "empty string")
			return
		}

		p.logger.Info("event", "CosmosMultiSigTx", "chain", multiSigTx.Chain, "txHash", multiSigTx.TxHash, "txType", multiSigTx.EventName, "message", "start to sign mutlsig tx")

		ctx := context.Background()

		eventLog := multiSigTx.EventLog
		c := p.client.ProxyClient(multiSigTx.Chain).(clienttypes.CosmosClient)
		c.SetAccountPrefix()

		multiSigPubKey := p.MultiSigPubKey(eventLog.ToChainName, validatortypes.Threshold)
		if len(multiSigPubKey.GetPubKeys()) == 0 {
			return
		}

		multiSignatureData := multisig.NewMultisig(len(multiSigPubKey.GetPubKeys()))
		multiSigAddress := cosmoscommon.FromAddress(eventLog.ToChainName, multiSigPubKey.Address().Bytes()).String()
		p.logger.Info("event", "CosmosMultiSigTx", "multiSigAddress", multiSigAddress)

		var msg *wasmtypes.MsgExecuteContract
		switch multiSigTx.EventName {
		case eventtypes.EventDeposit:
			minter, err := cache.ContractCache().GetContractByContractID(eventLog.ToChainName, bridgetypes.MinterContractID)
			if err != nil {
				p.logger.Error("event", "CosmosMultiSigTx", "err", err)
				return
			}

			msg = cosmostypes.NewMsgExecuteContract(
				multiSigAddress,
				minter.Address,
				bridge.MsgMint{
					ParentChainName: eventLog.FromChainName,
					ParentTokenAddr: strings.ToLower(eventLog.FromTokenAddr.String()),
					ParentTx:        eventLog.TxHash,
					FromAddr:        eventLog.From.String(),
					ToAddr:          cosmoscommon.FromAddress(eventLog.ToChainName, eventLog.ToAddr.Bytes()).String(),
					Amount:          eventLog.Amount.String(),
				},
				cosmossdk.NewCoins(),
			)
		case eventtypes.EventBurn:
			vault, err := cache.ContractCache().GetContractByContractID(eventLog.ToChainName, bridgetypes.VaultContractID)
			if err != nil {
				p.logger.Error("event", "CosmosMultiSigTx", "err", err)
				return
			}

			coinResponse := bridge.QueryCoinResponse{}
			c.CallWasm(context.Background(), vault.Address, bridge.QueryCoinRequest{
				ChildChainName: eventLog.FromChainName,
				ChildTokenAddr: strings.ToLower(eventLog.FromTokenAddr.String()),
			}, &coinResponse)

			if coinResponse.NativeDenom == "" {
				msg = cosmostypes.NewMsgExecuteContract(
					multiSigAddress,
					vault.Address,
					bridge.MsgWithdraw{
						ChildChainName: eventLog.FromChainName,
						ChildTokenAddr: strings.ToLower(eventLog.FromTokenAddr.String()),
						ChildTx:        eventLog.TxHash,
						FromAddr:       eventLog.From.String(),
						ToAddr:         cosmoscommon.FromAddress(eventLog.ToChainName, eventLog.ToAddr.Bytes()).String(),
						Amount:         eventLog.Amount.String(),
					},
					cosmossdk.NewCoins(),
				)
			} else {
				msg = cosmostypes.NewMsgExecuteContract(
					multiSigAddress,
					vault.Address,
					bridge.MsgWithdrawCoin{
						ChildChainName: eventLog.FromChainName,
						ChildTokenAddr: strings.ToLower(eventLog.FromTokenAddr.String()),
						ChildTx:        eventLog.TxHash,
						FromAddr:       eventLog.From.String(),
						ToAddr:         cosmoscommon.FromAddress(eventLog.ToChainName, eventLog.ToAddr.Bytes()).String(),
						Amount:         eventLog.Amount.String(),
					},
					cosmossdk.NewCoins(),
				)
			}
		}

		accountNumber, accountSequence, err := c.GetAccountNumberAndSequence(context.Background(), multiSigAddress)
		if err != nil {
			return
		}

		p.logger.Info("event", "CosmosMultiSigTx", "address", multiSigAddress, "accSeq", accountSequence, "accNum", accountNumber)

		for _, validator := range p.validatorList {
			// Request SignatureV2 for MultiSigTransaction
			signature, err := validator.GetSignatureV2(
				multiSigTx.Chain,
				multiSigTx.TxHash,
				multiSigAddress,
				int64(accountNumber),
				int64(accountSequence),
			)
			if err != nil {
				p.logger.Error("event", "CosmosMultiSigTx", "err", err)
				return
			}

			if signature.Sequence != accountSequence {
				p.logger.Error("event", "CosmosMultiSigTx", "err", "invalid sequence")
				return
			}

			p.logger.Info("event", "CosmosMultiSigTx", "validator", validator.Name, "signature", signature)
			if err := multisig.AddSignatureV2(multiSignatureData, signature, multiSigPubKey.GetPubKeys()); err != nil {
				p.logger.Error("event", "CosmosMultiSigTx", "err", err)
				return
			}
		}
		multiSignatureV2 := signing.SignatureV2{
			PubKey:   multiSigPubKey,
			Data:     multiSignatureData,
			Sequence: accountSequence,
		}

		minGasPrice, err := c.GetMinimumGasPrice(ctx)
		if err != nil {
			return
		}
		gasLimit := int64(500_000)

		gasPrice := util.ToDecimal(minGasPrice, 0).Mul(decimal.NewFromInt(gasLimit)).Round(0).IntPart()
		txBuilder := c.TxConfig().NewTxBuilder()
		txBuilder.SetMsgs(msg)
		txBuilder.SetFeeAmount(cosmossdk.NewCoins(cosmossdk.NewInt64Coin(token.DenomByChain(c.Chain()), gasPrice)))
		txBuilder.SetGasLimit(uint64(gasLimit))
		txBuilder.SetSignatures(multiSignatureV2)
		feeGranter, err := p.configService.FindOneAppSettingByKey(ctx, "finschia_fee_granter")
		if err == nil {
			txBuilder.SetFeeGranter(cosmoscommon.FromBech32UnSafe(c.Chain(), feeGranter.Value.(string)).Address())
		}

		txBytes, err := c.TxConfig().TxEncoder()(txBuilder.GetTx())
		if err != nil {
			p.logger.Error("event", "CosmosMultiSigTx", "err", err)
		}

		txResult, err := p.client.BroadcastRawTx(ctx, multiSigTx.Chain, txBytes, tx.BroadcastMode_BROADCAST_MODE_SYNC)
		if err != nil {
			p.logger.Error("event", "CosmosMultiSigTx", "err", err)
			return
		} else {
			p.logger.Info(
				"event", "CosmosMultiSigTx",
				"chain", multiSigTx.EventLog.ToChainName, "txHash", txResult.Hash,
				"message", "transaction broadcasted",
			)
		}

		err = retry.Do(
			func() error {
				_, currentSequence, err := c.GetAccountNumberAndSequence(context.Background(), multiSigAddress)
				if err != nil {
					return err
				}

				if accountSequence == currentSequence {
					return errors.New("same sequence")
				} else {
					return nil
				}
			},
			retry.Delay(time.Second*2),
			retry.Attempts(30),
			retry.DelayType(retry.FixedDelay),
			retry.OnRetry(func(n uint, err error) {
				p.logger.Trace(
					"event", "CosmosMultiSigTx",
					"chain", multiSigTx.EventLog.ToChainName, "txHash", txResult.Hash,
					"message", "waiting transaction",
				)
			}),
		)
		if err != nil {
			p.logger.Error(
				"event", "CosmosMultiSigTx",
				"chain", multiSigTx.EventLog.ToChainName, "txHash", txResult.Hash,
				"message", "transaction broadcasted but still not confirmed",
			)
		} else {
			p.logger.Info(
				"event", "CosmosMultiSigTx",
				"chain", multiSigTx.EventLog.ToChainName, "txHash", txResult.Hash,
				"message", "transaction confirmed",
			)
		}
	}
}
