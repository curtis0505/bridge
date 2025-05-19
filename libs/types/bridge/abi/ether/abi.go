package ether

import (
	common "github.com/curtis0505/bridge/libs/types/bridge/abi"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
)

var (
	ERC20, _              = abi.JSON(strings.NewReader(common.ERC20Abi))
	Minter, _             = abi.JSON(strings.NewReader(common.MinterAbi))
	Vault, _              = abi.JSON(strings.NewReader(common.VaultAbi))
	RestakeVaultAbi, _    = abi.JSON(strings.NewReader(common.RestakeVaultAbi))
	MultiSigWallet, _     = abi.JSON(strings.NewReader(common.MultiSigWalletAbi))
	FxERC20RootTunnel, _  = abi.JSON(strings.NewReader(common.FxERC20RootTunnelAbi))
	FxERC20ChildTunnel, _ = abi.JSON(strings.NewReader(common.FxERC20ChildTunnelAbi))
	RootChainProxy, _     = abi.JSON(strings.NewReader(common.RootChainProxyAbi))

	All = []abi.ABI{
		ERC20,
		Minter,
		Vault,
		RestakeVaultAbi,
		MultiSigWallet,
		FxERC20RootTunnel,
		FxERC20ChildTunnel,
		RootChainProxy,
	}
)

var (
	uint256Ty, _ = abi.NewType("uint256", "", nil)
	UInt256      = abi.Arguments{
		{Type: uint256Ty},
	}
)
