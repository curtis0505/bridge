package cache

import (
	gc "github.com/chenyahui/gin-cache"
	redisbackend "github.com/chenyahui/gin-cache-redis"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

type Gin interface {
	Register(cacheType GinCacheType, fn ...GinCacheConfigFunc) gin.HandlerFunc
}

type ginCache struct {
	store *redisbackend.RedisV9Store
}

func NewGinCache(client *redis.Client) Gin {
	return &ginCache{store: redisbackend.NewRedisV9Store(client)}
}

type GinCacheType string

const (
	GinCacheTypeNone          GinCacheType = "None"
	GinCacheTypeAuthorization GinCacheType = "Authorization"
)

const (
	serverApp         = "server_app"
	cacheKeyPrefix    = "cache"
	cacheKeySeparator = ":"
)

type GinCacheConfig struct {
	AppName string
	TimeOut time.Duration
}

func NewCacheConfig() *GinCacheConfig {
	return &GinCacheConfig{
		AppName: "",
		TimeOut: time.Second * 30,
	}
}

type GinCacheConfigFunc func(config *GinCacheConfig)

func GinCacheWithTimeOut(duration time.Duration) func(config *GinCacheConfig) {
	return func(config *GinCacheConfig) {
		config.TimeOut = duration
	}
}
func GinCacheWithAppName(appName string) func(config *GinCacheConfig) {
	return func(config *GinCacheConfig) {
		config.AppName = appName
	}
}

func (g *ginCache) Register(cacheType GinCacheType, fn ...GinCacheConfigFunc) gin.HandlerFunc {
	config := NewCacheConfig()
	for _, f := range fn {
		f(config)
	}

	if cacheType == GinCacheTypeNone {
		return func(c *gin.Context) {
			c.Set(serverApp, config.AppName)
			gc.Cache(g.store, config.TimeOut, gc.WithCacheStrategyByRequest(cacheStrategyNone))(c)
		}
	} else {
		return func(c *gin.Context) {
			c.Set(serverApp, config.AppName)
			gc.Cache(g.store, config.TimeOut, gc.WithCacheStrategyByRequest(cacheStrategyAuthorization))(c)
		}
	}
}

func cacheStrategyNone(c *gin.Context) (bool, gc.Strategy) {
	return true, gc.Strategy{
		CacheKey: strings.Join(
			[]string{
				cacheKeyPrefix,
				c.GetString(serverApp),
				c.Request.RequestURI,
			},
			cacheKeySeparator,
		),
	}
}

func cacheStrategyAuthorization(c *gin.Context) (bool, gc.Strategy) {
	var token string
	bearerToken := c.Request.Header.Get("Authorization")
	slice := strings.Split(bearerToken, " ")
	if len(slice) > 1 {
		token = slice[1]
	}
	return true, gc.Strategy{
		CacheKey: strings.Join(
			[]string{
				cacheKeyPrefix,
				c.GetString(serverApp),
				c.Request.RequestURI,
				token,
			},
			cacheKeySeparator,
		),
	}
}
