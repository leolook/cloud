package impl

import (
	"cloud/httpServer/bean"
	"cloud/httpServer/dao"
	"cloud/httpServer/dao/impl"
)

type ClassifyServiceImpl struct{}

var classifyDao dao.ClassifyDaoIntf

func init() {
	if classifyDao == nil {
		classifyDao = new(impl.ClassifyDaoImpl)
	}
}

//获取所有分类
func (impl ClassifyServiceImpl) All() []bean.ClassifyBean {
	return classifyDao.All()
}
