package boot

import (
	"fmt"
	"go-es/global"
	"go-es/internal/pkg/asynq"
)

func SetupAsynq() {
	cf := global.GConfig.Redis

	asynq.ConnectAsynq(
		fmt.Sprintf("%s:%d", cf.Host, cf.Port),
		cf.Username,
		cf.Password,
		cf.Database,
	)
}
