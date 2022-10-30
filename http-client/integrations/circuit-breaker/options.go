package circuitbreaker

// Options interface to extend circuit breaker client
type Options interface {
	Apply(*circuitBreakerClient)
}
