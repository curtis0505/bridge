package types

import (
	"encoding/json"
	"fmt"
	arbabi "github.com/curtis0505/arbitrum/accounts/abi"
	arbcommon "github.com/curtis0505/arbitrum/common"
	baseabi "github.com/curtis0505/base/accounts/abi"
	basecommon "github.com/curtis0505/base/common"
	"github.com/curtis0505/bridge/libs/logger/v2"
	etherabi "github.com/ethereum/go-ethereum/accounts/abi"
	ethercommon "github.com/ethereum/go-ethereum/common"
	klayabi "github.com/klaytn/klaytn/accounts/abi"
	klaycommon "github.com/klaytn/klaytn/common"
	"reflect"
	"strings"
)

type ABI struct {
	chain string
	inner any
}

func NewABIJson(chain string, abi []byte) (*ABI, error) {
	switch chain {
	case ChainKLAY:
		result, err := klayabi.JSON(strings.NewReader(string(abi)))
		if err != nil {
			return nil, err
		}
		return &ABI{chain: chain, inner: result}, nil
	case ChainETH, ChainMATIC:
		result, err := etherabi.JSON(strings.NewReader(string(abi)))
		if err != nil {
			return nil, err
		}
		return &ABI{chain: chain, inner: result}, nil
	case ChainBASE:
		result, err := baseabi.JSON(strings.NewReader(string(abi)))
		if err != nil {
			return nil, err
		}
		return &ABI{chain: chain, inner: result}, nil
	case ChainARB:
		result, err := arbabi.JSON(strings.NewReader(string(abi)))
		if err != nil {
			return nil, err
		}
		return &ABI{chain: chain, inner: result}, nil
	default:
		return nil, fmt.Errorf("invalid chain: %s", chain)
	}
}

func (a *ABI) EventByID(chain string, topic string) (*Event, error) {
	switch v := a.inner.(type) {
	case klayabi.ABI:
		result, err := v.EventByID(klaycommon.HexToHash(topic))
		if err != nil {
			return nil, err
		}
		return NewEvent(chain, result), nil
	case etherabi.ABI:
		result, err := v.EventByID(ethercommon.HexToHash(topic))
		if err != nil {
			return nil, err
		}
		return NewEvent(chain, result), nil
	case baseabi.ABI:
		result, err := v.EventByID(basecommon.HexToHash(topic))
		if err != nil {
			return nil, err
		}
		return NewEvent(chain, result), nil
	case arbabi.ABI:
		result, err := v.EventByID(arbcommon.HexToHash(topic))
		if err != nil {
			return nil, err
		}
		return NewEvent(chain, result), nil
	default:
		return nil, fmt.Errorf("invalid chain: %s", chain)
	}
}

func (a *ABI) MethodByID(chain string, topic []byte) (*Method, error) {
	switch v := a.inner.(type) {
	case klayabi.ABI:
		result, err := v.MethodById(topic)
		if err != nil {
			return nil, err
		}
		return NewMethod(chain, result), nil
	case etherabi.ABI:
		result, err := v.MethodById(topic)
		if err != nil {
			return nil, err
		}
		return NewMethod(chain, result), nil
	case baseabi.ABI:
		result, err := v.MethodById(topic)
		if err != nil {
			return nil, err
		}
		return NewMethod(chain, result), nil
	case arbabi.ABI:
		result, err := v.MethodById(topic)
		if err != nil {
			return nil, err
		}
		return NewMethod(chain, result), nil
	default:
		return nil, fmt.Errorf("invalid chain: %s", chain)
	}
}

