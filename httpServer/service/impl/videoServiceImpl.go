package impl

import (
	"cloud/httpServer/bean"
	"cloud/httpServer/dao"
	. "cloud/httpServer/dao/impl"
)

type VideoServiceImpl struct{}

var videoDao dao.VideoDaoIntf

func init() {
	if videoDao == nil {
		videoDao = new(VideoDaoImpl)
	}
}

//视频添加or修改
func (impl VideoServiceImpl) AddOrUpdate(userId string, video *bean.VideoBean) bool {
	if video.Id > 0 {
		return videoDao.Update(userId, video)
	}
	return videoDao.Insert(userId, video)
}

//video 是否已存在
func (impl VideoServiceImpl) IsExistByName(name string) int64 {
	return videoDao.IsExistByName(name)
}

//根据video id获取视频信息
func (impl VideoServiceImpl) GetById(id int64) *bean.VideoBean {
	video := videoDao.GetVideoInfo(id)
	if video == nil {
		return nil
	}
	paths := videoDao.GetVideoPathInfo(id)
	if paths == nil {
		return video
	}
	video.Path = paths
	return video
}
