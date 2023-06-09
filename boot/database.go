package boot

import (
	"errors"
	"fmt"
	"go-es/config"
	"go-es/internal/pkg/database"
	"go-es/internal/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func SetupDB() {
	cf := config.GlobalConfig.DataBase

	var dialector gorm.Dialector

	switch cf.Driver {
	case "mysql":
		// 构建 DSN 信息
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			cf.Username,
			cf.Password,
			cf.Host,
			cf.Port,
			cf.Database,
			cf.Charset,
		)
		dialector = mysql.New(mysql.Config{
			DSN: dsn,
		})
	default:
		panic(errors.New("database connection not supported"))
	}

	// 连接数据库，并设置 GORM 的日志模式
	database.Connect(dialector, logger.NewGormLogger())

	database.SqlDB.SetMaxOpenConns(cf.MaxOpenConnections)
	database.SqlDB.SetMaxIdleConns(cf.MaxIdleConnections)
	database.SqlDB.SetConnMaxLifetime(time.Duration(cf.MaxLifeSeconds) * time.Second)

	// global.GDB.AutoMigrate(&user.User{}) // 自动迁移
}
