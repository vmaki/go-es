package global

import (
	"go-es/internal/pkg/cache"
	"go-es/internal/pkg/redis"
	"gorm.io/gorm"
)

var (
	GDB    *gorm.DB
	GRedis *redis.RedisClient
	GCache *cache.CacheClient
)
