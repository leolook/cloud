package service

import "cloud/httpServer/bean"

type VideoServiceIntf interface {
	Add(userId string, video *bean.VideoBean) bool //视频添加
	IsExistByName(name string) bool                //video 是否已存在
}
