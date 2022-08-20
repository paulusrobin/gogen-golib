package httpclientmock

import (
	httpclient "github.com/paulusrobin/gogen-golib/http-client/interface"
	"net/http"
)

type (
	client struct {
		response *http.Response
		err      error
	}
)

// Do function implements httpclient.ClientDoer
func (c client) Do(request *http.Request) (*http.Response, error) {
	return c.response, c.err
}

// NewMockHTTPClient function to initialize httpclient.mock
func NewMockHTTPClient(response *http.Response, err error) httpclient.ClientDoer {
	return client{response: response, err: err}
}
