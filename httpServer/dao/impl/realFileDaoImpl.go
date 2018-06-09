package impl

import (
	"cloud/constants"
	"cloud/httpServer/bean"
)

type RealFilelImpl struct {
	BaseDaoImpl
	UserId string //用户id
}

//分页列表
func (r *RealFilelImpl) List(req bean.RealFilePage) []bean.RealFile {
	return nil
}

//保存
func (r *RealFilelImpl) Save(path string, fileType int) bool {
	adInfo := r.GetAdminSession(r.UserId)
	return r.SaveOrUpdate(constants.ADMIN_REAL_FILE_INSERT, path, adInfo.UserName, fileType)
}

//更新关联关系
func (r *RealFilelImpl) UpdateRelative(id, videoId int64) bool {
	return r.SaveOrUpdate(constants.ADMIN_REAL_FILE_UPDATE, videoId, id)
}

//批量删除
func (r *RealFilelImpl) Del(id []int64) bool {
	return false
}
