package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-es/app/api/handle"
	"go-es/app/mqueue/tasks"
	"go-es/common/responsex"
	"go-es/global"
	"go-es/internal/pkg/cache"
)

type TestHandle struct {
	handle.BaseAPIController
}

func (h *TestHandle) Index(ctx *gin.Context) {
	global.GRedis.Set("f1", "asasasasas", 64)
	cache.Set("f2", "asasasasas", 64)
	responsex.Success(ctx, nil)
}

func (h *TestHandle) SysErr(ctx *gin.Context) {
	panic("报错了")

	responsex.Success(ctx, nil)
}

func (h *TestHandle) Job(ctx *gin.Context) {
	err := tasks.SendSMSTask(tasks.SendSMSPayload{
		Phone: "15913395633",
		Code:  123456,
	})
	if err != nil {
		fmt.Println("创建任务失败, err: " + err.Error())
		return
	}

	err = tasks.SendSMSTask(tasks.SendSMSPayload{
		Phone:    "15913395644",
		Code:     654321,
		WorkTime: 30,
	})
	if err != nil {
		fmt.Println("创建任务2失败, err: " + err.Error())
		return
	}

	responsex.Success(ctx, nil)
}
