package boot

import (
	"fmt"
	"go-es/global"
	"go-es/internal/pkg/redis"
)

func SetupRedis() {
	cf := global.GConfig.Redis

	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", cf.Host, cf.Port),
		cf.Username,
		cf.Password,
		cf.Database,
	)
}
