package createaccount

import "time"

type ResponseBodyStruct struct {
	UserName       string    `json:"userName"`
	UserPasswd     string    `json:"userPasswd,omitempty"`
	UserEmail      string    `json:"userEmail"`
	UserPhone      string    `json:"userPhone"`
	InstiCode      string    `json:"instiCode"`
	LastName       string    `json:"lastName"`
	GivenName      string    `json:"givenName"`
	MiddleName     string    `json:"middleName"`
	UserId         int       `json:"userId,omitempty"`
	CreatedDate    string    `json:"createdDate"`
	UserStatus     string    `json:"userStatus"`
	Position       string    `json:"userPosition"`
	FailedAttempts int       `gorm:"default:0"`
	LockoutTime    time.Time `json:"lockoutTime"`
	IsLogin        string
}
