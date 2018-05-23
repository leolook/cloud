package controller

import (
	"cloud/common/flag"
	"cloud/common/logger"
	"cloud/constants"
	"cloud/httpServer/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type File struct{}

//上传文件
func (File) Upload(c *gin.Context) {
	tempFile, err := c.FormFile(constants.HTTP_ADMIN_FILE_FILE)
	if err != nil {
		logger.Error(fmt.Sprintf("File upload fail,err:%v", err))
		return
	}
	tempFileName := tempFile.Filename
	tempFileType := constants.OTHER_FILE_DIRECTORY
	if strings.Index(tempFileName, ".") != -1 {
		tempFileType = strings.Split(tempFileName, ".")[1]
	}
	folder := time.Now().Format(constants.TIME_FORMAT_Y_M_D)
	path := fmt.Sprintf("%s/%s/%s", flag.FilePath, folder, tempFileType)
	if err := os.MkdirAll(path, os.ModePerm); err != nil { //生成多级目录
		logger.Error(fmt.Sprintf("Mkdir folder fail,err:%v", err))
		return
	}
	fileName := fmt.Sprintf("%v", time.Now().Unix())
	if tempFileType != constants.OTHER_FILE_DIRECTORY {
		fileName = fmt.Sprintf("/%v.%s", fileName, tempFileType)
	}
	path += fileName
	newFile, err := os.Create(path)
	if err != nil {
		logger.Error(fmt.Sprintf("Create file fail,err:%v,path:%s", err, path))
		return
	}
	defer newFile.Close()
	temp, err := tempFile.Open()
	if err != nil {
		logger.Error(fmt.Sprintf("TempFile open fail,err:%v", err))
		return
	}
	defer temp.Close()
	_, err = io.Copy(newFile, temp)
	if err != nil {
		logger.Error(fmt.Sprintf("Copy file fail,err:%v", err))
		return
	}
	c.JSON(http.StatusOK, response.GetSuccessResponse(path))
}

//文件删除
func (File) DelFile(c *gin.Context) {
	path := c.Query(constants.HTTP_ADMIN_FILE_PATH)
	if path == constants.STR_IS_EMPTY {
		c.JSON(http.StatusOK, response.GetResponse(constants.CODE_PARAM_IS_NULL, constants.ERR_FILE_PATH_CAN_NOT_BE_EMPTY))
		return
	}
	if !checkFileIsExist(path) {
		c.JSON(http.StatusOK, response.GetResponse(constants.CODE_PARAM_IS_WRONG, constants.ERR_FILE_PATH_IS_NOT_EXIST))
		return
	}
	err := os.Remove(path)
	if err != nil {
		logger.Error(fmt.Sprintf("Del file fail,err:%v", err))
		c.JSON(http.StatusOK, response.GetResponse(constants.CODE_SYSTEM_ERROR, constants.ERR_SYSTEM_ERROR))
		return
	}
	c.JSON(http.StatusOK, response.GetSuccessResponse(constants.SUCCESS_DEL_FILE))
}

//判断文件是否存在
func checkFileIsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
