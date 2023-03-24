package v1

import (
	"github.com/gin-gonic/gin"
	"go-es/app/api/handle"
	"go-es/common/errorx"
)

type TestHandle struct {
	handle.BaseAPIController
}

func (h *TestHandle) Index(ctx *gin.Context) {
	errorx.Success(ctx, nil)
}

func (h *TestHandle) SysErr(ctx *gin.Context) {
	panic("报错了")

	errorx.Success(ctx, nil)
}
