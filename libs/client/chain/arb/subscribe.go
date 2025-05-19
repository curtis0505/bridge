package arb

import (
	"context"
	clienttypes "github.com/curtis0505/bridge/libs/client/chain/types"
	"github.com/curtis0505/bridge/libs/types"
)

func (c *client) Subscribe(ctx context.Context, cb func(eventLog types.Log), addresses ...string) error {
	return clienttypes.NotImplemented
}
