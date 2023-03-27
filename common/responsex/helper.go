package responsex

import (
	"github.com/gin-gonic/gin"
	"go-es/internal/pkg/paginator"
	"net/http"
)

func JSON(ctx *gin.Context, data *Response) {
	ctx.JSON(http.StatusOK, data)
}

func Success(ctx *gin.Context, data interface{}) {
	JSON(ctx, NewResponse(200, "success", data))
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

func Failure(ctx *gin.Context, status int, data *Response) {
	ctx.AbortWithStatusJSON(status, data)
}

func SysError(ctx *gin.Context) {
	Failure(ctx, http.StatusInternalServerError, NewResponse(500, "服务器内部错误，请稍后再试", nil))
}

func Unauthorized(ctx *gin.Context, msg string) {
	Failure(ctx, http.StatusOK, NewResponse(4005, msg, nil))
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
