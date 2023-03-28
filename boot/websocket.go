package boot

import (
	socketio "github.com/googollee/go-socket.io"
	"go-es/app/websocket"
)

func SetupWebsocket() *socketio.Server {
	ws := socketio.NewServer(nil)
	websocket.RegisterWebsocket(ws)

	return ws
}
