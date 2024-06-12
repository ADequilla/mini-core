package response

import "time"

type ResponseModel struct {
	RetCode string      `json:"retCode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type GetClientModel struct {
	Birthday          time.Time `json:"birthday"`
	Cid               string    `json:"cid"`
	Mobile            string    `json:"mobile"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	MiddleName        string    `json:"middle_name"`
	MaidenFName       string    `json:"maiden_f_name"`
	MaidenLName       string    `json:"maiden_l_name"`
	MaidenMName       string    `json:"maiden_m_name"`
	BirthPlace        string    `json:"birth_place"`
	Sex               string    `json:"sex"`
	CivilStatus       string    `json:"civil_status"`
	MemberMaidenFName string    `json:"member_maiden_f_name"`
	MemberMaidenLName string    `json:"member_maiden_l_name"`
	MemberMaidenMName string    `json:"member_maiden_m_name"`
	Email             string    `json:"email"`
	InstiCode         string    `json:"institution"`
	UnitCode          string    `json:"unit_code"`
	CenterCode        string    `json:"center_code"`
	CreatedDate       time.Time `json:"created_date"`
}

type GetAccountsModel struct {
	Account_number    string    `json:"account_number"`
	Acc               string    `json:"acc"`
	Accttype          int       `json:"accttype"`
	Accdesc           string    `json:"accdesc"`
	DOpen             time.Time `json:"dopen"`
	Statusdesc        string    `json:"statusdesc"`
	Iiid              int       `json:"iiid"`
	Status            int       `json:"status"`
	Title             int64     `json:"title"`
	Classification    int64     `json:"classification"`
	SubClassification int64     `json:"sub_classification"`
	Do_entry          string    `json:"do_entry"`
	Do_recognized     string    `json:"do_recognized"`
	Do_resigned       string    `json:"do_resigned"`
	Insti_code        int64     `json:"insti_code"`
	Branch_code       int64     `json:"branch_code"`
	Unit_code         int64     `json:"unit_code"`
	Center_code       int64     `json:"center_code"`
	Uuid              int       `json:"uuid"`
	Cid               int       `json:"cid"`
	Areacode          int       `json:"areacode"`
	Area              string    `json:"area"`
	Balance           float64   `json:"balance"`
	Withdrawable      float64   `json:"withdrawable"`
	Ledger_Balance    float64   `json:"ledger_balance"`
}
