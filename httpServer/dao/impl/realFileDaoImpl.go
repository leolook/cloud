package impl

import (
	"cloud/common/db"
	logger "github.com/alecthomas/log4go"
	"cloud/constants"
	"cloud/httpServer/bean"
	"fmt"
	"strings"
)

type RealFilelImpl struct {
	BaseDaoImpl
	UserId string //用户id
}

//分页列表
func (r *RealFilelImpl) List(req bean.RealFilePage) []bean.RealFile {
	joinSql := constants.STR_IS_EMPTY
	arg := make([]interface{}, 0)
	if req.FileType > constants.INT_IS_ZERO {
		joinSql += " and file_type=?"
		arg = append(arg, req.FileType)
	}
	if req.IsRelative {
		joinSql += " and video_id>0"
	} else {
		joinSql += " and video_id=0"
	}
	sql := fmt.Sprintf(constants.ADMIN_REAL_FILE_LIST, joinSql)
	arg = append(arg, req.GetOffset(), req.GetPageSize())
	realFiles := make([]bean.RealFile, 0, 10)
	eng := db.GetEngine()
	err := eng.SQL(sql, arg).Find(&realFiles)
	if err != nil {
		logger.Error(fmt.Sprintf("find real file fail,[sql=%s] [arg=%v] [err=%v]", sql, arg, err))
		return nil
	}
	return realFiles
}

//保存
func (r *RealFilelImpl) Save(path string) bool {
	fileType := constants.OTHER_FILE
	if strings.Index(path, ".jpg") != -1 || strings.Index(path, ".png") != -1 {
		fileType = constants.IMG_FILE
	}
	if strings.Index(path, ".mp4") != -1 || strings.Index(path, "wmv") != -1 {
		fileType = constants.VIDEO_FILE
	}
	adInfo := r.GetAdminSession(r.UserId)
	return r.Execute(constants.ADMIN_REAL_FILE_INSERT, path, adInfo.UserName, fileType)
}

//更新关联关系
func (r *RealFilelImpl) UpdateRelative(id, videoId int64) bool {
	return r.Execute(constants.ADMIN_REAL_FILE_UPDATE, videoId, id)
}

//批量删除
func (r *RealFilelImpl) Del(ids string) bool {
	return r.Execute(constants.ADMIN_REAL_FILE_DEL, ids)
}

//根据路径删除
func (r *RealFilelImpl) DelByPath(path string) bool {
	return r.Execute(constants.ADMIN_REAL_FILE_DEL_BY_PATH, path)
}
