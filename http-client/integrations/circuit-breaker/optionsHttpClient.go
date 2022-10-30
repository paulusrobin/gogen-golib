package circuitbreaker

import (
	"net/http"
)

type withHTTPClient struct{ *http.Client }

// Apply implements Options interface
func (w withHTTPClient) Apply(client *circuitBreakerClient) {
	client.doer = w.Client
}

// WithHTTPClient function to override http client doer
func WithHTTPClient(client *http.Client) Options {
	return withHTTPClient{client}
}
