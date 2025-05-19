package database

import (
	"context"
	"fmt"
	"github.com/curtis0505/bridge/libs/cache"
	"github.com/curtis0505/bridge/libs/database/conf"
	gormdb "github.com/curtis0505/bridge/libs/database/gorm"
	mongodb "github.com/curtis0505/bridge/libs/database/mongo"
	mongoAccountDB "github.com/curtis0505/bridge/libs/database/mongo/account_db"
	mongoSeviceDB "github.com/curtis0505/bridge/libs/database/mongo/service_db"
	redisdb "github.com/curtis0505/bridge/libs/database/redis"
	"github.com/curtis0505/bridge/libs/logger/v2"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"strconv"
	"sync"
)

type Registry struct {
	mongoClient *mongo.Client
	redisClient *redis.Client
	accountDB   *mongoAccountDB.AccountDB
	serviceDB   *mongoSeviceDB.ServiceDB

	redis *redisdb.Redis

	transactionManager *TransactionManager
	// ... other services
}

func (r *Registry) MongoClient() *mongo.Client { return r.mongoClient }

func (r *Registry) RedisClient() *redis.Client { return r.redisClient }

//func (r *Registry) RedisService() *RedisService { return r.redisService }

func (r *Registry) AccountDB() *mongoAccountDB.AccountDB { return r.accountDB }

func (r *Registry) ServiceDB() *mongoSeviceDB.ServiceDB { return r.serviceDB }

func (r *Registry) TransactionManager() *TransactionManager { return r.transactionManager }

type initializer struct {
	*Registry
	*sync.Once
	initialized bool
}

const (
	DBConfigKeyMongo = "mongodb"
	DBConfigKeyMySQL = "mysql"
	DBConfigKeyRedis = "redis"
)

var initRegistry = &initializer{
	Once:        new(sync.Once),
	initialized: false,
}

func Init(repositoryConfig conf.Config) {
	initRegistry.Once.Do(func() {
		ctx := context.TODO()
		var (
			mongoClient *mongo.Client
			gormClient  *gorm.DB
			redisClient *redis.Client
			err         error
		)

		for _, cfg := range repositoryConfig {
			switch cfg.Type {
			case DBConfigKeyMongo:
				mongoClient, err = mongodb.NewMongoClient(ctx, cfg.DataSource, cfg.User, cfg.Password)
				if err != nil {
					logger.Error("Models Init", logger.BuildLogInput().WithError(fmt.Errorf("failed to connect to mongodb: %v", err)))
				}
			case DBConfigKeyMySQL:
				gormClient, err = gormdb.NewGORMClient(cfg.DataSource, cfg.User, cfg.Password, "statistics")
				if err != nil {
					panic(err) //todo panic
				}
			case DBConfigKeyRedis:
				redisDB, err := strconv.Atoi(cfg.RedisDB)
				if err != nil {
					redisDB = 0
				}
				redisClient = redisdb.NewRedisV8Client(cfg.DataSource, cfg.Password, redisDB, cfg.TLS)
			default:
				logger.Error("Models Init", logger.BuildLogInput().WithError(fmt.Errorf("unknown db type")).WithData("type", cfg.Type))
			}
		}

		cache.Init(mongoClient, redisClient)

		initRegistry.Registry = &Registry{
			mongoClient:        mongoClient,
			redisClient:        redisClient,
			accountDB:          mongoAccountDB.NewAccountDB(mongoClient),
			serviceDB:          mongoSeviceDB.NewServiceDB(mongoClient),
			redis:              redisdb.NewRedis(redisClient),
			transactionManager: NewTransactionManager(mongoClient, gormClient),
		}

		initRegistry.initialized = true
	})
}

func GetRegistry() *Registry {
	if !initRegistry.initialized {
		panic("service registry is not initialized")
	}

	return initRegistry.Registry
}
