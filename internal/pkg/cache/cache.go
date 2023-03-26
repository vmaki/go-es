package cache

import (
	"sync"
)

type CacheClient struct {
	Store IStore
}

var (
	once        sync.Once
	cacheClient *CacheClient
)

func NewCacheClient(store IStore) {
	once.Do(func() {
		cacheClient = &CacheClient{
			Store: store,
		}
	})
}

func Set(key string, value interface{}, expireTime int64) {
	cacheClient.Store.Set(key, value, expireTime)
}

func Get(key string) string {
	return cacheClient.Store.Get(key)
}

func Has(key string) bool {
	return cacheClient.Store.Has(key)
}

func Incr(key string) bool {
	return cacheClient.Store.Incr(key)
}

func IncrBy(key string, value int64) bool {
	return cacheClient.Store.IncrBy(key, value)
}

func Decr(key string) bool {
	return cacheClient.Store.Decr(key)
}

func DecrBy(key string, value int64) bool {
	return cacheClient.Store.DecrBy(key, value)
}

func Del(key string) bool {
	return cacheClient.Store.Del(key)
}
