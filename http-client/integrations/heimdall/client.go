package heimdall

import (
	"github.com/gojek/heimdall/v7"
	heimdallhttpclient "github.com/gojek/heimdall/v7/httpclient"
	heimdallhystrixclient "github.com/gojek/heimdall/v7/hystrix"
	httpclient "github.com/paulusrobin/gogen-golib/http-client/interface"
	"net/http"
)

type client struct {
	c heimdall.Client
}

// Do function implements httpclient.ClientDoer
func (c client) Do(request *http.Request) (*http.Response, error) {
	return c.c.Do(request)
}

// NewHttpClient function to initialize circuit heimdall http.
func NewHttpClient(options ...heimdallhttpclient.Option) httpclient.ClientDoer {
	return &client{heimdallhttpclient.NewClient(options...)}
}

// NewHystrixClient function to initialize circuit heimdall hystrix.
func NewHystrixClient(options ...heimdallhystrixclient.Option) httpclient.ClientDoer {
	return &client{heimdallhystrixclient.NewClient(options...)}
}
