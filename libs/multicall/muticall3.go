package multicall

import (
	"encoding/json"
	"fmt"
	"github.com/curtis0505/bridge/libs/logger/v2"
	"github.com/curtis0505/bridge/libs/types"
	ethercommon "github.com/ethereum/go-ethereum/common"
	"reflect"
)

type MultiCall3 []Call3

func New3() *MultiCall3 {
	var multiCall MultiCall3
	multiCall = make([]Call3, 0)
	return &multiCall
}

func (m *MultiCall3) AddCall(chain, to string, methodName string, allowFailure bool, abi []map[string]interface{}, args ...interface{}) (int, error) {
	var multiCallData Call3
	var err error

	multiCallData, err = newCall3(chain, to, methodName, allowFailure, abi, args...)
	if err != nil {
		return 0, err
	}
	idx := len(*m)

	*m = append(*m, multiCallData)
	return idx, nil
}

func (m *MultiCall3) AddCallUnmarshaler(chain, to string, methodName string, allowFailure bool, abi []map[string]interface{}, v types.CallMsgUnmarshaler, args ...interface{}) error {
	var multiCallData Call3
	var err error

	if reflect.ValueOf(v).Kind() != reflect.Pointer {
		return fmt.Errorf("not pointer unmarshaler: %s", reflect.TypeOf(v).Name())
	}

	multiCallData, err = newCall3Unmarshaler(chain, to, methodName, allowFailure, abi, v, args...)
	if err != nil {
		return err
	}
	*m = append(*m, multiCallData)
	return nil
}

func (m MultiCall3) Len() int { return len(m) }
func (m MultiCall3) ChunkSize(size int) int {
	if size == 0 {
		size = DefaultBatchSize
	}
	return m.Len()/size + 1
}
func (m MultiCall3) Chunk(index, size int) *MultiCall3 {
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

func (m MultiCall3) Unpack(data []any) ([][]any, error) {
	unpacked := make([][]any, 0)

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

func (m *MultiCall3) Unmarshal(v []any) {
	bz, err := json.Marshal(v[0])
	if err != nil {
		return
	}

	var results []Result
	err = json.Unmarshal(bz, &results)
	if err != nil {
		return
	}

	_m := *m
	for i, result := range results {
		if result.Success {
			unpack, err := _m[i].Unpack(result.ReturnData)
			if err != nil {
				if _m[i].AllowFailure {
					logger.Warn("MultiCall3",
						logger.BuildLogInput().
							WithMethod(_m[i].MethodName).
							WithAddress(ethercommon.BytesToAddress(_m[i].Address[:]).String()).
							WithError(err).
							WithData("index", i),
					)
					continue
				} else {
					logger.Warn("MultiCall3",
						logger.BuildLogInput().
							WithMethod(_m[i].MethodName).
							WithAddress(ethercommon.BytesToAddress(_m[i].Address[:]).String()).
							WithError(err),
					)
					return
				}
			}
			_m[i].Unmarshal(unpack)
		} else {
			logger.Warn("MultiCall3",
				logger.BuildLogInput().
					WithMethod(_m[i].MethodName).
					WithAddress(ethercommon.BytesToAddress(_m[i].Address[:]).String()).
					WithError(err).
					WithData("index", i),
			)
		}
	}
}
