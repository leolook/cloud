package response

type Response struct {
	Code int       `json:"code"`
	Data interface{} `json:"data"`
}

func GetResponse(code int, data interface{}) *Response {
	res := new(Response)
	res.Code, res.Data = code, data
	return res
}

func GetSuccessResponse(data interface{}) *Response {
	return GetResponse(0, data)
}
