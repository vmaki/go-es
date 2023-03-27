package services

import (
	"github.com/gin-gonic/gin"
	"go-es/app/models/user"
	"go-es/common/responsex"
	"go-es/internal/pkg/paginator"
	"go-es/internal/services/dto"
)

type User struct {
}

func NewUserService() *User {
	return &User{}
}

// List 用户列表
func (s *User) List(ctx *gin.Context, pageSize int) (users []user.User, paging paginator.Paging) {
	return user.Paginate(ctx, pageSize)
}

// Info 用户信息
func (s *User) Info(uid int64) (*dto.UserInfoResp, error) {
	data := user.Info(uid)
	if data.ID == 0 {
		return nil, responsex.NewResponse(4003, "用户尚未注册", nil)
	}

	return &dto.UserInfoResp{
		Nickname: data.Nickname,
	}, nil
}
