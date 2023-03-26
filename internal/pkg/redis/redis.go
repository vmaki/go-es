package redis

import (
	"context"
	"go-es/internal/pkg/logger"
	"sync"
	"time"

	redisLib "github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Context context.Context
	Client  *redisLib.Client
}

var (
	once        sync.Once
	redisClient *RedisClient
)

func NewRedisClient(address string, username string, password string, db int) *RedisClient {
	once.Do(func() {
		redisClient = Connect(address, username, password, db)
	})

	return redisClient
}

func Connect(address string, username string, password string, db int) *RedisClient {
	rds := &RedisClient{}
	rds.Context = context.Background()
	rds.Client = redisLib.NewClient(&redisLib.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})

	err := rds.Ping()
	if err != nil {
		panic("Redis connection failure, err: " + err.Error())
	}

	return rds
}

func (rds RedisClient) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}

func (rds RedisClient) Set(key string, value interface{}, expireTime int64) bool {
	if err := rds.Client.Set(rds.Context, key, value, time.Duration(expireTime)*time.Second).Err(); err != nil {
		logger.ErrorString("Redis", "Set", err.Error())
		return false
	}

	return true
}

func (rds RedisClient) Get(key string) string {
	if result, err := rds.Client.Get(rds.Context, key).Result(); err != nil {
		if err != redisLib.Nil {
			logger.ErrorString("Redis", "Get", err.Error())
		}

		return ""
	} else {
		return result
	}
}

// Has 判断一个 key 是否存在，内部错误和 redis.Nil 都返回 false
func (rds RedisClient) Has(key string) bool {
	if _, err := rds.Client.Get(rds.Context, key).Result(); err != nil {
		if err != redisLib.Nil {
			logger.ErrorString("Redis", "Has", err.Error())
		}

		return false
	}

	return true
}

func (rds RedisClient) Del(keys ...string) bool {
	if err := rds.Client.Del(rds.Context, keys...).Err(); err != nil {
		logger.ErrorString("Redis", "Del", err.Error())
		return false
	}

	return true
}

func (rds RedisClient) Incr(key string) bool {
	if err := rds.Client.Incr(rds.Context, key).Err(); err != nil {
		logger.ErrorString("Redis", "Incr", err.Error())
		return false
	}

	return true
}

func (rds RedisClient) IncrBy(key string, value int64) bool {
	if err := rds.Client.IncrBy(rds.Context, key, value).Err(); err != nil {
		logger.ErrorString("Redis", "IncrBy", err.Error())
		return false
	}

	return true
}

func (rds RedisClient) Decr(key string) bool {
	if err := rds.Client.Decr(rds.Context, key).Err(); err != nil {
		logger.ErrorString("Redis", "Decr", err.Error())
		return false
	}

	return true
}

func (rds RedisClient) DecrBy(key string, value int64) bool {
	if err := rds.Client.DecrBy(rds.Context, key, value).Err(); err != nil {
		logger.ErrorString("Redis", "DecrBy", err.Error())
		return false
	}

	return true
}
