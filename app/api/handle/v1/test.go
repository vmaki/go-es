package v1

import (
	"github.com/gin-gonic/gin"
	"go-es/app/api/handle"
	"go-es/config"
	"net/http"
)

type TestHandle struct {
	handle.BaseAPIController
}

func (h *TestHandle) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Hello " + config.GlobalConfig.Name,
	})
}

func (h *TestHandle) SysErr(ctx *gin.Context) {
	panic("报错了")

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Hello " + config.GlobalConfig.Name,
	})
}
