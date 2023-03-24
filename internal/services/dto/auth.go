package dto

import (
	"github.com/thedevsaddam/govalidator"
	"go-es/common/requestx"
)

type AuthRegisterReq struct {
	Phone    string `json:"phone,omitempty" valid:"phone"`       // 手机号码
	Password string `json:"password,omitempty" valid:"password"` //  密码
}

func (s *AuthRegisterReq) Generate(data interface{}) error {
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

type AuthRegisterResp struct {
	Token string `json:"token"` // jwt-token
}
