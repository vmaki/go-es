package dto

import (
	"github.com/thedevsaddam/govalidator"
	"go-es/common/requestx"
	_ "go-es/internal/validators"
)

type AuthRegisterReq struct {
	Phone    string `json:"phone,omitempty" valid:"phone"`       // 手机号码
	Password string `json:"password,omitempty" valid:"password"` //  密码
}

func (s *AuthRegisterReq) Generate(data interface{}) error {
	rules := govalidator.MapData{
		"phone":    []string{"required", "digits:11", "not_exists:users,phone"},
		"password": []string{"required"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项",
			"digits:手机号长度必须为 11 位的数字",
			"not_exists:该手机已注册可直接登录",
		},
		"password": []string{
			"required:密码为必填项",
		},
	}

	return requestx.GoValidate(data, rules, messages)
}

type AuthRegisterResp struct {
	AccessToken  string `json:"access_token"`  // jwt-token
	AccessExpire int64  `json:"access_expire"` // 过期时间
}

type AuthLoginReq struct {
	Phone    string `json:"phone,omitempty" valid:"phone"`       // 手机号码
	Password string `json:"password,omitempty" valid:"password"` //  密码
}

func (s *AuthLoginReq) Generate(data interface{}) error {
	rules := govalidator.MapData{
		"phone":    []string{"required", "digits:11"},
		"password": []string{"required"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项",
			"digits:手机号长度必须为 11 位的数字",
		},
		"password": []string{
			"required:密码为必填项",
		},
	}

	return requestx.GoValidate(data, rules, messages)
}

type AuthLoginResp struct {
	AccessToken  string `json:"access_token"`  // jwt-token
	AccessExpire int64  `json:"access_expire"` // 过期时间
}

type AuthRefreshTokenResp struct {
	AccessToken  string `json:"access_token"`  // jwt-token
	AccessExpire int64  `json:"access_expire"` // 过期时间
}
