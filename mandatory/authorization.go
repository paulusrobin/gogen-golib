package mandatory

// Authorization object of mandatory request.
type Authorization struct {
	authorization string
	token         string
	apiKey        string
	serviceID     string
	serviceSecret string
}

// Authorization getter function to get mandatory Authorization.
func (a Authorization) Authorization() string {
	return a.authorization
}

// Token getter function to get mandatory Authorization Token.
func (a Authorization) Token() string {
	return a.token
}

// ApiKey getter function to get mandatory Authorization ApiKey.
func (a Authorization) ApiKey() string {
	return a.apiKey
}

// ServiceID getter function to get mandatory Authorization ServiceID.
func (a Authorization) ServiceID() string {
	return a.serviceID
}

// ServiceSecret getter function to get mandatory Authorization ServiceSecret.
func (a Authorization) ServiceSecret() string {
	return a.serviceSecret
}
