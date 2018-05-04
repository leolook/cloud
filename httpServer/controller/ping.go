package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Network struct{}

func (Network) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
