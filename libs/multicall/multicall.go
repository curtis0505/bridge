package multicall

import (
	"fmt"
	"github.com/curtis0505/bridge/libs/types"
	"reflect"
)

type MultiCall []Call

// Deprecated: NewMultiCall
// use New instead.
func NewMultiCall() *MultiCall {
	var multiCall MultiCall
	multiCall = make([]Call, 0)
	return &multiCall
}

func New() *MultiCall {
	var multiCall MultiCall
	multiCall = make([]Call, 0)
	return &multiCall
}

func (m *MultiCall) AddCall(chain, to string, methodName string, abi []map[string]interface{}, args ...interface{}) (int, error) {
	var multiCallData Call
	var err error

	multiCallData, err = newCall(chain, to, methodName, abi, args...)
	if err != nil {
		return 0, err
	}
	idx := len(*m)

	*m = append(*m, multiCallData)
	return idx, nil
}

func (m *MultiCall) AddCallUnmarshaler(chain, to string, methodName string, abi []map[string]interface{}, v types.CallMsgUnmarshaler, args ...interface{}) error {
	var multiCallData Call
	var err error

	if reflect.ValueOf(v).Kind() != reflect.Pointer {
		return fmt.Errorf("not pointer unmarshaler: %s", reflect.TypeOf(v).Name())
	}

	multiCallData, err = newCallUnmarshaler(chain, to, methodName, abi, v, args...)
	if err != nil {
		return err
	}
	*m = append(*m, multiCallData)
	return nil
}

func (m MultiCall) Len() int { return len(m) }
func (m MultiCall) ChunkSize(size int) int {
	if size == 0 {
		size = DefaultBatchSize
	}
	return m.Len()/size + 1
}
func (m MultiCall) Chunk(index, size int) *MultiCall {
	if size == 0 {
		size = DefaultBatchSize
	}
	if m.Len() > size {
		if index == m.ChunkSize(size)-1 {
			s := index * size

			call := m[s:]
			return &call
		} else {
			s := index * size
			e := s + size

			call := m[s:e]
			return &call
		}
	} else {
		return &m
	}
}

func (m MultiCall) Unpack(data []interface{}) ([][]interface{}, error) {
	unpacked := make([][]interface{}, 0)

	calls, ok := data[1].([][]byte)
	if !ok {
		return unpacked, fmt.Errorf("failed to cast multicall response")
	}

	for i, call := range calls {
		unpack, err := m[i].Unpack(call)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack %s : %s", m[i].MethodName, err.Error())
		}
		unpacked = append(unpacked, unpack)
	}
	return unpacked, nil
}

func (m *MultiCall) Unmarshal(v []interface{}) {
	calls, ok := v[1].([][]byte)
	if !ok {
		return
	}

	_m := *m
	for i, call := range calls {
		unpack, err := _m[i].Unpack(call)
		if err != nil {
			return
		}
		_m[i].Unmarshal(unpack)
	}
}
