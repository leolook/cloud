package dao

import (
	"cloud/httpServer/bean"
)

type BaseDaoIntf interface {
	GetAdminSession(userId string) *bean.AdminSession
	SaveOrUpdate(sqlStr string, args ...interface{}) bool
}
