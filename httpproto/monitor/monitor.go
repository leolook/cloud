package monitor

type RegReq struct {
	Name         string `json:"name"`
	Addr         string `json:"addr"`
	PhysicalAddr string `json:"physicalAddr"`
}

type RefRsp struct {
	Name string
	Code int64
	Msg  string
}
