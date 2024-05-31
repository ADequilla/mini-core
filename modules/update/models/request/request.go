package request

import "time"

type UpdateClientModel struct {
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
	InstiCode         string    `json:"insti_code"`
	UnitCode          string    `json:"unit_code"`
	CenterCode        string    `json:"center_code"`
}

type UpdateAccountsModel struct {
	SetAccountNumber     string    `json:"set_account_number"`
	SetAcc               string    `json:"set_acc"`
	SetAcctType          int       `json:"set_acct_type"`
	SetAccDesc           string    `json:"set_acc_desc"`
	SetDOpen             time.Time `json:"set_d_open"`
	SetStatusDesc        string    `json:"set_status_desc"`
	SetIIID              int       `json:"set_iiid"`
	SetStatus            int       `json:"set_status"`
	SetTitle             int64     `json:"set_title"`
	SetClassification    int64     `json:"set_classification"`
	SetSubClassification int64     `json:"set_sub_classification"`
	SetDoEntry           string    `json:"set_do_entry"`
	SetDoRecognized      string    `json:"set_do_recognized"`
	SetDoResigned        string    `json:"set_do_resigned"`
	SetInstiCode         int64     `json:"set_insti_code"`
	SetBranchCode        int64     `json:"set_branch_code"`
	SetUnitCode          int64     `json:"set_unit_code"`
	SetCenterCode        int64     `json:"set_center_code"`
	SetUUID              int       `json:"set_uuid"`
	SetCID               int       `json:"set_cid"`
	SetAreaCode          int       `json:"set_areacode"`
	SetArea              string    `json:"set_area"`
	SetBalance           float64   `json:"set_balance"`
	SetWithdrawable      float64   `json:"set_withdrawable"`
	SetLedgerBalance     float64   `json:"set_ledger_balance"`
}
