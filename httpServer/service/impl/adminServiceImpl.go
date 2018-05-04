package impl

import (
	"cloud/common/db"
	"cloud/common/logger"
	"cloud/common/util"
	"cloud/constants"
	"cloud/httpServer/bean"
	"fmt"
	"time"
)

type AdminServiceImpl struct{}

//登录验证
func (impl *AdminServiceImpl) CheckLogin(userName, password string) interface{} {
	logger.Info(fmt.Sprintf("userName:%s,password:%s", userName, password))
	engine := db.GetEngine()
	var userId string
	isExist, err := engine.SQL(constants.ADMIN_LOGIN_CHECK_SQL, userName, password).Get(&userId)
	if !isExist || err != nil {
		logger.Warn(fmt.Sprintf("Find admin fail,isExist:%verr:%v", isExist, err))
		return nil
	}
	token := impl.CreateSession(userId, userName, password)
	if token == constants.STR_IS_EMPTY {
		return nil
	}
	data := bean.AdminLogin{
		UserId: userId,
		Token:  token,
	}
	return data
}

//创建会话
func (impl *AdminServiceImpl) CreateSession(userId, userName, password string) string {
	client := db.GetClient()
	key := constants.ADMIN_LIST_SESSION_HASH_KEY
	sid := client.HGet(key, userId).Val()
	pipe := client.Pipeline()
	if sid != constants.STR_IS_EMPTY { //会话已存在,清理掉之前会话
		pipe.HDel(key, userId)
		key = fmt.Sprintf(constants.ADMIN_SESSION_HASH_KEY, sid)
		pipe.Del(key)
	}
	//创建会话
	sid = util.Sha1Md5(fmt.Sprintf("%v%v", userId, time.Now().UnixNano()))
	token := util.Sha1Md5(fmt.Sprintf("%v", time.Now().UnixNano()))
	logger.Info(fmt.Sprintf("sid:%s,token:%s", sid, token))

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
