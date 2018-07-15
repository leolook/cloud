package main

import (
	. "cloud/shelves/controller"
	. "cloud/shelves/middle"
	base "github.com/baseserver/gin"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	b := base.BaseServer{
		Engine:        r,
		Server:        &Srv{},
		Group:         "shelves",
		CommonReqFunc: GetCommonReq,
	}
	b.Register()

	r.Run(":3030")
}
