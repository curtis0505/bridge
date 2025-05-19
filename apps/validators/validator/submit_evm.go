package validator

import (
	"context"
	"fmt"
	"github.com/curtis0505/bridge/libs/common"
	commontypes "github.com/curtis0505/bridge/libs/types"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	"github.com/curtis0505/bridge/libs/types/bridge/abi"
	"math/big"
)

type Submission struct {
	TxHash      string
	Destination string
	Amount      *big.Int
	Proof       []byte
	Data        []byte
}

type ChainInfo struct {
	Client    interface{}
	ChainName string
}

type Method struct {
	EventMethod   string
	ExecuteMethod string
}

type PendingRequest struct {
	Chain  string `json:"chain"`
	From   string `json:"from"`
	TxHash string `json:"txHash"`
}

type TransactionHistory struct {
	ChainName  string
	TxHash     string
	Nonce      uint64
	GasPrice   *big.Int
	GasTip     *big.Int
	Status     int
	Submission Submission
	ToAddress  string
}

func (transactionHistory *TransactionHistory) ConvertPublicData() *PublicTransactionHistory {
	return &PublicTransactionHistory{
		ChainName: transactionHistory.ChainName,
		Nonce:     transactionHistory.Nonce,
		GasPrice:  transactionHistory.GasPrice,
		GasTip:    transactionHistory.GasTip,
	}
}

type PublicTransactionHistory struct {
	ChainName string
	Nonce     uint64
	GasPrice  *big.Int
	GasTip    *big.Int
}

func (p *Validator) MintToEVM(commonLog bridgetypes.CommonEventLog) error {
	err := p.SubmitTransaction(
		Method{
			EventMethod:   bridgetypes.EventNameDeposit,
			ExecuteMethod: bridgetypes.MethodNameBridgeBurn,
		},
		commonLog,
	)
	if err != nil {
		p.logger.Error("SubmitTransaction(Mint)", err, "commonLog", commonLog)
		return err
	}

	return nil
}

// npt 체인 마이그레이션을 위해 klay vault 에서 base vault 로 이동
func (p *Validator) NptFromKlayToBase(commonLog bridgetypes.CommonEventLog) error {
	err := p.SubmitTransaction(
		Method{
			EventMethod:   bridgetypes.EventNameDeposit,
			ExecuteMethod: bridgetypes.MethodNameBridgeWithdraw,
		},
		commonLog,
	)
	if err != nil {
		p.logger.Error("SubmitTransaction(Mint)", err, "commonLog", commonLog)
		return err
	}

	return nil
}

func (p *Validator) WithdrawToEVM(commonLog bridgetypes.CommonEventLog) error {
	err := p.SubmitTransaction(
		Method{
			EventMethod:   bridgetypes.EventNameBurn,
			ExecuteMethod: bridgetypes.MethodNameBridgeWithdraw,
		},
		commonLog,
	)
	if err != nil {
		p.logger.Error("SubmitTransaction(Withdraw)", err, "commonLog", commonLog)
		return fmt.Errorf("SubmitTransaction: %w", err)
	}

	return nil
}

