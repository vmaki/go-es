package make

import (
	"fmt"
	"go-es/internal/pkg/console"
	"strings"

	"github.com/spf13/cobra"
)

var CmdMakeAPIHandle = &cobra.Command{
	Use:   "apihandle",
	Short: "Create api handle，example: make apihandle v1/user",
	Run:   runMakeAPIHandle,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeAPIHandle(cmd *cobra.Command, args []string) {
	// 处理参数，要求附带 API 版本（v1 或者 v2）
	array := strings.Split(args[0], "/")
	if len(array) != 2 {
		console.Exit("api controller name format: v1/user")
	}

	// apiVersion 用来拼接目标路径
	// name 用来生成 cmd.Model 实例
	apiVersion, name := array[0], array[1]
	model := makeModelFromString(name)

	// 组建目标目录
	filePath := fmt.Sprintf("app/api/handle/%s/%s.go", apiVersion, model.TableName)

	// 基于模板创建文件（做好变量替换）
	createFileFromStub(filePath, "api_handle", model)
}
