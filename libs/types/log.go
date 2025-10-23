package types

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	baseabi "github.com/curtis0505/base/accounts/abi"
	basecommon "github.com/curtis0505/base/common"
	basetypes "github.com/curtis0505/base/core/types"
	"github.com/curtis0505/bridge/libs/types/bridge/abi/base"
	"github.com/curtis0505/bridge/libs/types/bridge/abi/ether"
	"github.com/curtis0505/bridge/libs/types/bridge/abi/klay"
	etherabi "github.com/ethereum/go-ethereum/accounts/abi"
	ethercommon "github.com/ethereum/go-ethereum/common"
	ethertypes "github.com/ethereum/go-ethereum/core/types"
	klayabi "github.com/kaiachain/kaia/accounts/abi"
	klaytypes "github.com/kaiachain/kaia/blockchain/types"
	klaycommon "github.com/kaiachain/kaia/common"
)

type Log struct {
	EventName string

	chain string
	abi   interface{}
	inner interface{}

	blockNumber     uint64
	index           int
	txHash          string
	contractAddress string
	logs            map[string]string
}

func NewLog(l interface{}, chain string) Log {
	//TODO : bridge chain 추가 시 체크 필요
	log := Log{
		inner: l,
		chain: chain,
	}

	switch v := log.inner.(type) {
	case klaytypes.Log:
		topic := log.Topics()[0]
		for _, a := range klay.All {
			event, err := a.EventByID(klaycommon.HexToHash(topic))
			if err != nil {
				continue
			}
			log.abi = a
			log.EventName = event.Name
			break
		}
	case ethertypes.Log:
		topic := log.Topics()[0]
		for _, a := range ether.All {
			event, err := a.EventByID(ethercommon.HexToHash(topic))
			if err != nil {
				continue
			}
			log.abi = a
			log.EventName = event.Name
			break
		}
	case basetypes.Log:
		topic := log.Topics()[0]
		for _, a := range base.All {
			event, err := a.EventByID(basecommon.HexToHash(topic))
			if err != nil {
				continue
			}
			log.abi = a
			log.EventName = event.Name
			break
		}
	case abcitypes.Event:
		log.logs = make(map[string]string)
		for _, attributes := range v.Attributes {
			key, errKey := base64.StdEncoding.DecodeString(attributes.Key)
			value, errValue := base64.StdEncoding.DecodeString(attributes.Value)
			if errKey == nil && errValue == nil {
				log.logs[string(key)] = string(value)
			} else {
				log.logs[attributes.Key] = attributes.Value
			}
		}
		if v.Type == wasmtypes.ModuleName {
			const contractAddressKey = "_contract_address"
			const actionKey = "action"
			log.contractAddress = log.logs[contractAddressKey]
			log.EventName = log.logs[actionKey]
			delete(log.logs, contractAddressKey)
			delete(log.logs, actionKey)
		}
	case Log:
		return l.(Log)
	}
	return log
}

func NewLogMetaData(l interface{}, chain, txHash string, index int, blockNumber int64) Log {
	log := NewLog(l, chain)
	log.txHash = txHash
	log.index = index
	log.blockNumber = uint64(blockNumber)

	return log
}

func NewLogWithAbi(chain string, l interface{}, abi interface{}) Log {
	log := NewLog(l, chain)
	log.abi = abi

	return log
}

func (log Log) SetTxHash(txHash string) Log {
	log.txHash = txHash
	return log
}

func (log Log) Abi(abi interface{}) (Log, error) {
	c, err := NewAbi(log.Chain(), abi)
	if err != nil {
		return log, err
	}

	topic := log.Topics()[0]
	switch v := c.(type) {
	case etherabi.ABI:
		event, err := v.EventByID(ethercommon.HexToHash(topic))
		if err != nil {
			return log, err
		}
		log.abi = v
		log.EventName = event.Name
	case klayabi.ABI:
		event, err := v.EventByID(klaycommon.HexToHash(topic))
		if err != nil {
			return log, err
		}
		log.abi = v
		log.EventName = event.Name
	case baseabi.ABI:
		event, err := v.EventByID(basecommon.HexToHash(topic))
		if err != nil {
			return log, err
		}
		log.abi = v
		log.EventName = event.Name
	}

	return log, nil
}

func (log Log) Chain() string {
	return log.chain
}

func (log Log) KlayLog() klaytypes.Log {
	switch v := log.inner.(type) {
	case klaytypes.Log:
		return v
	}
	return klaytypes.Log{}
}

func (log Log) EtherLog() ethertypes.Log {
	switch v := log.inner.(type) {
	case ethertypes.Log:
		return v
	}
	return ethertypes.Log{}
}

