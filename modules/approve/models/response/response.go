package response

type ResponseModel struct {
	RetCode string      `json:"retCode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ApproveClientModel struct {
	Message string `json:"message"`
}
