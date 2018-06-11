package impl

import (
	"cloud/httpServer/bean"
	"cloud/httpServer/dao"
	. "cloud/httpServer/dao/impl"
)

type RealFileServiceImpl struct{}

var realFileDao dao.RealFileDaoIntf

func init() {
	if realFileDao == nil {
		realFileDao = new(RealFilelImpl)
	}
}

//分页列表
func (impl RealFileServiceImpl) List(req bean.RealFilePage) []bean.RealFile {
	return realFileDao.List(req)
}

//批量删除
func (impl RealFileServiceImpl) Del(ids string) bool {
	return realFileDao.Del(ids)
}
