package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/curtis0505/bridge/apps/validators/conf"
	"github.com/curtis0505/bridge/apps/validators/validator"
	"github.com/curtis0505/bridge/libs/client/chain"
	"github.com/curtis0505/bridge/libs/client/chain/cosmos/types"
	cosmoscommon "github.com/curtis0505/bridge/libs/common/cosmos"
	"github.com/curtis0505/bridge/libs/logger"
	commontypes "github.com/curtis0505/bridge/libs/types"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	"github.com/gin-gonic/gin"
	"github.com/kaiachain/kaia/accounts/keystore"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strings"
)

type Controller struct {
	Config    *conf.Config
	Client    *chain.Client
	Validator *validator.Validator
}

func NewController(config *conf.Config, client *chain.Client, validator *validator.Validator) (*Controller, error) {
	controller := Controller{
		Config:    config,
		Client:    client,
		Validator: validator,
	}

	return &controller, nil
}

func (controller *Controller) HealthCheck(c *gin.Context) {
	Response200(c, NewBaseResponse(Success))
}

func (controller *Controller) ValidatorAddressByChain(c *gin.Context) {
	chainSymbol := c.Param("chain")

	if controller.Validator.Account[chainSymbol] == nil {
		Response404(c)
		return
	}

	Response200(c, ValidatorAddressByChainResponse{
		BaseResponse: NewBaseResponse(Success),
		Address:      controller.Validator.Account[chainSymbol].Address,
	})
}

func (controller *Controller) GetValidatorInfos(c *gin.Context) {
	var validatorInfo []ValidatorInfo

	for _, chain := range controller.Client.GetChains() {
		validatorInfo = append(validatorInfo, ValidatorInfo{
			ChainSymbol: chain,
			Address:     controller.Validator.Account[chain].Address,
		})
	}

	Response200(c, ValidatorAddressResponse{
		BaseResponse: NewBaseResponse(Success),
		Validator:    validatorInfo,
	})
}

func (controller *Controller) RecoverTransaction(c *gin.Context) {
	req := RetryTransactionRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response422(c, err)
		return
	}

	logger.Info("API.RecoverTransaction", "req", req)

	chainSymbol := c.Param("chain")
	chainSymbol = strings.ToUpper(chainSymbol)

	tx, isPending, err := controller.Client.GetTransactionWithReceipt(context.Background(), chainSymbol, req.TxHash)
	if err != nil {
		ResponseException(c, InvalidTransaction, err)
		return
	}

	if isPending {
		ResponseException(c, PendingTransaction, nil)
		return
	}

	for _, log := range tx.Receipt().Logs() {
		if log.TxHash() == "" {
			log = log.SetTxHash(tx.TxHash())
		}

		if bridgetypes.GetEventType(log.EventName) != bridgetypes.EventNameDeposit &&
			bridgetypes.GetEventType(log.EventName) != bridgetypes.EventNameDepositRestakeToken &&
			bridgetypes.GetEventType(log.EventName) != bridgetypes.EventNameBurn {
			continue
		}

		err = controller.Validator.LogHandler(log)
		if err != nil {
			if err.Error() == bridgetypes.TransactionAlreadyExecuted {
				ResponseException(c, TransactionAlreadyExecuted, nil)
				return
			}
			logger.Warn("RecoverTransaction", "txHash", log.TxHash(), "event", log.EventName, "LogHandler", err)
		}
	}

	Response200(c, NewBaseResponse(Success))
}

func (controller *Controller) GetCacheTransaction(c *gin.Context) {
	transactionHistory := map[string]*validator.PublicTransactionHistory{}

	for _, history := range controller.Validator.PendingTransaction.List() {
		transactionHistory[history.TxHash] = history.ConvertPublicData()
	}

	Response200(c, CacheTransactionResponse{
		BaseResponse:    NewBaseResponse(Success),
		TransactionList: transactionHistory,
	})
}

