package boot

import (
	"go-es/config"
	"go-es/internal/pkg/logger"
	"go-es/internal/tools"
)

func SetLogger() {
	cf := config.GlobalConfig.Log

	logger.InitLogger(cf.Level, cf.Type, cf.Filename, cf.MaxSize, cf.MaxAge, cf.MaxBackup, cf.Compress, tools.IsLocal())
}
