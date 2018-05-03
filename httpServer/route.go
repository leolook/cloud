package httpServer

import (
	"cloud/common/config"
	"cloud/common/flag"
	"cloud/httpServer/controller"
	"fmt"
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
	HEAD        = "/cloud"
	PING        = HEAD + "/ping"
	UPLOAD_FILE = HEAD + "/uploadFile" //上传文件
	DEL_FILE    = HEAD + "/delFile"    //文件删除
)

func init() {
	HttpRoute = make(map[string]*Controller)
	HttpRoute[PING] = addRoute(GET_REQUEST, controller.Network{}.Ping)
	HttpRoute[UPLOAD_FILE] = addRoute(POST_REQUEST, controller.File{}.Upload)
	HttpRoute[DEL_FILE] = addRoute(GET_REQUEST, controller.File{}.DelFile)
}

//启动server
func StartUpServer() {
	gin.SetMode(gin.ReleaseMode)
	route := gin.Default()
	route.Static(fmt.Sprintf("%s/%s", HEAD, flag.FilePath), flag.FilePath) //对外放开静态文件
	for url, val := range HttpRoute {
		route.Handle(val.Method, url, Filter)
	}
	route.Run(config.GetConf().Http.Addr)
}
