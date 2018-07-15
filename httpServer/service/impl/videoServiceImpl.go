package impl

import (
	logger "github.com/alecthomas/log4go"
	"cloud/common/util"
	"cloud/httpServer/bean"
	"cloud/httpServer/dao"
	. "cloud/httpServer/dao/impl"
	"fmt"
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

//视频分页
func (impl VideoServiceImpl) ListByPage(req bean.VideoPageReq) *bean.BasePageRes {
	req.Name = util.StrRemoveSpace(req.Name)
	res := new(bean.BasePageRes)
	list := videoDao.List(req)
	if list == nil {
		return res
	}
	logger.Info(fmt.Sprintf("%v", list))
	count := videoDao.Count(req)
	if len(list) > 0 && count <= 0 {
		return nil
	}
	res.Total, res.Rows = count, list
	return res
}

//用户侧视频分页
func (impl VideoServiceImpl) UserListByPage(req bean.VideoPageReq) *bean.BasePageRes {
	req.Name = util.StrRemoveSpace(req.Name)
	res := new(bean.BasePageRes)
	list := videoDao.UserList(req)
	if list == nil {
		return res
	}
	logger.Info(fmt.Sprintf("%v", list))
	count := videoDao.Count(req)
	if len(list) > 0 && count <= 0 {
		return nil
	}
	res.Total, res.Rows = count, list
	return res
}
