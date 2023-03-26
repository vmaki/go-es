package database

import (
	"database/sql"
	"fmt"
	"go-es/global"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	SqlDB *sql.DB
)

func Connect(dbConfig gorm.Dialector, _logger logger.Interface) {
	var err error
	global.GDB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})

	// 处理错误
	if err != nil {
		fmt.Println(err.Error())
	}

	// 获取底层的 sqlDB
	SqlDB, err = global.GDB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}
