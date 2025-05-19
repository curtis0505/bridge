package validator

import (
	"context"
	"fmt"
	"github.com/curtis0505/bridge/apps/managers/util"
	"github.com/curtis0505/bridge/libs/types"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
)

func (p *ValidatorHandler) CheckBalance() {
	for _, v := range p.validatorList {
		for _, chain := range bridgetypes.GetBridgeEVMChains() {
			balance, err := p.client.BalanceAt(context.Background(), chain, v.AddressInfo[chain].Address, nil)
			if err != nil {
				p.logger.Error("event", "CheckBalance", "chain", chain, "address", v.AddressInfo, "err", err)
				continue
			}
			v.AddressInfo[chain].Balance = util.ToEther(balance).String()
		}

		p.logger.Debug(
			"event", "CheckBalance", "name", v.Name,
			types.ChainKLAY, fmt.Sprintf("%s%s", util.ToEther(v.AddressInfo[types.ChainKLAY].Balance).String(), "KLAY"),
			types.ChainMATIC, fmt.Sprintf("%s%s", util.ToEther(v.AddressInfo[types.ChainMATIC]).String(), "MATIC"),
			types.ChainETH, fmt.Sprintf("%s%s", util.ToEther(v.AddressInfo[types.ChainETH]).String(), "ETH"),
		)
	}
}