func (controller *Controller) SpeedUpTransaction(c *gin.Context) {
	req := SpeedUpTransactionRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response422(c, err)
		return
	}

	logger.Info("API.SpeedUpTransaction", "req", req)

	chainSymbol := c.Param("chain")
	chainSymbol = strings.ToUpper(chainSymbol)

	history, err := controller.Validator.PendingTransaction.Get(req.TxHash)
	if history == nil || err != nil {
		logger.Error("API.SpeedUpTransaction", "history", "not found")
		Response404(c)
		return
	}

	if history.ChainName != chainSymbol {
		logger.Error("API.SpeedUpTransaction", "chainName", chainSymbol, "history.ChainName", history.ChainName)
		Response404(c)
		return
	}

	_, isPending, _ := controller.Client.GetTransactionWithReceipt(context.Background(), chainSymbol, req.TxHash)

	if isPending == false {
		ResponseException(c, TransactionAlreadyExecuted, nil)
		return
	}

	option, err := controller.Client.GetTransactionOption(context.Background(), history.ChainName, controller.Validator.Account[history.ChainName].Address)
	if err != nil {
		logger.Error("API.SpeedUpTransaction", "GetTransactionOption", err, "history", history)
		Response404(c)
		return
	}

	// nonce 는 기존에 사용했던 nonce 사용
	// GasPrice, GasTip 은 데이터 존재하면 입력받은 데이터 넣고, 아니면 네트워크에서 요청받은 데이터 사용
	option.Nonce = history.Nonce

	if req.GasPrice != nil {
		option.GasPrice = req.GasPrice
	}

	if req.GasFeeCap != nil {
		option.GasFeeCap = req.GasFeeCap
	}

	if req.GasTip != nil {
		option.GasTipCap = req.GasTip
	}

	submitTx, err := controller.Validator.SendSubmitTransaction(history.ChainName, history.ToAddress, history.Submission, option)
	if err != nil {
		logger.Error("API.SpeedUpTransaction", "SendSubmitTransaction", err, "history", history.TxHash)
		ResponseException(c, InvalidTransaction, err)
		return
	}

	logger.Info("API.SpeedUpTransaction", "tx", submitTx)

	Response200(c, NewBaseResponse(Success))
}

func (controller *Controller) CancelTransaction(c *gin.Context) {
	req := CancelTransactionRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response422(c, err)
		return
	}

	logger.Info("API.CancelTransaction", "req", req)

	chainSymbol := c.Param("chain")
	chainSymbol = strings.ToUpper(chainSymbol)

	ctx := context.Background()

	nonce, err := controller.Client.NonceAt(ctx, chainSymbol, controller.Validator.Account[chainSymbol].Address)
	if err != nil {
		logger.Error("API.CancelTransaction", "method", "NonceAt", "err", err,
			"chain", chainSymbol, "address", controller.Validator.Account[chainSymbol].Address)
		ResponseException(c, InvalidTransaction, err)
		return
	}

	logger.Info("API.CancelTransaction", "nonce", nonce, "req.Nonce", req)

	if req.Nonce < nonce {
		ResponseException(c, TransactionAlreadyExecuted, nil)
		return
	}

	transaction, err := controller.Client.GetTransactionData(chainSymbol, &commontypes.RequestTransaction{
		From:      controller.Validator.Account[chainSymbol].Address,
		To:        "0x0000000000000000000000000000000000000000",
		Nonce:     req.Nonce,
		GasPrice:  req.GasPrice,
		GasFeeCap: req.GasPrice,
		GasTipCap: req.GasTip,
		GasLimit:  uint64(commontypes.GasLimit),
		Value:     big.NewInt(0),
		Data:      []byte{},
	})
	if err != nil {
		logger.Error("API.CancelTransaction", "method", "GetTransactionData", "err", err)
		ResponseException(c, InvalidTransaction, err)
		return
	}

	chainID, err := controller.Client.GetChainID(context.Background(), chainSymbol)
	if err != nil {
		logger.Error("API.CancelTransaction", "method", "GetChainID", "err", err, "chain", chainSymbol)
		ResponseException(c, InvalidTransaction, err)
		return
	}

	signedTx, err := controller.Validator.Account[chainSymbol].Sign(transaction, chainID)
	if err != nil {
		logger.Error("API.CancelTransaction", "method", "Sign", "err", err, "chainID", chainID)
		ResponseException(c, InvalidTransaction, err)
		return
	}

	logger.Info("API.CancelTransaction", "tx", signedTx.TxHash())

	Response200(c, NewBaseResponse(Success))
}

