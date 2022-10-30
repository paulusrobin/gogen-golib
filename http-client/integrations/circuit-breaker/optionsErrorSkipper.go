package circuitbreaker

// CounterErrorSkipperFunc function skipper circuit breaker counter return true to skip
type CounterErrorSkipperFunc func(err error) bool

// Apply implements Options interface
func (c CounterErrorSkipperFunc) Apply(client *circuitBreakerClient) {
	client.cbConfig.IsSuccessful = c
}

// WithCounterErrorSkipper function to bring counter error skipper to circuit breaker
func WithCounterErrorSkipper(skipperFunc func(err error) bool) Options {
	return CounterErrorSkipperFunc(skipperFunc)
}
