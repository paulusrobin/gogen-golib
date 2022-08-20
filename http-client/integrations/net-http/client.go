package nethttp

import (
	httpclient "github.com/paulusrobin/gogen-golib/http-client/interface"
	"net/http"
)

type client struct {
	doer *http.Client
}

// NewClient function to instantiate http client.
func NewClient(c *http.Client) httpclient.ClientDoer {
	return &client{
		doer: c,
	}
}

// Do function to implement or run http request.
func (c client) Do(req *http.Request) (*http.Response, error) {
	return c.doer.Do(req)
}
