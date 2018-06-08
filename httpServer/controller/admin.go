package controller

import (
	"cloud/constants"
	"cloud/httpServer/rsp"
	"cloud/httpServer/service"
	"cloud/httpServer/service/impl"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Admin struct{}

var adminService service.AdminServiceIntf

func init() {
	if adminService == nil {
		adminService = new(impl.AdminServiceImpl)
	}
}

//登录
func (Admin) Login(c *gin.Context) {
	userName := c.PostForm(constants.HTTP_ADMIN_LOGIN_USER_NAME)
	password := c.PostForm(constants.HTTP_ADMIN_LOGIN_PASSWORD)
	if userName == constants.STR_IS_EMPTY || password == constants.STR_IS_EMPTY {
		c.JSON(http.StatusOK, rsp.NewRsp(constants.CODE_PARAM_IS_NULL, constants.ERR_USERNAME_OR_PASSWORD_CAN_NOT_BE_EMPTY))
		return
	}
	data := adminService.CheckLogin(userName, password)
	if data == nil {
		c.JSON(http.StatusOK, rsp.NewRsp(constants.CODE_PARAM_IS_WRONG, constants.ERR_USERNAME_OR_PASSWORD_IS_WRONG))
		return
	}
	c.JSON(http.StatusOK, rsp.NewSucRsp(data))
}

//退出登录
func (Admin) LoginOut(c *gin.Context) {
	userId := c.GetHeader(constants.HTTP_HEADER_USER_ID)
	success := adminService.LoginOut(userId)
	if success {
		c.JSON(http.StatusOK, rsp.NewSucRsp(nil))
	}
	c.JSON(http.StatusOK, rsp.NewRsp(constants.CODE_SYSTEM_ERROR, constants.ERR_LOGIN_OUT))
}
