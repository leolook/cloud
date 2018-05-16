package bean

type VideoBean struct {
	Id       int64           `json:"id"`
	Name     string          `json:"name"`
	Info     string          `json:"info"`
	Classify int64           `json:"classify"`
	Path     []VideoPathBean `json:"path"`
}

type VideoPathBean struct {
	Path       string `json:"path"`
	Number     int    `json:"number"`
	Info       string `json:"info"`
	CreateTime int64  `json:"createTime"`
}
