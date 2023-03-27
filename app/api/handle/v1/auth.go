package v1

import (
	"github.com/gin-gonic/gin"
	"go-es/app/api/handle"
	"go-es/common/jwt"
	"go-es/common/requestx"
	"go-es/common/responsex"
	"go-es/internal/services"
	"go-es/internal/services/dto"
)

type AuthHandle struct {
	handle.BaseAPIHandle
}

// Register
// @Tags      Auth
// @Summary   注册
// @accept    application/json
// @Produce   application/json
// @Param     data  body      dto.AuthRegisterReq true "手机, 密码"
// @Success   200   {object}  responsex.Response{data=dto.AuthRegisterResp}
// @Router    /auth/register [post]
func (h *AuthHandle) Register(ctx *gin.Context) {
	req := dto.AuthRegisterReq{}
	if err := requestx.Validate(ctx, &req); err != nil {
		responsex.Error(ctx, err)
		return
	}

	s := services.NewAuthService()
	data, err := s.Register(&req)
	if err != nil {
		responsex.Error(ctx, err)
		return
	}

	responsex.Success(ctx, data)
}

// Login
// @Tags      Auth
// @Summary   登录
// @accept    application/json
// @Produce   application/json
// @Param     data  body      dto.AuthLoginReq true  "手机, 密码"
// @Success   200   {object}  responsex.Response{data=dto.AuthLoginResp}
// @Router    /auth/login [post]
func (h *AuthHandle) Login(ctx *gin.Context) {
	req := dto.AuthLoginReq{}
	if err := requestx.Validate(ctx, &req); err != nil {
		responsex.Error(ctx, err)
		return
	}

	s := services.NewAuthService()
	data, err := s.Login(&req)
	if err != nil {
		responsex.Error(ctx, err)
		return
	}

	responsex.Success(ctx, data)
}

// RefreshToken
// @Tags      Auth
// @Summary   刷新 token
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  responsex.Response{data=dto.AuthRefreshTokenResp}
// @Router    /auth/refresh-token [post]
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
