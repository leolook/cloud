package bean

type ClassifyBean struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	CreateTime int64  `json:"createTime" xorm:"create_time"`
}
