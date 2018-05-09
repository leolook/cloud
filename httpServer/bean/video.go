package bean

type VideoBean struct {
	Id       int64       `json:"id"`
	Name     string      `json:"name"`
	Describe string      `json:"describe"`
	Classify interface{} `json:"classify"`
	Path     string      `json:"path"`
}

type VideoPath struct {
	Path       string `json:"path"`
	Order      int    `json:"order"`
	CreateTime int64  `json:"createTime"`
}
