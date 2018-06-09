package impl

import (
	"cloud/common/db"
	"cloud/common/logger"
	"cloud/constants"
	"cloud/httpServer/bean"
	"fmt"
	"github.com/go-xorm/xorm"
)

type BaseDaoImpl struct{}

//保存或者更新
func (impl BaseDaoImpl) SaveOrUpdate(sqlStr string, args ...interface{}) bool {
	result, err := db.GetEngine().Exec(sqlStr, args)
	if err != nil {
		logger.Error(fmt.Sprintf("SaveOrUpdate fail,[err=%v] [result=%v]", err, result))
		return false
	}
	rows, err := result.RowsAffected()
	if err != nil {
		logger.Error(fmt.Sprintf("RowsAffected fail,[err=%v] [rows=%v]", err, rows))
		return false
	}
	if rows <= 0 {
		logger.Warn(fmt.Sprintf("RowsAffected rows is zero,[rows=%v]", rows))
	}
	return true
}

//获取admin会话
func (impl BaseDaoImpl) GetAdminSession(userId string) *bean.AdminSession {
	client := db.GetClient()
	key := constants.ADMIN_LIST_SESSION_HASH_KEY
	sid := client.HGet(key, userId).Val()
	if sid == constants.STR_IS_EMPTY {
		logger.Error(fmt.Sprintf("Sid is empty,key:%s,sid:%s", key, sid))
		return nil
	}
	key = fmt.Sprintf(constants.ADMIN_SESSION_HASH_KEY, sid)
	data := client.HMGet(key, constants.ADMIN_SESSION_USER_NAME, constants.ADMIN_SESSION_CREATE_TIME, constants.ADMIN_SESSION_TOKEN).Val()
	if len(data) <= 0 {
		logger.Error(fmt.Sprintf("Get admin session fail,key:%s,data:%v", key, data))
		return nil
	}
	session := new(bean.AdminSession)
	session.UserName, _ = data[0].(string)
	session.CreateTime, _ = data[1].(int64)
	session.Token, _ = data[2].(string)
	return session
}

//事物提交: 插入数据,获取自增长生成的id
func (impl BaseDaoImpl) GetLastIdByTxInsert(tx *xorm.Session) int64 {
	var lastId int64
	isExist, err := tx.SQL(constants.GET_LAST_ID_BY_INSERT_SQL).Get(&lastId)
	if err != nil {
		logger.Error(fmt.Sprintf("Tx exec sql fail,err:%v", err))
		tx.Rollback()
		return 0
	}
	if !isExist || lastId <= 0 {
		return 0
	}
	return lastId
}

//非事物提交: 插入数据,获取自增长生成的id
func (impl BaseDaoImpl) GetLastIdByInsert(engine *xorm.Engine) int64 {
	var lastId int64
	isExist, err := engine.SQL(constants.GET_LAST_ID_BY_INSERT_SQL).Get(&lastId)
	if err != nil {
		logger.Error(fmt.Sprintf("Exec sql fail,err:%v", err))
		return 0
	}
	if !isExist || lastId <= 0 {
		return 0
	}
	return lastId
}
