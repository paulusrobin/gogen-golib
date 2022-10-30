package circuitbreaker

import "github.com/sony/gobreaker"

// ReadyToTripFunc on ready to trip
type ReadyToTripFunc func(counts gobreaker.Counts) bool

// Apply implements Options interface
func (o ReadyToTripFunc) Apply(client *circuitBreakerClient) {
	client.cbConfig.ReadyToTrip = o
}

// WithReadyToTripFunc function to run hook on state change
func WithReadyToTripFunc(hookFn func(counts gobreaker.Counts) bool) Options {
	return ReadyToTripFunc(hookFn)
}
