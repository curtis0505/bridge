package redis

import (
	"context"
	"crypto/tls"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

func NewRedisV8Client(host, password string, db int, tlsOpt bool) *redis.Client {
	var tlsConfig *tls.Config
	if tlsOpt {
		tlsConfig = &tls.Config{InsecureSkipVerify: true}
	} else {
		tlsConfig = nil
	}
	return redis.NewClient(&redis.Options{
		Addr:      host,
		Password:  password,
		DB:        db,
		TLSConfig: tlsConfig,
	})
}

type Redis struct {
	client *redis.Client
}

func NewRedis(client *redis.Client) *Redis {
	return &Redis{client: client}
}

func (r *Redis) Conn() *redis.Client {
	return r.client
}

func (r *Redis) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return r.client.Expire(ctx, key, expiration).Err()
}

func (r *Redis) Exists(ctx context.Context, key string) (bool, error) {
	val, err := r.client.Exists(ctx, key).Result()
	return val > 0, err
}

func (r *Redis) RPop(ctx context.Context, key string) (string, error) {
	return r.client.RPop(ctx, key).Result()
}

func (r *Redis) LPush(ctx context.Context, key string, values ...any) error {
	return r.client.LPush(ctx, key, values...).Err()
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *Redis) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *Redis) Del(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}

func (r *Redis) HGet(ctx context.Context, key, field string) (string, error) {
	return r.client.HGet(ctx, key, field).Result()
}

func (r *Redis) HSet(ctx context.Context, key string, values ...any) (int64, error) {
	return r.client.HSet(ctx, key, values...).Result()
}

func (r *Redis) HDel(ctx context.Context, key string, fields ...string) (int64, error) {
	return r.client.HDel(ctx, key, fields...).Result()
}

func (r *Redis) Publish(ctx context.Context, channel string, message any) (int64, error) {
	return r.client.Publish(ctx, channel, message).Result()
}

func (r *Redis) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
	return r.client.Subscribe(ctx, channels...)
}

func (r *Redis) SetNX(ctx context.Context, key string, value any, expiration time.Duration) (bool, error) {
	return r.client.SetNX(ctx, key, value, expiration).Result()
}

func (r *Redis) XGroupCreate(ctx context.Context, stream, group string) error {
	// ignore if group already exists
	if err := r.client.XGroupCreate(ctx, stream, group, "0").Err(); err != nil && !strings.Contains(err.Error(), "BUSYGROUP") {
		return err
	}

	return nil
}

func (r *Redis) XGroupCreateMkStream(ctx context.Context, stream, group string) error {
	// 스트림이 존재하지 않는 경우, MKSTREAM 옵션이 있으면 빈 스트림을 생성하고 그룹을 생성
	if err := r.client.XGroupCreateMkStream(ctx, stream, group, "0").Err(); err != nil && !strings.Contains(err.Error(), "BUSYGROUP") {
		return err
	}

	return nil
}

func (r *Redis) XReadGroup(ctx context.Context, stream, group, consumer string, block time.Duration, count int) ([]redis.XStream, error) {
	return r.client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    group,
		Consumer: consumer,
		Streams:  []string{stream, ">"},
		Count:    int64(count),
		Block:    block, // milliseconds
	}).Result()
}

func (r *Redis) XAdd(ctx context.Context, stream string, values ...any) error {
	return r.client.XAdd(ctx, &redis.XAddArgs{
		Stream: stream,
		Values: values,
	}).Err()
}

func (r *Redis) XDel(ctx context.Context, stream string, ids ...string) (int64, error) {
	return r.client.XDel(ctx, stream, ids...).Result()
}

func (r *Redis) XAck(ctx context.Context, stream, group string, ids ...string) error {
	return r.client.XAck(ctx, stream, group, ids...).Err()
}

func (r *Redis) XAutoClaim(ctx context.Context, stream, group, consumer string, minIdleTime time.Duration, count int, start string) ([]redis.XMessage, string, error) {
	return r.client.XAutoClaim(ctx, &redis.XAutoClaimArgs{
		Stream:   stream,
		Group:    group,
		Consumer: consumer,
		MinIdle:  minIdleTime,
		Start:    start,
		Count:    int64(count),
	}).Result()
}

func (r *Redis) XClaim(ctx context.Context, stream, group, consumer string, minIdleTime time.Duration, ids []string) ([]redis.XMessage, error) {
	return r.client.XClaim(ctx, &redis.XClaimArgs{
		Stream:   stream,
		Group:    group,
		Consumer: consumer,
		MinIdle:  minIdleTime,
		Messages: ids,
	}).Result()
}
