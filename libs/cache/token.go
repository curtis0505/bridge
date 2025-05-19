package cache

import (
	"context"
	"encoding/json"
	"fmt"
	mongoServiceDB "github.com/curtis0505/bridge/libs/database/mongo/service_db"
	mongorepository "github.com/curtis0505/bridge/libs/database/repository"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

type Token interface {
	Cache
	GetTokensByCurrencyID(ctx context.Context, currencyID string) (*mongoServiceDB.Tokens, error)
	GetTokensByAddress(ctx context.Context, chain, address string) (*mongoServiceDB.Tokens, error)
	GetTokensListByChain(ctx context.Context, chain string, tokenType mongoServiceDB.TokenType) []*mongoServiceDB.Tokens
	GetTokensListOnlyCoin(ctx context.Context) []*mongoServiceDB.Tokens
	GetTokensList(ctx context.Context) []*mongoServiceDB.Tokens
}

type tokenCache struct {
	redisClient         *redis.Client
	tokenInfoRepository mongorepository.MongoRepository[mongoServiceDB.Tokens]
}

func newTokenCache(mongoClient *mongo.Client, redisClient *redis.Client) Token {
	return &tokenCache{
		redisClient:         redisClient,
		tokenInfoRepository: mongorepository.NewMongoRepository[mongoServiceDB.Tokens](mongoClient.Database("ItemDB").Collection("token_info")),
	}
}

func (t *tokenCache) SetRedis(ctx context.Context) {
	tokens, err := t.tokenInfoRepository.Find(ctx, bson.M{})
	if err != nil {
		return
	}

	jsonData, err := json.Marshal(tokens)
	if err != nil {
		return
	}

	t.redisClient.SetNX(ctx, "tokens", jsonData, 0)

	return
}

func (t *tokenCache) GetRedis(ctx context.Context) interface{} {
	var tokenInfo []*mongoServiceDB.Tokens
	val, err := t.redisClient.Get(ctx, "tokens").Result()
	if err != nil {
		return nil
	}

	if err = json.Unmarshal([]byte(val), &tokenInfo); err != nil {
		return nil
	}

	return tokenInfo
}

func (t *tokenCache) GetTokensByCurrencyID(ctx context.Context, currencyID string) (*mongoServiceDB.Tokens, error) {
	var token *mongoServiceDB.Tokens
	tokens := t.GetRedis(ctx).([]*mongoServiceDB.Tokens)

	lo.ForEach(tokens, func(info *mongoServiceDB.Tokens, _ int) {
		if info.CurrencyID == currencyID {
			token = info
			return
		}
	})

	if token == nil {
		return nil, fmt.Errorf("token not found for currencyID: %s", currencyID)
	}

	return token, nil
}

func (t *tokenCache) GetTokensByAddress(ctx context.Context, chain, address string) (*mongoServiceDB.Tokens, error) {
	var token *mongoServiceDB.Tokens
	tokens := t.GetRedis(ctx).([]*mongoServiceDB.Tokens)

	lo.ForEach(tokens, func(info *mongoServiceDB.Tokens, _ int) {
		if info.Chain == chain && strings.EqualFold(info.Address, address) {
			token = info
			return
		}
	})

	if token == nil {
		return nil, fmt.Errorf("not found token given chain(%s) address(%s)", chain, address)
	}

	return token, nil
}

func (t *tokenCache) GetTokensList(ctx context.Context) []*mongoServiceDB.Tokens {
	return t.GetRedis(ctx).([]*mongoServiceDB.Tokens)
}

func (t *tokenCache) GetTokensListByChain(ctx context.Context, chain string, tokenType mongoServiceDB.TokenType) []*mongoServiceDB.Tokens {
	tokenList := make([]*mongoServiceDB.Tokens, 0)
	tokens := t.GetRedis(ctx).([]*mongoServiceDB.Tokens)

	lo.ForEach(tokens, func(info *mongoServiceDB.Tokens, _ int) {
		if info.Chain == chain {
			tokenList = append(tokenList, info)
		}
	})

	return tokenList
}

func (t *tokenCache) GetTokensListOnlyCoin(ctx context.Context) []*mongoServiceDB.Tokens {
	tokenList := make([]*mongoServiceDB.Tokens, 0)
	tokens := t.GetRedis(ctx).([]*mongoServiceDB.Tokens)

	lo.ForEach(tokens, func(info *mongoServiceDB.Tokens, _ int) {
		if info.Type == mongoServiceDB.TokenTypeCoin {
			tokenList = append(tokenList, info)
		}
	})

	return tokenList
}
