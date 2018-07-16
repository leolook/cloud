package fileserver

import (
	"github.com/gin-gonic/gin/json"
	"net/http"
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

func ReplyJson(w http.ResponseWriter, data interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	header := w.Header()
	header["Content-Type"] = []string{"application/json; charset=utf-8"}
	_, err = w.Write(bytes)
	return err
}
