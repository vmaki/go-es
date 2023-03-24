package user

import "go-es/internal/pkg/database"

func IsPhoneExist(phone string) bool {
	var count int64
	database.GlobalDB.Model(User{}).Where("phone = ?", phone).Count(&count)

	return count > 0
}

func GetByPhone(phone string) (userModel *User) {
	database.GlobalDB.Where("phone = ?", phone).First(&userModel)
	return
}
