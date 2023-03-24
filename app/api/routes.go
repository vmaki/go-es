package api

import (
	"github.com/gin-gonic/gin"
	v1 "go-es/app/api/handle/v1"
)

func RegisterAPIRoutes(r *gin.Engine) {
	apiV1 := r.Group("/v1")
	{
		authGroup := apiV1.Group("/auth")
		{
			handle := new(v1.AuthHandle)

			authGroup.POST("/register", handle.Register)
			authGroup.POST("/login", handle.Login)
		}

		// 测试接口
		testGroup := apiV1.Group("/test")
		{
			handle := new(v1.TestHandle)

			testGroup.GET("/", handle.Index)
			testGroup.GET("/500", handle.SysErr)
		}
	}
}
