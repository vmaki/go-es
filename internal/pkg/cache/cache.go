package cache

import (
	"sync"
)

type Service struct {
	Store IStore
}

var (
	once  sync.Once
	Cache *Service
)

func NewService(store IStore) {
	once.Do(func() {
		Cache = &Service{
			Store: store,
		}
	})
}

func Set(key string, value interface{}, expireTime int64) {
	Cache.Store.Set(key, value, expireTime)
}

func Get(key string) string {
	return Cache.Store.Get(key)
}

func Has(key string) bool {
	return Cache.Store.Has(key)
}

func Incr(key string) bool {
	return Cache.Store.Incr(key)
}

func IncrBy(key string, value int64) bool {
	return Cache.Store.IncrBy(key, value)
}

func Decr(key string) bool {
	return Cache.Store.Decr(key)
}

func DecrBy(key string, value int64) bool {
	return Cache.Store.DecrBy(key, value)
}

func Del(key string) bool {
	return Cache.Store.Del(key)
}
