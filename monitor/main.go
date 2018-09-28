package main

import (
	"cloud/common/conf"
	"cloud/monitor/service"
	"github.com/gin-gonic/gin"
)

func main() {

	eng := gin.Default()
	r := eng.Group("monitor")
	r.POST("/register", service.Register)
	eng.Run(conf.Monitor().Monitor.Addr)
}
