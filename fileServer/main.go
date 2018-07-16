package main

import (
	"github.com/gin-gonic/gin"
	. "github.com/gitbubhwt/fileserver"
)

func main() {
	gin.SetMode(gin.DebugMode)
	g := gin.Default()

	g.POST("file/upload", func(c *gin.Context) {
		Upload(c.Writer, c.Request)
	})

	g.POST("file/remove", func(c *gin.Context) {
		Remove(c.Writer, c.Request)
	})

	g.Run(":8090")
}
