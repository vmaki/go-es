package services

import (
	"go-es/app/models/user"
	"go-es/common"
	"go-es/common/encryption"
	"go-es/common/errorx"
	"go-es/config"
	"go-es/internal/services/dto"
)

type Auth struct {
}

// Register 注册
func (s *Auth) Register(req *dto.AuthRegisterReq) (*dto.AuthRegisterResp, error) {
	if isExist := user.IsPhoneExist(req.Phone); isExist {
		return nil, errorx.NewResponse(4002, "用户已存在", nil)
	}

	data := &user.User{
		Nickname: common.MaskPhone(req.Phone),
		Phone:    req.Phone,
		Password: encryption.Md5(req.Password, config.GlobalConfig.Name),
	}
	data.Create()

	if data.ID < 1 {
		return nil, errorx.NewResponse(500, "注册失败", nil)
	}

	token := "1234567890" // todo jwt-token
	return &dto.AuthRegisterResp{Token: token}, nil
}
