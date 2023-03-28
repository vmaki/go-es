package websocket

import (
	socketio "github.com/googollee/go-socket.io"
	"log"
)

func Connect(s socketio.Conn) error {
	s.SetContext("")
	log.Println("connected:", s.ID())

	return nil
}

func Error(s socketio.Conn, e error) {
	log.Println("meet error:", e)
}

func Disconnect(s socketio.Conn, msg string) {
	log.Println("closed", msg)
}

func OnEvent(s socketio.Conn) string {
	last := s.Context().(string)
	s.Emit("bye", last)
	s.Close()

	return last
}

func OnNotice(s socketio.Conn, msg string) {
	log.Println("notice:", msg)
	s.Emit("reply", "用户"+s.ID()+", "+msg)
}
