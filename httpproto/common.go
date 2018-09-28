package httpproto

type Rsp struct {
	Code    int64       `json:"code"`
	Message string      `json:"msg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Fail(code int64, msg string) *Rsp {
	return &Rsp{
		Code: code, Message: msg,
	}
}

func Suc(data interface{}) *Rsp {
	return &Rsp{
		Code: 200, Message: "", Data: data,
	}
}
