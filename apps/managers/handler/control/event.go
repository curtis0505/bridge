package control

import (
	"github.com/curtis0505/bridge/apps/managers/util"
	bridge "github.com/curtis0505/bridge/libs/types"
)

func (p *ControlHandler) UnPaused(log bridge.Log) error {
	pending := p.CheckPendingTx(log.Chain(), log.EventName)
	p.logger.Info("event", log.EventName, "chain", log.Chain(), "txHash", log.TxHash(), "pending", pending)

	util.NewMessage().
		SetZone(p.cfg.Server.ServiceId).
		SetTitle("Bridge UnPaused").
		SetMessageType(util.MessageTypeAlert).
		AddKeyValueWidget("Chain", log.Chain()).
		AddKeyValueWidget("TxHash", log.TxHash()).
		SendMessage()

	if p.CheckPendingTx(log.Chain(), log.EventName) {
		p.RemovePendingTx(log.Chain(), log.EventName)
	}

	return nil
}

func (p *ControlHandler) Paused(log bridge.Log) error {
	pending := p.CheckPendingTx(log.Chain(), log.EventName)
	p.logger.Info("event", log.EventName, "chain", log.Chain(), "txHash", log.TxHash(), "pending", pending)

	util.NewMessage().
		SetZone(p.cfg.Server.ServiceId).
		SetTitle("Bridge Paused").
		SetMessageType(util.MessageTypeAlert).
		AddKeyValueWidget("Chain", log.Chain()).
		AddKeyValueWidget("TxHash", log.TxHash()).
		SendMessage()

	if p.CheckPendingTx(log.Chain(), log.EventName) {
		p.RemovePendingTx(log.Chain(), log.EventName)
	}

	return nil
}
