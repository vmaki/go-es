package v1

import (
	"github.com/gin-gonic/gin"
	"go-es/app/api/handle"
	"go-es/common/responsex"
	"go-es/internal/pkg/cache"
	"go-es/internal/pkg/redis"
)

type TestHandle struct {
	handle.BaseAPIController
}

func (h *TestHandle) Index(ctx *gin.Context) {
	redis.GlobalRedis.Set("f1", "asasasasas", 64)
	cache.Set("f2", "asasasasas", 64)
	responsex.Success(ctx, nil)
}

func (h *TestHandle) SysErr(ctx *gin.Context) {
	panic("报错了")

	responsex.Success(ctx, nil)
}
