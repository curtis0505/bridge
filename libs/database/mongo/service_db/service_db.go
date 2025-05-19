package mongo

import (
	mongorepository "github.com/curtis0505/bridge/libs/database/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

const serviceDB = "serviceDB"

type ServiceDB struct {
	mongoClient *mongo.Client
	nodes       mongorepository.MongoRepository[Nodes]
	tokens      mongorepository.MongoRepository[Tokens]
	contracts   mongorepository.MongoRepository[Contracts]
}

func NewServiceDB(
	client *mongo.Client,
) *ServiceDB {
	return &ServiceDB{
		mongoClient: client,
		nodes:       mongorepository.NewMongoRepository[Nodes](client.Database(serviceDB).Collection("nodes")),
		tokens:      mongorepository.NewMongoRepository[Tokens](client.Database(serviceDB).Collection("tokens")),
		contracts:   mongorepository.NewMongoRepository[Contracts](client.Database(serviceDB).Collection("contracts")),
	}
}
