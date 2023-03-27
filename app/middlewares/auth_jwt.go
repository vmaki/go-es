package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-es/common/ctxdata"
	"go-es/common/jwt"
	"go-es/common/responsex"
)

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, err := jwt.NewJWT().ParserToken(ctx)
		if err != nil {
			responsex.Unauthorized(ctx, err.Error())
			return
		}

		ctx.Set(ctxdata.CtxKeyJwtUserId, claims.UserID)

		ctx.Next()
	}
}
