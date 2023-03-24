package services

import (
	"go-es/app/models/user"
	"go-es/common"
	"go-es/common/encryption"
	"go-es/common/errorx"
	"go-es/config"
	"go-es/internal/pkg/jwt"
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

	token, expireTime := jwt.NewJWT().GenerateToken(data.ID)

	return &dto.AuthRegisterResp{
		AccessToken:  token,
		AccessExpire: expireTime,
		RefreshAfter: expireTime / 2,
	}, nil
}

// Login  登录
func (s *Auth) Login(req *dto.AuthLoginReq) (*dto.AuthLoginResp, error) {
	data := user.GetByPhone(req.Phone)
	if data == nil {
		return nil, errorx.NewResponse(4003, "用户尚未注册", nil)
	}

	if data.Password != encryption.Md5(req.Password, config.GlobalConfig.Name) {
		return nil, errorx.NewResponse(4004, "账户或密码错误", nil)
	}

	token, expireTime := jwt.NewJWT().GenerateToken(data.ID)

	return &dto.AuthLoginResp{
		AccessToken:  token,
		AccessExpire: expireTime,
		RefreshAfter: expireTime / 2,
	}, nil
}
