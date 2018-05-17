package service

import "cloud/httpServer/bean"

type VideoServiceIntf interface {
	AddOrUpdate(userId string, video *bean.VideoBean) bool  //视频添加
	IsExistByName(name string) int64                        //video 是否已存在
	GetById(id int64) *bean.VideoBean                       //根据video id获取视频信息
	ListByPage(req bean.VideoPageReq) *bean.BasePageRes     //视频分页
	UserListByPage(req bean.VideoPageReq) *bean.BasePageRes //用户侧视频分页
}
