package v1

import (
	"github.com/gin-gonic/gin"
	"go-es/app/api/handle"
	"go-es/common/requestx"
	"go-es/common/responsex"
	"go-es/internal/pkg/jwt"
	"go-es/internal/services"
	"go-es/internal/services/dto"
)

type AuthHandle struct {
	handle.BaseAPIController
}

func (h *AuthHandle) Register(ctx *gin.Context) {
	req := dto.AuthRegisterReq{}
	if err := requestx.Validate(ctx, &req); err != nil {
		responsex.Error(ctx, err)
		return
	}

	s := services.Auth{}
	data, err := s.Register(&req)
	if err != nil {
		responsex.Error(ctx, err)
		return
	}

	responsex.Success(ctx, data)
}

func (h *AuthHandle) Login(ctx *gin.Context) {
	req := dto.AuthLoginReq{}
	if err := requestx.Validate(ctx, &req); err != nil {
		responsex.Error(ctx, err)
		return
	}

	s := services.Auth{}
	data, err := s.Login(&req)
	if err != nil {
		responsex.Error(ctx, err)
		return
	}

	responsex.Success(ctx, data)
}

func (h *AuthHandle) RefreshToken(ctx *gin.Context) {
	token, expire, err := jwt.NewJWT().RefreshToken(ctx)
	if err != nil {
		responsex.Unauthorized(ctx, err.Error())
		return
	}

	data := dto.AuthRefreshTokenResp{
		AccessToken:  token,
		AccessExpire: expire,
	}

	responsex.Success(ctx, data)
}
