package types

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/cometbft/cometbft/crypto/tmhash"
	cosmostxtypes "github.com/cosmos/cosmos-sdk/types/tx"
	arbabi "github.com/curtis0505/arbitrum/accounts/abi"
	arbcommon "github.com/curtis0505/arbitrum/common"
	arbhexutil "github.com/curtis0505/arbitrum/common/hexutil"
	arbtypes "github.com/curtis0505/arbitrum/core/types"
	arbrlp "github.com/curtis0505/arbitrum/rlp"
	baseabi "github.com/curtis0505/base/accounts/abi"
	basecommon "github.com/curtis0505/base/common"
	basehexutil "github.com/curtis0505/base/common/hexutil"
	basetypes "github.com/curtis0505/base/core/types"
	baserlp "github.com/curtis0505/base/rlp"
	cosmoscommon "github.com/curtis0505/bridge/libs/common/cosmos"
	"github.com/curtis0505/bridge/libs/logger/v2"
	troncore "github.com/curtis0505/grpc-idl/tron/core"
	etherabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethercommon "github.com/ethereum/go-ethereum/common"
	etherhexutil "github.com/ethereum/go-ethereum/common/hexutil"
	ethertypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	etherrlp "github.com/ethereum/go-ethereum/rlp"
	klayabi "github.com/klaytn/klaytn/accounts/abi"
	klaytypes "github.com/klaytn/klaytn/blockchain/types"
	klaycommon "github.com/klaytn/klaytn/common"
	klayhexutil "github.com/klaytn/klaytn/common/hexutil"
	klayrlp "github.com/klaytn/klaytn/rlp"
	"google.golang.org/protobuf/proto"
	"math/big"
	"reflect"
)

type Transaction struct {
	inner   interface{}
	receipt *Receipt
	chain   string
}

type RequestTransaction struct {
	From      string
	To        string
	Nonce     uint64
	GasPrice  *big.Int
	GasFeeCap *big.Int // a.k.a. maxFeePerGas
	GasTipCap *big.Int // a.k.a. maxPriorityFeePerGas
	GasLimit  uint64
	Value     *big.Int
	Data      []byte
}

type TransactionOption struct {
	Nonce     uint64
	GasPrice  *big.Int
	GasFeeCap *big.Int // a.k.a. maxFeePerGas
	GasTipCap *big.Int // a.k.a. maxPriorityFeePerGas
}

func NewTx(chain string, txData interface{}) *Transaction {
	switch v := txData.(type) {
	case *ethertypes.LegacyTx:
		tx := ethertypes.NewTx(v)
		return NewTransaction(tx, chain)

	case *ethertypes.DynamicFeeTx:
		tx := ethertypes.NewTx(v)
		return NewTransaction(tx, chain)

	case *basetypes.DynamicFeeTx:
		tx := basetypes.NewTx(v)
		return NewTransaction(tx, chain)

	case map[klaytypes.TxValueKeyType]interface{}:
		tx, err := klaytypes.NewTransactionWithMap(klaytypes.TxTypeSmartContractExecution, v)
		if err != nil {
			return NewTransaction(nil, chain)
		}
		return NewTransaction(tx, chain)

	default:
		return NewTransaction(nil, chain)
	}
}

func NewTransaction(tx interface{}, chain string) *Transaction {
	return &Transaction{
		inner: tx,
		chain: chain,
	}
}

func NewTransactionWithReceipt(tx interface{}, receipt interface{}, chain string) *Transaction {
	return &Transaction{
		inner:   tx,
		chain:   chain,
		receipt: NewReceipt(receipt, chain),
	}
}

