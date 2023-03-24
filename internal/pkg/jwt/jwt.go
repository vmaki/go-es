package jwt

import (
	jwtPkg "github.com/golang-jwt/jwt"
	"go-es/config"
	"go-es/internal/pkg/logger"
	"go-es/internal/tools"
	"time"
)

type JWT struct {
	SignKey    []byte
	MaxRefresh time.Duration
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte(config.GlobalConfig.JWT.Secret),
		MaxRefresh: time.Duration(config.GlobalConfig.JWT.MaxRefreshTime) * time.Second,
	}
}

type Claims struct {
	UserID       int64 `json:"user_id"`
	ExpireAtTime int64 `json:"expire_time"`

	jwtPkg.StandardClaims
}

// createToken 创建Token，内部使用
func (jwt *JWT) createToken(claims Claims) (string, error) {
	token := jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.SignKey)
}

// expireAtTime 过期时间
func (jwt *JWT) expireAtTime() int64 {
	expire := time.Duration(config.GlobalConfig.JWT.ExpireTime) * time.Second

	return tools.TimeNowByTimezone().Add(expire).Unix()
}

// GenerateToken 生成Token，在登录成功时调用
func (jwt *JWT) GenerateToken(userID int64) (string, int64) {
	expireAtTime := jwt.expireAtTime()

	claims := Claims{
		userID,
		expireAtTime,

		jwtPkg.StandardClaims{
			NotBefore: tools.TimeNowByTimezone().Unix(), // 签名生效时间
			IssuedAt:  tools.TimeNowByTimezone().Unix(), // 首次签名时间（后续刷新 Token 不会更新）
			ExpiresAt: expireAtTime,                     // 签名过期时间
			Issuer:    config.GlobalConfig.Name,         // 签名颁发者
		},
	}

	token, err := jwt.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return "", 0
	}

	return token, expireAtTime
}
