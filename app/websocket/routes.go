package websocket

import (
	socketio "github.com/googollee/go-socket.io"
	"go-es/app/websocket/handle/chat"
	"go-es/app/websocket/handle/websocket"
)

func RegisterWebsocket(ws *socketio.Server) {
	ws.OnConnect("/", websocket.Connect)       // 有新用户
	ws.OnError("/", websocket.Error)           // 连接报错
	ws.OnDisconnect("/", websocket.Disconnect) // 断开链接

	ws.OnEvent("/", "bye", websocket.OnEvent)     // 用户主动离开
	ws.OnEvent("/", "notice", websocket.OnNotice) // 收到用户发来的消息

	// chat 频道
	ws.OnEvent("/chat", "msg", chat.OnNotice) // 收到用户发来的消息
}
