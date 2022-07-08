package nethttp

import (
	"context"
	httpclient "github.com/paulusrobin/gogen-lib/http-client/interface"
	"net/http"
)

type client struct {
	doer *http.Client
}

// NewClient function to instantiate http client.
func NewClient(c *http.Client) httpclient.Client {
	return &client{
		doer: c,
	}
}

// Do function to implement or run http request.
func (c client) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	return c.doer.Do(req)
}
