package routes

import (
	"github.com/gin-gonic/gin"
	"go-es/config"
	"net/http"
)

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "Hello " + config.GlobalConfig.Name,
			})
		})

		v1.GET("/500", func(c *gin.Context) {
			panic("报错了")

			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "Hello " + config.GlobalConfig.Name,
			})
		})
	}
}
