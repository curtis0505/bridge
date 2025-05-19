package chain

import (
	"github.com/curtis0505/bridge/libs/client/chain/conf"
	"github.com/curtis0505/bridge/libs/logger/v2"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testAddress = map[string]string{
		types.ChainKLAY:  "0xc7a3C10B87BDEefc09ABF5a957A1f575c39bF6d7",
		types.ChainETH:   "0xc7a3C10B87BDEefc09ABF5a957A1f575c39bF6d7",
		types.ChainMATIC: "0xc7a3C10B87BDEefc09ABF5a957A1f575c39bF6d7",
		types.ChainARB:   "0xc7a3C10B87BDEefc09ABF5a957A1f575c39bF6d7",
		types.ChainTRX:   "TTwMLbrGur4r9PvmAPtNuM2Xn3eSZFjET8",
		types.ChainATOM:  "cosmos1e28nz29md9ukjqnnf00vmvm3rex9s334q7mkfl",
		types.ChainKAVA:  "kava13kk5pjkfpde0dlnc4xtutxz2ewns5y5phd24xu",
		types.ChainFNSA:  "link1k3e2yekshkyuzdcx5sfjena3da7rh87tklff6y",
	}

	testTxHash = map[string]string{
		types.ChainKLAY:  "0x5a61190acf3d0c07b3a8963173c956becf3687b2c6d7f4777b20b3621d962d02",
		types.ChainETH:   "0xd4815dc81e074ce6d6c3ff6e390dfa0abde596a3cebbe4799e470368b519573a",
		types.ChainMATIC: "0x65368d7081f50fe993ad71663338bbad6b6892fca7c594ae6f33da51cd2cef4e",
		types.ChainARB:   "0xf552e6402bdcfc70469f922099a488ebd56743b2f9ca202ddb370c082e889c57",
		types.ChainBASE:  "0x6ff58b30e85a2fdf753efbf4c7181a5e11444bceeae0f39c19e526b95a307edc",
		//types.ChainTRX:   "e5726ede8220fba7472eb59e1ceef16e6471755288105f078471661e2f78dfd0",
		//types.ChainATOM:  "7AC43CF368795A8920E2224E45799A38509BEC675687CF4454FE193C8AB9C43D",
		//types.ChainKAVA:  "A4C45C7B56FB37DDD916CCF224B9C2975D1E4C660C9DCBF053D0BB896E041376",
		//types.ChainFNSA:  "135BEC4C35C5E6D9FCAB3286E30D8BD28FF682EBCD4F3DA64689CBDC249E1538",
	}

	testProxy = map[string]conf.ClientConfig{
		types.ChainETH: {
			Chain:     types.ChainETH,
			ChainName: "ethereum",
			Url:       "https://sepolia.dq.neopin.pmang.cloud",
		},
		types.ChainMATIC: {
			Chain:     types.ChainMATIC,
			ChainName: "a-polygon",
			Url:       "https://amoy.dq.neopin.pmang.cloud",
		},
		types.ChainKLAY: {
			Chain:     types.ChainKLAY,
			ChainName: "a-klay",
			Url:       "https://node1.dq.neopin.pmang.cloud",
		},
		types.ChainARB: {
			Chain:     types.ChainARB,
			ChainName: "arbitrum",
			Url:       "https://arb-sepolia.g.alchemy.com/v2/w_6Kf_v-FaWG4BSE6WFkgttsLmz-xlPV",
		},
		types.ChainBASE: {
			Chain:     types.ChainBASE,
			ChainName: "base",
			Url:       "https://base-sepolia.g.alchemy.com/v2/WPsU39WfFnTnB_eA9ff45GHyfHzv8oR5",
		},
		types.ChainTRX: {
			Chain:     types.ChainTRX,
			ChainName: types.ChainTRX,
			Url:       "grpc.trongrid.io:50051",
		},
		types.ChainATOM: {
			Chain:     types.ChainATOM,
			ChainName: types.ChainATOM,
			Url:       "cosmos-grpc.polkachu.com:14990",
		},
		types.ChainKAVA: {
			Chain:     types.ChainKAVA,
			ChainName: types.ChainKAVA,
			Url:       "kava-grpc.polkachu.com:13990",
		},
		types.ChainFNSA: {
			Chain:     types.ChainFNSA,
			ChainName: types.ChainFNSA,
			Url:       "10.30.172.141:32300",
		},
	}
)

func NewTestProxy(t *testing.T) *Client {
	proxy := NewClient(nil)
	logger.InitLog(logger.Config{VerbosityTerminal: 5, UseTerminal: true})

	for _, config := range testProxy {
		err := proxy.AddClient(config)
		assert.NoError(t, err)
	}

	return proxy
}
