package controller

import (
	logger "github.com/alecthomas/log4go"
	"cloud/constants"
	"cloud/httpServer/rsp"
	"cloud/httpServer/service"
	"cloud/httpServer/service/impl"
	"fmt"
	"github.com/gin-gonic/gin"
)

type File struct{}

var fileSrv service.FileSrvIntf

func init() {
	fileSrv = new(impl.FileSrvImpl)
}

//上传文件
func (File) Upload(c *gin.Context) {
	tempFile, err := c.FormFile(constants.HTTP_ADMIN_FILE_FILE)
	if err != nil {
		logger.Error(fmt.Sprintf("File upload fail,err:%v", err))
		rsp.NewRsp(constants.CODE_SYSTEM_ERROR, "service err").Reply(c)
		return
	}
	uploadFile, err := fileSrv.Upload(tempFile)
	if err != nil {
		rsp.NewRsp(constants.CODE_SYSTEM_ERROR, "service err").Reply(c)
		return
	}
	rsp.NewSucRsp(uploadFile).Reply(c)
}

//文件删除
func (File) DelFile(c *gin.Context) {
	path := c.Query(constants.HTTP_ADMIN_FILE_PATH)
	if path == constants.STR_IS_EMPTY {
		rsp.NewRsp(constants.CODE_PARAM_IS_NULL, constants.ERR_FILE_PATH_CAN_NOT_BE_EMPTY).Reply(c)
		return
	}
	if !fileSrv.IsExist(path) {
		rsp.NewSucRsp(constants.SUCCESS_DEL_FILE).Reply(c)
		return
	}
	err := fileSrv.Del(path)
	if err != nil {
		rsp.NewRsp(constants.CODE_SYSTEM_ERROR, constants.ERR_SYSTEM_ERROR).Reply(c)
		return
	}
	rsp.NewSucRsp(constants.SUCCESS_DEL_FILE).Reply(c)
}
