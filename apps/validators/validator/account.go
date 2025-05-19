package validator

import (
	"context"
	"fmt"
	"github.com/curtis0505/bridge/apps/validators/conf"
	"github.com/curtis0505/bridge/libs/client/chain"
	commonconf "github.com/curtis0505/bridge/libs/client/chain/conf"
	commontypes "github.com/curtis0505/bridge/libs/types"
)

func NewAccount(client *chain.Client, config *conf.Config) (map[string]*commontypes.Account, error) {
	account := map[string]*commontypes.Account{}

	for _, chainSymbol := range client.GetChains() {
		chainConfig, ok := getConfig(config.Client, chainSymbol)
		if ok == false {
			continue
		}

		if len(chainConfig.Url) == 0 {
			continue
		}

		ctx := context.Background()
		chainId, err := client.GetChainID(ctx, chainSymbol)
		if err != nil {
			return nil, fmt.Errorf("chain: %s, GetChainID: %w", chainSymbol, err)
		}

		newAccount, err := commontypes.NewAccount(config.Account, chainId)
		if err != nil {
			return nil, fmt.Errorf("NewAccount: %w", err)
		}

		account[chainSymbol] = newAccount
	}

	return account, nil
}

func getConfig(c commonconf.Config, chain string) (commonconf.ClientConfig, bool) {
	for _, config := range c {
		if config.Chain == chain {
			return config, true
		}
	}
	return commonconf.ClientConfig{}, false
}