func (a *ABI) UnpackIntoMap(out map[string]interface{}, name string, data []byte, topics []string) error {
	switch v := a.inner.(type) {
	case klayabi.ABI:
		err := v.UnpackIntoMap(out, name, data)
		if err != nil {
			return err
		}

		var indexed klayabi.Arguments
		for _, arg := range v.Events[name].Inputs {
			if arg.Indexed {
				indexed = append(indexed, arg)
			}
		}

		var topicsHash []klaycommon.Hash
		for _, topic := range topics[1 : 1+len(indexed)] {
			topicsHash = append(topicsHash, klaycommon.HexToHash(topic))
		}

		err = klayabi.ParseTopicsIntoMap(out, indexed, topicsHash)
		if err != nil {
			logger.Error("Handler", logger.BuildLogInput().
				WithEvent("parseLogs").
				WithError(fmt.Errorf("ParseTopicsIntoMap: %w", err)))
		}
	}

	return nil
}
func (a *ABI) UnpackIntoInterface(out any, name string, data []byte, topics []string) error {
	switch v := a.inner.(type) {
	case klayabi.ABI:
		if len(data) > 0 {
			err := v.UnpackIntoInterface(out, name, data)
			if err != nil {
				return err
			}
		}
		var indexed klayabi.Arguments
		for _, arg := range v.Events[name].Inputs {
			if arg.Indexed {
				indexed = append(indexed, arg)
			}
		}

		var topicsHash []klaycommon.Hash
		for _, topic := range topics[1 : 1+len(indexed)] {
			topicsHash = append(topicsHash, klaycommon.HexToHash(topic))
		}
		err := klayabi.ParseTopics(out, indexed, topicsHash)
		if err != nil {
			return err
		}
	case etherabi.ABI:
		if len(data) > 0 {
			err := v.UnpackIntoInterface(out, name, data)
			if err != nil {
				return err
			}
		}
		var indexed etherabi.Arguments
		for _, arg := range v.Events[name].Inputs {
			if arg.Indexed {
				indexed = append(indexed, arg)
			}
		}

		var topicsHash []ethercommon.Hash
		for _, topic := range topics[1 : 1+len(indexed)] {
			topicsHash = append(topicsHash, ethercommon.HexToHash(topic))
		}
		err := etherabi.ParseTopics(out, indexed, topicsHash)
		if err != nil {
			return err
		}
	case baseabi.ABI:
		if len(data) > 0 {
			err := v.UnpackIntoInterface(out, name, data)
			if err != nil {
				return err
			}
		}
		var indexed baseabi.Arguments
		for _, arg := range v.Events[name].Inputs {
			if arg.Indexed {
				indexed = append(indexed, arg)
			}
		}

		var topicsHash []basecommon.Hash
		for _, topic := range topics[1 : 1+len(indexed)] {
			topicsHash = append(topicsHash, basecommon.HexToHash(topic))
		}
		err := baseabi.ParseTopics(out, indexed, topicsHash)
		if err != nil {
			return err
		}
	case arbabi.ABI:
		if len(data) > 0 {
			err := v.UnpackIntoInterface(out, name, data)
			if err != nil {
				return err
			}
		}
		var indexed arbabi.Arguments
		for _, arg := range v.Events[name].Inputs {
			if arg.Indexed {
				indexed = append(indexed, arg)
			}
		}

		var topicsHash []arbcommon.Hash
		for _, topic := range topics[1 : 1+len(indexed)] {
			topicsHash = append(topicsHash, arbcommon.HexToHash(topic))
		}
		err := arbabi.ParseTopics(out, indexed, topicsHash)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewAbi(chain string, abi interface{}) (interface{}, error) {
	// TODO: 체인 추가시 체크 필요
	switch v := abi.(type) {
	case etherabi.ABI:
		return v, nil
	case klayabi.ABI:
		return v, nil
	case *etherabi.ABI:
		return *v, nil
	case *klayabi.ABI:
		return *v, nil
	case *baseabi.ABI:
		return *v, nil
	case []map[string]interface{}:
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		switch chain {
		case ChainKLAY:
			abi, err = klayabi.JSON(strings.NewReader(string(b)))
			if err != nil {
				return nil, err
			}
		case ChainETH, ChainMATIC:
			abi, err = etherabi.JSON(strings.NewReader(string(b)))
			if err != nil {
				return nil, err
			}
		case ChainARB:
			abi, err = arbabi.JSON(strings.NewReader(string(b)))
			if err != nil {
				return nil, err
			}
		case ChainBASE:
			abi, err = baseabi.JSON(strings.NewReader(string(b)))
			if err != nil {
				return nil, err
			}
		}
		return abi, nil
	}
	return abi, fmt.Errorf("invalid abi type: %s", reflect.TypeOf(abi))
}

func PackAbi(chain string, abi interface{}, method string, v ...interface{}) ([]byte, error) {
	switch abi.(type) {
	case etherabi.ABI:
		return abi.(etherabi.ABI).Pack(method, v...)
	case klayabi.ABI:
		return abi.(klayabi.ABI).Pack(method, v...)
	case arbabi.ABI:
		return abi.(arbabi.ABI).Pack(method, v...)
	case baseabi.ABI:
		return abi.(baseabi.ABI).Pack(method, v...)
	default:
		a, err := NewAbi(chain, abi)
		if err != nil {
			return nil, err
		}
		return PackAbi(chain, a, method, v...)
	}
}

func UnpackValuesAbi(chain string, abi interface{}, method string, b []byte) ([]interface{}, error) {
	switch abi.(type) {
	case etherabi.ABI:
		return abi.(etherabi.ABI).Methods[method].Inputs.UnpackValues(b)
	case klayabi.ABI:
		return abi.(klayabi.ABI).Methods[method].Inputs.UnpackValues(b)
	case arbabi.ABI:
		return abi.(arbabi.ABI).Methods[method].Inputs.UnpackValues(b)
	case baseabi.ABI:
		return abi.(baseabi.ABI).Methods[method].Inputs.UnpackValues(b)
	default:
		a, err := NewAbi(chain, abi)
		if err != nil {
			return nil, err
		}
		return UnpackValuesAbi(chain, a, method, b)
	}
}

type Method struct {
	chain string
	inner any
}

func NewMethod(chain string, inner any) *Method {
	return &Method{
		chain: chain,
		inner: inner,
	}
}

func (m *Method) UnpackInputData(value any, data []byte) error {
	switch v := m.inner.(type) {
	case *klayabi.Method:
		unpack, err := v.Inputs.Unpack(data[4:])
		if err != nil {
			return err
		}
		err = v.Inputs.Copy(value, unpack)
		if err != nil {
			return err
		}
	case *etherabi.Method:
		unpack, err := v.Inputs.Unpack(data[4:])
		if err != nil {
			return err
		}
		err = v.Inputs.Copy(value, unpack)
		if err != nil {
			return err
		}
	case *arbabi.Method:
		unpack, err := v.Inputs.Unpack(data[4:])
		if err != nil {
			return err
		}
		err = v.Inputs.Copy(value, unpack)
		if err != nil {
			return err
		}
	case *baseabi.Method:
		unpack, err := v.Inputs.Unpack(data[4:])
		if err != nil {
			return err
		}
		err = v.Inputs.Copy(value, unpack)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Method) Name() string {
	switch v := m.inner.(type) {
	case *klayabi.Method:
		return v.Name
	case *etherabi.Method:
		return v.Name
	case *arbabi.Method:
		return v.Name
	case *baseabi.Method:
		return v.Name
	}

	return ""
}

func (m *Method) RawName() string {
	switch v := m.inner.(type) {
	case *klayabi.Method:
		return v.RawName
	case *etherabi.Method:
		return v.RawName
	case *arbabi.Method:
		return v.RawName
	case *baseabi.Method:
		return v.RawName
	}

	return ""
}
