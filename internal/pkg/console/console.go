package console

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

func Success(msg string) {
	colorOut(msg, "green")
}

func Error(msg string) {
	colorOut(msg, "red")
}

func Warning(msg string) {
	colorOut(msg, "yellow")
}

// colorOut 内部使用，设置高亮颜色
func colorOut(message, color string) {
	_, _ = fmt.Fprintln(os.Stdout, ansi.Color(message, color))
}
