package main

import (
	"github.com/gin-gonic/gin"
	"go-es/init"
)

func main() {
	r := gin.New()

	init.SetupRoute(r)

	// 运行服务
	err := r.Run(":7001")
	if err != nil {
		panic("启动服务失败, err:" + err.Error())
	}
}
