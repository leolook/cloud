package service

type AdminServiceIntf interface {
	CheckLogin(userName, password string) interface{} //登录验证
}
