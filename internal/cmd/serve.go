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
	"go-es/internal/pkg/asynq"
	"go-es/internal/pkg/logger"
	"log"
	"net/http"
	"sync"
)

var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	boot.SetupRoute(r)

	server := http.Server{
		Addr:    ":" + cast.ToString(config.GlobalConfig.Port),
		Handler: r,
	}

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server listen err:%s", err)
		}
	}()

	// 延迟任务
	go func() {
		defer wg.Done()

		tasks := mqueue.NewMQueue(context.Background()).Register()
		if err := asynq.Srv.Run(tasks); err != nil {
			logger.ErrorString("CMD", "serve", err.Error())
		}
	}()

	// 定时任务
	go func() {
		defer wg.Done()

		c := cronx.NewCron()
		c.Register()
		c.Start()
	}()

	wg.Wait()
}
