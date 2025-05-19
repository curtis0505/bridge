package mongo

import (
	mongorepository "github.com/curtis0505/bridge/libs/database/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountDB struct {
	mongoClient *mongo.Client
	accounts    mongorepository.MongoRepository[Accounts]
	addresses   mongorepository.MongoRepository[Addresses]
}

func NewAccountDB(
	client *mongo.Client,
) *AccountDB {
	return &AccountDB{
		mongoClient: client,
		accounts:    mongorepository.NewMongoRepository[Accounts](client.Database("accountDB").Collection("account_info")),
		addresses:   mongorepository.NewMongoRepository[Addresses](client.Database("accountDB").Collection("address_info")),
	}
}
