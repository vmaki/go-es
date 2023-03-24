package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
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

	// 任务
	go func() {
		tasks := mqueue.NewMQueue(context.Background()).Register()
		if err := asynq.Srv.Run(tasks); err != nil {
			logger.ErrorString("CMD", "serve", err.Error())
		}
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
