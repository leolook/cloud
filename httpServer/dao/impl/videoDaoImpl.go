package impl

import (
	"cloud/common/db"
	"cloud/common/logger"
	"cloud/constants"
	"cloud/httpServer/bean"
	"fmt"
	"github.com/go-xorm/xorm"
)

type VideoDaoImpl struct {
	Tx *xorm.Session
	BaseDaoImpl
}

//根据名称判断video是否已经存在
func (impl VideoDaoImpl) IsExistByName(name string) int64 {
	engine := db.GetEngine()
	var id int64
	isExist, err := engine.SQL(constants.ADMIN_VIDEO_IS_EXIST_BY_NAME_SQL, name).Get(&id)
	if err != nil {
		logger.Warn(fmt.Sprintf("Not found this video:%s,err:%v", name, err))
		return 0
	}
	if !isExist || id <= 0 {
		return 0
	}
	return id
}

//插入video
func (impl VideoDaoImpl) Insert(useId string, video *bean.VideoBean) bool {
	tx := db.GetEngine().NewSession()
	impl.Tx = tx
	if !impl.insertVideo(useId, video) {
		return false
	}
	lastId := impl.GetLastIdByTxInsert(impl.Tx)
	if lastId == 0 {
		return false
	}
	video.Id = lastId
	if !impl.insertVideoPath(video) {
		return false
	}
	err := tx.Commit()
	if err != nil {
		logger.Error(fmt.Sprintf("Tx commit fail,err:%v", err))
		return false
	}
	return true
}

//修改video
func (impl VideoDaoImpl) Update(useId string, video *bean.VideoBean) bool {
	tx := db.GetEngine().NewSession()
	impl.Tx = tx
	if !impl.updateVideo(useId, video) {
		return false
	}
	if !impl.delVideoPath(video.Id) {
		return false
	}
	if !impl.insertVideoPath(video) {
		return false
	}
	err := tx.Commit()
	if err != nil {
		logger.Error(fmt.Sprintf("Tx commit fail,err:%v", err))
		return false
	}
	return true
}

//插入video信息
func (impl VideoDaoImpl) insertVideo(useId string, video *bean.VideoBean) bool {
	session := impl.GetAdminSession(useId)
	if session == nil {
		return false
	}
	_, err := impl.Tx.Exec(constants.ADMIN_VIDEO_FILE_INSERT_SQL, video.Name, video.Info, video.Classify, session.UserName)
	if err != nil {
		logger.Error(fmt.Sprintf("Tx exec fail,err:%v", err))
		impl.Tx.Rollback()
		return false
	}
	return true
}

//更新video信息
func (impl VideoDaoImpl) updateVideo(useId string, video *bean.VideoBean) bool {
	session := impl.GetAdminSession(useId)
	if session == nil {
		return false
	}
	_, err := impl.Tx.Exec(constants.ADMIN_VIDEO_FILE_UPDATE_SQL, video.Name, video.Info, video.Classify, session.UserName, video.Id)
	if err != nil {
		logger.Error(fmt.Sprintf("Tx exec fail,err:%v", err))
		impl.Tx.Rollback()
		return false
	}
	return true
}

//删除video视频文件路径
func (impl VideoDaoImpl) delVideoPath(id int64) bool {
	_, err := impl.Tx.Exec(constants.ADMIN_VIDEO_PATH_DELETE_SQL, id)
	if err != nil {
		logger.Error(fmt.Sprintf("Tx exec fail,err:%v", err))
		impl.Tx.Rollback()
		return false
	}
	return true
}

//video视频文件路径插入
func (impl VideoDaoImpl) insertVideoPath(video *bean.VideoBean) bool {
	for _, v := range video.Path {
		_, err := impl.Tx.Exec(constants.ADMIN_VIDEO_PATH_INSERT_SQL, video.Id, v.Path, v.Number, v.Info, v.CreateTime)
		if err != nil {
			logger.Error(fmt.Sprintf("Tx exec fail,err:%v", err))
			impl.Tx.Rollback()
			return false
		}
	}
	return true
}