func (log Log) BaseLog() basetypes.Log {
	switch v := log.inner.(type) {
	case basetypes.Log:
		return v
	}
	return basetypes.Log{}
}

func (log Log) Data() []byte {
	switch v := log.inner.(type) {
	case klaytypes.Log:
		return v.Data
	case ethertypes.Log:
		return v.Data
	case basetypes.Log:
		return v.Data
	}
	return nil
}

func (log Log) Address() string {
	switch v := log.inner.(type) {
	case klaytypes.Log:
		return v.Address.String()
	case ethertypes.Log:
		return v.Address.String()
	case basetypes.Log:
		return v.Address.String()
	case abcitypes.Event:
		return log.contractAddress
	}
	return ""
}

func (log Log) BlockNumber() uint64 {
	switch log.inner.(type) {
	case klaytypes.Log:
		return log.KlayLog().BlockNumber
	case ethertypes.Log:
		return log.EtherLog().BlockNumber
	case basetypes.Log:
		return log.BaseLog().BlockNumber
	case abcitypes.Event:
		return log.blockNumber
	}
	return 0
}

func (log Log) TxHash() string {
	switch v := log.inner.(type) {
	case klaytypes.Log:
		return v.TxHash.String()
	case ethertypes.Log:
		return v.TxHash.String()
	case basetypes.Log:
		return v.TxHash.String()
	case *abcitypes.Event:
		return log.txHash
	}
	return log.txHash
}

func (log Log) Topics() []string {
	topics := make([]string, 0)
	switch v := log.inner.(type) {
	case klaytypes.Log:
		for _, topic := range v.Topics {
			topics = append(topics, topic.String())
		}
	case ethertypes.Log:
		for _, topic := range v.Topics {
			topics = append(topics, topic.String())
		}
	case basetypes.Log:
		for _, topic := range v.Topics {
			topics = append(topics, topic.String())
		}
	case *abcitypes.Event:
		return topics
	}
	return topics
}

func (log Log) Index() uint {
	switch v := log.inner.(type) {
	case klaytypes.Log:
		return v.Index
	case ethertypes.Log:
		return v.Index
	case basetypes.Log:
		return v.Index
	case abcitypes.Event:
		return uint(log.index)
	}
	return 0
}

func (log Log) TxIndex() uint {
	switch v := log.inner.(type) {
	case klaytypes.Log:
		return v.TxIndex
	case ethertypes.Log:
		return v.TxIndex
	case basetypes.Log:
		return v.TxIndex
	}
	return 0
}

func (log Log) Unmarshal(value interface{}) error {
	if GetChainType(log.chain) == ChainTypeEVM {
		if log.abi == nil {
			return fmt.Errorf("unknown log")
		}
	}

	var err error
	switch v := log.inner.(type) {
	case klaytypes.Log:
		abi := log.abi.(klayabi.ABI)
		var indexed klayabi.Arguments

		for _, arg := range abi.Events[log.EventName].Inputs {
			if arg.Indexed {
				indexed = append(indexed, arg)
			}
		}

		if len(log.Data()) > 0 {
			err = abi.UnpackIntoInterface(value, log.EventName, log.Data())
			if err != nil {
				return err
			}
		}

		err = klayabi.ParseTopics(value, indexed, v.Topics[1:])
		if err != nil {
			return err
		}
	case ethertypes.Log:
		abi := log.abi.(etherabi.ABI)
		var indexed etherabi.Arguments

		for _, arg := range abi.Events[log.EventName].Inputs {
			if arg.Indexed {
				indexed = append(indexed, arg)
			}
		}

		if len(log.Data()) > 0 {
			err = abi.UnpackIntoInterface(value, log.EventName, log.Data())
			if err != nil {
				return err
			}
		}

		err = etherabi.ParseTopics(value, indexed, v.Topics[1:])
		if err != nil {
			return err
		}
	case basetypes.Log:
		abi := log.abi.(baseabi.ABI)
		var indexed baseabi.Arguments

		for _, arg := range abi.Events[log.EventName].Inputs {
			if arg.Indexed {
				indexed = append(indexed, arg)
			}
		}

		if len(log.Data()) > 0 {
			err = abi.UnpackIntoInterface(value, log.EventName, log.Data())
			if err != nil {
				return err
			}
		}

		err = baseabi.ParseTopics(value, indexed, v.Topics[1:])
		if err != nil {
			return err
		}
	case abcitypes.Event:
		bz, _ := json.Marshal(log.logs)
		return json.Unmarshal(bz, &value)
	}

	return nil
}

func (log Log) UnmarshalABI(abi interface{}, value interface{}) error {
	abiLog, err := log.Abi(abi)
	if err != nil {
		return err
	}

	return abiLog.Unmarshal(value)
}
