package createaccount

type RequestBodyStruct struct {
	UserLogin    string `json:"userLogin,omitempty"`
	UserPassword string `json:"userPasswd"`
	UserName     string `json:"userName"`
	UserEmail    string `json:"userEmail"`
	UserPhone    string `json:"userPhone"`
	InstiCode    string `json:"instiCode"`
	LastName     string `json:"lastName"`
	GivenName    string `json:"givenName"`
	MiddleName   string `json:"middleName"`
}
