package database

import (
	"context"
	"database/sql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

type TransactionManager struct {
	mongoClient *mongo.Client
	gormClient  *gorm.DB
}

func NewTransactionManager(mongoClient *mongo.Client, gormClient *gorm.DB) *TransactionManager {
	return &TransactionManager{mongoClient: mongoClient, gormClient: gormClient}
}

func (t *TransactionManager) WithGormTransaction(fn func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return t.gormClient.Transaction(fn, opts...)
}

func (t *TransactionManager) StartMongoTransaction(opts ...*options.SessionOptions) (mongo.Session, error) {
	return t.mongoClient.StartSession(opts...)
}

func (t *TransactionManager) CommitMongoTransaction(session mongo.Session, sessCtx context.Context) error {
	return session.CommitTransaction(sessCtx)
}

func (t *TransactionManager) AbortMongoTransaction(session mongo.Session, sessCtx context.Context) error {
	return session.AbortTransaction(sessCtx)
}

func (t *TransactionManager) EndMongoSession(session mongo.Session, sessCtx context.Context) {
	session.EndSession(sessCtx)
}

func (t *TransactionManager) WithMongoTransaction(sessCtx context.Context, fn func(ctx mongo.SessionContext) (interface{}, error), opts ...*options.SessionOptions) (interface{}, error) {
	session, err := t.StartMongoTransaction(opts...)
	if err != nil {
		return nil, err
	}
	defer t.EndMongoSession(session, sessCtx)
	return session.WithTransaction(sessCtx, fn)
}
