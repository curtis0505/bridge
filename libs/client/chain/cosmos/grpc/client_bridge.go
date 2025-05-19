package cosmos

import "github.com/curtis0505/bridge/libs/types/bridge"

func (c *client) NewMinter(address string, abi []map[string]interface{}) (bridge.Minter, error) {
	panic("implement me")
}

func (c *client) NewVault(address string, abi []map[string]interface{}) (bridge.Vault, error) {
	panic("implement me")
}

func (c *client) NewMultiSigWallet(address string) (bridge.MultiSigWallet, error) {
	panic("implement me")
}
