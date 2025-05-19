package cosmos

import (
	"fmt"
	"github.com/curtis0505/bridge/libs/client/chain/conf"
	grpcclient "github.com/curtis0505/bridge/libs/client/chain/cosmos/grpc"
	restclient "github.com/curtis0505/bridge/libs/client/chain/cosmos/rest"
	clienttypes "github.com/curtis0505/bridge/libs/client/chain/types"
	"strings"
)

func NewClient(config conf.ClientConfig) (clienttypes.CosmosClient, error) {
	if strings.HasPrefix(config.Url, "http") {
		return restclient.NewClient(config)
	} else {
		return grpcclient.NewClient(config)
	}
}

// common

func ProxyClient(proxy clienttypes.Proxy, chain string) (clienttypes.CosmosClient, error) {
	c := proxy.ProxyClient(chain)
	if c == nil {
		return nil, fmt.Errorf("not found proxy")
	}

	client, ok := c.(clienttypes.CosmosClient)
	if !ok {
		return nil, fmt.Errorf("failed to casting client")
	}

	return client, nil
}
