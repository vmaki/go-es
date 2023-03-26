package user

import (
	"go-es/app/models"
	"go-es/global"
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
	global.GDB.Create(&u)
}
