package bean

type VideoBean struct {
	Id       int64           `json:"id"`
	Name     string          `json:"name"`
	Describe string          `json:"describe"`
	Classify int64           `json:"classify"`
	Path     []VideoPathBean `json:"path"`
}

type VideoPathBean struct {
	Path       string `json:"path"`
	Order      int    `json:"order"`
	CreateTime int64  `json:"createTime"`
}
