package boot

import (
	"go-es/config"
	"go-es/internal/pkg/logger"
)

func SetLogger() {
	config := config.GlobalConfig.Log

	logger.InitLogger(config.Level, config.Type, config.Filename, config.MaxSize, config.MaxAge, config.MaxBackup, config.Compress)
}
