package main

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"go-es/boot"
	"go-es/internal/cmd"
	"go-es/internal/pkg/console"
	"os"
)

func init() {
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()

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

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)
	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