func (p *Validator) SubmitTransaction(contractMethod Method, commonLog bridgetypes.CommonEventLog) error {
	// getSubTransactionData
	data, err := p.getSubTransactionData(
		commonLog,
		contractMethod.ExecuteMethod,
	)
	if err != nil {
		p.logger.Error("SubmitTransactionData", err, "fromChain", commonLog.FromChainName, "txHash", commonLog.TxHash)
		return err
	}

	//proof
	proof, err := p.getProof(commonLog, contractMethod)
	if err != nil {
		p.logger.Error("getProof", err, "fromChain", commonLog.FromChainName, "txHash", commonLog.TxHash)
		return err
	}

	// SendSubmitTransaction
	option, err := p.client.GetTransactionOption(context.Background(), commonLog.ToChainName, p.Account[commonLog.ToChainName].Address)
	if err != nil {
		p.logger.Error("GetTransactionOption", err, "fromChain", commonLog.FromChainName, "txHash", commonLog.TxHash)
		return err
	}

	contractAddress, err := p.Contract.GetChain(commontypes.Chain(commonLog.ToChainName))
	if err != nil {
		return err
	}

	var toAddress string
	switch commonLog.Version {
	case bridgetypes.VersionBridge:
		toAddress = contractAddress.MultiSigWallet
	case bridgetypes.VersionRestakeBridge:
		toAddress = contractAddress.RestakeMultiSigWallet
	default:
		return fmt.Errorf("invalid version")
	}

	submitTx, err := p.SendSubmitTransaction(
		commonLog.ToChainName,
		toAddress,
		Submission{
			TxHash:      commonLog.TxHash,
			Destination: commonLog.Destination,
			Amount:      big.NewInt(0), // MultiSigWallet 에서 다른 컨트랙트 call 할때 보내는 수량 ex) _destination.call{value : _value}(_data)
			Proof:       proof,
			Data:        data,
		},
		option,
	)
	if err != nil {
		p.logger.Error("sendSubmitTransaction", err, "fromChain", commonLog.FromChainName, "txHash", commonLog.TxHash, "option", option)
		return err
	}

	p.logger.Info(
		"success", "SubmitTransaction",
		"contractAddress", toAddress,
		"toChainName", commonLog.ToChainName,
		"eventMethod", contractMethod.EventMethod,
		"txHash", commonLog.TxHash,
		"fromChainName", commonLog.FromChainName,
		"executeMethod", contractMethod.ExecuteMethod,
		"submit txHash", submitTx,
		"option", option,
	)

	return nil
}

func (p *Validator) SendSubmitTransaction(chainName string, toAddress string, submit Submission, option *commontypes.TransactionOption) (string, error) {
	submitTransactionData, err := commontypes.PackAbi(
		chainName,
		abi.GetAbiToMap(abi.MultiSigWalletAbi),
		"submitTransaction",
		common.HexToHash(submit.TxHash),
		common.HexToAddress(chainName, submit.Destination),
		submit.Amount,
		submit.Proof,
		submit.Data,
	)
	if err != nil {
		p.logger.Error("method", "Pack", "err", err, "packMethod", "submitTransaction")
		return "", err
	}

	signedTx, err := p.getSignedTx(
		chainName,
		&commontypes.RequestTransaction{
			From:      p.Account[chainName].Address,
			To:        toAddress,
			Nonce:     option.Nonce,
			GasPrice:  option.GasPrice,
			GasFeeCap: option.GasFeeCap,
			GasTipCap: option.GasTipCap,
			GasLimit:  commontypes.GasLimit,
			Value:     big.NewInt(0),
			Data:      submitTransactionData,
		},
	)
	if err != nil {
		p.logger.Error("method", "getSignedTx", "err", err)
		return "", err
	}

	// TODO :
	//body := PendingRequest{chainName, address, signedTx.TxHash()}
	//_, err = util.Post(p.cfg.Server.Monitor+"/validator/pending/add", body, nil, nil)
	//if err != nil {
	//	p.logger.Warn("sendSubmitTransaction", "Post", err)
	//	//return "", err
	//}

	//config := conf.ClientConfig{
	//	Klay: conf.NetworkInfo{
	//		Chain:               "MATIC",
	//		URL:                 []string{"ws://node2.dq.neopin.pmang.cloud:16200"},
	//		FinalizedBlockCount: 0,
	//		ContractInfo:        conf.ContractInfo{},
	//	},
	//	Polygon:     conf.NetworkInfo{},
	//	Ethereum:    conf.NetworkInfo{},
	//	Finschia:    conf.NetworkInfo{},
	//	AccountInfo: commontypes.AccountConfig{},
	//}
	//
	//maticClient := ether.NewClient("ws://node2.dq.neopin.pmang.cloud:16200", config, "MATIC")
	//etx, _ := signedTx.EthereumTransaction()
	//err = maticClient.SendTransaction(context.Background(), etx)
	//if err != nil {
	//	return "", err
	//}

	_, err = p.client.RawSendTxAsyncByTx(context.Background(), chainName, signedTx)
	if err != nil {
		p.logger.Error("method", "SendTx", "err", err)
		return "", err
	}

	usedNonce, err := signedTx.Nonce()
	if err != nil {
		p.logger.Error("method", "tx.Nonce()", "err", err)
		return "", err
	}

	p.PendingTransaction.Set(&TransactionHistory{
		ChainName:  chainName,
		TxHash:     signedTx.TxHash(),
		GasPrice:   signedTx.GasPrice(),
		GasTip:     signedTx.GasTipCap(),
		Nonce:      usedNonce,
		Status:     0,
		Submission: submit,
		ToAddress:  toAddress,
	})

	return signedTx.TxHash(), nil
}

