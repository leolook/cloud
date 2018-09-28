package service

import (
	pro "cloud/httpproto/monitor"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var req pro.RegReq
	err := c.Bind(&req)
	if err != nil {
		
	}
}
