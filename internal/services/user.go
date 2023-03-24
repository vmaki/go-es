package services

import (
	"go-es/app/models/user"
	"go-es/common/errorx"
	"go-es/internal/services/dto"
)

type User struct {
}

// Info 用户信息
func (s *User) Info(uid int64) (*dto.UserInfoResp, error) {
	data := user.Info(uid)
	if data.ID == 0 {
		return nil, errorx.NewResponse(4003, "用户尚未注册", nil)
	}

	return &dto.UserInfoResp{
		Nickname: data.Nickname,
	}, nil
}
