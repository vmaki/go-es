package boot

import (
	"go-es/app/api"
	"go-es/internal/middlewares"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {
	registerGlobalMiddleWare(router)

	//  注册 API 路由
	api.RegisterAPIRoutes(router)

	setup404Handler(router)
}

// registerGlobalMiddleWare 注册全局中间件
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
	)
}

// setup404Handler 404路由处理器
func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		acceptString := c.GetHeader("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"code": 404,
				"msg":  "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
