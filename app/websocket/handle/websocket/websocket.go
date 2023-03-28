package websocket

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"log"
)

func Connect(s socketio.Conn) error {
	s.SetContext("")
	log.Println("connected:", s.ID())
	s.Join("UserRoom")

	return nil
}

func Error(s socketio.Conn, e error) {
	log.Println("meet error:", e)
}

func Disconnect(s socketio.Conn, msg string) {
	log.Println("closed", msg)
}

func OnBye(s socketio.Conn) {
	last := s.Context().(string)
	s.Emit("bye", s.ID())
	fmt.Println("最后一条消息：" + last)

	_ = s.Close()
}

func OnNotice(s socketio.Conn, msg string) {
	s.Emit("reply", "用户"+s.ID()+", "+msg)
}
