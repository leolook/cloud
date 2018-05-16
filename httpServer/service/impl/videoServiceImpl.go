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

//视频添加
func (impl VideoServiceImpl) Add(userId string, video *bean.VideoBean) bool {
	return videoDao.Insert(userId, video)
}

//video 是否已存在
func (impl VideoServiceImpl) IsExistByName(name string) bool {
	return videoDao.IsExistByName(name)
}