func NewTransactionFromRLP(rlpTx string, chain string) (*Transaction, error) {
	switch chain {
	// TODO: 체인 추가시 체크 필요
	case ChainBASE:
		decoded, err := basehexutil.Decode(rlpTx)
		if err != nil {
			return nil, err
		}

		tx := new(basetypes.Transaction)
		err = baserlp.DecodeBytes(decoded, tx)
		if err != nil {
			return nil, fmt.Errorf("decode rlp error %w", err)
		}

		return NewTransaction(tx, chain), nil
	case ChainKLAY:
		decodeRlpTx, err := klayhexutil.Decode(rlpTx)
		if err != nil {
			return nil, err
		}

		tx := new(klaytypes.Transaction)
		err = klayrlp.DecodeBytes(decodeRlpTx, tx)
		if err != nil {
			return nil, fmt.Errorf("decode rlp error %w", err)
		}

		return NewTransaction(tx, chain), nil
	case ChainTRX, ChainATOM:
		return nil, NotImplemented

	default:
		decodeRlpTx, err := etherhexutil.Decode(rlpTx)
		if err != nil {
			return nil, err
		}

		tx := new(ethertypes.Transaction)
		err = etherrlp.DecodeBytes(decodeRlpTx, tx)
		if err != nil {
			return nil, fmt.Errorf("decode rlp error %w", err)
		}

		return NewTransaction(tx, chain), nil
	}
}

func (tx *Transaction) Inner() interface{} {
	return tx.inner
}

func (tx *Transaction) Chain() string {
	return tx.chain
}

func (tx *Transaction) From() string {
	switch v := tx.inner.(type) {
	// TODO: 체인 추가시 체크 필요
	case *basetypes.Transaction:
		sender, err := basetypes.LatestSignerForChainID(tx.ChainId()).Sender(v)
		if err != nil {
			return basecommon.Address{}.String()
		}
		return sender.String()
	case *klaytypes.Transaction:
		sender, err := klaytypes.LatestSignerForChainID(tx.ChainId()).Sender(v)
		if err != nil {
			return klaycommon.Address{}.String()
		}
		return sender.String()
	case *ethertypes.Transaction:
		sender, err := ethertypes.LatestSignerForChainID(tx.ChainId()).Sender(v)
		if err != nil {
			return ethercommon.Address{}.String()
		}
		return sender.String()
	case *arbtypes.Transaction:
		sender, err := arbtypes.LatestSignerForChainID(tx.ChainId()).Sender(v)
		if err != nil {
			return arbcommon.Address{}.String()
		}
		return sender.String()

	case *troncore.Transaction:
		return NotSupported.Error()
	case *cosmostxtypes.Tx:
		if len(v.GetAuthInfo().GetSignerInfos()) > 1 {
			return "MultiSig"
		} else {
			acc, err := cosmoscommon.FromPublicKey(tx.Chain(), v.GetAuthInfo().GetSignerInfos()[0].GetPublicKey().GetValue())
			if err != nil {
				return ""
			}
			return acc.String()
		}
	}
	return ethercommon.Address{}.String()
}

func (tx *Transaction) To() string {
	switch v := tx.inner.(type) {
	case *klaytypes.Transaction:
		to := v.To()
		if to == nil {
			return klaycommon.Address{}.String()
		}
		return to.String()
	case *ethertypes.Transaction:
		to := v.To()
		if to == nil {
			return ethercommon.Address{}.String()
		}
		return to.String()
	case *arbtypes.Transaction:
		to := v.To()
		if to == nil {
			return arbcommon.Address{}.String()
		}
		return to.String()
	case *basetypes.Transaction:
		to := v.To()
		if to == nil {
			return basecommon.Address{}.String()
		}
		return to.String()
	}
	return ethercommon.Address{}.String()
}

func (tx *Transaction) TxHash() string {
	switch v := tx.inner.(type) {
	case *klaytypes.Transaction:
		return v.Hash().String()
	case *ethertypes.Transaction:
		return v.Hash().String()
	case *arbtypes.Transaction:
		return v.Hash().String()
	case *basetypes.Transaction:
		return v.Hash().String()
	case *troncore.Transaction:
		raw, err := proto.Marshal(v.GetRawData())
		if err != nil {
			return ""
		}

		sha256hash := sha256.New()
		sha256hash.Write(raw)
		hash := sha256hash.Sum(nil)
		return hex.EncodeToString(hash)
	case *cosmostxtypes.Tx:
		raw, err := v.Marshal()
		if err != nil {
			return ""
		}
		return fmt.Sprintf("%X", tmhash.Sum(raw))
	}

	return ""
}

