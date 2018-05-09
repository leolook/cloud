package service

import "cloud/httpServer/bean"

type VideoServiceIntf interface {
	Add(video *bean.VideoBean) bool //视频添加
}
