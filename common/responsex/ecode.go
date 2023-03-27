package responsex

const (
	ErrSystem          ResCode = 500
	ErrNotFound        ResCode = 404
	ErrTooManyRequests ResCode = 429
)

const (
	ErrBadRequest ResCode = 40000 + iota
	ErrBadValidation
	ErrJWT
	ErrDataExist
	ErrDataNotExist
)

var codeMsgMap = map[ResCode]string{
	ErrBadRequest:    "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式",
	ErrBadValidation: "请求参数校验失败",
	ErrJWT:           "令牌有误",
	ErrDataExist:     "数据已存在",
	ErrDataNotExist:  "数据不存在",

	ErrSystem:          "服务器出错，请稍后再试",
	ErrNotFound:        "请确认 url 和请求方法是否正确",
	ErrTooManyRequests: "请求过于频繁",
}
