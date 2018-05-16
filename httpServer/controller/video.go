package controller

import (
	"cloud/common/logger"
	"cloud/common/util"
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

//添加
func (this Video) Add(c *gin.Context) {
	var videoBean bean.VideoBean
	err := c.Bind(&videoBean)
	if err != nil {
		logger.Error(fmt.Sprintf("Bind fail,err:%v", err))
		c.JSON(http.StatusOK, response.GetResponse(constants.CODE_SYSTEM_ERROR, constants.ERR_ADD_VIDEO_FAIL))
		return
	}
	video, res := this.checkParams(c, &videoBean)
	if res != nil {
		c.JSON(http.StatusOK, res)
		return
	}
	if videoService.IsExistByName(video.Name) > 0 {
		c.JSON(http.StatusOK, response.GetResponse(constants.CODE_PARAM_IS_REPEAT, constants.ERR_ADD_VIDEO_REPEAT))
		return
	}
	userId := c.GetHeader(constants.HTTP_HEADER_USER_ID)
	success := videoService.AddOrUpdate(userId, video)
	if !success {
		c.JSON(http.StatusOK, response.GetResponse(constants.CODE_SYSTEM_ERROR, constants.ERR_ADD_VIDEO_FAIL))
		return
	}
	c.JSON(http.StatusOK, response.GetSuccessResponse(nil))
}

//编辑
func (Video) Get(c *gin.Context) {

}

//修改
func (this Video) Update(c *gin.Context) {
	var videoBean bean.VideoBean
	err := c.Bind(&videoBean)
	if err != nil {
		logger.Error(fmt.Sprintf("Bind fail,err:%v", err))
		c.JSON(http.StatusOK, response.GetResponse(constants.CODE_SYSTEM_ERROR, constants.ERR_ADD_VIDEO_FAIL))
		return
	}
	video, res := this.checkParams(c, &videoBean)
	if res != nil {
		c.JSON(http.StatusOK, res)
		return
	}
	tempId := videoService.IsExistByName(video.Name)
	if tempId != 0 && tempId != video.Id {
		c.JSON(http.StatusOK, response.GetResponse(constants.CODE_PARAM_IS_REPEAT, constants.ERR_ADD_VIDEO_REPEAT))
		return
	}
	userId := c.GetHeader(constants.HTTP_HEADER_USER_ID)
	success := videoService.AddOrUpdate(userId, video)
	if !success {
		c.JSON(http.StatusOK, response.GetResponse(constants.CODE_SYSTEM_ERROR, constants.ERR_UPDATE_VIDEO_FAIL))
		return
	}
	c.JSON(http.StatusOK, response.GetSuccessResponse(nil))
}

//分页列表
func (Video) List(c *gin.Context) {

}

//参数校验
func (Video) checkParams(c *gin.Context, video *bean.VideoBean) (*bean.VideoBean, *response.Response) {
	video.Name = util.StrRemoveSpace(video.Name)
	video.Info = util.StrRemoveSpace(video.Info)
	if video.Name == constants.STR_IS_EMPTY {
		return nil, response.GetResponse(constants.CODE_PARAM_IS_NULL, constants.ERR_VIDEO_NAME_CAN_NOT_BE_EMPTY)
	}
	if video.Info == constants.STR_IS_EMPTY {
		return nil, response.GetResponse(constants.CODE_PARAM_IS_NULL, constants.ERR_VIDEO_DESCRIBE_CAN_NOT_BE_EMPTY)
	}
	if video.Classify == constants.INT_IS_ZERO {
		return nil, response.GetResponse(constants.CODE_PARAM_IS_NULL, constants.ERR_VIDEO_CLASSIFY_CAN_NOT_BE_EMPTY)
	}
	if len(video.Path) == constants.INT_IS_ZERO {
		return nil, response.GetResponse(constants.CODE_PARAM_IS_NULL, constants.ERR_VIDEO_FILE_CAN_NOT_BE_EMPTY)
	}
	return video, nil
}
