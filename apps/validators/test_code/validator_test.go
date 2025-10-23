package test_code

import (
	"context"
	"github.com/curtis0505/bridge/apps/validators/conf"
	"github.com/curtis0505/bridge/apps/validators/validator"
	"github.com/curtis0505/bridge/libs/client/chain"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	"testing"
)

func TestNewValidator(t *testing.T) {
	config, err := conf.NewConfig("../conf/config2.toml")
	if err != nil {
		panic(err)
	}

	config2, err := conf.NewConfig("../conf/config3.toml")
	if err != nil {
		panic(err)
	}

	logger.InitLog(config.Log)
	logger.SetAppName("validator")

	clientInstance := chain.NewClientByConfig(config.Client)

	account, err := validator.NewAccount(clientInstance, config)
	if err != nil {
		panic(err)
	}

	validatorInstance, err := validator.New(clientInstance, *config, account)
	if err != nil {
		panic(err)
	}

	account2, err := validator.NewAccount(clientInstance, config2)
	if err != nil {
		panic(err)
	}

	validatorInstance2, err := validator.New(clientInstance, *config2, account2)
	if err != nil {
		panic(err)
	}

	t.Log("good", validatorInstance.Account)

	txHash := "0xa95eb2843838c156a94e54ccf565680114ea9329cd4dc545c3b522eeadf31626"
	//txInfo, err := clientInstance.Ethereum.Client().GetTransactionReceipt(context.Background(), txHash)
	//if err != nil {
	//	t.Error("GetTransactionReceipt", err)
	//	return
	//}
	//for _, log := range txInfo.Logs {
	//	commonLog := commontypes.NewLog(log, "ETH")
	//	err := validatorInstance.LogHandler(commonLog)
	//	if err != nil {
	//		t.Error("LogHandler", err)
	//		continue
	//	}
	//
	//	err = validatorInstance.Deposit(commonLog)
	//	if err != nil {
	//		t.Error("Deposit", err)
	//		return
	//	}
	//}
	//
	//for _, log := range txInfo.Logs {
	//	t.Log("log", log)
	//	commonLog := commontypes.NewLog(log, "ETH")
	//	err := validatorInstance.Deposit(commonLog)
	//	if err != nil {
	//		t.Error("Deposit", err)
	//		continue
	//	}
	//}

	tx, isPending, err := clientInstance.GetTransactionWithReceipt(context.Background(), "ETH", txHash)
	if err != nil {
		t.Error("GetTransactionByHash", err)
		return
	}

	if isPending {
		t.Error("GetTransactionByHash", "pending")
		return
	}

	for _, log := range tx.Receipt().Logs() {
		if log.TxHash() == "" {
			log = log.SetTxHash(tx.TxHash())
		}

		if bridgetypes.GetEventType(log.EventName) != bridgetypes.EventNameDeposit &&
			bridgetypes.GetEventType(log.EventName) != bridgetypes.EventNameBurn {
			continue
		}

		err = validatorInstance.LogHandler(log)
		if err != nil {
			if err.Error() == bridgetypes.TransactionAlreadyExecuted {
				t.Log("TransactionAlreadyExecuted")
				return
			}
			t.Log("RecoverTransaction", "txHash", log.TxHash(), "event", log.EventName, "LogHandler", err)
		}

		err = validatorInstance2.LogHandler(log)
		if err != nil {
			if err.Error() == bridgetypes.TransactionAlreadyExecuted {
				t.Log("TransactionAlreadyExecuted")
				return
			}
			t.Log("RecoverTransaction", "txHash", log.TxHash(), "event", log.EventName, "LogHandler", err)
		}
	}

	t.Log("end", txHash)

}

func TestNewValidator2(t *testing.T) {
	config, err := conf.NewConfig("../conf/config2.toml")
	if err != nil {
		panic(err)
	}

	config2, err := conf.NewConfig("../conf/config3.toml")
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

	logger.InitLog(config.Log)
	logger.SetAppName("validator")

	clientInstance := chain.NewClientByConfig(config.Client)

	account, err := validator.NewAccount(clientInstance, config)
	if err != nil {
		panic(err)
	}

	validatorInstance, err := validator.New(clientInstance, *config, account)
	if err != nil {
		panic(err)
	}

	account2, err := validator.NewAccount(clientInstance, config2)
	if err != nil {
		panic(err)
	}

	validatorInstance2, err := validator.New(clientInstance, *config2, account2)
	if err != nil {
		panic(err)
	}

	t.Log("good", validatorInstance.Account)

	chain := "KLAY"
	txHash := "0x38fdfb08c21bc5fd62fe03d1fc3d22d0671adda59592661ae64a59423d15af30"

	tx, isPending, err := clientInstance.GetTransactionWithReceipt(context.Background(), chain, txHash)
	if err != nil {
		t.Error("GetTransactionByHash", err)
		return
	}

	if isPending {
		t.Error("GetTransactionByHash", "pending")
		return
	}

	for _, log := range tx.Receipt().Logs() {
		if log.TxHash() == "" {
			log = log.SetTxHash(tx.TxHash())
		}

		if bridgetypes.GetEventType(log.EventName) != bridgetypes.EventNameDeposit &&
			bridgetypes.GetEventType(log.EventName) != bridgetypes.EventNameBurn {
			continue
		}

		err = validatorInstance.LogHandler(log)
		if err != nil {
			if err.Error() == bridgetypes.TransactionAlreadyExecuted {
				t.Log("TransactionAlreadyExecuted")
				return
			}
			t.Log("RecoverTransaction", "txHash", log.TxHash(), "event", log.EventName, "LogHandler", err)
		}

		err = validatorInstance2.LogHandler(log)
		if err != nil {
			if err.Error() == bridgetypes.TransactionAlreadyExecuted {
				t.Log("TransactionAlreadyExecuted")
				return
			}
			t.Log("RecoverTransaction", "txHash", log.TxHash(), "event", log.EventName, "LogHandler", err)
		}
	}

	t.Log("end", txHash)

}

// 0x4e7744623e344dcfb1b714a614088b67fa6b780e28d5d8bd681ee32a047dffb8
//
