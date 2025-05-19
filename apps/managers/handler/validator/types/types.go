package types

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/curtis0505/bridge/apps/managers/util"
	"github.com/curtis0505/bridge/libs/client/chain/cosmos/types"
	"github.com/curtis0505/bridge/libs/common"
	"github.com/curtis0505/bridge/libs/elog"
	mongoentity "github.com/curtis0505/bridge/libs/entity/mongo"
	commontypes "github.com/curtis0505/bridge/libs/types"
	util2 "github.com/curtis0505/bridge/libs/util"
	"math/big"
	"time"
)

type ValidatorInfo struct {
	*mongoentity.ValidatorInfo
	AddressInfo map[string]*AddressInfo
	Active      bool `json:"active"`
}

type ValidatorSummary struct {
	Name    string
	Address string
}

type AddressInfo struct {
	Chain   string `json:"chain"`
	Address string `json:"address"`
	Balance string `json:"balance"`
}

const (
	addressPath = "%s/%s/validator/address"
	Threshold   = 3
)

func NewValidatorInfo(info *mongoentity.ValidatorInfo) *ValidatorInfo {
	v := ValidatorInfo{
		ValidatorInfo: info,
		AddressInfo: map[string]*AddressInfo{
			commontypes.ChainMATIC: {
				Chain: commontypes.ChainMATIC,
			},
			commontypes.ChainKLAY: {
				Chain: commontypes.ChainKLAY,
			},
			commontypes.ChainETH: {
				Chain: commontypes.ChainETH,
			},
		},
	}

	v.initAddressInfo()

	return &v
}

func (v *ValidatorInfo) initAddressInfo() {
	for _, addressInfo := range v.AddressInfo {
		address, err := v.getAddress(addressInfo.Chain)
		if err != nil {
			elog.Error("initAddressInfo", "name", v.Name, "desc", v.Description, "getAddress", err)
		}

		addressInfo.Address = address
	}
}

func (v *ValidatorInfo) getAddress(chain string) (string, error) {
	type Response struct {
		Address string `json:"address"`
	}

	var response Response

	resp, err := util.Get(fmt.Sprintf(addressPath, v.Url, chain), nil, nil)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		return "", err
	}

	return response.Address, nil
}

func (v *ValidatorInfo) RecoverTx(chain, txHash string) error {
	type requestBody struct {
		TxHash string `json:"txHash"`
	}

	type requestResponse struct {
		Result int `json:"result"`
	}

	request := requestBody{
		TxHash: txHash,
	}
	resp, err := util.Post(
		fmt.Sprintf("%s/%s/recover", v.Url, chain),
		request, nil, nil,
	)
	if err != nil {
		return err
	}

	var response requestResponse
	if err = json.Unmarshal([]byte(resp), &response); err != nil {
		return err
	}
	if response.Result != 0 {
		return fmt.Errorf("%d", response.Result)
	}

	return nil
}

func (v *ValidatorInfo) SpeedUpTx(chain, txHash string) error {
	type requestBody struct {
		TxHash string `json:"txHash"`
	}

	type requestResponse struct {
		Result int `json:"result"`
	}

	request := requestBody{
		TxHash: txHash,
	}
	resp, err := util.Post(
		fmt.Sprintf("%s/%s/speedup", v.Url, chain),
		request, nil, nil,
	)
	if err != nil {
		return err
	}

	var response requestResponse
	if err = json.Unmarshal([]byte(resp), &response); err != nil {
		return err
	}
	if response.Result != 0 {
		return fmt.Errorf("%d", response.Result)
	}

	return nil
}

func (v *ValidatorInfo) GetPubKey(chain string) (cryptotypes.PubKey, error) {
	resp, err := util2.GetRetry(context.Background(), fmt.Sprintf("%s/%s/pubkey", v.Url, chain), nil, nil, 5)
	if err != nil {
		return nil, err
	}

	cdc := codec.NewProtoCodec(types.NewInterfaceRegistry(types.WithCosmosRegistry()))
	var pubKey secp256k1.PubKey
	if err = cdc.UnmarshalJSON(resp, &pubKey); err != nil {
		return nil, commontypes.WrapError(string(resp), err)
	}

	return &pubKey, nil
}

func (v *ValidatorInfo) GetSignatureV2(chain, txHash, multiSigAddress string, accountNumber, accountSequence int64) (signing.SignatureV2, error) {
	type requestBody struct {
		Chain           string `json:"chain"`
		TxHash          string `json:"txHash"`
		Address         string `json:"address"`
		AccountNumber   int64  `json:"accountNumber"`
		AccountSequence int64  `json:"accountSequence"`
	}

	bz, err := util2.PostRetry(fmt.Sprintf("%s/sign/multisig", v.Url), requestBody{
		Chain:           chain,
		TxHash:          txHash,
		Address:         multiSigAddress,
		AccountNumber:   accountNumber,
		AccountSequence: accountSequence,
	}, nil, nil, 5)
	if err != nil {
		return signing.SignatureV2{}, err
	}

	cdc := codec.NewProtoCodec(types.NewInterfaceRegistry(types.WithCosmosRegistry()))
	signatures, err := authtx.NewTxConfig(cdc, authtx.DefaultSignModes).UnmarshalSignatureJSON(bz)
	if err != nil {
		return signing.SignatureV2{}, err
	}

	return signatures[0], nil
}

type BridgeStatus int

const (
	BridgeUnknown BridgeStatus = iota
	BridgeRequest
	BridgeSubmission
	BridgeConfirm
	BridgeExecute
	BridgeExecuteFailure
)

func (status BridgeStatus) String() string {
	switch status {
	case BridgeRequest:
		return "Request"
	case BridgeSubmission:
		return "Submission"
	case BridgeConfirm:
		return "Confirm"
	case BridgeExecute:
		return "Execute"
	}
	return "Unknown"
}

type PendingTx struct {
	Time   time.Time `json:"time"`
	Chain  string    `json:"chain"`
	TxHash string    `json:"txHash"`

	ValidatorInfo    *ValidatorInfo    `json:"-"`
	ValidatorSummary *ValidatorSummary `json:"validator,omitempty"`

	SubmitTransaction     *InputSubmitTransaction     `json:"-"`
	SubmitTransactionData *InputSubmitTransactionData `json:"-"`
}

type InputSubmitTransaction struct {
	TxHash      common.Hash
	Destination common.Address
	Value       *big.Int
	Data        common.Bytes
}

type InputSubmitTransactionData struct {
	FromChainName string
	From          common.Bytes
	To            common.Bytes
	ToAddr        common.Address

	Token     common.Bytes
	TokenAddr common.Address

	TxInfo    []common.Hash
	TokenInfo []*big.Int
}
