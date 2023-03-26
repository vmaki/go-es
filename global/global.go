package global

import (
	"go-es/config"
	"gorm.io/gorm"
)

var (
	GConfig = new(config.Config)
	GDB     *gorm.DB
)
