package types

import (
	"fmt"
	"github.com/curtis0505/bridge/apps/managers/util"
	protocol "github.com/curtis0505/bridge/libs/dto"
	mongoentity "github.com/curtis0505/bridge/libs/entity/mongo"
	"math/big"
	"time"
)

const (
	EventWithdraw = "Withdraw"
	EventDeposit  = "Deposit"
	EventBurn     = "Burn"
	EventMint     = "Mint"

	EventPaused   = "Paused"
	EventUnpaused = "Unpaused"

	EventConfirmation     = "Confirmation"
	EventSubmission       = "Submission"
	EventExecution        = "Execution"
	EventExecutionFailure = "ExecutionFailure"

	EventFxDepositERC20       = "FxDepositERC20"
	EventSyncDeposit          = "SyncDeposit"
	EventFxChildWithdrawERC20 = "FxChildWithdrawERC20"
	EventFxWithdrawERC20      = "FxWithdrawERC20"
	EventNewHeaderBlock       = "NewHeaderBlock"
)

const (
	LogEventBridgeTransfer = "BridgeTransfer"
	LogEventTransfer       = "Transfer"
	LogEventFeeTransfer    = "FeeTransfer"

	TxHistoryBridgeChainFee = "BRIDGE-CHAINFEE"
)

type BridgeTransferInfo struct {
	Token  mongoentity.TokenInfo
	Chain  string
	From   string
	To     string
	TxHash string
	Cate   string
	Amount *big.Int

	FromChain  string
	FromTxHash string
	ToChain    string
	ToTxHash   string
	CreateAt   time.Time
}

func NewTransferInfo() *BridgeTransferInfo {
	t := BridgeTransferInfo{}
	return &t
}

func (t *BridgeTransferInfo) SetChain(chain string) *BridgeTransferInfo {
	t.Chain = chain
	return t
}

func (t *BridgeTransferInfo) SetToken(token mongoentity.TokenInfo) *BridgeTransferInfo {
	t.Token = token
	return t
}

func (t *BridgeTransferInfo) SetCate(cate string) *BridgeTransferInfo {
	t.Cate = cate
	return t
}

func (t *BridgeTransferInfo) SetFrom(from string) *BridgeTransferInfo {
	t.From = from
	return t
}

func (t *BridgeTransferInfo) SetTo(to string) *BridgeTransferInfo {
	t.To = to
	return t
}

func (t *BridgeTransferInfo) SetTxHash(txHash string) *BridgeTransferInfo {
	t.TxHash = txHash
	return t
}

func (t *BridgeTransferInfo) SetFromChain(chain string) *BridgeTransferInfo {
	t.FromChain = chain
	return t
}

func (t *BridgeTransferInfo) SetFromTxHash(txHash string) *BridgeTransferInfo {
	t.FromTxHash = txHash
	return t
}

func (t *BridgeTransferInfo) SetToChain(chain string) *BridgeTransferInfo {
	t.ToChain = chain
	return t
}

func (t *BridgeTransferInfo) SetToTxHash(txHash string) *BridgeTransferInfo {
	t.ToTxHash = txHash
	return t
}

func (t *BridgeTransferInfo) SetAmount(amount *big.Int) *BridgeTransferInfo {
	t.Amount = amount
	return t
}

func (t *BridgeTransferInfo) String() string {
	return fmt.Sprintf("➡️ %s %s%s",
		t.Chain, util.ToEtherWithDecimal(t.Amount, uint64(t.Token.Decimal)).String(), t.Token.Symbol,
	)
}

func (t *BridgeTransferInfo) HistoryData() *protocol.ScBridgeTransferHistoryData {
	return &protocol.ScBridgeTransferHistoryData{
		Chain:      t.Chain,
		From:       t.From,
		To:         t.To,
		TxHash:     t.TxHash,
		Cate:       t.Cate,
		CurrencyId: t.Token.CurrencyID,
		GroupId:    t.Token.GroupID,
		Symbol:     t.Token.Symbol,
		Amount:     util.SafeBigIntToDecimal128(t.Amount),
		FromChain:  t.FromChain,
		FromTxHash: t.FromTxHash,
		ToChain:    t.ToChain,
		ToTxHash:   t.ToTxHash,
		CreateAt:   time.Now(),
	}
}
