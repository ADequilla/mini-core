package response

type ResponseModel struct {
	RetCode string      `json:"retCode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UpdatetClient struct {
	Message string `json:"message"`
}

type UpdateAccount struct {
	Message string `json:"message"`
}