func (p *Validator) getSignedTx(chainName string, request *commontypes.RequestTransaction) (*commontypes.Transaction, error) {
	transaction, err := p.client.GetTransactionData(chainName, request)
	if err != nil {
		p.logger.Error("method", "GetTransaction", "err", err)
		return nil, err
	}

	chainId, err := p.client.GetChainID(context.Background(), chainName)
	if err != nil {
		p.logger.Error("sendSubmitTransaction", "GetChainId", err)
		return nil, err
	}

	signedTx, err := p.Account[chainName].Sign(transaction, chainId)
	if err != nil {
		p.logger.Error("sendSubmitTransaction", "Sign", err)
		return nil, err
	}
	return signedTx, nil

}

func (p *Validator) getSubTransactionData(commonLog bridgetypes.CommonEventLog, executeMethod string) ([]byte, error) {
	var data []byte
	var err error

	switch executeMethod {
	case bridgetypes.MethodNameBridgeBurn:
		data, err = p.mintData(commonLog)
		if err != nil {
			p.logger.Error("method", "mintData", "err", err)
			return data, err
		}

	case bridgetypes.MethodNameBridgeWithdraw:
		data, err = p.withdrawData(commonLog)
		if err != nil {
			p.logger.Error("method", "burnData", "err", err)
			return data, err
		}
	}

	return data, nil
}

func (p *Validator) mintData(commonLog bridgetypes.CommonEventLog) ([]byte, error) {
	ctx := context.Background()
	minterAbi := abi.GetAbiToMap(abi.MinterAbi)

	msg, err := p.client.CallMsg(ctx,
		commonLog.ToChainName,
		"",
		commonLog.Destination,
		"getChainId",
		minterAbi,
		commonLog.FromChainName,
	)
	if err != nil {
		p.logger.Error("method", "CallMsg", "err", err, "callMethodName", "getChainId(fromChainId)")
		return nil, err
	}

	fromChainID := msg[0].([32]byte)

	msg, err = p.client.CallMsg(ctx,
		commonLog.ToChainName,
		"",
		commonLog.Destination,
		"isSupportChain",
		minterAbi,
		fromChainID,
	)
	if err != nil {
		p.logger.Error("method", "CallMsg", "err", err, "callMethodName", "isSupportChain")
		return nil, err
	}

	isSupportChain := msg[0].(bool)
	if !isSupportChain {
		p.logger.Error("mintData", "IsSupportChain", "invalid chain from "+commonLog.FromChainName)
		return nil, fmt.Errorf("invalid chain from " + commonLog.FromChainName)
	}

	msg, err = p.client.CallMsg(ctx,
		commonLog.ToChainName,
		"",
		commonLog.Destination,
		"getChainId",
		minterAbi,
		commonLog.ToChainName,
	)
	if err != nil {
		p.logger.Error("method", "CallMsg", "err", err, "callMethodName", "getChainId(toChainId")
		return nil, err
	}

	toChainID := msg[0].([32]byte)

	// 컨트랙트에서 코인의 경우 소스의 경우 address(0) 으로 관리하지만
	// 0x0000000000000000000000000000000000000001 로 저장되어 있기에 예외처리 하드코딩
	if commonLog.FromTokenAddr.String() == bridgetypes.ZeroAddress {
		commonLog.FromTokenAddr = common.HexToAddress(commonLog.FromChainName, bridgetypes.CoinAddress)
	}

	txInfo := [][32]byte{toChainID, common.HexToHash(commonLog.TxHash)}

	tokenInfo := make([]*big.Int, 0)
	tokenInfo = append(tokenInfo, commonLog.Amount)
	tokenInfo = append(tokenInfo, commonLog.Decimal)

	data, err := commontypes.PackAbi(
		commonLog.ToChainName,
		abi.GetAbiToMap(abi.MinterAbi),
		"mint",
		commonLog.FromChainName,
		commonLog.From.Bytes(),
		commonLog.ToAddr,
		commonLog.FromTokenAddr.Bytes(),
		txInfo,
		tokenInfo,
	)
	if err != nil {
		p.logger.Error("method", "Pack", "err", err)
		return nil, err
	}

	return data, nil
}

