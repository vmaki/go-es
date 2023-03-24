package requestx

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"go-es/common/responsex"
)

func Validate(ctx *gin.Context, req IRequest) error {
	if err := ctx.ShouldBind(req); err != nil {
		return responsex.NewResponse(4000, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式", nil)
	}

	// 表单验证
	err := req.Generate(req)
	if err != nil {
		return err
	}

	return nil
}

func GoValidate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) error {
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	// 开始验证
	errs := govalidator.New(opts).ValidateStruct()
	if len(errs) > 0 {
		str := ""
		for _, v := range errs {
			str = v[0]
			break
		}

		return responsex.NewResponse(4001, str, nil)
	}

	return nil
}
