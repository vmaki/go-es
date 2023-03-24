package user

import "go-es/app/models"

type User struct {
	models.BaseModel

	Nickname string `gorm:"column:nickname;type:varchar(64);comment:姓名" json:"nickname,omitempty"`
	Phone    string `gorm:"column:phone;uniqueIndex;type:varchar(20);comment:手机号码" json:"-"`
	Password string `gorm:"column:password;type:char(20);comment:密码" json:"-"`

	models.CommonTimestampsField
}
