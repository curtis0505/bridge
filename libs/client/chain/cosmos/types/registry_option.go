package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govtypesv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	// ibc
	icacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ibcconnectiontypes "github.com/cosmos/ibc-go/v7/modules/core/03-connection/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	ibccommitmenttypes "github.com/cosmos/ibc-go/v7/modules/core/23-commitment/types"
	ibcsolomachinetypes "github.com/cosmos/ibc-go/v7/modules/light-clients/06-solomachine"
	ibclighttenderminttypes "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
	icsprovidertypes "github.com/cosmos/interchain-security/v3/x/ccv/provider/types"

	//icahosttypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/types"
	//ibcporttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"

	// finschia
	"github.com/curtis0505/grpc-idl/finschia/collection"
	"github.com/curtis0505/grpc-idl/finschia/foundation"
	"github.com/curtis0505/grpc-idl/finschia/token"

	// cosmwasm
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

// RegistryOption interface registrar
type RegistryOption interface {
	registerInterfaces(registry codectypes.InterfaceRegistry)
}

type funcRegistryOption struct {
	f func(codectypes.InterfaceRegistry)
}

func (fro *funcRegistryOption) registerInterfaces(registry codectypes.InterfaceRegistry) {
	fro.f(registry)
}

func newFuncRegistryOption(f func(codectypes.InterfaceRegistry)) *funcRegistryOption {
	return &funcRegistryOption{
		f: f,
	}
}

func WithCosmosRegistry() RegistryOption {
	return newFuncRegistryOption(func(registry codectypes.InterfaceRegistry) {
		std.RegisterInterfaces(registry)
		banktypes.RegisterInterfaces(registry)
		authtypes.RegisterInterfaces(registry)
		authz.RegisterInterfaces(registry)
		stakingtypes.RegisterInterfaces(registry)
		minttypes.RegisterInterfaces(registry)
		evidencetypes.RegisterInterfaces(registry)
		distributiontypes.RegisterInterfaces(registry)
		crisistypes.RegisterInterfaces(registry)
		upgradetypes.RegisterInterfaces(registry)
		consensustypes.RegisterInterfaces(registry)
		slashingtypes.RegisterInterfaces(registry)
		vestingtypes.RegisterInterfaces(registry)
		govtypesv1beta1.RegisterInterfaces(registry)
		govtypesv1.RegisterInterfaces(registry)
		nft.RegisterInterfaces(registry)
		feegrant.RegisterInterfaces(registry)
	})
}

func WithFinschiaRegistry() RegistryOption {
	return newFuncRegistryOption(func(registry codectypes.InterfaceRegistry) {
		collection.RegisterInterfaces(registry)
		foundation.RegisterInterfaces(registry)
		token.RegisterInterfaces(registry)
	})
}

func WithIBCRegistry() RegistryOption {
	return newFuncRegistryOption(func(registry codectypes.InterfaceRegistry) {
		icatypes.RegisterInterfaces(registry)
		ibcchanneltypes.RegisterInterfaces(registry)
		icsprovidertypes.RegisterInterfaces(registry)
		icacontrollertypes.RegisterInterfaces(registry)
		ibctransfertypes.RegisterInterfaces(registry)
		ibcclienttypes.RegisterInterfaces(registry)
		ibccommitmenttypes.RegisterInterfaces(registry)
		ibcfeetypes.RegisterInterfaces(registry)
		ibcconnectiontypes.RegisterInterfaces(registry)
		ibcsolomachinetypes.RegisterInterfaces(registry)
		ibclighttenderminttypes.RegisterInterfaces(registry)
	})
}

func WithWasmRegistry() RegistryOption {
	return newFuncRegistryOption(func(registry codectypes.InterfaceRegistry) {
		wasmtypes.RegisterInterfaces(registry)
	})
}
