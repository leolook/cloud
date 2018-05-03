package controller

import (
	"cloud/common/flag"
	"cloud/common/logger"
	"cloud/constants"
	"cloud/httpServer/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strings"
	"time"
)

type File struct{}

//上传文件
func (File) Upload(c *gin.Context) {
	tempFileName := c.PostForm("name")
	tempFile, err := c.FormFile("file")
	if err != nil {
		logger.Error(fmt.Sprintf("File upload fail,err:%v", err))
		return
	}
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
	temp, err := tempFile.Open()
	if err != nil {
		logger.Error(fmt.Sprintf("TempFile open fail,err:%v", err))
		return
	}
	_, err = io.Copy(newFile, temp)
	if err != nil {
		logger.Error(fmt.Sprintf("Copy file fail,err:%v", err))
		return
	}
	c.JSON(200, response.GetSuccessResponse(path))
}
