# cloud
my cloud

#本地装govendeor tool
go get -u -v github.com/kardianos/govendor

#初始化自动创建vendor文件夹
govendor init

#vendor中添加第三方包
govendor add +e
