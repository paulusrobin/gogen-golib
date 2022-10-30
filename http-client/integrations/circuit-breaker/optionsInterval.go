package circuitbreaker

import "time"

type withInterval time.Duration

// Apply implements Options interface
func (w withInterval) Apply(client *circuitBreakerClient) {
	client.cbConfig.Interval = time.Duration(w)
}

// WithInterval function to override circuit breaker interval
func WithInterval(interval time.Duration) Options {
	return withInterval(interval)
}
