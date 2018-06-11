package dao

import "cloud/httpServer/bean"

type RealFileDaoIntf interface {
	List(req bean.RealFilePage) []bean.RealFile //分页列表
	Save(path string) bool                      //保存
	UpdateRelative(id, videoId int64) bool      //更新关联关系
	Del(ids string) bool                        //批量删除
	DelByPath(path string) bool                 //根据路径删除
}
