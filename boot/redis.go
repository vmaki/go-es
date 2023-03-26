package boot

import (
	"fmt"
	"go-es/config"
	"go-es/global"
	"go-es/internal/pkg/redis"
)

func SetupRedis() {
	cf := config.GlobalConfig.Redis

	global.GRedis = redis.NewRedisClient(
		fmt.Sprintf("%v:%v", cf.Host, cf.Port),
		cf.Username,
		cf.Password,
		cf.Database,
	)
}