func (p *Validator) withdrawData(commonLog bridgetypes.CommonEventLog) ([]byte, error) {
	ctx := context.Background()

	vaultABI := abi.GetAbiToMap(abi.VaultAbi)

	// brun 할 vault 를 기준으로 from chain 이 유효한지 체크
	msg, err := p.client.CallMsg(ctx,
		commonLog.ToChainName,
		"",
		commonLog.Destination,
		"getChainId",
		vaultABI,
		commonLog.FromChainName,
	)
	if err != nil {
		p.logger.Error("method", "CallMsg", "err", err, "callMethodName", "getChainId")
		return nil, err
	}

	fromChainID := msg[0].([32]byte)

	msg, err = p.client.CallMsg(ctx,
		commonLog.ToChainName,
		"",
		commonLog.Destination,
		"isSupportChain",
		vaultABI,
		fromChainID,
	)
	if err != nil {
		p.logger.Error("method", "CallMsg", "err", err, "callMethodName", "isSupportChain")
		return nil, err
	}

	isSupportChain := msg[0].(bool)

	if !isSupportChain {
		p.logger.Error("burnData", "IsSupportChain", fmt.Errorf("invalid chain from "+commonLog.FromChainName))
		return nil, fmt.Errorf("invalid chain from " + commonLog.FromChainName)
	}

	// Data 구하기
	msg, err = p.client.CallMsg(ctx,
		commonLog.ToChainName,
		"",
		commonLog.Destination,
		"getChainId",
		vaultABI,
		commonLog.ToChainName,
	)
	if err != nil {
		p.logger.Error("method", "CallMsg", "err", err, "callMethodName", "getChainId")
		return nil, err
	}

	toChainID := msg[0].([32]byte)

	if commonLog.ToTokenAddr.String() == bridgetypes.CoinAddress {
		commonLog.ToTokenAddr = common.HexToAddress(commonLog.ToChainName, bridgetypes.ZeroAddress)
	}

	txInfo := [][32]byte{toChainID, common.HexToHash(commonLog.TxHash)}

	tokenInfo := make([]*big.Int, 0)
	tokenInfo = append(tokenInfo, commonLog.Amount)
	tokenInfo = append(tokenInfo, commonLog.Decimal)

	data, err := commontypes.PackAbi(
		commonLog.ToChainName,
		abi.GetAbiToMap(abi.VaultAbi),
		"withdraw",
		commonLog.FromChainName,
		commonLog.From.Bytes(),
		commonLog.ToAddr,
		commonLog.ToTokenAddr,
		txInfo,
		tokenInfo,
	)
	if err != nil {
		p.logger.Error("method", "abi.Pack", "err", err, "packMethod", "withdraw")
		return nil, err
	}

	return data, nil
}
