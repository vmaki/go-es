package user

import (
	"go-es/global"
)

func IsPhoneExist(phone string) bool {
	var count int64
	global.GDB.Model(User{}).Where("phone = ?", phone).Count(&count)

	return count > 0
}

func GetByPhone(phone string) (userModel *User) {
	global.GDB.Where("phone = ?", phone).First(&userModel)
	return
}

func Info(id int64) (userModel User) {
	global.GDB.Where("id", id).First(&userModel)
	return
}
