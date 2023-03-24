package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	jwtPkg "github.com/golang-jwt/jwt"
	"go-es/config"
	"go-es/internal/pkg/logger"
	"go-es/internal/tools"
	"strings"
	"time"
)

var (
	ErrTokenExpired    = errors.New("4101-令牌已过期")
	ErrTokenMalformed  = errors.New("4102-请求令牌格式有误")
	ErrTokenInvalid    = errors.New("4103-请求令牌无效")
	ErrHeaderEmpty     = errors.New("4104-需要认证才能访问！")
	ErrHeaderMalformed = errors.New("4105-请求头中 Authorization 格式有误")
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

// getTokenFromHeader 使用 jwtPkg.ParseWithClaims 解析 Token
func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}

	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}

	return parts[1], nil
}

// parseTokenString 使用 jwtpkg.ParseWithClaims 解析 Token
func (jwt *JWT) parseTokenString(tokenString string) (*jwtPkg.Token, error) {
	return jwtPkg.ParseWithClaims(tokenString, &Claims{}, func(token *jwtPkg.Token) (interface{}, error) {
		return jwt.SignKey, nil
	})
}

// ParserToken 解析 Token，中间件中调用
func (jwt *JWT) ParserToken(c *gin.Context) (*Claims, error) {
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}

	// 1. 调用 jwt 库解析用户传参的 Token
	token, err := jwt.parseTokenString(tokenString)

	// 2. 解析出错
	if err != nil {
		validationErr, ok := err.(*jwtPkg.ValidationError)
		if ok {
			if validationErr.Errors == jwtPkg.ValidationErrorMalformed {
				return nil, ErrTokenMalformed
			} else if validationErr.Errors == jwtPkg.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}

		return nil, ErrTokenInvalid
	}

	// 3. 将 token 中的 claims 信息解析出来和 JWTCustomClaims 数据结构进行校验
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}
