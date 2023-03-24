package database

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	GlobalDB *gorm.DB
	SqlDB    *sql.DB
)

func Connect(dbConfig gorm.Dialector, _logger logger.Interface) {
	var err error
	GlobalDB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})

	// 处理错误
	if err != nil {
		fmt.Println(err.Error())
	}

	// 获取底层的 sqlDB
	SqlDB, err = GlobalDB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}
