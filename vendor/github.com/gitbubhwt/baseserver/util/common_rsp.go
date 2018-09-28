package util

import (
	"cloud/lib/log"
	"github.com/gin-gonic/gin/json"
)

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

func FailByte(code int64, message string) []byte {
	rsp := FailRsp(code, message)
	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error(err)
	}
	return bytes
}

func SucByte(code int64, data interface{}) []byte {
	rsp := SucRsp(code, data)
	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error(err)
	}
	return bytes
}
