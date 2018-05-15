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
		logger.Warn(fmt.Sprintf("Find admin fail,isExist:%v,err:%v", isExist, err))
		return nil
	}
	token := impl.createSession(userId, userName, password)
	if token == constants.STR_IS_EMPTY {
		return nil
	}
	data := bean.AdminLogin{
		UserId: userId,
		Token:  token,
	}
	return data
}

//会话验证
func (impl *AdminServiceImpl) CheckSession(userId, token string) bool {
	client := db.GetClient()
	key := constants.ADMIN_LIST_SESSION_HASH_KEY
	sid := client.HGet(key, userId).Val()
	if sid == constants.STR_IS_EMPTY {
		return false
	}
	key = fmt.Sprintf(constants.ADMIN_SESSION_HASH_KEY, sid)
	state := client.TTL(key).Val()
	if state == -2*time.Second { //不存在
		client.HDel(constants.ADMIN_LIST_SESSION_HASH_KEY, userId)
		return false
	}
	serverToken := client.HGet(key, constants.ADMIN_SESSION_TOKEN).Val()
	if serverToken != token {
		return false
	}
	client.Expire(key, constants.ADMIN_SESSION_EXPIRE_TIME*time.Second)
	return true
}

//退出登录
func (impl *AdminServiceImpl) LoginOut(userId string) bool {
	client := db.GetClient()
	key := constants.ADMIN_LIST_SESSION_HASH_KEY
	sid := client.HGet(key, userId).Val()
	if sid != constants.STR_IS_EMPTY { //会话已存在,清理掉之前会话
		return impl.clearSession(userId, sid)
	}
	return true
}

//创建会话
func (impl *AdminServiceImpl) createSession(userId, userName, password string) string {
	client := db.GetClient()
	key := constants.ADMIN_LIST_SESSION_HASH_KEY
	sid := client.HGet(key, userId).Val()
	if sid != constants.STR_IS_EMPTY { //会话已存在,清理掉之前会话
		impl.clearSession(userId, sid)
	}
	pipe := client.Pipeline()
	//创建会话
	sid = util.Sha1Md5(fmt.Sprintf("%v%v", userId, time.Now().UnixNano()))
	token := util.Sha1Md5(fmt.Sprintf("%v", time.Now().UnixNano()))
	logger.Info(fmt.Sprintf("sid:%s,token:%s", sid, token))

	key = fmt.Sprintf(constants.ADMIN_SESSION_HASH_KEY, sid)
	pipe.HSet(key, constants.ADMIN_SESSION_USER_NAME, userName)
	pipe.HSet(key, constants.ADMIN_SESSION_PASSWORD, password)
	pipe.HSet(key, constants.ADMIN_SESSION_CREATE_TIME, time.Now().Unix())
	pipe.HSet(key, constants.ADMIN_SESSION_TOKEN, token)
	pipe.Expire(key, constants.ADMIN_SESSION_EXPIRE_TIME*time.Second)
	key = constants.ADMIN_LIST_SESSION_HASH_KEY
	pipe.HSet(key, userId, sid)
	_, err := pipe.Exec()
	if err != nil {
		logger.Error(fmt.Sprintf("Create admin session fail,err:%v", err))
		return constants.STR_IS_EMPTY
	}
	return token
}

//清除会话
func (impl *AdminServiceImpl) clearSession(userId, sid string) bool {
	client := db.GetClient()
	key := constants.ADMIN_LIST_SESSION_HASH_KEY
	pipe := client.Pipeline()
	pipe.HDel(key, userId)
	key = fmt.Sprintf(constants.ADMIN_SESSION_HASH_KEY, sid)
	pipe.Del(key)
	result, err := pipe.Exec()
	if err != nil {
		logger.Error(fmt.Sprintf("Clear session err:%v,result:%v", err, result))
		return false
	}
	return true
}
