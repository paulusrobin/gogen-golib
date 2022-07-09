package httpclientmock

import (
	"context"
	httpclient "github.com/paulusrobin/gogen-golib/http-client"
	"net/http"
)

type (
	client struct {
		response *http.Response
		err      error
	}
)

// Do function implements httpclient.Client.
func (c client) Do(ctx context.Context, request *http.Request) (*http.Response, error) {
	return c.response, c.err
}

// NewMockHTTPClient function to initialize httpclient.mock
func NewMockHTTPClient(response *http.Response, err error) httpclient.Client {
	return client{response: response, err: err}
}
