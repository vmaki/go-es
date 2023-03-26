package database

import (
	"database/sql"
	"errors"
	"fmt"
	"go-es/config"
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

func CurrentDatabase() (dbname string) {
	dbname = global.GDB.Migrator().CurrentDatabase()
	return
}

func DeleteAllTables() error {
	var err error

	switch config.GlobalConfig.DataBase.Driver {
	case "mysql":
		err = deleteMySQLTables()
	default:
		panic(errors.New("database connection not supported"))
	}

	return err
}

func deleteMySQLTables() error {
	dbname := CurrentDatabase()
	var tables []string

	// 读取所有数据表
	err := global.GDB.Table("information_schema.tables").
		Where("table_schema = ?", dbname).
		Pluck("table_name", &tables).
		Error
	if err != nil {
		return err
	}

	// 暂时关闭外键检测
	global.GDB.Exec("SET foreign_key_checks = 0;")

	// 删除所有表
	for _, table := range tables {
		err := global.GDB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}

	// 开启 MySQL 外键检测
	global.GDB.Exec("SET foreign_key_checks = 1;")

	return nil
}
