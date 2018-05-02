package response

type Response struct {
	Code uint8       `json:"code"`
	Data interface{} `json:"data"`
}

func GetResponse(code uint8, data interface{}) *Response {
	res := new(Response)
	res.Code, res.Data = code, data
	return res
}

func GetSuccessResponse(data interface{}) *Response {
	return GetResponse(0, data)
}
