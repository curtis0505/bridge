package validator

import (
	validatortypes "github.com/curtis0505/bridge/apps/managers/handler/validator/types"
	bridgetypes "github.com/curtis0505/bridge/libs/types"
	"time"
)

func (p *ValidatorHandler) PendingTx(tx *bridgetypes.Transaction) {
	p.pendingTxMapLock.Lock()
	defer p.pendingTxMapLock.Unlock()

	pendingTx := &validatortypes.PendingTx{
		ValidatorInfo:    p.GetValidatorInfo(tx.Chain(), tx.From()),
		Chain:            tx.Chain(),
		TxHash:           tx.TxHash(),
		Time:             time.Now(),
		ValidatorSummary: p.validatorSummary(tx.Chain(), tx.From()),
	}

	if pendingTx.ValidatorInfo == nil || pendingTx.ValidatorSummary == nil {
		p.logger.Error(
			"event", "PendingTx", "chain", pendingTx.Chain, "validator", "unknown", "txHash", pendingTx.TxHash,
			"pending", len(p.pendingTxMap),
		)
		return
	}

	p.pendingTxMap[pendingTx.TxHash] = pendingTx

	p.logger.Info(
		"event", "PendingTx", "chain", pendingTx.Chain, "validator", pendingTx.ValidatorInfo.Name, "txHash", pendingTx.TxHash,
		"pending", len(p.pendingTxMap),
	)
}

func (p *ValidatorHandler) ConfirmTx(tx *bridgetypes.Transaction) {
	p.pendingTxMapLock.Lock()
	defer p.pendingTxMapLock.Unlock()

	pendingTx, ok := p.pendingTxMap[tx.TxHash()]
	if !ok {
		p.logger.Warn("event", "PendingTx", "chain", tx.Chain(), "msg", "not exist tx", "txHash", tx.TxHash())
		return
	}
	delete(p.pendingTxMap, tx.TxHash())

	p.logger.Info(
		"event", "ConfirmTx", "chain", tx.Chain(), "validator", pendingTx.ValidatorInfo.Name, "txHash", tx.TxHash(),
		"pending", len(p.pendingTxMap),
	)
}

func (p *ValidatorHandler) CheckSpeedUpTx() {
	p.pendingTxMapLock.Lock()
	defer p.pendingTxMapLock.Unlock()

	p.logger.Trace("event", "CheckSpeedUpTx", "pending", len(p.pendingTxMap))

	now := time.Now()
	for txHash, pending := range p.pendingTxMap {
		diff := now.Sub(pending.Time).Seconds()

		// 설정에 있는 timeout 시간이 지났을 경우
		if int(diff) > p.cfg.Validator.SpeedUpTimeOut {
			// SpeedUp 하도록 호출함
			err := pending.ValidatorInfo.SpeedUpTx(
				pending.Chain, pending.TxHash,
			)
			if err != nil {
				p.logger.Error("event", "CheckSpeedUpTx", "chain", pending.Chain, "validator", pending.ValidatorInfo.Name, "tx", pending.TxHash, "diff", diff, "err", err)
			} else {
				p.logger.Info("event", "CheckSpeedUpTx", "chain", pending.Chain, "validator", pending.ValidatorInfo.Name, "tx", pending.TxHash, "diff", diff)
			}

			// SpeedUp 하라고 보냈음으로 pendingMap 에서 삭제
			delete(p.pendingTxMap, txHash)
			continue
		}

		// timeout 시간이 지나지 않았으나 validator 정보가 없는 경우
		if pending.ValidatorInfo == nil || pending.ValidatorSummary == nil {
			// 에러 로그 출력 후 Tx 삭제
			p.logger.Error("event", "CheckSpeedUpTx", "chain", pending.Chain, "validator", "unknown", "tx", pending.TxHash, "diff", diff)
			delete(p.pendingTxMap, txHash)
			continue
		}

		// confirm 까지 대기
		p.logger.Trace("event", "CheckSpeedUpTx", "chain", pending.Chain, "validator", pending.ValidatorInfo.Name, "tx", pending.TxHash, "diff", diff)
	}
}
