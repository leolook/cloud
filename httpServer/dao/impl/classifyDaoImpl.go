package impl

import (
	"cloud/common/db"
	"cloud/constants"
	"cloud/httpServer/bean"
)

type ClassifyDaoImpl struct{}

//获取所有分类
func (impl ClassifyDaoImpl) All() []bean.ClassifyBean {
	engine := db.GetEngine()
	classifys := make([]bean.ClassifyBean, 0)
	engine.SQL(constants.ADMIN_CLASSIFY_ALL_SQL).Find(&classifys)
	return classifys
}
