package httpServer

import (
	"cloud/common/config"
	"cloud/common/flag"
	"cloud/httpServer/controller"
	"cloud/httpServer/controller/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//请求url
const (
	HEAD               = "/cloud"
	ADMIN              = HEAD + "/admin"
	USER               = HEAD + "/user"
	PING               = HEAD + "/ping"
	ADMIN_UPLOAD_FILE  = "/uploadFile"   //上传文件
	ADMIN_DEL_FILE     = "/delFile"      //文件删除
	ADMIN_LOGIN        = "/login"        //登录
	ADMIN_LOGIN_OUT    = "/loginOut"     //推出登录
	ADMIN_VIDEO_ADD    = "/video/add"    //视频添加
	ADMIN_VIDEO_UPDATE = "/video/update" //视频修改
	ADMIN_VIDEO_GET    = "/video/get"    //视频获取
	ADMIN_VIDEO_LIST   = "/video/list"   //视频列表
	ADMIN_CLASSIFY_ALL = "/classify/all" //所有分类
	USER_VIDEO_LIST    = "/video/list"   //视频列表

)

//启动server
func StartUpServer() {
	gin.SetMode(gin.ReleaseMode)
	route := gin.Default()
	route.Static(fmt.Sprintf("%s/%s", HEAD, flag.FilePath), flag.FilePath) //对外放开静态文件
	route.GET(PING, controller.Network{}.Ping)
	route.GET("/metrics", func(c *gin.Context) {
		h := promhttp.Handler()
		h.ServeHTTP(c.Writer, c.Request)
	})
	//后端
	admin := route.Group(ADMIN)
	admin.POST(ADMIN_LOGIN, controller.Admin{}.Login)
	admin.GET(ADMIN_LOGIN_OUT, controller.Admin{}.LoginOut)
	//后端token验证
	admin.Use(CheckToken)
	admin.POST(ADMIN_UPLOAD_FILE, controller.File{}.Upload)
	admin.GET(ADMIN_DEL_FILE, controller.File{}.DelFile)
	admin.POST(ADMIN_VIDEO_ADD, controller.Video{}.Add)
	admin.POST(ADMIN_VIDEO_UPDATE, controller.Video{}.Update)
	admin.GET(ADMIN_VIDEO_GET, controller.Video{}.Get)
	admin.POST(ADMIN_VIDEO_LIST, controller.Video{}.List)
	admin.GET(ADMIN_CLASSIFY_ALL, controller.Classify{}.All)
	//前端
	font := route.Group(USER)
	font.POST(USER_VIDEO_LIST, user.Video{}.List)

	route.Run(config.GetConf().Http.Addr)
}
