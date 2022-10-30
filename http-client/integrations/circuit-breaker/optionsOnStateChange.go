package circuitbreaker

import "github.com/sony/gobreaker"

// OnStateChangeFunc on state change hook
type OnStateChangeFunc func(name string, from gobreaker.State, to gobreaker.State)

// Apply implements Options interface
func (o OnStateChangeFunc) Apply(client *circuitBreakerClient) {
	client.cbConfig.OnStateChange = o
}

// WithOnStateChangeFunc function to run hook on state change
func WithOnStateChangeFunc(hookFn func(name string, from gobreaker.State, to gobreaker.State)) Options {
	return OnStateChangeFunc(hookFn)
}
