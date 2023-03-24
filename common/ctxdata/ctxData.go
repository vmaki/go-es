package ctxdata

import "github.com/gin-gonic/gin"

var CtxKeyJwtUserId = "current_uid"

func CurrentUID(c *gin.Context) int64 {
	return c.GetInt64(CtxKeyJwtUserId)
}
