package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go-es/boot"
	"go-es/config"
)

func init() {
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()

	boot.SetupConfig(env)
	boot.SetLogger()
	boot.SetupDB()
	boot.SetupRedis()
	boot.SetupCache()
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	boot.SetupRoute(r)

	// 运行服务
	err := r.Run(":" + cast.ToString(config.GlobalConfig.Port))
	if err != nil {
		panic("启动服务失败, err:" + err.Error())
	}
}
