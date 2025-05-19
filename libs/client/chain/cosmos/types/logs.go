package types

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/curtis0505/bridge/libs/logger/v2"
)

type ReceiptLogs []ReceiptLog

type ReceiptLog struct {
	LogType         string
	EventName       string
	ContractAddress string
	Logs            map[string]string
}

func (logs ReceiptLogs) Len() int { return len(logs) }

func (logs ReceiptLogs) GetLogsByEvent(events ...string) ReceiptLogs {
	receiptLogs := ReceiptLogs{}
	for _, log := range logs {
		for _, event := range events {
			if log.EventName == event {
				receiptLogs = append(receiptLogs, log)
			}
		}
	}

	return receiptLogs
}

func (logs ReceiptLogs) GetLogByEvent(event string) ReceiptLog {
	for _, log := range logs {
		if log.EventName == event {
			return log
		}
	}

	return ReceiptLog{}
}

func (logs ReceiptLogs) GetLogsByType(logTypes ...string) ReceiptLogs {
	receiptLogs := ReceiptLogs{}
	for _, log := range logs {
		for _, logType := range logTypes {
			if log.LogType == logType {
				receiptLogs = append(receiptLogs, log)
			}
		}
	}

	return receiptLogs
}

func (logs ReceiptLogs) GetLogByType(logType string) ReceiptLog {
	for _, log := range logs {
		if log.LogType == logType {
			return log
		}
	}

	return ReceiptLog{}
}

func (log ReceiptLog) Unmarshal(value interface{}) error {
	defer func() {
		if v := recover(); v != nil {
			logger.Error("Unmarshal", logger.BuildLogInput().WithError(fmt.Errorf("panic recovered: ReceiptLog")).WithData("Unmarshal", v))
		}
	}()

	b, _ := json.Marshal(log.Logs)
	return json.Unmarshal(b, &value)
}

func ParseLogs(txResponse *types.TxResponse) ReceiptLogs {
	receiptLogs := ReceiptLogs{}

	for _, event := range txResponse.Events {
		logs := make(map[string]string)

		for _, attributes := range event.GetAttributes() {
			key, errKey := base64.StdEncoding.DecodeString(attributes.Key)
			value, errValue := base64.StdEncoding.DecodeString(attributes.Value)
			if errKey == nil && errValue == nil {
				logs[string(key)] = string(value)
			} else {
				logs[attributes.Key] = attributes.Value
			}
		}

		if event.Type == wasmtypes.ModuleName {
			const contractAddressKey = "_contract_address"
			const actionKey = "action"

			contractAddress := logs[contractAddressKey]
			eventName := logs[actionKey]
			delete(logs, contractAddressKey)
			delete(logs, actionKey)

			receiptLog := ReceiptLog{
				LogType:         event.Type,
				EventName:       eventName,
				ContractAddress: contractAddress,
				Logs:            logs,
			}

			receiptLogs = append(receiptLogs, receiptLog)
		} else {
			receiptLog := ReceiptLog{
				LogType:   "module",
				EventName: event.Type,
				Logs:      logs,
			}

			receiptLogs = append(receiptLogs, receiptLog)
		}
	}

	return receiptLogs
}
