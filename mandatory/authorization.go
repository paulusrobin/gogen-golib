package mandatory

type Authorization struct {
	authorization string
	token         string
	apiKey        string
	serviceID     string
	serviceSecret string
}

func (a Authorization) Authorization() string {
	return a.authorization
}

func (a Authorization) Token() string {
	return a.token
}

func (a Authorization) ApiKey() string {
	return a.apiKey
}

func (a Authorization) ServiceID() string {
	return a.serviceID
}

func (a Authorization) ServiceSecret() string {
	return a.serviceSecret
}
