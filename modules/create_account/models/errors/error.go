package errors

type ErrorModel struct {
	Message   string `json:"message"`
	IsSuccess bool   `json:"isSuccess"`
	Error     error  `json:"error"`
}

type ResponseModel struct {
	RetCode string      `json:"retCode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type LoginResponseModel struct {
	RetCode string      `json:"retCode"`
	Message string      `json:"message"`
	Token   string      `json:"token"`
	Data    interface{} `json:"data"`
}
