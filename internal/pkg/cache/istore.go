package cache

type IStore interface {
	Set(key string, value interface{}, expireTime int64)
	Get(key string) string
	Has(key string) bool
	Incr(key string) bool
	IncrBy(key string, value int64) bool
	Decr(key string) bool
	DecrBy(key string, value int64) bool
	Del(key string) bool
}
