package lib

type ResponsesSuccess struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type ResponsesFail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//type Responses struct{}

func Success(data interface{}) ResponsesSuccess {
	var responsesSuccess ResponsesSuccess
	responsesSuccess.Code = 0
	responsesSuccess.Data = data
	return responsesSuccess
}

func Fail(code int, message string) ResponsesFail {
	var responsesFail ResponsesFail
	responsesFail.Code = code
	responsesFail.Message = message
	return responsesFail
}
