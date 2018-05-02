package controller

import (
	"cloud/common/flag"
	"cloud/common/logger"
	"cloud/httpServer/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strings"
	"time"
)

type File struct{}

type FileUpload struct {
	Code uint8  `json:"code"`
	Path string `json:"path"`
}

//上传文件
func (File) Upload(c *gin.Context) {
	tempFileName := c.PostForm("name")
	tempFile, err := c.FormFile("file")
	if err != nil {
		logger.Error(fmt.Sprintf("File upload fail,err:%v", err))
		return
	}
	var tempFileType string
	if strings.Index(tempFileName, ".") != -1 {
		tempFileType = strings.Split(tempFileName, ".")[1]
	}
	folder := time.Now().Format("2006-01-02")
	path := fmt.Sprintf("%s/%s/%s", flag.FilePath, folder, tempFileType)
	if err := os.MkdirAll(path, os.ModePerm); err != nil { //生成多级目录
		logger.Error(fmt.Sprintf("Mkdir folder fail,err:%v", err))
		return
	}
	path += fmt.Sprintf("/%v.%s", time.Now().Unix(), tempFileType)
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
