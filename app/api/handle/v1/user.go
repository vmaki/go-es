package v1

import (
	"github.com/gin-gonic/gin"
	"go-es/app/api/handle"
	"go-es/common/ctxdata"
	"go-es/common/errorx"
	"go-es/internal/services"
)

type UserHandle struct {
	handle.BaseAPIController
}

func (h *UserHandle) Info(ctx *gin.Context) {
	uid := ctxdata.CurrentUID(ctx)

	s := services.User{}
	data, err := s.Info(uid)
	if err != nil {
		errorx.Error(ctx, err)
		return
	}

	errorx.Success(ctx, data)
}
