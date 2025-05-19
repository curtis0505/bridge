package types

import (
	"encoding/hex"
	"fmt"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	arbtypes "github.com/curtis0505/arbitrum/core/types"
	basetypes "github.com/curtis0505/base/core/types"
	troncommon "github.com/curtis0505/bridge/libs/common/tron"
	"github.com/curtis0505/bridge/libs/logger/v2"
	troncore "github.com/curtis0505/grpc-idl/tron/core"
	ethertypes "github.com/ethereum/go-ethereum/core/types"
	klaytypes "github.com/klaytn/klaytn/blockchain/types"
	klaycommon "github.com/klaytn/klaytn/common"
	"math/big"
)

type Receipt struct {
	inner any
	chain string
}

// NewReceipt TODO
func NewReceipt(receipt any, chain string) *Receipt {
	return &Receipt{
		inner: receipt,
		chain: chain,
	}
}

func (r *Receipt) BlockNumber() *big.Int {
	if r.inner == nil {
		return big.NewInt(0)
	}

	switch v := r.inner.(type) {
	case *klaytypes.Receipt:
		logger.Warn("BlockNumber", logger.BuildLogInput().WithError(fmt.Errorf("not support klaytn receipt")))
		return big.NewInt(0)
	case *ethertypes.Receipt:
		return v.BlockNumber
	case *arbtypes.Receipt:
		return v.BlockNumber
	case *basetypes.Receipt:
		return v.BlockNumber
	case *troncore.TransactionInfo:
		return big.NewInt(v.GetBlockNumber())
	case *cosmossdk.TxResponse:
		return big.NewInt(v.Height)
	}
	return big.NewInt(0)
}

func (r *Receipt) Status() uint64 {
	if r.inner == nil {
		return 0
	}

	switch v := r.inner.(type) {
	case *klaytypes.Receipt:
		if v == nil {
			return 0
		}
		return uint64(v.Status)
	case *ethertypes.Receipt:
		if v == nil {
			return 0
		}
		return v.Status
	case *arbtypes.Receipt:
		if v == nil {
			return 0
		}
		return v.Status
	case *basetypes.Receipt:
		if v == nil {
			return 0
		}
		return v.Status
	case *troncore.TransactionInfo:
		if v == nil {
			return 0
		}
		return uint64(v.GetResult().Number())
	case *cosmossdk.TxResponse:
		if v == nil {
			return 0
		}
		return uint64(v.Code)
	}
	return 0
}

func (r *Receipt) Success() bool {
	switch v := r.inner.(type) {
	case *troncore.TransactionInfo:
		if int(v.GetResult().Number()) == int(troncore.Transaction_Result_SUCCESS) {
			return true
		}
	case *cosmossdk.TxResponse:
		if v.Code == 0 {
			return true
		}
	default:
		if r.Status() == ReceiptStatusSuccessful {
			return true
		}
	}
	return false
}

func (r *Receipt) ContractAddress() string {
	if r.inner == nil {
		return ""
	}

	switch v := r.inner.(type) {
	case *klaytypes.Receipt:
		return v.ContractAddress.String()
	case *ethertypes.Receipt:
		return v.ContractAddress.String()
	case *arbtypes.Receipt:
		return v.ContractAddress.String()
	case *basetypes.Receipt:
		return v.ContractAddress.String()
	case *troncore.TransactionInfo:
		return troncommon.FromBytes(v.GetContractAddress()).String()
	case *cosmossdk.TxResponse:
		return ""
	}
	return ""
}

func (r *Receipt) GasUsed() uint64 {
	if r.inner == nil {
		return 0
	}

	switch v := r.inner.(type) {
	case *klaytypes.Receipt:
		return v.GasUsed
	case *ethertypes.Receipt:
		return v.GasUsed
	case *arbtypes.Receipt:
		return v.GasUsed
	case *basetypes.Receipt:
		return v.GasUsed
	case *troncore.TransactionInfo:
		return uint64(v.GetFee())
	case *cosmossdk.TxResponse:
		return uint64(v.GasWanted)
	}
	return 0
}

func (r *Receipt) Logs() []Log {
	logs := make([]Log, 0)
	switch v := r.inner.(type) {
	case *klaytypes.Receipt:
		for _, log := range v.Logs {
			logs = append(logs, NewLog(*log, r.chain))
		}
	case *ethertypes.Receipt:
		for _, log := range v.Logs {
			logs = append(logs, NewLog(*log, r.chain))
		}
	case *arbtypes.Receipt:
		for _, log := range v.Logs {
			logs = append(logs, NewLog(*log, r.chain))
		}
	case *basetypes.Receipt:
		for _, log := range v.Logs {
			logs = append(logs, NewLog(*log, r.chain))
		}
	case *troncore.TransactionInfo:
		for _, log := range v.GetLog() {
			logs = append(logs, NewLog(log, r.chain))
		}
	case *cosmossdk.TxResponse:
		for _, log := range v.Events {
			logs = append(logs, NewLog(log, r.chain))
		}
	}
	return logs
}

func (r *Receipt) MarshalBinary() ([]byte, error) {
	if r.inner == nil {
		return nil, fmt.Errorf("empty receipt")
	}

	switch v := r.inner.(type) {
	case *klaytypes.Receipt:
		return v.MarshalJSON()
	case *ethertypes.Receipt:
		return v.MarshalBinary()
	case *arbtypes.Receipt:
		return v.MarshalBinary()
	case *basetypes.Receipt:
		return v.MarshalBinary()
	case *cosmossdk.TxResponse:
		return v.Marshal()
	}
	return []byte{}, NotSupported
}

func (r *Receipt) Unmarshal(chain string, input []byte) error {
	switch chain {
	case ChainKLAY:
		receipt := &klaytypes.Receipt{}
		err := receipt.UnmarshalJSON(input)
		if err != nil {
			return err
		}
		r.inner = receipt
		r.chain = chain
		return nil
	case ChainETH, ChainMATIC:
		receipt := &ethertypes.Receipt{}
		err := receipt.UnmarshalBinary(input)
		if err != nil {
			return err
		}
		r.inner = receipt
		r.chain = chain
		return nil
	case ChainBASE:
		receipt := &basetypes.Receipt{}
		err := receipt.UnmarshalBinary(input)
		if err != nil {
			return err
		}
		r.inner = receipt
		r.chain = chain
		return nil
	case ChainARB:
		receipt := &arbtypes.Receipt{}
		err := receipt.UnmarshalBinary(input)
		if err != nil {
			return err
		}
		r.inner = receipt
		r.chain = chain
		return nil
	default:
		return fmt.Errorf("not support chain %s", chain)
	}
}

func (r *Receipt) BlockHash() (string, error) {
	switch v := r.inner.(type) {
	case *klaytypes.Receipt:
		if len(v.Logs) == 0 {
			return klaycommon.Hash{}.Hex(), nil
		}
		return v.Logs[0].BlockHash.String(), nil
	case *ethertypes.Receipt:
		return v.BlockHash.String(), nil
	case *arbtypes.Receipt:
		return v.BlockHash.String(), nil
	case *basetypes.Receipt:
		return v.BlockHash.String(), nil
	case *troncore.TransactionInfo:
		return hex.EncodeToString(v.GetId()), nil
	case *cosmossdk.TxResponse:
		return "", nil
	}
	return "", NotSupported
}

func (r *Receipt) Inner() any { return r.inner }
