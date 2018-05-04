package service

type AdminServiceIntf interface {
	CheckLogin(userName, password string) interface{} //登录验证
	CheckSession(userId, token string) bool           //会话验证
	LoginOut(userId string) bool                      //推出登录
}
