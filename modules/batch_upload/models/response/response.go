package response

type ResponseModel struct {
	RetCode string      `json:"retCode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type InsertClientModel struct {
	Message string `json:"message"`
}

type InsertAccountModel struct {
	Message string `json:"message"`
}
