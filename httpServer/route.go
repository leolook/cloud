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
	HEAD              = "/cloud"
	ADMIN             = HEAD + "/admin"
	PING              = HEAD + "/ping"
	ADMIN_UPLOAD_FILE = ADMIN + "/uploadFile" //上传文件
	ADMIN_DEL_FILE    = ADMIN + "/delFile"    //文件删除
	ADMIN_LOGIN       = ADMIN + "/login"      //登录
	ADMIN_LOGIN_OUT   = ADMIN + "／loginOut"   //推出登录

)

func init() {
	HttpRoute = make(map[string]*Controller)
	HttpRoute[PING] = addRoute(GET_REQUEST, controller.Network{}.Ping)
	HttpRoute[ADMIN_UPLOAD_FILE] = addRoute(POST_REQUEST, controller.File{}.Upload)
	HttpRoute[ADMIN_DEL_FILE] = addRoute(GET_REQUEST, controller.File{}.DelFile)
	HttpRoute[ADMIN_LOGIN] = addRoute(POST_REQUEST, controller.Admin{}.Login)
	HttpRoute[ADMIN_LOGIN_OUT] = addRoute(GET_REQUEST, controller.Admin{}.LoginOut)
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
