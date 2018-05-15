package dao

import "cloud/httpServer/bean"

type VideoDaoIntf interface {
	IsExistByName(name string) bool    //根据名称判断video是否已经存在
	Insert(video *bean.VideoBean) bool //插入video
}
