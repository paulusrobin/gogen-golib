package circuitbreaker

type withMaxRequests uint32

// Apply implements Options interface
func (w withMaxRequests) Apply(client *circuitBreakerClient) {
	client.cbConfig.MaxRequests = uint32(w)
}

// WithMaxRequests function to override circuit breaker max requests
func WithMaxRequests(maxRequests uint32) Options {
	return withMaxRequests(maxRequests)
}
