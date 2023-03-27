package user

import (
	"github.com/gin-gonic/gin"
	"go-es/common/paginator"
	"go-es/global"
)

func GetByPhone(phone string) (userModel *User) {
	global.GDB.Where("phone = ?", phone).First(&userModel)
	return
}

func Info(id int64) (userModel User) {
	global.GDB.Where("id", id).First(&userModel)
	return
}

// Paginate 分页内容
func Paginate(ctx *gin.Context, pageSize int) (users []User, paging paginator.Paging) {
	paging = paginator.Paginate(
		ctx,
		global.GDB.Model(User{}),
		&users,
		pageSize,
	)

	return
}
