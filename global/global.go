package global

import (
	socketio "github.com/googollee/go-socket.io"
	"go-es/internal/pkg/redis"
	"gorm.io/gorm"
)

var (
	GDB        *gorm.DB
	GRedis     *redis.RedisClient
	GWebsocket *socketio.Server
)
