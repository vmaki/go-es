package boot

import (
	"fmt"
	"go-es/global"
	"go-es/internal/pkg/cache"
)

func SetupCache() {
	cf := global.GConfig.Redis

	rds := cache.NewRedisStore(
		fmt.Sprintf("%v:%v", cf.Host, cf.Port),
		cf.Username,
		cf.Password,
		cf.Database,
	)

	cache.NewService(rds)
}
