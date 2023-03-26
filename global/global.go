package global

import (
	"github.com/redis/go-redis/v9"
	"go-es/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GConfig   = new(config.Config)
	GDB       *gorm.DB
	GLog      *zap.Logger
	GVA_REDIS *redis.Client
)
