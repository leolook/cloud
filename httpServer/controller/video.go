package controller

import (
	"cloud/common/util"
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
func (Video) Add(c *gin.Context) {
	name := c.PostForm("name")
	describe := c.PostForm("describe")
	classify := c.PostForm("classify")
	path := c.PostForm("path")

	name = util.StrRemoveSpace(name)
	describe = util.StrRemoveSpace(describe)
	path = util.StrRemoveSpace(path)



	video := bean.VideoBean{
		Name:     name,
		Describe: describe,
		Classify: classify,
		Path:     path,
	}

	success := videoService.Add(&video)
	if !success {
		c.JSON(http.StatusOK, response.GetResponse())
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
