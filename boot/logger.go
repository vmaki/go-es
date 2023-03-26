package boot

import (
	"go-es/global"
	"go-es/internal/pkg/logger"
)

func SetLogger() {
	cf := global.GConfig.Log

	logger.InitLogger(cf.Level, cf.Type, cf.Filename, cf.MaxSize, cf.MaxAge, cf.MaxBackup, cf.Compress)
}
