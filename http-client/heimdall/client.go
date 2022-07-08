package heimdall

import (
	"context"
	"github.com/gojek/heimdall/v7"
	heimdallhttpclient "github.com/gojek/heimdall/v7/httpclient"
	heimdallhystrixclient "github.com/gojek/heimdall/v7/hystrix"
	httpclient "github.com/paulusrobin/gogen-lib/http-client"
	"net/http"
)

type client struct {
	c heimdall.Client
}

func (c client) Do(ctx context.Context, request *http.Request) (*http.Response, error) {
	return c.c.Do(request)
}

// NewHttpClient function to initialize circuit heimdall http.
func NewHttpClient(options ...heimdallhttpclient.Option) httpclient.Client {
	return &client{heimdallhttpclient.NewClient(options...)}
}

// NewHystrixClient function to initialize circuit heimdall hystrix.
func NewHystrixClient(options ...heimdallhystrixclient.Option) httpclient.Client {
	return &client{heimdallhystrixclient.NewClient(options...)}
}
