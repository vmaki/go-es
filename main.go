package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	ser := gin.Default()

	ser.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})

	ser.Run()
}
