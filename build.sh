#!/bin/sh

###获取服务名
name=""
len=${#1}
for ((i=0;i<$len;i++))
do
   tmp=${1:i:1}
   if [ "$tmp" == "/" ]
   then
      break
   fi

   name+="$tmp"
done

path=$1/main.go

###判断文件是否存在
if [ ! -f "$path" ]
then
  echo "not found file,[path=$path]"
  exit
fi

###进入该目录
cd $1

###判断是否执行成功
if [ $? != 0 ]
then
  echo "fail do:$?"
  exit
fi


###编译
echo "building:$path"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

###判断是否执行成功
if [ $? != 0 ]
then
  echo "failed to build:$path"
  exit
fi

echo "succeed to build:$path"

###上传文件
scp main root@39.105.89.132:/root/hwt/project/$1/$name

