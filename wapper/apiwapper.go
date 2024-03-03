package wapper

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Error  error       `json:"error"`
}

// type apiHandlerFunc func(c *gin.Context) *Response

func SuccessWithData(data interface{}) *Response {
	return &Response{
		Error:  nil,
		Data:   data,
		Status: "Success",
	}
}

func FailWithErr(err error) *Response {
	return &Response{
		Error:  err,
		Data:   nil,
		Status: "Fail",
	}
}

func Success() *Response {
	return &Response{
		Error:  nil,
		Data:   nil,
		Status: "Success",
	}
}

// func SuccessRes() gin.HandlerFunc{
//
// return func(ctx *gin.Context) {
//
//   }
