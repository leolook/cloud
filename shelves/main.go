package main

import (
	. "cloud/shelves/controller"
	. "cloud/shelves/middle"
	"github.com/gin-gonic/gin"
	base "github.com/gitbubhwt/baseserver/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	b := base.BaseServer{
		Server:        &Srv{},
		Group:         r.Group("shelves"),
		CommonReqFunc: GetCommonReq,
	}
	b.Register()

	r.Run(":3030")
}
