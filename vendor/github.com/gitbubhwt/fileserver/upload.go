package fileserver

import (
	"fmt"
	log "github.com/alecthomas/log4go"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

//上传文件
func Upload(w http.ResponseWriter, req *http.Request) {

	//解析表单
	err := req.ParseForm()
	if err != nil {
		log.Error(fmt.Sprintf("Parse form fail,[err=%v]", err))
		ReplyJson(w, FailRsp(ERR_SERVER_ERR, "server err"))
		return
	}

	//获取文件
	_, head, err := req.FormFile("file")
	if head != nil {
		log.Error(fmt.Sprintf("Form file fail,[err=%v]", err))
		ReplyJson(w, FailRsp(ERR_SERVER_ERR, "server err"))
		return
	}

	bucket := req.FormValue("bucket")

	if bucket == "" {
		ReplyJson(w, FailRsp(ERR_IS_EMPTY, "bucket is empty"))
		return
	}

	dir, path := createPath(bucket, head.Filename)

	err = createFile(head, dir, path)
	if err != nil {
		log.Error(fmt.Sprintf("Create file fail,[err=%v] [dir=%s] [path=%s]", err, dir, path))
		ReplyJson(w, FailRsp(ERR_SERVER_ERR, "server error"))
		return
	}

	ReplyJson(w, SucRsp(http.StatusOK, path))
}

//创建路径
func createPath(bucket, name string) (string, string) {

	dir, fileType := "", "other"
	now := time.Now()
	date := fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day())

	if strings.Index(name, ".") != -1 {
		fileType = strings.Split(name, ".")[1]
	}

	dir = fmt.Sprintf("%s/%s/%s/", bucket, fileType, date)

	return dir, fmt.Sprintf("%s/%d.%s", dir, now.UnixNano(), fileType)
}

//创建文件
func createFile(tempFile *multipart.FileHeader, dir, path string) error {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil { //生成多级目录
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	temp, err := tempFile.Open()
	if err != nil {
		return err
	}
	defer temp.Close()
	_, err = io.Copy(file, temp)
	if err != nil {
		return err
	}
	return nil
}
