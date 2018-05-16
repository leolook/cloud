package httpServer

import (
	"cloud/common/logger"
	"cloud/constants"
	"cloud/httpServer/response"
	"cloud/httpServer/service"
	"cloud/httpServer/service/impl"
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
)

var adminService service.AdminServiceIntf

func init() {
	if adminService == nil {
		adminService = new(impl.AdminServiceImpl)
	}
}

//过滤器
func Filter(c *gin.Context) {
	url := c.Request.URL.String()
	if strings.Index(url, "?") != -1 {
		url = strings.Split(url, "?")[0]
	}
	if !adminSessionIsExist(c, url) { //管理员会话校验不通过
		c.JSON(200, response.GetResponse(constants.CODE_TOKEN_INVALID, nil))
		return
	}
	value, ok := HttpRoute[url]
	if !ok {
		log := fmt.Sprintf("Not found this action,ok:%v,url:%s", ok, url)
		logger.Error(log)
		c.String(200, log)
		return
	}
	method := strings.ToUpper(c.Request.Method)
	if method != value.Method {
		logger.Error(fmt.Sprintf("Request method is wrong,method:%s", method))
		return
	}
	value.Action(c) //执行方法
}

func adminSessionIsExist(c *gin.Context, url string) bool {
	match, _ := regexp.MatchString("/cloud/admin/([a-z]+)", url)
	if !match {
		return true
	}
	if url == ADMIN_LOGIN || url == ADMIN_LOGIN_OUT { //登录和推出 不走会话校验
		return true
	}
	userId := c.GetHeader(constants.HTTP_HEADER_USER_ID)
	token := c.GetHeader(constants.HTTP_HEADER_TOKEN)
	logger.Info(fmt.Sprintf("userId:%s,token:%s", userId, token))
	if userId == constants.STR_IS_EMPTY || token == constants.STR_IS_EMPTY {
		return false
	}
	return adminService.CheckSession(userId, token)
}
