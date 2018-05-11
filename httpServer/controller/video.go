package controller

import (
	"cloud/common/util"
	"cloud/constants"
	"cloud/httpServer/bean"
	"cloud/httpServer/response"
	"cloud/httpServer/service"
	"cloud/httpServer/service/impl"
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
	video, res := this.checkParams(c)
	if res != nil {
		c.JSON(http.StatusOK, res)
		return
	}
	success := videoService.Add(video)
	if !success {
		c.JSON(http.StatusOK, response.GetResponse(constants.CODE_SYSTEM_ERROR, constants.ERR_ADD_VIDEO_FAIL))
	}
	c.JSON(http.StatusOK, response.GetSuccessResponse(nil))
}

//编辑
func (Video) Get(c *gin.Context) {

}

//修改
func (Video) Update(c *gin.Context) {

}

//分页列表
func (Video) List(c *gin.Context) {

}

//参数校验
func (Video) checkParams(c *gin.Context) (*bean.VideoBean, *response.Response) {
	name := c.PostForm("name")
	describe := c.PostForm("describe")
	classify := c.PostForm("classify")
	path := c.PostForm("path")

	name = util.StrRemoveSpace(name)
	describe = util.StrRemoveSpace(describe)
	path = util.StrRemoveSpace(path)

	if name == constants.STR_IS_EMPTY {
		return nil, response.GetResponse(constants.CODE_PARAM_IS_NULL, constants.ERR_VIDEO_NAME_CAN_NOT_BE_EMPTY)
	}
	if describe == constants.STR_IS_EMPTY {
		return nil, response.GetResponse(constants.CODE_PARAM_IS_NULL, constants.ERR_VIDEO_DESCRIBE_CAN_NOT_BE_EMPTY)
	}
	if classify == constants.STR_IS_EMPTY {
		return nil, response.GetResponse(constants.CODE_PARAM_IS_NULL, constants.ERR_VIDEO_CLASSIFY_CAN_NOT_BE_EMPTY)
	}

	video := bean.VideoBean{
		Name:     name,
		Describe: describe,
		Classify: classify,
		Path:     path,
	}

	return &video, nil
}