func (controller *Controller) GenerateKey(c *gin.Context) {

	logger.Info("API.GenerateKey")
	req := GenerateKeyRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response422(c, err)
		return
	}

	// TODO key 생성시 체인 체크하는 부분 남겨놓을지 체크
	chainSymbol := c.Param("chain")
	chainName := ""
	switch strings.ToUpper(chainSymbol) {
	case commontypes.ChainKLAY:
		chainName = commontypes.ChainKLAY
	case commontypes.ChainMATIC:
		chainName = commontypes.ChainMATIC
	case commontypes.ChainETH:
		chainName = commontypes.ChainETH
	default:
		Response404(c)
		return
	}

	if req.PassName == "" {
		Response400(c)
		return
	}

	if chainName == "" {
		Response404(c)
		return
	}

	password := controller.Config.Account.KeystoreInfo.GetPassword()

	d, err := ioutil.TempDir("", "bridge-ks")
	if err != nil {
		ResponseException(c, Failed, err)
	}
	newKs := keystore.NewKeyStore(d, keystore.StandardScryptN, keystore.StandardScryptP)
	defer os.RemoveAll(d)
	a, err := newKs.NewAccount(password)
	if err != nil {
		ResponseException(c, Failed, err)
	}
	_, err = os.Stat(a.URL.Path)
	if err != nil {
		ResponseException(c, Failed, err)
	}

	if !newKs.HasAddress(a.Address) {
		ResponseException(c, Failed, errors.New("address not found"))
	}

	keyJson, err := newKs.Export(a, password, password)
	if err != nil {
		ResponseException(c, Failed, err)
	}
	fmt.Println("v4 keyJson:", string(keyJson), password)

	key, err := keystore.DecryptKey(keyJson, password)
	if err != nil {
		ResponseException(c, Failed, err)
	}
	fmt.Println("key addr:", key.GetAddress().String())

	keyJsonV3, err := keystore.EncryptKeyV3(key, password, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		ResponseException(c, Failed, err)
	}
	fmt.Println("v3 keyjson:", string(keyJsonV3))

	Response200(c, GenerateKeyResponse{
		BaseResponse: NewBaseResponse(Success),
		Keystore:     string(keyJsonV3),
	})
}

func (controller *Controller) PubKey(c *gin.Context) {
	JSONMarshal(c, http.StatusOK, func() ([]byte, error) {
		cdc := codec.NewProtoCodec(types.NewInterfaceRegistry(types.WithCosmosRegistry()))
		account := controller.Validator.Account[c.Param("chain")]
		if controller.Validator.Account[c.Param("chain")].Type == commontypes.KMS {
			pubKeyUncompressed, err := account.KMS.PublicKey(c)
			if err != nil {
				return nil, err
			} else {
				pubKey, err := cosmoscommon.NewPubKey(pubKeyUncompressed)
				if err != nil {
					return nil, err
				}

				return cdc.MarshalJSON(pubKey)
			}
		} else {
			return cdc.MarshalJSON(controller.Validator.Account[c.Param("chain")].Secp256k1().PubKey())
		}
	})
}

func (controller *Controller) SignMultiSig(c *gin.Context) {
	var sigs []signing.SignatureV2
	var sig signing.SignatureV2
	req := MultiSigRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		JSONMarshal(c, http.StatusBadRequest, nil)
		return
	}

	tx, isPending, err := controller.Client.GetTransactionWithReceipt(c, req.Chain, req.TxHash)
	if err != nil || isPending {
		JSONMarshal(c, http.StatusBadRequest, nil)
		return
	}

	for _, log := range tx.Receipt().Logs() {
		if log.TxHash() == "" {
			log = log.SetTxHash(req.TxHash)
		}

		//tx method check
		var commonLog bridgetypes.CommonEventLog
		switch bridgetypes.GetEventType(log.EventName) {
		case bridgetypes.EventNameDeposit:
			commonLog, err = bridgetypes.GetDepositCommonLog(log, controller.Validator.Contract)
			if err != nil {
				JSONMarshal(c, http.StatusBadRequest, nil)
				return
			}
		case bridgetypes.EventNameBurn:
			commonLog, err = bridgetypes.GetBurnCommonLog(log, controller.Validator.Contract)
			if err != nil {
				JSONMarshal(c, http.StatusBadRequest, nil)
				return
			}
		}

		//to, from chain cosmos check
		if commontypes.GetChainType(commonLog.ToChainName) == commontypes.ChainTypeCOSMOS {
			switch bridgetypes.GetEventType(log.EventName) {
			case bridgetypes.EventNameDeposit:
				sig, err = controller.Validator.MintToCOSMOS(commonLog, req.Address)
			case bridgetypes.EventNameBurn:
				sig, err = controller.Validator.WithdrawToCOSMOS(commonLog, req.Address)
			}
		}
	}
	sigs = append(sigs, sig)

	JSONMarshal(c, http.StatusOK, func() ([]byte, error) {
		cdc := codec.NewProtoCodec(types.NewInterfaceRegistry(types.WithCosmosRegistry()))
		return authtx.NewTxConfig(cdc, authtx.DefaultSignModes).MarshalSignatureJSON(sigs)
	})
}

func JSONMarshal(ctx *gin.Context, code int, marshal func() ([]byte, error)) {
	bz, err := marshal()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	} else {
		ctx.Data(code, http.DetectContentType(bz), bz)
	}
}
