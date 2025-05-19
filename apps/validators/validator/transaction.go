package validator

import (
	"context"
	"errors"
	"sync"
	"time"
)

type PendingTransaction struct {
	rWMutex sync.RWMutex
	History map[string]*TransactionHistory
}

func NewPendingTransaction() PendingTransaction {
	return PendingTransaction{
		History: make(map[string]*TransactionHistory),
	}
}

func (pendingTransaction *PendingTransaction) Get(txHash string) (*TransactionHistory, error) {
	pendingTransaction.rWMutex.RLock()
	defer pendingTransaction.rWMutex.RUnlock()

	result, ok := pendingTransaction.History[txHash]
	if ok == false {
		return nil, errors.New("not found txHash")
	}

	return result, nil
}
func (pendingTransaction *PendingTransaction) List() map[string]*TransactionHistory {
	pendingTransaction.rWMutex.RLock()
	defer pendingTransaction.rWMutex.RUnlock()

	return pendingTransaction.History
}

func (pendingTransaction *PendingTransaction) Set(transaction *TransactionHistory) {
	pendingTransaction.rWMutex.Lock()
	defer pendingTransaction.rWMutex.Unlock()

	pendingTransaction.History[transaction.TxHash] = transaction
}
func (pendingTransaction *PendingTransaction) Delete(txHash string) {
	pendingTransaction.rWMutex.Lock()
	defer pendingTransaction.rWMutex.Unlock()

	delete(pendingTransaction.History, txHash)
}

/*
CheckTransactionHistory
1분마다 한번씩 펜딩 트랜잭션들을 조회하여
온체인되고, 에러가 없으면 트랜잭션 삭제
*/
func (p *Validator) CheckTransactionHistory() {
	defer func() {
		if v := recover(); v != nil {
			p.logger.Error("panic recovered", v)
			return
		}
	}()

	ctx := context.Background()
	for {
		for txHash, history := range p.PendingTransaction.List() {
			_, isPending, err := p.client.GetTransactionWithReceipt(ctx, history.ChainName, txHash)
			if err == nil && isPending == false {
				p.PendingTransaction.Delete(txHash)
				p.logger.Info("CheckTransactionHistory delete txHash", txHash)
			}

			nonce, err := p.client.NonceAt(ctx, history.ChainName, p.Account[history.ChainName].Address)
			if err != nil {
				break
			}

			if nonce > history.Nonce {
				p.PendingTransaction.Delete(txHash)
				p.logger.Info("CheckTransactionHistory delete txHash", txHash, ", network nonce", nonce)
			}

		}
		time.Sleep(time.Minute)
	}
}
