package service

import "cloud/httpServer/bean"

type ClassifyServiceIntf interface {
	All() []bean.ClassifyBean //获取所有分类
}
