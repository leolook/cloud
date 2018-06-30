package rsp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Rsp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func NewRsp(code int, data interface{}) *Rsp {
	res := new(Rsp)
	res.Code, res.Data = code, data
	return res
}

func NewSucRsp(data interface{}) *Rsp {
	return NewRsp(0, data)
}

func (r *Rsp) Reply(c *gin.Context) {
	c.JSON(http.StatusOK, r)
}
