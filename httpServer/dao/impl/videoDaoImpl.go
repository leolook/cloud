package impl

import (
	"cloud/common/db"
	"cloud/common/logger"
	"cloud/constants"
	"cloud/httpServer/bean"
	"fmt"
)

type VideoDaoImpl struct {
	BaseDaoImpl
}

//根据名称判断video是否已经存在
func (impl VideoDaoImpl) IsExistByName(name string) bool {
	engine := db.GetEngine()
	var count int64
	isExist, err := engine.SQL(constants.ADMIN_VIDEO_IS_EXIST_BY_NAME_SQL, name).Get(&count)
	if err != nil {
		logger.Warn(fmt.Sprintf("Not found this video:%s,err:%v", name, err))
		return false
	}
	if !isExist || count <= 0 {
		return false
	}
	return true
}

//插入video
func (impl VideoDaoImpl) Insert(useId string, video *bean.VideoBean) bool {
	session := impl.GetAdminSession(useId)
	if session == nil {
		return false
	}
	tx := db.GetEngine().NewSession()
	_, err := tx.Exec(constants.ADMIN_VIDEO_FILE_INSERT_SQL, video.Name, video.Info, video.Classify, session.UserName)
	if err != nil {
		logger.Error(fmt.Sprintf("Tx exec fail,err:%v", err))
		tx.Rollback()
		return false
	}
	lastId := impl.GetLastIdByTxInsert(tx)
	for _, v := range video.Path {
		_, err = tx.Exec(constants.ADMIN_VIDEO_PATH_INSERT_SQL, lastId, v.Path, v.Number, v.Info, v.CreateTime)
		if err != nil {
			logger.Error(fmt.Sprintf("Tx exec fail,err:%v", err))
			tx.Rollback()
			return false
		}
	}
	logger.Info(lastId)
	err = tx.Commit()
	if err != nil {
		logger.Error(fmt.Sprintf("Tx commit fail,err:%v", err))
		return false
	}
	return true
}
