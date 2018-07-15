package httpServer

import (
	"bytes"
	logger "github.com/alecthomas/log4go"
	"cloud/constants"
	"cloud/httpServer/rsp"
	"cloud/httpServer/service"
	"cloud/httpServer/service/impl"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

var adminService service.AdminServiceIntf

func init() {
	if adminService == nil {
		adminService = new(impl.AdminServiceImpl)
	}
}

//请求参数日志打印
func logRequest(c *gin.Context) {
	url := c.Request.URL.Path
	if url == ADMIN+ADMIN_UPLOAD_FILE || url == ADMIN+ADMIN_DEL_FILE {
		return
	}
	var contentType string
	if len(c.Request.Header[constants.CONTENT_TYPE]) >= 1 {
		contentType = c.Request.Header[constants.CONTENT_TYPE][0]
	}
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error(fmt.Sprintf("Read body err:%v", err))
		return
	}
	logger.Info(fmt.Sprintf("Request before,URL:%s,contentType:%s,body:%v", c.Request.URL, contentType, string(buf)))
	rc := ioutil.NopCloser(bytes.NewBuffer(buf))
	c.Request.Body = rc
}

//后端token验证
func CheckToken(c *gin.Context) {
	userId := c.GetHeader(constants.HTTP_HEADER_USER_ID)
	token := c.GetHeader(constants.HTTP_HEADER_TOKEN)
	logger.Info(fmt.Sprintf("userId:%s,token:%s", userId, token))
	if userId == constants.STR_IS_EMPTY || token == constants.STR_IS_EMPTY {
		rsp.NewRsp(constants.CODE_TOKEN_INVALID, nil).Reply(c)
		c.Abort()
		return
	}
	suc := adminService.CheckSession(userId, token)
	if !suc {
		rsp.NewRsp(constants.CODE_TOKEN_INVALID, nil).Reply(c)
		c.Abort()
		return
	}
	logRequest(c)
	c.Next()
}
