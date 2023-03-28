package cmd

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"go-es/app/cronx"
	"go-es/app/mqueue"
	"go-es/boot"
	"go-es/config"
	"go-es/global"
	"go-es/internal/pkg/asynq"
	"go-es/internal/pkg/logger"
	"go-es/internal/tools"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

var allowOriginFunc = func(r *http.Request) bool {
	return true
}

func runWeb(cmd *cobra.Command, args []string) {
	ctx, channel := context.WithCancel(context.Background())

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	boot.SetupRoute(r)

	server := http.Server{
		Addr:    ":" + cast.ToString(config.GlobalConfig.Port),
		Handler: r,
	}

	// 主服务
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server listen err:%s", err)
		}
	}()

	// websocket
	go func() {
		if err := global.GWebsocket.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	// 延迟任务
	go func() {
		aq := asynq.Server
		tasks := mqueue.NewMQueue(context.Background()).Register()

		go func() {
			if err := aq.Run(tasks); err != nil {
				logger.ErrorString("CMD", "serve", err.Error())
			}
		}()

		for {
			select {
			case <-ctx.Done():
				log.Println("关闭延迟任务")
				aq.Shutdown()
				return
			}
		}
	}()

	// 定时任务，但是本地环境就不启动了
	if !tools.IsLocal() {
		go func() {
			c := cronx.NewCron()
			c.Register()

			go func() {
				c.Start()
			}()

			for {
				select {
				case <-ctx.Done():
					log.Println("关闭定时任务")
					c.C.Stop()
					return
				}
			}
		}()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit // 在此阻塞

	// 开始关闭服务
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("服务关闭失败")
	}

	global.GWebsocket.Close()

	channel()

	if !tools.IsLocal() {
		time.Sleep(time.Second * 1)
	} else {
		time.Sleep(time.Second * 5)
	}

	log.Println("服务关闭成功，正在退出...")
}
