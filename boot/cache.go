package boot

import (
	"fmt"
	"go-es/config"
	"go-es/internal/pkg/cache"
)

func SetupCache() {
	cf := config.GlobalConfig.Redis

	rds := cache.NewRedisStore(
		fmt.Sprintf("%v:%v", cf.Host, cf.Port),
		cf.Username,
		cf.Password,
		cf.Database,
	)

	cache.NewService(rds)
}
