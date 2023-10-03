package response

type ResponseResult struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseOutput struct {
	Code    int32          `json:"code"`
	Message string         `json:"message"`
	Result  ResponseResult `json:"result"`
}
