package service

import "cloud/httpServer/bean"

type RealFileServiceIntf interface {
	List(req bean.RealFilePage) []bean.RealFile //分页列表
	Del(ids string) bool                        //批量删除
}
