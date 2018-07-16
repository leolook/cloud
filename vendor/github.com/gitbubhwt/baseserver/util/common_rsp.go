package util

type Rsp struct {
	Code    int64       `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func FailRsp(code int64, message string) Rsp {
	return Rsp{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

func SucRsp(code int64, data interface{}) Rsp {
	return Rsp{
		Code:    code,
		Message: "",
		Data:    data,
	}
}
