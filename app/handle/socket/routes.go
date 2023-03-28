package socket

import (
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func RegisterSocketRoutes(r *gin.Engine, ws *socketio.Server) {
	r.GET("/socket.io/*any", gin.WrapH(ws))
	r.POST("/socket.io/*any", gin.WrapH(ws))
}
