package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

type Cache interface {
	SetRedis(ctx context.Context)
	GetRedis(ctx context.Context) interface{}
}

var c = &cache{Once: new(sync.Once)}

type cache struct {
	*sync.Once
	initialized   bool
	tokenCache    Token
	priceCache    Price
	contractCache Contract
	ginCache      Gin
	caches        []Cache
}

//func (c *cache) subscribe(redisClient *redis.Client) {
//	ctx := context.Background()
//	ch := redisClient.Subscribe(ctx, "cache:map").Channel()
//	ticker := time.Tick(5 * time.Minute)
//	for {
//		select {
//		case <-ch:
//			refresh()
//		case <-ticker:
//			refresh()
//		}
//	}
//}

//func refresh() {
//	lop.ForEach(c.caches, func(c Cache, _ int) {
//		c.Refresh()
//	})
//}

func Init(mongoClient *mongo.Client, redisClient *redis.Client) {
	c.Once.Do(func() {
		c.tokenCache = newTokenCache(mongoClient, redisClient)
		c.priceCache = newPriceCache(redisClient)
		c.contractCache = newContractCache(mongoClient, redisClient)
		c.ginCache = NewGinCache(redisClient)
		c.initialized = true
		c.caches = append(c.caches,
			c.tokenCache,
			c.priceCache,
			c.contractCache,
			// and so on ...
		)

		//refresh()
		//go c.subscribe(redisClient)
	})
}

func GinCache() Gin {
	checkInitialized()
	return c.ginCache
}

func TokenCache() Token {
	checkInitialized()
	return c.tokenCache
}

func PriceCache() Price {
	checkInitialized()
	return c.priceCache
}

func ContractCache() Contract {
	checkInitialized()
	return c.contractCache
}

func checkInitialized() {
	if !c.initialized {
		panic("cache not initialized")
	}
}
