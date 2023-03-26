package user

import (
	"fmt"
	"go-es/app/models"
	"go-es/internal/pkg/database"
)

type User struct {
	models.BaseModel

	Nickname string `gorm:"column:nickname;type:varchar(64);comment:姓名" json:"nickname,omitempty"`
	Phone    string `gorm:"column:phone;uniqueIndex;type:varchar(32);comment:手机号码" json:"-"`
	Password string `gorm:"column:password;type:char(32);comment:密码" json:"-"`

	models.CommonTimestampsField
}

// Create 创建用户
func (u *User) Create() {
	fmt.Println(u)
	database.GlobalDB.Create(&u)
}
