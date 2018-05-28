package bean

type VideoBean struct {
	Id         int64           `json:"key"`
	Name       string          `json:"name"`
	Info       string          `json:"info"`
	Cover      string          `json:"cover"`
	ClassifyId int64           `json:"classifyId" xorm:"classify_id"`
	Classify   string          `json:"classify" xorm:"classify"`
	CreateTime int64           `json:"createTime" xorm:"create_time"`
	UpdateUser string          `json:"updateUser" xorm:"update_user"`
	UpdateTime int64           `json:"updateTime" xorm:"update_time"`
	Path       []VideoPathBean `json:"path"`
}

type VideoPathBean struct {
	Path       string `json:"path"`
	Number     int    `json:"number"`
	Info       string `json:"info"`
	CreateTime int64  `json:"createTime" xorm:"create_time"`
}

//数据分页请求参数
type VideoPageReq struct {
	BasePageReq
	Name string `json:"name"`
}
