package dao

import "cloud/httpServer/bean"

type VideoDaoIntf interface {
	IsExistByName(name string) int64                 //根据名称判断video是否已经存在
	Insert(useId string, video *bean.VideoBean) bool //插入video
	Update(useId string, video *bean.VideoBean) bool //修改video
	GetVideoInfo(id int64) *bean.VideoBean           //获取视频信息
	GetVideoPathInfo(id int64) []bean.VideoPathBean  //获取视频文件路径信息
}
