package limiter

import (
	"go-es/global"
	"go-es/internal/pkg/logger"
	"go-es/internal/pkg/redis"
	"go-es/internal/tools"
	"strings"

	"github.com/gin-gonic/gin"
	limiterPkg "github.com/ulule/limiter/v3"
	limiterRedis "github.com/ulule/limiter/v3/drivers/store/redis"
)

func GetKeyIP(ctx *gin.Context) string {
	return tools.GetClientIP(ctx)
}

func GetKeyRouteWithIP(ctx *gin.Context) string {
	return routeToKeyString(ctx.FullPath()) + "" + tools.GetClientIP(ctx)
}

// CheckRate 检测请求是否超额
func CheckRate(ctx *gin.Context, key string, formatted string) (limiterPkg.Context, error) {
	var context limiterPkg.Context

	rate, err := limiterPkg.NewRateFromFormatted(formatted)
	if err != nil {
		logger.LogIf(err)
		return context, err
	}

	// 初始化
	store, err := limiterRedis.NewStoreWithOptions(redis.GlobalRedis.Client, limiterPkg.StoreOptions{
		Prefix: global.GConfig.Name + ":limiter",
	})
	if err != nil {
		logger.LogIf(err)
		return context, err
	}

	limiterObj := limiterPkg.New(store, rate)

	// 获取限流的结果
	if ctx.GetBool("limiter-once") {
		// Peek() 取结果，不增加访问次数
		return limiterObj.Peek(ctx, key)
	} else {
		// 确保多个路由组里调用 LimitIP 进行限流时，只增加一次访问次数。
		ctx.Set("limiter-once", true)

		// Get() 取结果且增加访问次数
		return limiterObj.Get(ctx, key)
	}
}

// routeToKeyString 辅助方法，将 URL 中的 / 格式为 -
func routeToKeyString(routeName string) string {
	routeName = strings.ReplaceAll(routeName, "/", "-")
	routeName = strings.ReplaceAll(routeName, ":", "_")

	return routeName
}
