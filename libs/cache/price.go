package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/curtis0505/bridge/libs/util/storage"
	"github.com/redis/go-redis/v9"
	"strings"
)

var _ Price = (*priceCache)(nil)

type Price interface {
	Cache
	GetPriceBySymbol(symbol string) (float64, error)
	GetPriceBySymbolWithNoErr(symbol string) float64
}

type priceCache struct {
	redisClient  *redis.Client
	priceStorage storage.Storage[string, float64]
}

func newPriceCache(client *redis.Client) Price {
	return &priceCache{
		redisClient:  client,
		priceStorage: storage.New[string, float64](),
	}
}

func (p *priceCache) SetRedis(ctx context.Context) {

	return
}

func (p *priceCache) GetRedis(ctx context.Context) interface{} {
	var tokenPriceSymbol map[string]float64
	val, err := p.redisClient.Get(ctx, "token_price_symbol").Result()
	if err != nil {
		return nil
	}

	if err = json.Unmarshal([]byte(val), &tokenPriceSymbol); err != nil {
		return nil
	}

	return tokenPriceSymbol
}

func (p *priceCache) GetPriceBySymbol(symbol string) (float64, error) {
	price, ok := p.priceStorage.Load(strings.ToLower(symbol))
	if !ok {
		return 0.0, fmt.Errorf("not found token price given symbol(%s)", symbol)
	}

	return price, nil
}

func (p *priceCache) GetPriceBySymbolWithNoErr(symbol string) float64 {
	price, _ := p.GetPriceBySymbol(symbol)
	return price
}
