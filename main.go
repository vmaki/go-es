package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go-es/app/cronx"
	"go-es/app/mqueue"
	"go-es/boot"
	"go-es/config"
	"go-es/internal/pkg/asynq"
	"go-es/internal/pkg/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	boot.SetupAsynq()
}

// @title           go-es 的 API 文档
// @version         1.0
// @description     这是 go-es 的 API 文档

// @host      localhost:7001/api/v1
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	boot.SetupRoute(r)

	server := http.Server{
		Addr:    ":" + cast.ToString(config.GlobalConfig.Port),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server listen err:%s", err)
		}
	}()

	// 延迟任务
	go func() {
		tasks := mqueue.NewMQueue(context.Background()).Register()
		if err := asynq.Srv.Run(tasks); err != nil {
			logger.ErrorString("CMD", "serve", err.Error())
		}
	}()

	// 定时任务
	go func() {
		c := cronx.NewCron()
		c.Register()
		c.Start()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 在此阻塞
	<-quit

	ctx, channel := context.WithTimeout(context.Background(), 5*time.Second)

	defer channel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown error")
	}
	log.Println("server exiting...")
}
