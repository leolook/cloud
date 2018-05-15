package impl

import (
	"cloud/common/db"
	"cloud/common/logger"
	"cloud/constants"
	"cloud/httpServer/bean"
	"fmt"
)

type VideoDaoImpl struct{}

//根据名称判断video是否已经存在
func (dao VideoDaoImpl) IsExistByName(name string) bool {
	engine := db.GetEngine()
	var count int64
	isExist, err := engine.SQL(constants.ADMIN_VIDEO_IS_EXIST_BY_NAME, name).Get(&count)
	if isExist || err != nil {
		logger.Warn(fmt.Sprintf("Not found this video:%s,err:%v", name, err))
		return false
	}
	if count <= 0 {
		return false
	}
	return true
}

//插入video
func (dao VideoDaoImpl) Insert(video *bean.VideoBean) bool {

}
