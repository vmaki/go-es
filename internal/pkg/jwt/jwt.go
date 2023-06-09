package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	jwtPkg "github.com/golang-jwt/jwt/v4"
	"go-es/internal/pkg/logger"
	"strings"
	"time"
)

var (
	ErrTokenCreate                  = errors.New("创建令牌失败")
	ErrTokenExpired                 = errors.New("令牌已过期")
	ErrTokenMalformed               = errors.New("请求令牌格式有误")
	ErrTokenInvalid                 = errors.New("请求令牌无效")
	ErrHeaderEmpty                  = errors.New("需要认证才能访问！")
	ErrHeaderMalformed              = errors.New("请求头中 Authorization 格式有误")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
)

type JWT struct {
	Name       string
	SignKey    []byte
	ExpireTime time.Duration
	MaxRefresh time.Duration
	Timezone   string
}

func NewJWT(name string, secret string, expireTime, maxRefreshTime int64, timezone string) *JWT {
	return &JWT{
		Name:       name,
		SignKey:    []byte(secret),
		ExpireTime: time.Duration(expireTime) * time.Second,
		MaxRefresh: time.Duration(maxRefreshTime) * time.Second,
		Timezone:   timezone,
	}
}

type Claims struct {
	UserID       int64 `json:"user_id"`
	ExpireAtTime int64 `json:"expire_time"`

	jwtPkg.RegisteredClaims
}

// createToken 创建Token，内部使用
func (jwt *JWT) createToken(claims Claims) (string, error) {
	token := jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.SignKey)
}

func (jwt *JWT) timeNowByTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(jwt.Timezone)
	return time.Now().In(chinaTimezone)
}

// expireAtTime 过期时间
func (jwt *JWT) expireAtTime() time.Time {
	timezone := jwt.timeNowByTimezone()

	return timezone.Add(jwt.ExpireTime)
}

// GenerateToken 生成Token，在登录成功时调用
func (jwt *JWT) GenerateToken(userID int64) (string, int64) {
	expireAtTime := jwt.expireAtTime()

	claims := Claims{
		userID,
		expireAtTime.Unix(),

		jwtPkg.RegisteredClaims{
			NotBefore: jwtPkg.NewNumericDate(jwt.timeNowByTimezone()), // 签名生效时间
			IssuedAt:  jwtPkg.NewNumericDate(jwt.timeNowByTimezone()), // 首次签名时间（后续刷新 Token 不会更新）
			ExpiresAt: jwtPkg.NewNumericDate(expireAtTime),            // 签名过期时间
			Issuer:    jwt.Name,                                       // 签名颁发者
		},
	}

	token, err := jwt.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return "", 0
	}

	return token, expireAtTime.Unix()
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

// parseTokenString 使用 jwtPkg.ParseWithClaims 解析 Token
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

// RefreshToken 更新 Token，用以提供 refresh token 接口
func (jwt *JWT) RefreshToken(c *gin.Context) (string, int64, error) {
	// 1. 从 Header 里获取 token
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return "", 0, parseErr
	}

	// 2. 调用 jwt 库解析用户传参的 Token
	token, err := jwt.parseTokenString(tokenString)
	if err != nil {
		validationErr, ok := err.(*jwtPkg.ValidationError)
		if !ok || validationErr.Errors != jwtPkg.ValidationErrorExpired {
			return "", 0, err
		}
	}

	// 3. 解析 JWTCustomClaims 的数据
	claims := token.Claims.(*Claims)

	// 4. 检查是否过了『最大允许刷新的时间』
	expireAt := jwt.expireAtTime()
	x := jwt.timeNowByTimezone().Add(-jwt.MaxRefresh).Unix()
	if claims.IssuedAt.Unix() > x {
		claims.RegisteredClaims.ExpiresAt = jwtPkg.NewNumericDate(expireAt)
		newToken, err := jwt.createToken(*claims)
		if err != nil {
			return "", 0, ErrTokenCreate
		}

		return newToken, expireAt.Unix(), nil
	}

	return "", 0, ErrTokenExpiredMaxRefresh
}
