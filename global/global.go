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
	GVA_REDIS *redis.Client
	GVA_LOG   *zap.Logger
)
