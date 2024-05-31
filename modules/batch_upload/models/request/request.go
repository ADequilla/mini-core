package request

import "time"

// Person represents the structure of the data
type Clients struct {
	DOB               time.Time
	CID               string
	Mobile            string
	FirstName         string
	LastName          string
	MiddleName        string
	MaidenFName       string
	MaidenLName       string
	MaidenMName       string
	BirthPlace        string
	Sex               string
	CivilStatus       string
	MemberMaidenFName string
	MemberMaidenLName string
	MemberMaidenMName string
	Email             string
	InstitutionCode   string
	UnitCode          string
	CenterCode        string
}

// Person represents the structure of the data
type Accounts struct {
	AccountNumber      string
	Account            string
	AccountType        int
	AccountDescription string
	DateOpen           time.Time
	StatusDescription  string
	IiID               string
	Status             int
	Title              int
	Classification     int
	SubClassification  int
	DateEntry          time.Time
	DateRecognized     time.Time
	DateResigned       time.Time
	InstitutionCode    int
	BranchCode         int
	UnitCode           int
	CenterCode         int
	UUID               int
	SetCid             int
	AreaCode           int
	Area               string
	Balance            float64
	Withdrawable       float64
	LeagerBalance      float64
}
