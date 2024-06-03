package request

// Person represents the structure of the data
type Search struct {
	Search_input string `json:"search_input"`
}

type ViewClientAccount struct {
	Id_input     int    `json:"id_input"`
	Cid_input    int    `json:"cid_input"`
	Mobile_input string `json:"mobile_input"`
}