func (tx *Transaction) Data() []byte {
	switch v := tx.Inner().(type) {
	case *klaytypes.Transaction:
		return v.Data()
	case *ethertypes.Transaction:
		return v.Data()
	case *arbtypes.Transaction:
		return v.Data()
	case *basetypes.Transaction:
		return v.Data()
	case *troncore.Transaction:
		return nil
	case *cosmostxtypes.Tx:
		fmt.Println(v.GetMsgs())
		return nil
	default:
		fmt.Println(reflect.TypeOf(v))
	}
	return nil
}

func (tx *Transaction) MethodName(abi interface{}) string {
	id := tx.Data()[:4]

	c, err := NewAbi(tx.Chain(), abi)
	if err != nil {
		return ""
	}

	switch tx.inner.(type) {
	case *klaytypes.Transaction:
		a := c.(klayabi.ABI)
		m := klayabi.Method{}
		for _, method := range a.Methods {
			if bytes.Equal(id, method.ID) {
				m = method
				break
			}
		}
		return m.RawName
	case *ethertypes.Transaction:
		a := c.(etherabi.ABI)
		m := etherabi.Method{}
		for _, method := range a.Methods {
			if bytes.Equal(id, method.ID) {
				m = method
				break
			}
		}
		return m.RawName
	case *arbtypes.Transaction:
		a := c.(arbabi.ABI)
		m := arbabi.Method{}
		for _, method := range a.Methods {
			if bytes.Equal(id, method.ID) {
				m = method
				break
			}
		}
		return m.RawName
	case *basetypes.Transaction:
		a := c.(arbabi.ABI)
		m := arbabi.Method{}
		for _, method := range a.Methods {
			if bytes.Equal(id, method.ID) {
				m = method
				break
			}
		}
		return m.RawName
	}
	return ""
}

