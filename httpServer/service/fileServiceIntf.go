package service

import (
	"cloud/httpServer/bean"
	"mime/multipart"
)

type FileSrvIntf interface {
	Upload(tempFile *multipart.FileHeader) (*bean.UploadFile, error)
	Del(path string) error
	IsExist(path string) bool
}
