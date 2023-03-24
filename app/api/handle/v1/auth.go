package v1

import (
	"fmt"
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
	user, err := s.Register(&req)
	if err != nil {
		return
	}

	// todo jwt
	fmt.Println("user id", user.ID)

	errorx.Success(ctx, nil)
}
