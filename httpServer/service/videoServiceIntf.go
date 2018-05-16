package service

import "cloud/httpServer/bean"

type VideoServiceIntf interface {
	AddOrUpdate(userId string, video *bean.VideoBean) bool //视频添加
	IsExistByName(name string) int64                       //video 是否已存在
	GetById(id int64) *bean.VideoBean                      //根据video id获取视频信息
}
