package chat

import (
	socketio "github.com/googollee/go-socket.io"
	"time"
)

func OnNotice(s socketio.Conn, msg string) string {
	s.SetContext(msg)

	return "当前服务器时间：" + time.Now().Format("2006-01-02 15:04:05")
}
