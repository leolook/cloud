package protocol

type ConnectInfoReq struct {
	Addr string `json:"addr" empty:"no"`
	User string `json:"user" empty:"no"`
	Pwd  string `json:"pwd" empty:"no"`
	Db   string `json:"db" empty:"no"`
}

type ConnectInfoRsp struct {
	Url  string   `json:"url"`
	List []string `json:"list"`
}
