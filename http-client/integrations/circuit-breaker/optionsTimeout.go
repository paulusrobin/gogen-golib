package circuitbreaker

import "time"

type withTimeout time.Duration

// Apply implements Options interface
func (w withTimeout) Apply(client *circuitBreakerClient) {
	client.doer.Timeout = time.Duration(w)
	client.cbConfig.Timeout = time.Duration(w)
}

// WithTimeout function to override http client timeout
func WithTimeout(timeout time.Duration) Options {
	return withTimeout(timeout)
}
