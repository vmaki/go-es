package responsex

type ResCode int64

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = "未知错误"
	}

	return msg
}

type Response struct {
	Code ResCode     `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func NewResponse(code ResCode, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func NewResponseErr(code ResCode, msg ...string) *Response {
	m := code.Msg()
	if len(msg) > 0 {
		m = msg[0]
	}

	return &Response{
		Code: code,
		Msg:  m,
		Data: nil,
	}
}

func (r *Response) Error() string {
	return r.Msg
}
