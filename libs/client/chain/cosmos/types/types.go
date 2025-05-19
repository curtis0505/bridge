package types

import (
	"fmt"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	querytypes "github.com/cosmos/cosmos-sdk/types/query"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
)

type ProtoTxProvider interface {
	GetProtoTx() *txtypes.Tx
}

func NewInterfaceRegistry(options ...RegistryOption) codectypes.InterfaceRegistry {
	interfaceRegistry := codectypes.NewInterfaceRegistry()

	for _, option := range options {
		option.registerInterfaces(interfaceRegistry)
	}

	return interfaceRegistry
}

func DefaultPageRequest() *querytypes.PageRequest {
	return &querytypes.PageRequest{
		Key:        []byte(""),
		Offset:     0,
		Limit:      500,
		CountTotal: true,
	}
}

func NewQueryTxEvent(event, key, value string) string {
	return fmt.Sprintf("%s.%s='%s'", event, key, value)
}
