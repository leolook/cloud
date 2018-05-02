package httpServer

import (
	"cloud/common/config"
	"cloud/httpServer/controller"
	"github.com/gin-gonic/gin"
)

type Action func(c *gin.Context)

type Controller struct {
	Method string
	Action Action
}

func addRoute(method string, action func(c *gin.Context)) *Controller {
	con := new(Controller)
	con.Method, con.Action = method, action
	return con
}

var HttpRoute map[string]*Controller

//请求方式
const (
	GET_REQUEST  = "GET"
	POST_REQUEST = "POST"
)

//请求url
const (
	HEAD = "/cloud"
	PING = HEAD + "/ping"
)

func init() {
	HttpRoute = make(map[string]*Controller)
	HttpRoute[PING] = addRoute(GET_REQUEST, controller.Network{}.Ping)
}

//启动server
func StartUpServer() {
	gin.SetMode(gin.ReleaseMode)
	route := gin.Default()
	for url, val := range HttpRoute {
		route.Handle(val.Method, url, Filter)
	}
	route.Run(config.GetConf().Http.Addr)
}
