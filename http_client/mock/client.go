package httpclientmock

import (
	"context"
	httpclient "github.com/paulusrobin/gogen-lib/http_client"
	"net/http"
)

type (
	response struct {
		response *http.Response
		err      error
	}
	client struct {
		response *http.Response
		err      error
	}
)

func (c client) Do(ctx context.Context, request *http.Request) (*http.Response, error) {
	return c.response, c.err
}

func NewMockHTTPClient(response *http.Response, err error) httpclient.Client {
	return client{response: response, err: err}
}
