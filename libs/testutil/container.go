package testutil

import (
	"context"
	"fmt"
	gormdb "github.com/curtis0505/bridge/libs/database/gorm"
	mongodb2 "github.com/curtis0505/bridge/libs/database/mongo"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"testing"
)

func MySQLContainer(t *testing.T) *gorm.DB {
	ctx := context.Background()
	mysqlContainer, err := mysql.Run(ctx,
		"mysql:5.7",
		mysql.WithDatabase("foo"),
		mysql.WithUsername("root"),
		mysql.WithPassword("password"),
	)
	require.NoError(t, err)
	mysqlContainer.Start(ctx)
	endpoint, err := mysqlContainer.Endpoint(ctx, "")
	require.NoError(t, err)
	db, err := gormdb.NewGORMClient(endpoint, "root", "password", "foo")
	require.NoError(t, err)

	sqlDB, err := db.DB()
	require.NoError(t, err)

	err = sqlDB.Ping()
	require.NoError(t, err)

	return db
}

func MongoContainer(t *testing.T) *mongo.Client {
	ctx := context.Background()
	mongoContainer, err := mongodb.Run(ctx, "mongo:4.4.10",
		mongodb.WithUsername("admin"),
		mongodb.WithPassword("password"),
	)

	require.NoError(t, err)
	mongoContainer.Start(ctx)
	endpoint, err := mongoContainer.Endpoint(ctx, "")
	require.NoError(t, err)
	db, err := mongodb2.NewMongoClient(ctx, fmt.Sprintf("mongodb://%s", endpoint), "admin", "password")
	require.NoError(t, err)

	err = db.Ping(ctx, nil)
	require.NoError(t, err)
	return db
}
