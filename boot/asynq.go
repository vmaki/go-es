package boot

import (
	"fmt"
	"go-es/config"
	"go-es/internal/pkg/asynq"
)

func SetupAsynq() {
	cf := config.GlobalConfig.Redis

	asynq.ConnectAsynq(
		fmt.Sprintf("%s:%d", cf.Host, cf.Port),
		cf.Username,
		cf.Password,
		cf.Database,
	)
}
