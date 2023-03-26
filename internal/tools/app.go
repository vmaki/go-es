package tools

import (
	"github.com/gin-gonic/gin"
	"go-es/global"
	"net"
)

func IsLocal() bool {
	return global.GConfig.Mode == "local"
}

func GetClientIP(ctx *gin.Context) string {
	clientIP := ctx.Request.RemoteAddr

	if ip := ctx.GetHeader("X-Real-IP"); ip != "" {
		clientIP = ip
	} else if ip = ctx.GetHeader("X-Forward-For"); ip != "" {
		clientIP = ip
	} else {
		clientIP, _, _ = net.SplitHostPort(clientIP)
	}

	if clientIP == "::1" {
		clientIP = "127.0.0.1"
	}

	return clientIP
}
