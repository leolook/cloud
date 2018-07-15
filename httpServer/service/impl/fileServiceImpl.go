package impl

import (
	"cloud/common/config"
	logger "github.com/alecthomas/log4go"
	"cloud/constants"
	"cloud/httpServer/bean"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

type FileSrvImpl struct{}

func (impl FileSrvImpl) Upload(tempFile *multipart.FileHeader) (*bean.UploadFile, error) {
	tempFileName := tempFile.Filename
	tempFileType := constants.OTHER_FILE_DIRECTORY
	if strings.Index(tempFileName, ".") != -1 {
		tempFileType = strings.Split(tempFileName, ".")[1]
	}
	folder := time.Now().Format(constants.TIME_FORMAT_Y_M_D)
	path := fmt.Sprintf("%s/%s/%s", config.GetConf().UploadFile.Path, folder, tempFileType)
	if err := os.MkdirAll(path, os.ModePerm); err != nil { //生成多级目录
		logger.Error(fmt.Sprintf("Mkdir folder fail,err:%v", err))
		return nil, err
	}
	fileName := fmt.Sprintf("%v", time.Now().Unix())
	if tempFileType != constants.OTHER_FILE_DIRECTORY {
		fileName = fmt.Sprintf("/%v.%s", fileName, tempFileType)
	}
	path += fileName
	newFile, err := os.Create(path)
	if err != nil {
		logger.Error(fmt.Sprintf("Create file fail,err:%v,path:%s", err, path))
		return nil, err
	}
	defer newFile.Close()
	temp, err := tempFile.Open()
	if err != nil {
		logger.Error(fmt.Sprintf("TempFile open fail,err:%v", err))
		return nil, err
	}
	defer temp.Close()
	_, err = io.Copy(newFile, temp)
	if err != nil {
		logger.Error(fmt.Sprintf("Copy file fail,err:%v", err))
		return nil, err
	}
	suc := realFileDao.Save(path)
	if !suc {
		go func() {
			time.Sleep(1 * time.Second)
			impl.Del(path)
		}()
		return nil, errors.New(fmt.Sprintf("save path fail,[suc=%v]", suc))
	}
	return &bean.UploadFile{Path: path, CreateTime: time.Now().Unix()}, nil
}

func (impl FileSrvImpl) Del(path string) error {
	err := os.Remove(path)
	if err != nil {
		logger.Error(fmt.Sprintf("Del file fail,err:%v", err))
		return err
	}
	realFileDao.DelByPath(path)
	return nil
}

func (impl FileSrvImpl) IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
