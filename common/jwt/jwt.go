package jwt

import (
	"go-es/config"
	"go-es/internal/pkg/jwt"
)

func NewJWT() *jwt.JWT {
	return jwt.NewJWT(
		config.GlobalConfig.Name,
		config.GlobalConfig.JWT.Secret,
		config.GlobalConfig.JWT.ExpireTime,
		config.GlobalConfig.JWT.MaxRefreshTime,
		config.GlobalConfig.Timezone,
	)
}
