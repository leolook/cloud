package impl

import (
	"cloud/common/db"
	"cloud/common/logger"
	"cloud/common/util"
	"cloud/constants"
	"fmt"
	"time"
)

type AdminServiceImpl struct{}

//登录验证
func (impl *AdminServiceImpl) CheckLogin(userName, password string) interface{} {
	logger.Info(fmt.Sprintf("userName:%s,password:%s", userName, password))
	impl.CreateSession("1", userName, password)
	return nil
}

//创建会话
func (impl *AdminServiceImpl) CreateSession(userId, userName, password string) string {
	client := db.GetClient()
	key := constants.ADMIN_LIST_SESSION_HASH_KEY
	sid := client.HGet(key, userId).Val()
	if sid != constants.STR_IS_EMPTY { //会话已存在,清理掉之前会话
		err := client.HDel(key, userId).Err()
		if err != nil {
			logger.Error(fmt.Sprintf("Remove admin session fail,err:%v,key:%s,userId:%s", err, key, userId))
			return constants.STR_IS_EMPTY
		}
	}
	//创建会话
	sid = util.Sha1Md5(userId)
	token := util.Sha1Md5(fmt.Sprintf("%v", time.Now().UnixNano()))
	logger.Info(fmt.Sprintf("sid:%s,token:%s", sid, token))

	pipe := client.Pipeline()
	key = fmt.Sprintf(constants.ADMIN_SESSION_HASH_KEY, sid)
	pipe.HSet(key, constants.ADMIN_SESSION_USER_NAME, userName)
	pipe.HSet(key, constants.ADMIN_SESSION_PASSWORD, password)
	pipe.HSet(key, constants.ADMIN_SESSION_CREATE_TIME, time.Now().Unix())
	pipe.HSet(key, constants.ADMIN_SESSION_TOKEN, token)
	key = constants.ADMIN_LIST_SESSION_HASH_KEY
	pipe.HSet(key, userId, sid)
	_, err := pipe.Exec()
	if err != nil {
		logger.Error(fmt.Sprintf("Create admin session fail,err:%v", err))
		return constants.STR_IS_EMPTY
	}
	return token
}
