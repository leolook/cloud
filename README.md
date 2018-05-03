# cloud
my cloud

#本地装govendeor tool
go get -u -v github.com/kardianos/govendor

#初始化自动创建vendor文件夹
govendor init

#vendor中添加第三方包
govendor add +e

#文件url访问格式
get http://127.0.0.1:8010/cloud/file/2018-05-03/png/1525312817.png
#上传文件
post http://127.0.0.1:8010/cloud/uploadFile
参数: name file
#删除文件
get http://127.0.0.1:8010/cloud/delFile?path=file/2018-05-03/png/1525316454.png

