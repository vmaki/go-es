package boot

import (
	"fmt"
	"go-es/config"
	"go-es/internal/pkg/redis"
)

func SetupRedis() {
	cf := config.GlobalConfig.Redis

	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", cf.Host, cf.Port),
		cf.Username,
		cf.Password,
		cf.Database,
	)
}
