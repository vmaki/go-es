package boot

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-es/app/api"
	"go-es/common/responsex"
	_ "go-es/docs"
	"go-es/internal/middlewares"
	"net/http"
	"strings"
)

func SetupRoute(router *gin.Engine) {
	registerGlobalMiddleWare(router)
	registerSwagger(router)

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

func registerSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// setup404Handler 404路由处理器
func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(ctx *gin.Context) {
		accept := ctx.GetHeader("Accept")

		if strings.Contains(accept, "text/html") {
			ctx.String(http.StatusNotFound, "页面返回 404")
		}

		responsex.NotFound(ctx)
	})
}
