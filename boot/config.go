package boot

import "go-es/internal/pkg/config"

func SetupConfig(env string) {
	config.LoadConfig(env)
}
