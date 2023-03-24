package api

import (
	"github.com/gin-gonic/gin"
	v1 "go-es/app/api/handle/v1"
)

func RegisterAPIRoutes(r *gin.Engine) {
	apiV1 := r.Group("/v1")
	{

		// 测试接口
		testGroup := apiV1.Group("/test")
		{
			handle := new(v1.TestHandle)

			testGroup.GET("/", handle.Index)
			testGroup.GET("/500", handle.SysErr)
		}
	}
}
