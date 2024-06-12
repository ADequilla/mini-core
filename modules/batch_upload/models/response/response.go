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

type DuplicateClient struct {
	CID        string `json:"cid"`
	Mobile     string `json:"mobile"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
}