// Deprecated: Unmarshal
func (tx *Transaction) Unmarshal(abi interface{}, value interface{}) error {
	if len(tx.Data()) <= 4 {
		return fmt.Errorf("empty data")
	}

	c, err := NewAbi(tx.Chain(), abi)
	if err != nil {
		return err
	}

	id := tx.Data()[:4]
	switch tx.inner.(type) {
	case *klaytypes.Transaction:
		a := c.(klayabi.ABI)
		m := klayabi.Method{}
		for _, method := range a.Methods {
			if bytes.Equal(id, method.ID) {
				m = method
				break
			}
		}

		if value, err = m.Inputs.Unpack(tx.Data()[4:]); err != nil {
			return err
		}
		return nil
	case *ethertypes.Transaction:
		a := c.(etherabi.ABI)
		m := etherabi.Method{}
		for _, method := range a.Methods {
			if bytes.Equal(id, method.ID) {
				m = method
				break
			}
		}

		values, err := m.Inputs.Unpack(tx.Data()[4:])
		if err != nil {
			return err
		}

		if err = m.Inputs.Copy(value, values); err != nil {
			return err
		}
		return nil
	case *arbtypes.Transaction:
		a := c.(arbabi.ABI)
		m := arbabi.Method{}
		for _, method := range a.Methods {
			if bytes.Equal(id, method.ID) {
				m = method
				break
			}
		}

		values, err := m.Inputs.Unpack(tx.Data()[4:])
		if err != nil {
			return err
		}

		if err = m.Inputs.Copy(value, values); err != nil {
			return err
		}
		return nil
	case *basetypes.Transaction:
		a := c.(arbabi.ABI)
		m := arbabi.Method{}
		for _, method := range a.Methods {
			if bytes.Equal(id, method.ID) {
				m = method
				break
			}
		}

		values, err := m.Inputs.Unpack(tx.Data()[4:])
		if err != nil {
			return err
		}

		if err = m.Inputs.Copy(value, values); err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("invalid tx")
}

func (tx *Transaction) UnmarshalABI(abi interface{}, value interface{}) error {
	if len(tx.Data()) <= 4 {
		return fmt.Errorf("empty data")
	}

	c, err := NewAbi(tx.Chain(), abi)
	if err != nil {
		return err
	}

	id := tx.Data()[:4]
	switch tx.inner.(type) {
	case *klaytypes.Transaction:
		a := c.(klayabi.ABI)
		m := klayabi.Method{}
		for _, method := range a.Methods {
			if bytes.Equal(id, method.ID) {
				m = method
				break
			}
		}

		values, err := m.Inputs.Unpack(tx.Data()[4:])
		if err != nil {
			return err
		}

		if err = m.Inputs.Copy(value, values); err != nil {
			return err
		}

		return nil
	case *ethertypes.Transaction:
		a := c.(etherabi.ABI)
		m := etherabi.Method{}
		for _, method := range a.Methods {
			if bytes.Equal(id, method.ID) {
				m = method
				break
			}
		}

		values, err := m.Inputs.Unpack(tx.Data()[4:])
		if err != nil {
			return err
		}

		if err = m.Inputs.Copy(value, values); err != nil {
			return err
		}
		return nil
	case *arbtypes.Transaction:
		a := c.(arbabi.ABI)
		m := arbabi.Method{}
		for _, method := range a.Methods {
			if bytes.Equal(id, method.ID) {
				m = method
				break
			}
		}

		values, err := m.Inputs.Unpack(tx.Data()[4:])
		if err != nil {
			return err
		}

		if err = m.Inputs.Copy(value, values); err != nil {
			return err
		}
		return nil
	case *basetypes.Transaction:
		a := c.(baseabi.ABI)
		m := baseabi.Method{}
		for _, method := range a.Methods {
			if bytes.Equal(id, method.ID) {
				m = method
				break
			}
		}

		values, err := m.Inputs.Unpack(tx.Data()[4:])
		if err != nil {
			return err
		}

		if err = m.Inputs.Copy(value, values); err != nil {
			return err
		}
		return nil
	case *troncore.Transaction:
		a := c.(etherabi.ABI)
		m := etherabi.Method{}
		for _, method := range a.Methods {
			if bytes.Equal(id, method.ID) {
				m = method
				break
			}
		}

		values, err := m.Inputs.Unpack(tx.Data()[4:])
		if err != nil {
			return err
		}

		if err = m.Inputs.Copy(value, values); err != nil {
			return err
		}
		return nil
	}

	return InvalidTransaction
}

func (tx *Transaction) UnmarshalProto(message proto.Message) error {
	switch v := tx.inner.(type) {
	case *troncore.Transaction:
		contracts := make([]string, 0)
		for _, contract := range v.GetRawData().GetContract() {
			contracts = append(contracts, contract.GetType().String())

			err := contract.GetParameter().UnmarshalTo(message)
			if err != nil {
				continue
			}
			return nil
		}
		return fmt.Errorf("not found contract event: (%s) exists: %v", reflect.TypeOf(message).String(), contracts)
	case *cosmostxtypes.Tx:

	}
	return NotSupported
}

func (tx *Transaction) ChainId() *big.Int {
	switch v := tx.inner.(type) {
	case *klaytypes.Transaction:
		return v.ChainId()
	case *ethertypes.Transaction:
		return v.ChainId()
	case *arbtypes.Transaction:
		return v.ChainId()
	case *basetypes.Transaction:
		return v.ChainId()
	case *troncore.Transaction:
		return big.NewInt(0)
	}
	return big.NewInt(0)
}

func (tx *Transaction) Type() TxType {
	switch v := tx.inner.(type) {
	case *klaytypes.Transaction:
		txType := v.Type()
		if txType.IsLegacyTransaction() {
			return TxTypeKlaytnTransaction
		} else if txType.IsFeeDelegatedTransaction() {
			return TxTypeKlaytnFeeDelegatedTransaction
		} else {
			return TxTypeKlaytnTransaction
		}
	case *ethertypes.Transaction:
		switch v.Type() {
		case ethertypes.LegacyTxType:
			return TxTypeLegacyTransaction
		case ethertypes.AccessListTxType:
			return TxTypeEthereumAccessList
		case ethertypes.DynamicFeeTxType:
			return TxTypeEthereumDynamicFee
		}
	case *arbtypes.Transaction:
	// TODO: implement me
	case *basetypes.Transaction:
		switch v.Type() {
		case basetypes.LegacyTxType:
			return TxTypeLegacyTransaction
		case basetypes.AccessListTxType:
			return TxTypeEthereumAccessList
		case basetypes.DynamicFeeTxType:
			return TxTypeEthereumDynamicFee
		}
	case *troncore.Transaction:
		return TxTypeTronTransaction
	case *cosmostxtypes.Tx:
		return TxTypeCosmosTransaction
	}
	return TxTypeLegacyTransaction
}

func (tx *Transaction) Gas() uint64 {
	switch v := tx.inner.(type) {
	case *klaytypes.Transaction:
		return v.Gas()
	case *ethertypes.Transaction:
		return v.Gas()
	case *arbtypes.Transaction:
		return v.Gas()
	case *basetypes.Transaction:
		return v.Gas()
	case *troncore.Transaction:
		return 0
	case *cosmostxtypes.Tx:
		return 0
	}
	return 0
}

func (tx *Transaction) GasPrice() *big.Int {
	switch v := tx.inner.(type) {
	case *klaytypes.Transaction:
		return v.GasPrice()
	case *ethertypes.Transaction:
		return v.GasPrice()
	case *arbtypes.Transaction:
		return v.GasPrice()
	case *basetypes.Transaction:
		return v.GasPrice()
	case *troncore.Transaction:
		return big.NewInt(0)
	case *cosmostxtypes.Tx:
		return big.NewInt(0)
	}
	return big.NewInt(0)
}

// TxFee : blockBaseFee - header.BaseFee
func (tx *Transaction) TxFee(blockBaseFee *big.Int) *big.Int {
	gasUsed := big.NewInt(int64(tx.receipt.GasUsed()))
	switch tx.Type() {
	case TxTypeLegacyTransaction, TxTypeEthereumAccessList:
		return new(big.Int).Mul(tx.GasPrice(), gasUsed)
	case TxTypeKlaytnFeeDelegatedTransaction, TxTypeKlaytnTransaction:
		if blockBaseFee != nil {
			return new(big.Int).Mul(blockBaseFee, gasUsed)
		}

		if tx.GasFeeCap() != nil {
			return new(big.Int).Mul(tx.GasFeeCap(), gasUsed)
		}

		return new(big.Int).Mul(tx.GasPrice(), gasUsed)
	case TxTypeEthereumDynamicFee:
		return new(big.Int).Mul(new(big.Int).Add(blockBaseFee, tx.GasTipCap()), gasUsed)
	case TxTypeTronTransaction:
		return big.NewInt(0)
	}

	return big.NewInt(0)
}

func (tx *Transaction) GasFeeCap() *big.Int {
	switch v := tx.inner.(type) {
	case *klaytypes.Transaction:
		return v.GasFeeCap()
	case *ethertypes.Transaction:
		return v.GasFeeCap()
	case *arbtypes.Transaction:
		return v.GasFeeCap()
	case *basetypes.Transaction:
		return v.GasFeeCap()
	}
	return big.NewInt(0)
}

func (tx *Transaction) GasTipCap() *big.Int {
	switch v := tx.inner.(type) {
	case *klaytypes.Transaction:
		return v.GasTipCap()
	case *ethertypes.Transaction:
		return v.GasTipCap()
	case *arbtypes.Transaction:
		return v.GasTipCap()
	case *basetypes.Transaction:
		return v.GasTipCap()
	}
	return big.NewInt(0)
}

func (tx *Transaction) Receipt() *Receipt {
	return tx.receipt
}

func (tx *Transaction) SignTx(account *Account, chainId *big.Int) (*Transaction, error) {
	switch v := tx.inner.(type) {
	case *klaytypes.Transaction:
		signed, err := klaytypes.SignTx(v, klaytypes.LatestSignerForChainID(chainId), account.PrivateKey)
		if err != nil {
			return nil, err
		}
		tx.inner = signed
		return tx, nil
	case *ethertypes.Transaction:
		signed, err := ethertypes.SignTx(v, ethertypes.LatestSignerForChainID(chainId), account.PrivateKey)
		if err != nil {
			return nil, err
		}
		tx.inner = signed
		return tx, nil
	case *arbtypes.Transaction:
		signed, err := arbtypes.SignTx(v, arbtypes.LatestSignerForChainID(chainId), account.PrivateKey)
		if err != nil {
			return nil, err
		}
		tx.inner = signed
		return tx, nil
	case *basetypes.Transaction:
		signed, err := basetypes.SignTx(v, basetypes.LatestSignerForChainID(chainId), account.PrivateKey)
		if err != nil {
			return nil, err
		}
		tx.inner = signed
		return tx, nil
	case *troncore.Transaction:
		raw, err := proto.Marshal(v.GetRawData())
		if err != nil {
			return nil, err
		}

		sha256hash := sha256.New()
		sha256hash.Write(raw)
		hash := sha256hash.Sum(nil)
		sig, err := crypto.Sign(hash[:], account.PrivateKey)
		if err != nil {
			return nil, err
		}

		v.Signature = append(v.Signature, sig)
		tx.inner = v
		return tx, nil
	case *cosmostxtypes.Tx:
	}

	return nil, fmt.Errorf("invalid tx")
}

func (tx *Transaction) SignTxKms(signer *bind.TransactOpts, chainId *big.Int) (*Transaction, error) {
	logger.Debug("SignTxKms", logger.BuildLogInput().WithData("signer", signer, "chainId", chainId, "inner", tx.inner, "tx", tx))
	switch v := tx.inner.(type) {
	case *klaytypes.Transaction:
		nonce, err := tx.Nonce()
		if err != nil {
			return nil, fmt.Errorf("tx.Nonce: %w", err)
		}
		toAddress := ethercommon.HexToAddress(tx.To())
		newTx := NewTx(tx.Chain(), &ethertypes.DynamicFeeTx{
			Nonce:     nonce,
			GasFeeCap: tx.GasPrice(),
			GasTipCap: tx.GasPrice(),
			Gas:       GasLimit,
			To:        &toAddress,
			Value:     big.NewInt(0),
			Data:      tx.Data(),
		})
		ethTx, _ := newTx.EthereumTransaction()

		signedTx, err := signer.Signer(signer.From, ethTx)
		if err != nil {
			return nil, fmt.Errorf("signer.Signer: %w", err)
		}

		klayTypeTo := klaycommon.HexToAddress(newTx.To())
		accessList := klaytypes.AccessList{}

		txData := map[klaytypes.TxValueKeyType]interface{}{
			klaytypes.TxValueKeyChainID:    chainId,
			klaytypes.TxValueKeyTo:         &klayTypeTo,
			klaytypes.TxValueKeyNonce:      nonce,
			klaytypes.TxValueKeyGasFeeCap:  tx.GasPrice(),
			klaytypes.TxValueKeyGasTipCap:  tx.GasPrice(),
			klaytypes.TxValueKeyGasLimit:   uint64(GasLimit),
			klaytypes.TxValueKeyAmount:     big.NewInt(0),
			klaytypes.TxValueKeyData:       tx.Data(),
			klaytypes.TxValueKeyAccessList: accessList,
		}
		// only for compatible return value
		retTx, err := klaytypes.NewTransactionWithMap(klaytypes.TxTypeEthereumDynamicFee, txData)
		if err != nil {
			return nil, fmt.Errorf("NewTransactionWithMap: %w", err)
		}

		vv, r, s := signedTx.RawSignatureValues()
		signatures := klaytypes.TxSignatures{
			&klaytypes.TxSignature{V: vv, R: r, S: s},
		}
		retTx.SetSignature(signatures)
		tx.inner = retTx
		return tx, nil

	case *ethertypes.Transaction:
		signed, err := signer.Signer(signer.From, v)
		if err != nil {
			return nil, err
		}
		tx.inner = signed

		return tx, nil
	case *arbtypes.Transaction:
		// TODO: implement me
	case *basetypes.Transaction:
		nonce, err := tx.Nonce()
		if err != nil {
			return nil, fmt.Errorf("tx.Nonce: %w", err)
		}

		toAddress := ethercommon.HexToAddress(tx.To())
		newTx := NewTx(tx.Chain(), &ethertypes.DynamicFeeTx{
			Nonce:     nonce,
			GasFeeCap: tx.GasPrice(),
			GasTipCap: tx.GasTipCap(),
			Gas:       GasLimit,
			To:        &toAddress,
			Value:     big.NewInt(0),
			Data:      tx.Data(),
		})

		etherTx, err := newTx.EthereumTransaction()
		if err != nil {
			return nil, fmt.Errorf("newTx.EthereumTransaction: %w", err)
		}

		signed, err := signer.Signer(signer.From, etherTx)
		if err != nil {
			return nil, fmt.Errorf("signer.Signer: %w", err)
		}

		value, err := tx.Value()
		if err != nil {
			return nil, fmt.Errorf("tx.Value: %w", err)
		}

		baseTypeTo := basecommon.HexToAddress(tx.To())
		baseTx := basetypes.NewTx(&basetypes.DynamicFeeTx{
			ChainID:    chainId,
			Nonce:      nonce,
			GasTipCap:  tx.GasTipCap(),
			GasFeeCap:  tx.GasPrice(),
			Gas:        tx.Gas(),
			To:         &baseTypeTo,
			Value:      value,
			Data:       tx.Data(),
			AccessList: basetypes.AccessList{},
		})

		vv, r, s := signed.RawSignatureValues()
		signature := append(r.Bytes(), s.Bytes()...)
		// v 값이 0인 경우 bytes 로 변환시 빈배열 이므로 0을 넣어 자릿수 맞춰줌
		if vv.Cmp(big.NewInt(0)) == 0 {
			signature = append(signature[:], []byte{0}...)
		} else {
			signature = append(signature[:], vv.Bytes()...)
		}

		logger.Debug("SignTxKms_WithSignature_Before",
			logger.BuildLogInput().WithData(
				"signer", signer,
				"etherTx", etherTx,
				"signed", signed,
				"vv", vv,
				"r", r,
				"s", s,
				"signature", signature,
				"tx.ChainId", tx.ChainId(),
				"paramChainId", chainId,
				"tx", tx,
			),
		)

		baseTx, err = baseTx.WithSignature(basetypes.LatestSignerForChainID(chainId), signature)
		if err != nil {
			return nil, fmt.Errorf("WithSignature: %w", err)
		}

		tx.inner = baseTx
		return tx, nil
	}

	return nil, fmt.Errorf("invalid tx")
}

func (tx *Transaction) RlpTx() (string, error) {
	var data []byte
	var err error

	switch v := tx.inner.(type) {
	case *klaytypes.Transaction:
		data, err = v.MarshalBinary()
		if err != nil {
			return "", err
		}
	case *ethertypes.Transaction:
		data, err = v.MarshalBinary()
		if err != nil {
			return "", err
		}
	case *arbtypes.Transaction:
		data, err = v.MarshalBinary()
		if err != nil {
			return "", err
		}
	case *basetypes.Transaction:
		data, err = baserlp.EncodeToBytes(v)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("invalid tx")
	}

	return HexUtilEncode(tx.Chain(), data), nil
}

func (tx *Transaction) KlaytnTransaction() (*klaytypes.Transaction, error) {
	rlpTx, err := tx.RlpTx()
	if err != nil {
		return nil, err
	}

	var raw []byte
	var fn func(string) ([]byte, error)
	switch v := tx.inner.(type) {
	case *ethertypes.Transaction:
		fn = etherhexutil.Decode
	case *arbtypes.Transaction:
		fn = arbhexutil.Decode
	case *basetypes.Transaction:
		fn = basehexutil.Decode
	case *klaytypes.Transaction:
		return v, nil
	}

	raw, err = fn(rlpTx)
	if err != nil {
		return nil, err
	}

	klayTx := new(klaytypes.Transaction)
	if err = klayrlp.DecodeBytes(raw, klayTx); err != nil {
		return nil, err
	}

	return klayTx, nil
}

func (tx *Transaction) EthereumTransaction() (*ethertypes.Transaction, error) {
	rlpTx, err := tx.RlpTx()
	if err != nil {
		return nil, err
	}

	var raw []byte
	var fn func(string) ([]byte, error)
	switch v := tx.inner.(type) {
	case *klaytypes.Transaction:
		fn = klayhexutil.Decode
	case *arbtypes.Transaction:
		fn = arbhexutil.Decode
	case *basetypes.Transaction:
		fn = basehexutil.Decode
	case *ethertypes.Transaction:
		return v, nil
	}

	raw, err = fn(rlpTx)
	if err != nil {
		return nil, err
	}

	etherTx := new(ethertypes.Transaction)
	if err = etherrlp.DecodeBytes(raw, etherTx); err != nil {
		return nil, err
	}

	return etherTx, nil
}

func (tx *Transaction) ArbitrumTransaction() (*arbtypes.Transaction, error) {
	rlpTx, err := tx.RlpTx()
	if err != nil {
		return nil, err
	}

	var raw []byte
	var fn func(string) ([]byte, error)
	switch v := tx.inner.(type) {
	case *klaytypes.Transaction:
		fn = klayhexutil.Decode
	case *ethertypes.Transaction:
		fn = etherhexutil.Decode
	case *basetypes.Transaction:
		fn = basehexutil.Decode
	case *arbtypes.Transaction:
		return v, nil
	}

	raw, err = fn(rlpTx)
	if err != nil {
		return nil, err
	}

	arbTx := new(arbtypes.Transaction)
	if err = arbrlp.DecodeBytes(raw, arbTx); err != nil {
		return nil, err
	}

	return arbTx, nil
}

func (tx *Transaction) BaseTransaction() (*basetypes.Transaction, error) {
	rlpTx, err := tx.RlpTx()
	if err != nil {
		return nil, err
	}

	//TODO 해당 부분 테스트
	raw, err := HexUtilDecode(tx.Chain(), rlpTx)
	if err != nil {
		return nil, err
	}

	baseTx := new(basetypes.Transaction)
	if err = baserlp.DecodeBytes(raw, baseTx); err != nil {
		return nil, err
	}

	return baseTx, nil
}

func (tx *Transaction) CosmosTransaction() (*cosmostxtypes.Tx, error) {
	switch v := tx.inner.(type) {
	case *cosmostxtypes.Tx:
		return v, nil
	}
	return nil, fmt.Errorf("invalid tx")
}

func (tx *Transaction) TronTransaction() (*troncore.Transaction, error) {
	switch v := tx.inner.(type) {
	case *troncore.Transaction:
		return v, nil
	}
	return nil, fmt.Errorf("invalid tx")
}

func (tx *Transaction) MarshalJSON() ([]byte, error) {
	switch v := tx.inner.(type) {
	case *klaytypes.Transaction:
		return json.MarshalIndent(v, "", " ")
	case *ethertypes.Transaction:
		return json.MarshalIndent(v, "", " ")
	case *arbtypes.Transaction:
		return json.MarshalIndent(v, "", " ")
	case *basetypes.Transaction:
		return json.MarshalIndent(v, "", " ")
	case *troncore.Transaction:
		return json.MarshalIndent(v, "", " ")
	case *cosmostxtypes.Tx:
		return json.MarshalIndent(v, "", " ")
	}
	return nil, fmt.Errorf("invalid tx")
}

func (tx *Transaction) Nonce() (uint64, error) {
	switch v := tx.inner.(type) {
	case *klaytypes.Transaction:
		return v.Nonce(), nil
	case *ethertypes.Transaction:
		return v.Nonce(), nil
	case *arbtypes.Transaction:
		return v.Nonce(), nil
	case *basetypes.Transaction:
		return v.Nonce(), nil
	case *cosmostxtypes.Tx:
		return v.GetAuthInfo().GetSignerInfos()[0].Sequence, nil
	}
	return 0, fmt.Errorf("invalid tx")
}

func (tx *Transaction) Value() (*big.Int, error) {
	switch v := tx.inner.(type) {
	case *klaytypes.Transaction:
		return v.Value(), nil
	case *ethertypes.Transaction:
		return v.Value(), nil
	case *arbtypes.Transaction:
		return v.Value(), nil
	case *basetypes.Transaction:
		return v.Value(), nil
	}
	return big.NewInt(0), fmt.Errorf("invalid tx")
}
