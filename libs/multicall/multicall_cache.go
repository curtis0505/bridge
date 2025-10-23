package multicall

import (
	"encoding/json"
	"fmt"
	"github.com/curtis0505/bridge/libs/logger/v2"
	"github.com/curtis0505/bridge/libs/types"
	etherabi "github.com/ethereum/go-ethereum/accounts/abi"
	klayabi "github.com/kaiachain/kaia/accounts/abi"
	"strings"
	"sync"
	"time"
)

type abiCache struct {
	abi     map[string]*abiExpire
	abiLock *sync.Mutex
}

type abiExpire struct {
	chain    string
	klayAbi  *klayabi.ABI
	etherAbi *etherabi.ABI
	expire   time.Time
}

var c *abiCache

func getAbiCache() *abiCache {
	if c == nil {
		c = &abiCache{
			abi:     make(map[string]*abiExpire),
			abiLock: &sync.Mutex{},
		}
	}
	return c
}

func (cache *abiCache) getCacheKey(chain, to string) string {
	return fmt.Sprintf("%s/%s", chain, to)
}

func (cache *abiCache) getAbi(chain, to string, abi []map[string]interface{}) (*abiExpire, error) {
	cache.abiLock.Lock()
	defer cache.abiLock.Unlock()

	now := time.Now()
	if cached, ok := cache.abi[cache.getCacheKey(chain, to)]; ok {
		if cached.expire.After(now) {
			return cached, nil
		}
	}

	bz, err := json.Marshal(abi)
	if err != nil {
		return nil, err
	}

	var cached *abiExpire
	expire := now.Add(ExpireDuration)

	switch chain {
	// TODO: 체인 추가시 체크 필요
	case types.ChainKLAY:
		klayAbi, err := klayabi.JSON(strings.NewReader(string(bz)))
		if err != nil {
			return nil, err
		}

		cached = &abiExpire{
			chain:   chain,
			klayAbi: &klayAbi,
			expire:  expire,
		}

	default:
		etherAbi, err := etherabi.JSON(strings.NewReader(string(bz)))
		if err != nil {
			return nil, err
		}

		cached = &abiExpire{
			chain:    chain,
			etherAbi: &etherAbi,
			expire:   expire,
		}
	}

	cache.abi[cache.getCacheKey(chain, to)] = cached

	logger.Debug("MultiCall CacheAbi",
		logger.BuildLogInput().
			WithChain(chain).
			WithAddress(to).
			WithData("expire", expire),
	)
	return cached, nil
}
