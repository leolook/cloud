package controller

import "github.com/gin-gonic/gin"

type Network struct{}

func (Network) Ping(c *gin.Context) {
	c.String(200, "pong")
}
