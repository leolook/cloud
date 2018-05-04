package controller

import (
	"cloud/constants"
	"cloud/httpServer/response"
	"cloud/httpServer/service"
	"cloud/httpServer/service/impl"
	"github.com/gin-gonic/gin"
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
	userName := c.PostForm("userName")
	password := c.PostForm("password")
	if userName == constants.STR_IS_EMPTY || password == constants.STR_IS_EMPTY {
		c.JSON(200, response.GetResponse(constants.CODE_PARAM_IS_NULL, constants.ERR_USERNAME_OR_PASSWORD_CAN_NOT_BE_EMPTY))
		return
	}
	data := adminService.CheckLogin(userName, password)
	if data == nil {
		c.JSON(200, response.GetResponse(constants.CODE_PARAM_IS_WRONG, constants.ERR_USERNAME_OR_PASSWORD_IS_WRONG))
		return
	}
}

//推出登录
func (Admin) LoginOut(c *gin.Context) {

}
