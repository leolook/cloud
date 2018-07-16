package protocol

type TestReq struct {
	Id   int64    `json:"id" form:"id" empty:"no"`
	Name string   `json:"name" form:"name" empty:"no"`
	Arr  []string `json:"arr" form:"arr" empty:"no"`
	Data struct {
		Id int64 `json:"id" empty:"no"`
	} `json:"data" form:"data" empty:"no"`
}

type TestRsp struct {
	Name string `json:"name,omitempty"`
	Id   int64  `json:"id"`
}
