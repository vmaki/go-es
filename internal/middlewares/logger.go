package middlewares

import (
	"bytes"
	"go-es/internal/pkg/logger"
	"go-es/internal/tools"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: ctx.Writer}
		ctx.Writer = w

		// 获取请求数据
		var requestBody []byte
		if ctx.Request.Body != nil {
			// c.Request.Body 是一个 buffer 对象，只能读取一次
			requestBody, _ = io.ReadAll(ctx.Request.Body)

			// 读取后，重新赋值 c.Request.Body ，以供后续的其他操作
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		start := time.Now()
		ctx.Next()

		// 开始记录日志
		cost := time.Since(start)
		status := ctx.Writer.Status()

		logFields := []zap.Field{
			zap.Int("status", status),
			zap.String("request", ctx.Request.Method+" "+ctx.Request.URL.String()),
			zap.String("query", ctx.Request.URL.RawQuery),
			zap.String("ip", tools.GetClientIP(ctx)),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", tools.MicrosecondsStr(cost)),
		}

		if ctx.Request.Method == "POST" || ctx.Request.Method == "PUT" || ctx.Request.Method == "DELETE" {
			logFields = append(logFields, zap.String("Request Body", string(requestBody)))
			logFields = append(logFields, zap.String("Response Body", w.body.String()))
		}

		if status > 400 && status <= 499 {
			logger.Warn("HTTP Warning "+cast.ToString(status), logFields...)
		} else if status >= 500 && status <= 599 {
			logger.Error("HTTP Error "+cast.ToString(status), logFields...)
		} else {
			logger.Debug("HTTP Access Log", logFields...)
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
