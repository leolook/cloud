package user

import (
	"cloud/common/logger"
	"cloud/constants"
	"cloud/httpServer/bean"
	"cloud/httpServer/response"
	"cloud/httpServer/service"
	"cloud/httpServer/service/impl"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Video struct{}

var videoService service.VideoServiceIntf

func init() {
	if videoService == nil {
		videoService = new(impl.VideoServiceImpl)
	}
}

//分页列表
func (Video) List(c *gin.Context) {
	var req bean.VideoPageReq
	err := c.Bind(&req)
	if err != nil {
		logger.Error(fmt.Sprintf("Bind fail,err:%v", err))
		c.JSON(http.StatusOK, response.GetResponse(constants.CODE_PARAM_IS_WRONG, constants.ERR_PARAM_IS_WRONG))
		return
	}
	data := videoService.UserListByPage(req)
	if data == nil {
		c.JSON(http.StatusOK, response.GetResponse(constants.CODE_SYSTEM_ERROR, constants.ERR_LIST_BY_PAGE_FAIL))
		return
	}
	c.JSON(http.StatusOK, response.GetSuccessResponse(data))
}
