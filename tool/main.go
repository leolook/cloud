package main

import (
	. "cloud/tool/controller"
	"cloud/tool/middle"
	"github.com/gin-gonic/gin"
	base "github.com/gitbubhwt/baseserver/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.Static("/static/", "static")
	b := base.BaseServer{
		Server:        &Srv{},
		Group:         r.Group("tool"),
		CommonReqFunc: middle.Common,
	}
	b.Register()

	r.Run(":12020")
}
