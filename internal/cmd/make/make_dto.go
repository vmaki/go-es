package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeDto = &cobra.Command{
	Use:   "dto",
	Short: "Create dto file, example: make dto user",
	Run:   runMakeDto,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeDto(cmd *cobra.Command, args []string) {
	// 格式化模型名称，返回一个 Model 对象
	model := makeModelFromString(args[0])

	// 拼接目标文件路径
	filePath := fmt.Sprintf("internal/services/dto/%s.go", model.PackageName)

	// 基于模板创建文件（做好变量替换）
	createFileFromStub(filePath, "dto", model)
}
