package user

import "go-es/internal/pkg/database"

func IsPhoneExist(phone string) bool {
	var count int64
	database.GlobalDB.Model(User{}).Where("phone = ?", phone).Count(&count)

	return count > 0
}
