package socket

import (
	"github.com/gin-gonic/gin"
	"go-es/app/middlewares"
	"go-es/global"
)

func RegisterSocketRoutes(r *gin.Engine) {
	r.GET("/socket.io/*any", middlewares.GinMiddleware("*"), gin.WrapH(global.GWebsocket))
	r.POST("/socket.io/*any", middlewares.GinMiddleware("*"), gin.WrapH(global.GWebsocket))
}
