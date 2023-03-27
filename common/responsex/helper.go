package responsex

import (
	"github.com/gin-gonic/gin"
	"go-es/internal/pkg/paginator"
	"net/http"
)

func JSON(ctx *gin.Context, data *Response) {
	ctx.JSON(http.StatusOK, data)
}

func Failure(ctx *gin.Context, status int, data *Response) {
	ctx.AbortWithStatusJSON(status, data)
}

// Success 请求成功
func Success(ctx *gin.Context, data interface{}) {
	JSON(ctx, NewResponse(200, "success", data))
}

// SysError 请求失败-未知错误
func SysError(ctx *gin.Context) {
	Failure(ctx, http.StatusInternalServerError, NewResponseErr(ErrSystem))
}

// NotFound 请求失败-404
func NotFound(ctx *gin.Context) {
	Failure(ctx, http.StatusNotFound, NewResponseErr(ErrNotFound))
}

// TooManyRequests 请求失败-请求过于频繁
func TooManyRequests(ctx *gin.Context) {
	Failure(ctx, http.StatusTooManyRequests, NewResponseErr(ErrTooManyRequests))
}

// Unauthorized 请求失败-未授权
func Unauthorized(ctx *gin.Context, msg string) {
	Failure(ctx, http.StatusOK, NewResponseErr(ErrJWT, msg))
}

// Error 自定义错误
func Error(ctx *gin.Context, err error) {
	switch e := err.(type) {
	case *Response:
		JSON(ctx, e)
	default:
		SysError(ctx)
	}
}

type Paginate struct {
	List   interface{}      `json:"list"`
	Paging paginator.Paging `json:"paging"`
}

func List(ctx *gin.Context, list interface{}, paging paginator.Paging) {
	data := Paginate{
		List:   list,
		Paging: paging,
	}

	Success(ctx, data)
}
