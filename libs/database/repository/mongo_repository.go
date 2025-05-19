package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
)

type MongoRepository[T Document] interface {
	FindByID(ctx context.Context, id primitive.ObjectID) (*T, error)
	FindOne(ctx context.Context, filter any, opts ...*options.FindOneOptions) (*T, error)
	FindOneAndUpdate(ctx context.Context, filter any, update any, opts ...*options.FindOneAndUpdateOptions) (*T, error)
	Find(ctx context.Context, filter any, opts ...*options.FindOptions) ([]*T, error)
	Insert(ctx context.Context, document *T, opts ...*options.InsertOneOptions) error
	InsertMany(ctx context.Context, documents []*T, opts ...*options.InsertManyOptions) error
	UpdateOne(ctx context.Context, filter any, update any, opts ...*options.UpdateOptions) error
	UpdateMany(ctx context.Context, filter any, update any, opts ...*options.UpdateOptions) error
	UpdateByID(ctx context.Context, id primitive.ObjectID, update any, opts ...*options.UpdateOptions) error
	UpdateOneWithResult(ctx context.Context, filter any, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter any, opts ...*options.DeleteOptions) error
	DeleteMany(ctx context.Context, filter any, opts ...*options.DeleteOptions) error
	Aggregate(ctx context.Context, pipeline any, opts ...*options.AggregateOptions) ([]*T, error)
	AggregateWithSlice(ctx context.Context, pipeline any, dest any, opts ...*options.AggregateOptions) error
	CountDocuments(ctx context.Context, filter any, opts ...*options.CountOptions) (int64, error)
	UpsertOne(ctx context.Context, filter any, update any, opts ...*options.UpdateOptions) error
	UpsertMany(ctx context.Context, filter bson.M, update bson.M, opts ...*options.UpdateOptions) error
	BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error)
	ReplaceOne(ctx context.Context, filter bson.M, update any, opts ...*options.ReplaceOptions) error
	Drop(ctx context.Context) error
}

func NewMongoRepository[T any](conn *mongo.Collection) MongoRepository[T] {
	return &mongoRepository[T]{conn}
}

type Document any

type mongoRepository[T Document] struct {
	conn *mongo.Collection
}

func (m *mongoRepository[T]) AggregateWithSlice(ctx context.Context, pipeline any, dest any, opts ...*options.AggregateOptions) error {
	var (
		cursor *mongo.Cursor
		err    error
	)

	val := reflect.TypeOf(dest)
	if val.Kind() != reflect.Ptr {
		return errors.New("destination must be a pointer to a slice")
	}

	if val.Elem().Kind() != reflect.Slice {
		return errors.New("destination must be a pointer to a slice")
	}

	cursor, err = m.conn.Aggregate(ctx, pipeline, opts...)
	if err != nil {
		return err
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)

	if err = cursor.All(ctx, dest); err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository[T]) UpsertOne(ctx context.Context, filter, update any, opts ...*options.UpdateOptions) error {
	var (
		err error
	)

	opts = append(opts, options.Update().SetUpsert(true))

	if _, err = m.conn.UpdateOne(ctx, filter, update, opts...); err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository[T]) UpsertMany(ctx context.Context, filter, update bson.M, opts ...*options.UpdateOptions) error {
	var (
		err error
	)

	opts = append(opts, options.Update().SetUpsert(true))

	if _, err = m.conn.UpdateMany(ctx, filter, update, opts...); err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository[T]) BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	return m.conn.BulkWrite(ctx, models, opts...)
}

func (m *mongoRepository[T]) ReplaceOne(ctx context.Context, filter bson.M, update any, opts ...*options.ReplaceOptions) error {
	var (
		err error
	)

	if _, err = m.conn.ReplaceOne(ctx, filter, update, opts...); err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository[T]) FindByID(ctx context.Context, id primitive.ObjectID) (*T, error) {
	var (
		result = new(T)
		err    error
	)

	if err = m.conn.FindOne(ctx, bson.D{{"_id", id}}).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (m *mongoRepository[T]) FindOne(ctx context.Context, filter any, opts ...*options.FindOneOptions) (*T, error) {
	var (
		result = new(T)
		err    error
	)

	if err = m.conn.FindOne(ctx, filter, opts...).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (m *mongoRepository[T]) FindOneAndUpdate(ctx context.Context, filter any, update any, opts ...*options.FindOneAndUpdateOptions) (*T, error) {
	var (
		result = new(T)
		err    error
	)

	if err = m.conn.FindOneAndUpdate(ctx, filter, update, opts...).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (m *mongoRepository[T]) Find(ctx context.Context, filter any, opts ...*options.FindOptions) ([]*T, error) {
	var (
		result = make([]*T, 0)
		cursor *mongo.Cursor
		err    error
	)

	if cursor, err = m.conn.Find(ctx, filter, opts...); err != nil {
		return nil, err
	}
	defer func(ctx context.Context) {
		_ = cursor.Close(ctx)
	}(ctx)

	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (m *mongoRepository[T]) Insert(ctx context.Context, document *T, opts ...*options.InsertOneOptions) error {
	var err error

	if _, err = m.conn.InsertOne(ctx, document, opts...); err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository[T]) InsertMany(ctx context.Context, documents []*T, opts ...*options.InsertManyOptions) error {
	var (
		slice []any
		err   error
	)

	for _, document := range documents {
		slice = append(slice, document)
	}

	if _, err = m.conn.InsertMany(ctx, slice, opts...); err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository[T]) UpdateOne(ctx context.Context, filter any, update any, opts ...*options.UpdateOptions) error {
	var (
		err error
	)

	if _, err = m.conn.UpdateOne(ctx, filter, update, opts...); err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository[T]) UpdateOneWithResult(ctx context.Context, filter any, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return m.conn.UpdateOne(ctx, filter, update, opts...)
}

func (m *mongoRepository[T]) UpdateMany(ctx context.Context, filter any, update any, opts ...*options.UpdateOptions) error {
	var (
		err error
	)

	if _, err = m.conn.UpdateMany(ctx, filter, update, opts...); err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository[T]) UpdateByID(ctx context.Context, id primitive.ObjectID, update any, opts ...*options.UpdateOptions) error {
	var (
		err error
	)

	if _, err = m.conn.UpdateByID(ctx, id, update, opts...); err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository[T]) DeleteOne(ctx context.Context, filter any, opts ...*options.DeleteOptions) error {
	var (
		err error
	)

	if _, err = m.conn.DeleteOne(ctx, filter, opts...); err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository[T]) DeleteMany(ctx context.Context, filter any, opts ...*options.DeleteOptions) error {
	var (
		err error
	)

	if _, err = m.conn.DeleteMany(ctx, filter, opts...); err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository[T]) Aggregate(ctx context.Context, pipeline any, opts ...*options.AggregateOptions) ([]*T, error) {
	var (
		result []*T
		cursor *mongo.Cursor
		err    error
	)

	if cursor, err = m.conn.Aggregate(ctx, pipeline, opts...); err != nil {
		return nil, err
	}
	defer func(ctx context.Context) {
		_ = cursor.Close(ctx)
	}(ctx)

	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (m *mongoRepository[T]) CountDocuments(ctx context.Context, filter any, opts ...*options.CountOptions) (int64, error) {
	var (
		result int64
		err    error
	)

	if result, err = m.conn.CountDocuments(ctx, filter, opts...); err != nil {
		return 0, err
	}

	return result, nil
}

func (m *mongoRepository[T]) Drop(ctx context.Context) error {
	return m.conn.Drop(ctx)
}
