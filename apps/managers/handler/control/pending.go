package control

import (
	"fmt"
	"time"
)

func (p *ControlHandler) AddPendingTx(chain, event, txHash string) {
	p.pendingLock.Lock()
	defer p.pendingLock.Unlock()

	p.pending[p.pendingKey(chain, event)] = PendingTx{
		TxHash:    txHash,
		Chain:     chain,
		Timestamp: time.Now(),
	}
}

func (p *ControlHandler) RemovePendingTx(chain, event string) {
	p.pendingLock.Lock()
	defer p.pendingLock.Unlock()

	delete(p.pending, p.pendingKey(chain, event))
}

func (p *ControlHandler) CheckPendingTx(chain, event string) bool {
	p.pendingLock.Lock()
	defer p.pendingLock.Unlock()

	tx := p.pending[p.pendingKey(chain, event)]
	return tx.TxHash != ""
}

func (p *ControlHandler) CheckConfirmTx() {
	p.pendingLock.Lock()
	defer p.pendingLock.Unlock()

	for key, pendingTx := range p.pending {
		diff := time.Now().Sub(pendingTx.Timestamp)

		//tx, pending, err := p.client.GetTransactionWithReceipt(context.Background(), pendingTx.Chain, pendingTx.TxHash)
		//if err != nil {
		//	p.logger.Warn("event", "CheckConfirmTx", "chain", pendingTx.Chain, "event", key, "txHash", pendingTx.TxHash, "warn", "notfound", "time", diff.String())
		//	continue
		//}
		//
		//if pending {
		//	p.logger.Info("event", "CheckConfirmTx", "chain", pendingTx.Chain, "event", key, "txHash", pendingTx.TxHash, "time", diff.String())
		//	continue
		//}
		//
		//if tx != nil {
		//	p.logger.Info("event", "CheckConfirmTx", "chain", pendingTx.Chain, "event", key, "txHash", pendingTx.TxHash, "message", "confirmed", "time", diff.String())
		//	delete(p.pending, key)
		//	continue
		//}

		if diff > time.Second*time.Duration(p.cfg.Control.BridgePauseTimeOut) {
			p.logger.Warn("event", "CheckConfirmTx", "chain", pendingTx.Chain, "event", key, "txHash", pendingTx.TxHash, "warn", "timeout", "time", diff.String())
			delete(p.pending, key)
			continue
		}

		p.logger.Info(key, pendingTx)
	}

	p.logger.Trace("event", "CheckConfirmTx", "pendingTx", len(p.pending))
}

func (p *ControlHandler) pendingKey(chain, event string) string {
	return fmt.Sprintf("%s/%s", chain, event)
}
