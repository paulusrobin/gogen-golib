package nethttp

import (
	httpclient "github.com/paulusrobin/gogen-lib/http_client"
	"net/http"
	"time"
)

type (
	Builder interface {
		WithTimeout(duration time.Duration) Builder
		Build() httpclient.Client
	}
)

// NewBuilder function to make http net builder.
func NewBuilder() Builder {
	return &builder{
		client: &http.Client{},
	}
}

type builder struct {
	client *http.Client
}

// WithTimeout function to set http net request timeout.
func (b *builder) WithTimeout(duration time.Duration) Builder {
	b.client.Timeout = duration
	return b
}

// Build function to instantiate http net http client.
func (b builder) Build() httpclient.Client {
	return NewClient(b.client)
}
