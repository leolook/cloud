package bean

type AdminLogin struct {
	UserId   string `json:"userId"`
	Token    string `json:"token"`
	FullName string `json:"fullName"`
	ImgPath  string `json:"imgPath"`
}

type AdminUser struct {
	FullName      string `json:"fullName" xorm:"full_name"`
	ImgPath       string `json:"imgPath" xorm:"img_path"`
	Id            int64  `json:"id"`
	IsLogin       string `json:"isLogin" xorm:"is_login"`
	LastLoginIp   string `json:"lastLoginIp" xorm:"last_login_ip"`
	LastLoginTime int64  `json:"lastLoginTime" xorm:"last_login_time"`
	CreateTime    int64  `json:"createTime" xorm:"create_time"`
}

type AdminSession struct {
	UserName   string `json:"userName"`
	CreateTime int64  `json:"createTime"`
	Token      string `json:"token"`
}
