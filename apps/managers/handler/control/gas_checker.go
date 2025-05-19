package control

import (
	"context"
	"fmt"
	"github.com/curtis0505/bridge/apps/managers/util"
	"github.com/curtis0505/bridge/libs/cache"
	commontypes "github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/types/base"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	"math/big"
)

func (p *ControlHandler) CheckGasPrice(chain string) {
	switch chain {
	case commontypes.ChainETH:
		p.CheckETHGasPrice()
	}
}

func (p *ControlHandler) CheckETHGasPrice() {
	ctx := context.Background()

	blockNumber, err := p.client.BlockNumber(ctx, commontypes.ChainETH)
	if err != nil {
		p.logger.Error("event", "BlockNumber", "chain", commontypes.ChainETH, "error", err)
		return
	}

	header, err := p.client.HeaderByNumber(ctx, commontypes.ChainETH, blockNumber)
	if err != nil {
		p.logger.Error("event", "GetBlockHeader", "chain", commontypes.ChainETH, "error", err)
		return
	}

	gasTipCap, err := p.client.SuggestGasTipCap(ctx, commontypes.ChainETH)
	if err != nil {
		p.logger.Error("event", "SuggestGasTipCap", "chain", commontypes.ChainETH, "error", err)
		return
	}

	if new(big.Int).Add(header.BaseFee(), gasTipCap).Cmp(big.NewInt(int64(p.cfg.Control.BridgePauseGasPrice))) == 1 {
		if p.pauseCount.Load() < int32(p.cfg.Control.BridgePauseCountMax) {
			p.pauseCount.Add(1)
		}
	} else {
		if p.pauseCount.Load() > 0 {
			p.pauseCount.Add(-1)
		}
	}

	p.logger.Info("event", "CheckGasPrice", "chain", commontypes.ChainETH,
		"number", header.BlockNumber(),
		"baseFee", fmt.Sprintf("%v%s", util.ToEtherWithDecimal(header.BaseFee(), 9), "Gwei"),
		"gasTipCap", fmt.Sprintf("%v%s", util.ToEtherWithDecimal(gasTipCap, 9), "Gwei"),
		"pauseCount", p.pauseCount.Load(),
	)

	contract, err := cache.ContractCache().GetContractByContractID(commontypes.ChainETH, bridgetypes.MultisigWalletContractID)
	if err != nil {
		p.logger.Error("event", "GetContractByContractID", "chain", commontypes.ChainETH, "error", err)
		return
	}

	var paused base.OutputBool
	err = p.client.CallMsgUnmarshalContract2(ctx, contract, "paused", &paused)
	if err != nil {
		p.logger.Error("event", "CallMsg", "chain", commontypes.ChainETH, "error", err)
		return
	}

	if p.pauseCount.Load() >= int32(p.cfg.Control.BridgePauseCount) {
		//pause
		if !paused.Value {
			_, err := p.Pause(ctx, commontypes.ChainETH)
			if err != nil {
				p.logger.Error("event", "Pause", "chain", commontypes.ChainETH, "error", err)
			}
		}
	} else if p.pauseCount.Load() <= 0 {
		//check paused
		if paused.Value {
			_, err := p.Unpause(ctx, commontypes.ChainETH)
			if err != nil {
				p.logger.Error("event", "Unpause", "chain", commontypes.ChainETH, "error", err)
			}
		}
	}
}
