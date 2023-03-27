package migrations

import (
	"database/sql"
	"go-es/app/models"
	"go-es/internal/pkg/migrate"

	"gorm.io/gorm"
)

func init() {
	type User struct {
		models.BaseModel

		Nickname string `gorm:"column:nickname;type:varchar(64);comment:姓名"`
		Phone    string `gorm:"column:phone;uniqueIndex;type:varchar(32);comment:手机号码"`
		Password string `gorm:"column:password;type:char(32);comment:密码"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&User{})
	}

	migrate.Add("2023_03_27_020433_add_users_table", up, down)
}
