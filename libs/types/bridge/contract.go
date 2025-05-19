package bridge

import (
	"github.com/curtis0505/bridge/libs/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type Minter interface {
	GetChainId(chainName string) (string, error)
	GetChainFee(chainId string) (*big.Int, error)
	GetTaxRateBP() (*big.Int, error)
	IsSupportChain(chainId string) (bool, error)
	GetInputDataMint(fromChainName string, from []byte, toAddr string, token []byte, chainId string, txHash string, amount, decimal *big.Int) ([]byte, error)
	Burn(tokenAddr string, toChainName string, to []byte, amount *big.Int, coinAmount *big.Int, proof []byte, account *types.Account) (*types.Transaction, error)
	GetMinterAddressByProof() []byte
	GetFeeTax(toChainName string) (*big.Int, *big.Int, error)
	GetAddress() string
}

type MultiSigWallet interface {
	GetConfirmation(transactionId *big.Int, address string) (bool, error)
	GetTxId(txHash string) (*big.Int, error)
	GetTransaction(transactionId *big.Int) (*TransactionResponse, error)
	GetTransactionCount(pending, executed bool) (*big.Int, error)
	GetValidatorAddresses() ([]string, error)
	GetConfirmationCount(transactionId *big.Int) (*big.Int, error)
	GetProof(txHash string, minterAddress string, value *big.Int, mintInputData []byte) []byte
	GetSubmitTransactionData(txHash, destination string, value *big.Int, proof, data []byte) ([]byte, error)
	SubmitTransaction(txHash, destination string, value *big.Int, proof, data []byte, account *types.Account) (*types.Transaction, error)
	SubmitTransactionWithNonce(txHash, destination string, value *big.Int, proof, data []byte, account *types.Account, nonce uint64) (*types.Transaction, error)
	SubmitTransactionKMS(txHash, destination string, value *big.Int, proof, data []byte) (*types.Transaction, error)
	SubmitTransactionKMSWithNonce(txHash, destination string, value *big.Int, proof, data []byte, nonce uint64) (*types.Transaction, error)
	ConfirmTransaction(txHash [32]byte, transactionId *big.Int, proof [32]byte, account *types.Account) (*types.Transaction, error)
	ConfirmTransactionKMS(txHash [32]byte, transactionId *big.Int, proof [32]byte) (*types.Transaction, error)

	GetAddress() string
}

type Vault interface {
	GetChainId(chainName string) (string, error)
	GetChainFee(chainId string) (*big.Int, error)
	GetTaxRateBP() (*big.Int, error)
	IsSupportChain(chainId string) (bool, error)
	Deposit(toChainName string, to []byte, amount *big.Int, proof []byte, account *types.Account) (*types.Transaction, error)
	DepositToken(tokenAddr string, toChainName string, to []byte, amount *big.Int, proof []byte, account *types.Account) (*types.Transaction, error)
	GetInputDataWithdraw(fromChainName string, from []byte, toAddr string, token []byte, chainId string, txHash string, amount, decimal *big.Int) ([]byte, error)
	GetVaultAddressByProof() []byte
	GetFeeTax(toChainName string) (*big.Int, *big.Int, error)
	GetAddress() string
}

type ERC20 interface {
	Approve(spender string, amount *big.Int, account *types.Account) (*types.Transaction, error)
	ApproveWithNonce(spender string, amount *big.Int, nonce uint64, account *types.Account) (*types.Transaction, error)
	IncreaseAllowance(spender string, amount *big.Int, account *types.Account) (*types.Transaction, error)
	Transfer(recipient string, amount *big.Int, account *types.Account) (*types.Transaction, error)
	BalanceOf(address string) (*big.Int, error)
	Allowance(owner, spender string) (*big.Int, error)
	Symbol() (symbol string, err error)
}

type FxERC20RootTunnel interface {
	Deposit(tokenAddr string, user []byte, amount *big.Int, data []byte, account *types.Account) (*types.Transaction, error)
	ReceiveMessage(inputData []byte, account *types.Account) (*types.Transaction, error)
	GetChainFee() (*big.Int, error)
	GetTaxRateBP() (*big.Int, error)
}

type FxERC20ChildTunnel interface {
	Withdraw(childTokenAddr string, amount *big.Int, account *types.Account) (*types.Transaction, error)
	WithdrawTo(childTokenAddr string, receiverAddr string, amount *big.Int, account *types.Account) (*types.Transaction, error)
	MapToken() (*types.Transaction, error)
	GetChainFee() (*big.Int, error)
	GetTaxRateBP() (*big.Int, error)
}

type TransactionResponse struct {
	Proof    [32]byte
	Executed bool
}

func (t *TransactionResponse) ProofString() string {
	return hexutil.Encode(t.Proof[:])
}

type MultiSigOperation interface {
	GetConfirmation(transactionId *big.Int, address string) (bool, error)
	GetTxId(txHash string) (*big.Int, error)
	GetTransaction(transactionId *big.Int) (*TransactionResponse, error)
	GetTransactionCount(pending, executed bool) (*big.Int, error)
	GetValidatorAddresses() ([]string, error)
	GetConfirmationCount(transactionId *big.Int) (*big.Int, error)
	GetProof(txHash string, minterAddress string, value *big.Int, mintInputData []byte) []byte
	GetSubmitTransactionData(txHash, destination string, value *big.Int, proof, data []byte) ([]byte, error)
	SubmitTransaction(txHash, destination string, value *big.Int, proof, data []byte, account *types.Account) (*types.Transaction, error)
	SubmitTransactionWithNonce(txHash, destination string, value *big.Int, proof, data []byte, account *types.Account, nonce uint64) (*types.Transaction, error)
	SubmitTransactionKMS(txHash, destination string, value *big.Int, proof, data []byte) (*types.Transaction, error)
	SubmitTransactionKMSWithNonce(txHash, destination string, value *big.Int, proof, data []byte, nonce uint64) (*types.Transaction, error)
	ConfirmTransaction(txHash [32]byte, transactionId *big.Int, proof [32]byte, account *types.Account) (*types.Transaction, error)
	ConfirmTransactionKMS(txHash [32]byte, transactionId *big.Int, proof [32]byte) (*types.Transaction, error)

	GetAddress() string
}
