package util

import (
	"errors"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
)

// MongoError skip for mongo.ErrNoDocuments
func MongoError(err error) error {
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}
	return nil
}

// MongoNotFound check for mongo.ErrNoDocuments
func MongoNotFound(err error) bool {
	return err != nil && errors.Is(err, mongo.ErrNoDocuments)
}

// RedisNil check for redis.Nil
func RedisNil(err error) bool {
	return err != nil && errors.Is(err, redis.Nil)
}

// IsNil
// check nil
func IsNil(v any) bool {
	return v == nil || reflect.ValueOf(v).IsNil()
}
