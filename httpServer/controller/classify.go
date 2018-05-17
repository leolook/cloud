package controller

import (
	"cloud/httpServer/response"
	"cloud/httpServer/service"
	"cloud/httpServer/service/impl"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Classify struct{}

var classifyService service.ClassifyServiceIntf

func init() {
	if classifyService == nil {
		classifyService = new(impl.ClassifyServiceImpl)
	}
}

//获取所有分类
func (the Classify) All(c *gin.Context) {
	data := classifyService.All()
	c.JSON(http.StatusOK, response.GetSuccessResponse(data))
}
