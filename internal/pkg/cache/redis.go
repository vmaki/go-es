package cache

import (
	"go-es/config"
	"go-es/internal/pkg/redis"
)

type RedisStore struct {
	KeyPrefix   string
	RedisClient *redis.Server
}

func NewRedisStore(address string, username string, password string, db int) *RedisStore {
	rs := &RedisStore{}
	rs.RedisClient = redis.Connect(address, username, password, db)
	rs.KeyPrefix = config.GlobalConfig.Name + ":cache:"

	return rs
}

func (s *RedisStore) Set(key string, value interface{}, expireTime int64) {
	s.RedisClient.Set(s.KeyPrefix+key, value, expireTime)
}

func (s *RedisStore) Get(key string) string {
	return s.RedisClient.Get(key)
}

func (s *RedisStore) Has(key string) bool {
	return s.RedisClient.Has(s.KeyPrefix + key)
}

func (s *RedisStore) Incr(key string) bool {
	return s.RedisClient.Incr(key)
}

func (s *RedisStore) IncrBy(key string, value int64) bool {
	return s.RedisClient.Incr(key)
}

func (s *RedisStore) Decr(key string) bool {
	return s.RedisClient.Decr(key)
}

func (s *RedisStore) DecrBy(key string, value int64) bool {
	return s.RedisClient.DecrBy(key, value)
}

func (s *RedisStore) Del(key string) bool {
	return s.RedisClient.Del(key)
}
