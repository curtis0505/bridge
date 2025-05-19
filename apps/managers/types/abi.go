package types

import (
	"github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/types/bridge/abi/ether"
	"github.com/curtis0505/bridge/libs/types/bridge/abi/klay"
)

var All = []interface{}{
	klay.Vault,
	ether.Vault,
	klay.Minter,
	ether.Minter,
}

var VaultAbi = map[string]interface{}{
	types.ChainKLAY:  klay.Vault,
	types.ChainMATIC: ether.Vault,
	types.ChainETH:   ether.Vault,
}

var MinterAbi = map[string]interface{}{
	types.ChainKLAY:  klay.Minter,
	types.ChainMATIC: ether.Minter,
	types.ChainETH:   ether.Minter,
}

var MultiSigAbi = map[string]interface{}{
	types.ChainKLAY:  klay.MultiSigWallet,
	types.ChainMATIC: ether.MultiSigWallet,
	types.ChainETH:   ether.MultiSigWallet,
}

var ERC20 = map[string]interface{}{
	types.ChainKLAY:  klay.ERC20,
	types.ChainMATIC: ether.ERC20,
	types.ChainETH:   ether.ERC20,
}
