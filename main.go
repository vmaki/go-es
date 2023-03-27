package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-es/boot"
	"go-es/internal/cmd"
	"go-es/internal/pkg/console"
	"os"
)

// @title           go-es
// @version         1.0
// @description     这是 go-es 的 API 文档
// @host      localhost:7001
// @BasePath  /api/v1
// @securityDefinitions.basic  BasicAuth
func main() {
	var rootCmd = &cobra.Command{
		Use:   "GoEs",
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		PersistentPreRun: func(command *cobra.Command, args []string) {
			boot.SetupConfig(cmd.Env)
			boot.SetLogger()
			boot.SetupDB()
			boot.SetupRedis()
			boot.SetupCache()
			boot.SetupAsynq()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdMake,
		cmd.CmdMigrate,
	)

	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe) // 配置默认运行 Web 服务
	cmd.RegisterGlobalFlags(rootCmd)              // 注册全局参数，--env

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
