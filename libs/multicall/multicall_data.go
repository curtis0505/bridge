package multicall

import (
	"fmt"
	"github.com/curtis0505/bridge/libs/types"
	etherabi "github.com/ethereum/go-ethereum/accounts/abi"
	ethercommon "github.com/ethereum/go-ethereum/common"
	klayabi "github.com/kaiachain/kaia/accounts/abi"
)

type Call struct {
	Chain       string
	Abi         interface{}
	MethodName  string
	Address     [20]byte `abi:"target"`
	Data        []byte   `abi:"callData"`
	Unmarshaler types.CallMsgUnmarshaler
}

type Call3 struct {
	Chain        string
	Abi          interface{}
	MethodName   string
	Address      [20]byte `abi:"target"`
	Data         []byte   `abi:"callData"`
	AllowFailure bool     `abi:"allowFailure"`
	Unmarshaler  types.CallMsgUnmarshaler
}

type Result struct {
	Success    bool   `abi:"success" json:"success"`
	ReturnData []byte `abi:"returnData" json:"returnData"`
}

func newCall(chain, to string, methodName string, abi []map[string]interface{}, args ...interface{}) (Call, error) {
	call := Call{
		Chain:      chain,
		MethodName: methodName,
	}

	switch chain {
	// TODO: 체인 추가시 체크 필요
	case types.ChainKLAY:
		abiParsed, err := getAbiCache().getAbi(chain, to, abi)
		if err != nil {
			return call, err
		}

		input, err := abiParsed.klayAbi.Pack(methodName, args...)
		if err != nil {
			return call, err
		}

		call.Abi = abiParsed.klayAbi
		call.Address = ethercommon.HexToAddress(to)
		call.Data = input
	default:
		abiParsed, err := getAbiCache().getAbi(chain, to, abi)
		if err != nil {
			return call, err
		}

		input, err := abiParsed.etherAbi.Pack(methodName, args...)
		if err != nil {
			return call, err
		}

		call.Abi = abiParsed.etherAbi
		call.Address = ethercommon.HexToAddress(to)
		call.Data = input
	}

	return call, nil
}

func newCall3(chain, to string, methodName string, allowFailure bool, abi []map[string]interface{}, args ...interface{}) (Call3, error) {
	call, err := newCall(chain, to, methodName, abi, args...)
	if err != nil {
		return Call3{}, err
	}

	call3 := Call3{
		Chain:        call.Chain,
		MethodName:   call.MethodName,
		Address:      call.Address,
		Abi:          call.Abi,
		Unmarshaler:  call.Unmarshaler,
		Data:         call.Data,
		AllowFailure: allowFailure,
	}

	return call3, nil
}

func newCallUnmarshaler(chain, to string, methodName string, abi []map[string]interface{}, v types.CallMsgUnmarshaler, args ...interface{}) (Call, error) {
	multiCall := Call{
		Chain:       chain,
		MethodName:  methodName,
		Unmarshaler: v,
	}

	switch chain {
	// TODO: 체인 추가시 체크 필요
	case types.ChainKLAY:
		abiParsed, err := getAbiCache().getAbi(chain, to, abi)
		if err != nil {
			return Call{}, err
		}

		input, err := abiParsed.klayAbi.Pack(methodName, args...)
		if err != nil {
			return Call{}, err
		}

		multiCall.Abi = abiParsed.klayAbi
		multiCall.Address = ethercommon.HexToAddress(to)
		multiCall.Data = input
	default:
		abiParsed, err := getAbiCache().getAbi(chain, to, abi)
		if err != nil {
			return Call{}, err
		}

		input, err := abiParsed.etherAbi.Pack(methodName, args...)
		if err != nil {
			return Call{}, err
		}

		multiCall.Abi = abiParsed.etherAbi
		multiCall.Address = ethercommon.HexToAddress(to)
		multiCall.Data = input
	}

	return multiCall, nil
}

func newCall3Unmarshaler(chain, to string, methodName string, allowFailure bool, abi []map[string]interface{}, v types.CallMsgUnmarshaler, args ...interface{}) (Call3, error) {
	call, err := newCallUnmarshaler(chain, to, methodName, abi, v, args...)
	if err != nil {
		return Call3{}, err
	}

	call3 := Call3{
		Chain:        call.Chain,
		MethodName:   call.MethodName,
		Address:      call.Address,
		Abi:          call.Abi,
		Unmarshaler:  call.Unmarshaler,
		Data:         call.Data,
		AllowFailure: allowFailure,
	}

	return call3, nil
}

func (m Call) Unpack(call []byte) ([]any, error) {
	switch m.Chain {
	// TODO: 체인 추가시 체크 필요
	case types.ChainKLAY:
		result, err := m.Abi.(*klayabi.ABI).Methods[m.MethodName].Outputs.UnpackValues(call)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack %s : %s", m.MethodName, err.Error())
		}
		return result, nil
	default:
		result, err := m.Abi.(*etherabi.ABI).Methods[m.MethodName].Outputs.UnpackValues(call)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack %s : %s", m.MethodName, err.Error())
		}
		return result, nil
	}
}

func (m Call) Unmarshal(v []any) {
	m.Unmarshaler.Unmarshal(v)
}

func (m Call3) Unpack(call []byte) ([]any, error) {
	switch m.Chain {
	// TODO: 체인 추가시 체크 필요
	case types.ChainKLAY:
		result, err := m.Abi.(*klayabi.ABI).Methods[m.MethodName].Outputs.UnpackValues(call)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack %s : %s", m.MethodName, err.Error())
		}
		return result, nil
	default:
		result, err := m.Abi.(*etherabi.ABI).Methods[m.MethodName].Outputs.UnpackValues(call)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack %s : %s", m.MethodName, err.Error())
		}
		return result, nil
	}
}

func (m Call3) Unmarshal(v []any) {
	m.Unmarshaler.Unmarshal(v)
}
