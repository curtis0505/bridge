package cache

import (
	"context"
	"encoding/json"
	mongoServiceDB "github.com/curtis0505/bridge/libs/database/mongo/service_db"
	mongorepository "github.com/curtis0505/bridge/libs/database/repository"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

type Contract interface {
	Cache
	GetContractByAddress(ctx context.Context, chain, address string) (*mongoServiceDB.Contracts, error)
	GetContractByContractID(ctx context.Context, chain, contractID string) (*mongoServiceDB.Contracts, error)
}

type contractCache struct {
	redisClient        *redis.Client
	contractRepository mongorepository.MongoRepository[mongoServiceDB.Contracts]
}

func newContractCache(mongoClient *mongo.Client, redisClient *redis.Client) Contract {
	return &contractCache{
		redisClient:        redisClient,
		contractRepository: mongorepository.NewMongoRepository[mongoServiceDB.Contracts](mongoClient.Database("contractDB").Collection("contracts")),
	}
}

func (p *contractCache) SetRedis(ctx context.Context) {
	contracts, err := p.contractRepository.Find(ctx, bson.M{})
	if err != nil {
		return
	}

	jsonData, err := json.Marshal(contracts)
	if err != nil {
		return
	}

	p.redisClient.SetNX(ctx, "contracts", jsonData, 0)

	return
}

func (p *contractCache) GetRedis(ctx context.Context) interface{} {
	var contracts []*mongoServiceDB.Contracts
	val, err := p.redisClient.Get(ctx, "contracts").Result()
	if err != nil {
		return nil
	}

	if err = json.Unmarshal([]byte(val), &contracts); err != nil {
		return nil
	}

	return contracts
}

func (p *contractCache) GetContractByAddress(ctx context.Context, chain, address string) (*mongoServiceDB.Contracts, error) {
	var contract *mongoServiceDB.Contracts
	contracts := p.GetRedis(ctx).([]*mongoServiceDB.Contracts)

	lo.ForEach(contracts, func(info *mongoServiceDB.Contracts, _ int) {
		if info.Chain == chain && strings.EqualFold(info.Address, address) {
			contract = info
			return
		}
	})

	return contract, nil
}

func (p *contractCache) GetContractByContractID(ctx context.Context, chain, contractID string) (*mongoServiceDB.Contracts, error) {
	var contract *mongoServiceDB.Contracts
	contracts := p.GetRedis(ctx).([]*mongoServiceDB.Contracts)

	lo.ForEach(contracts, func(info *mongoServiceDB.Contracts, _ int) {
		if info.Chain == chain && info.ContractID == contractID {
			contract = info
			return
		}
	})

	return contract, nil
}
