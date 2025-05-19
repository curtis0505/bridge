package rest

import (
	"fmt"
	commontypes "github.com/curtis0505/bridge/libs/types"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	"github.com/ethereum/go-ethereum/common"

	"math/big"
)

type MultiSigWallet struct {
	Client  *client
	Address string
}

func newMultiSigWallet(client *client, address string) *MultiSigWallet {
	return &MultiSigWallet{
		Client:  client,
		Address: address,
	}
}

func (m *MultiSigWallet) GetAddress() string {
	return m.Address
}

func (m *MultiSigWallet) GetValidatorAddresses() ([]string, error) {
	return nil, fmt.Errorf("")
}

func (m *MultiSigWallet) GetTransactionIds(from, to *big.Int, pending, executed bool) ([]*big.Int, error) {
	return nil, fmt.Errorf("not supported")
}

func (m *MultiSigWallet) GetTransaction(transactionId *big.Int) (*bridgetypes.TransactionResponse, error) {
	return nil, fmt.Errorf("not supported")
}

func (m *MultiSigWallet) GetConfirmation(transactionId *big.Int, address string) (bool, error) {
	return false, fmt.Errorf("not supported")
}

func (m *MultiSigWallet) GetTxId(txHash string) (*big.Int, error) {
	return nil, fmt.Errorf("not supported")
}

func (m *MultiSigWallet) GetTransactionCount(pending, executed bool) (*big.Int, error) {
	return nil, fmt.Errorf("not supported")
}

func (m *MultiSigWallet) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return nil, fmt.Errorf("not supported")
}

func (m *MultiSigWallet) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return nil, fmt.Errorf("not supported")
}

func (m *MultiSigWallet) GetProof(txHash string, minterAddress string, value *big.Int, mintInputData []byte) []byte {
	return nil
}

func (m *MultiSigWallet) GetSubmitTransactionData(txHash, destination string, value *big.Int, proof, data []byte) ([]byte, error) {
	return nil, fmt.Errorf("not supported")
}

func (m *MultiSigWallet) SubmitTransaction(txHash, destination string, value *big.Int, proof, data []byte, account *commontypes.Account) (*commontypes.Transaction, error) {
	return nil, fmt.Errorf("not supported")
}

func (m *MultiSigWallet) SubmitTransactionWithNonce(txHash, destination string, value *big.Int, proof, data []byte, account *commontypes.Account, nonce uint64) (*commontypes.Transaction, error) {
	return nil, fmt.Errorf("not supported")
}

func (m *MultiSigWallet) SubmitTransactionKMS(txHash, destination string, value *big.Int, proof, data []byte) (*commontypes.Transaction, error) {
	return nil, fmt.Errorf("not supported")
}

func (m *MultiSigWallet) SubmitTransactionKMSWithNonce(txHash, destination string, value *big.Int, proof, data []byte, nonce uint64) (*commontypes.Transaction, error) {
	return nil, fmt.Errorf("not supported")
}

func (m *MultiSigWallet) ConfirmTransaction(txHash [32]byte, transactionId *big.Int, proof [32]byte, account *commontypes.Account) (*commontypes.Transaction, error) {
	return nil, fmt.Errorf("not supported")
}

func (m *MultiSigWallet) ConfirmTransactionKMS(txHash [32]byte, transactionId *big.Int, proof [32]byte) (*commontypes.Transaction, error) {
	return nil, fmt.Errorf("not supported")
}
