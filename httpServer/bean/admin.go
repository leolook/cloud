package bean

type AdminLogin struct {
	UserId string `json:"userId"`
	Token  string `json:"token"`
}

type AdminSession struct {
	UserName   string `json:"userName"`
	CreateTime int64  `json:"createTime"`
	Token      string `json:"token"`
}
