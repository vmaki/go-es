package v1

import (
	"github.com/gin-gonic/gin"
	"go-es/app/api/handle"
	"go-es/common/errorx"
	"go-es/common/requestx"
	"go-es/internal/services"
	"go-es/internal/services/dto"
)

type AuthHandle struct {
	handle.BaseAPIController
}

func (h *AuthHandle) Register(ctx *gin.Context) {
	req := dto.AuthRegisterReq{}
	if err := requestx.Validate(ctx, &req); err != nil {
		errorx.Error(ctx, err)
		return
	}

	s := services.Auth{}
	data, err := s.Register(&req)
	if err != nil {
		errorx.Error(ctx, err)
		return
	}

	errorx.Success(ctx, data)
}

func (h *AuthHandle) Login(ctx *gin.Context) {
	req := dto.AuthLoginReq{}
	if err := requestx.Validate(ctx, &req); err != nil {
		errorx.Error(ctx, err)
		return
	}

	s := services.Auth{}
	data, err := s.Login(&req)
	if err != nil {
		errorx.Error(ctx, err)
		return
	}

	errorx.Success(ctx, data)
}
