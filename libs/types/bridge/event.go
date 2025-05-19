package bridge

import (
	"fmt"
	"github.com/curtis0505/bridge/libs/common"
	"math/big"
)

type EventDeposit struct {
	ToChainName  string         `json:"toChainName"`
	FromAddr     common.Address `json:"fromAddr"`
	To           common.Bytes   `json:"to"`
	TokenAddr    common.Address `json:"tokenAddr"`
	Decimal      uint8          `json:"decimal"`
	Amount       *big.Int       `json:"amount"`
	DepositNonce *big.Int       `json:"depositNonce"`

	//restake bridge
	RestakeTokenAddr   common.Address `json:"restakeTokenAddr"`
	RestakeTokenAmount *big.Int       `json:"restakeTokenAmount"`
}

func (event EventDeposit) BridgeVersion() Version {
	if event.TokenAddr != nil {
		return VersionBridge
	}

	if event.RestakeTokenAddr != nil {
		return VersionRestakeBridge
	}

	return ""
}

func (event EventDeposit) SafeTokenAddr() (common.Address, error) {
	if event.TokenAddr != nil {
		return event.TokenAddr, nil
	}

	if event.RestakeTokenAddr != nil {
		return event.RestakeTokenAddr, nil
	}

	return nil, fmt.Errorf("token address is empty")
}

func (event EventDeposit) SafeAmount() (*big.Int, error) {
	if event.Amount != nil {
		return event.Amount, nil
	}

	if event.RestakeTokenAmount != nil {
		return event.RestakeTokenAmount, nil
	}

	return nil, fmt.Errorf("token amount is empty")
}

type EventBurn struct {
	ToChainName string         `json:"toChainName"`
	FromAddr    common.Address `json:"fromAddr"`
	To          common.Bytes   `json:"to"`
	Token       common.Bytes   `json:"token"`
	TokenAddr   common.Address `json:"tokenAddr"`
	Decimal     uint8          `json:"decimal"`
	Amount      *big.Int       `json:"amount"`
	BurnNonce   *big.Int       `json:"burnNonce"`

	//restake bridge
	RestakeToken common.Bytes   `json:"restakeToken"`
	WTokenAddr   common.Address `json:"wTokenAddr"`
	IsDirect     bool           `json:"isDirect"`
}

func (event EventBurn) SafeToken() (common.Bytes, error) {
	if len(event.Token) > 0 {
		return event.Token, nil
	}

	if len(event.RestakeToken) > 0 {
		return event.RestakeToken, nil
	}

	return nil, fmt.Errorf("token is empty")
}

func (event EventBurn) SafeTokenAddr() (common.Address, error) {
	if event.TokenAddr != nil {
		return event.TokenAddr, nil
	}

	if event.WTokenAddr != nil {
		return event.WTokenAddr, nil
	}

	return nil, fmt.Errorf("token address is empty")
}

type EventMint struct {
	FromChainName string
	From          common.Bytes
	To            common.Bytes
	TokenAddr     common.Address
	Bytes32s      []common.Hash
	Uints         []*big.Int
}

type EventWithdraw struct {
	FromChainName string
	From          common.Bytes
	To            common.Bytes
	Token         common.Bytes
	Bytes32s      []common.Hash
	Uints         []*big.Int
}

type EventSubmission struct {
	TransactionId *big.Int
	TxHash        common.Hash
	Required      *big.Int
}

type EventConfirmation struct {
	Sender        common.Address
	TransactionId *big.Int
}

type EventExecution struct {
	TransactionId *big.Int
	TxHash        common.Hash
}

type EventExecutionFailure struct {
	TransactionId *big.Int
	TxHash        common.Hash
	Err           common.Bytes
}

type EventFxDepositERC20 struct {
	RootToken    common.Address
	Depositor    common.Address
	UserAddress  common.Address
	Amount       *big.Int
	DepositNonce *big.Int
}

type EventFxSyncDeposit struct {
	MsgSender    common.Address
	RootToken    common.Address
	Depositor    common.Address
	To           common.Address
	Amount       *big.Int
	DepositData  []byte
	DepositNonce *big.Int
}

type EventFxChildWithdrawERC20 struct {
	MsgSender  common.Address
	RootToken  common.Address
	ChildToken common.Address
	Receiver   common.Address
	Amount     *big.Int
}

type EventFxWithdrawERC20 struct {
	RootToken   common.Address
	ChildToken  common.Address
	UserAddress common.Address
	Amount      *big.Int
}

type EventNewHeaderBlock struct {
	Proposer      common.Address
	HeaderBlockId *big.Int
	Reward        *big.Int
	Start         *big.Int
	End           *big.Int
}
