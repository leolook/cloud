package dao

import "cloud/httpServer/bean"

type ClassifyDaoIntf interface {
	All() []bean.ClassifyBean //获取所有分类
}
