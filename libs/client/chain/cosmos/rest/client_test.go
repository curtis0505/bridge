package rest

import (
	"context"
	"github.com/curtis0505/bridge/libs/client/chain/conf"
	"github.com/curtis0505/bridge/libs/types"
	"testing"
)

func TestClient(t *testing.T) {
	c, _ := NewClient(conf.ClientConfig{
		Chain:     types.ChainFNSA,
		ChainName: types.ChainFNSA,
		Url:       "https://dsvt-finschia.line-apps.com",
	})
	address := "link1nddwnkc47p9v9apruhp9xtq4rquf2y49xcw67r"
	ctx := context.Background()
	t.Log(c.Balance(ctx, address, "cony"))
	t.Log(c.Balances(ctx, address))
	t.Log(c.GetMinimumGasPrice(ctx))
	t.Log(c.GetAccountNumberAndSequence(ctx, address))
	t.Log(c.GetReward(ctx, address))
	t.Log(c.GetStaking(ctx, address))
	t.Log(c.GetStakingParams(ctx))
	t.Log(c.GetStaking(ctx, address))
	t.Log(c.BlockNumber(ctx))
	t.Log(c.GetValidatorApr(ctx, "linkvaloper1nddwnkc47p9v9apruhp9xtq4rquf2y495vv8ss"))
}
