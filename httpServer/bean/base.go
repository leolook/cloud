package bean

type BasePageReq struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}

func (the BasePageReq) GetPageNo() int {
	if the.PageNo <= 0 {
		the.PageNo = 1
	}
	return the.PageNo
}

func (the BasePageReq) GetPageSize() int {
	if the.PageSize <= 0 {
		the.PageSize = 10
	}
	return the.PageSize
}

func (the BasePageReq) GetOffset() int {
	return (the.GetPageNo() - 1) * the.GetPageSize()
}

type BasePageRes struct {
	Total int64       `json:"total"`
	Rows  interface{} `json:"rows"`
}
