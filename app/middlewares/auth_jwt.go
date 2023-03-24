package middlewares

import (
	"go-es/common/ctxdata"
	"go-es/common/errorx"
	"go-es/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, err := jwt.NewJWT().ParserToken(ctx)
		if err != nil {
			errorx.Unauthorized(ctx, err.Error())
			return
		}

		ctx.Set(ctxdata.CtxKeyJwtUserId, claims.UserID)

		ctx.Next()
	}
}
