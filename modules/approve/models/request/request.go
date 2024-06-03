package request

type ApproveClients struct {
	Id_input     int    `json:"id_input"`
	Cid_input    int    `json:"cid_input"`
	Mobile_input string `json:"mobile_input"`
}
