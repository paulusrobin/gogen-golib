package circuitbreaker

import (
	"fmt"
	httpclient "github.com/paulusrobin/gogen-golib/http-client/interface"
	"github.com/sony/gobreaker"
	"net/http"
	"time"
)

const (
	defaultMaxRequests         = 50
	defaultInterval            = 10 * time.Second
	defaultTimeout             = 30 * time.Second
	defaultConsecutiveFailures = 10
	defaultRequestsThreshold   = 100
	defaultFailureRatio        = 0.1
)

type circuitBreakerClient struct {
	doer     *http.Client
	cb       *gobreaker.CircuitBreaker
	cbConfig gobreaker.Settings
}

// Do function to implement or run http request.
func (c *circuitBreakerClient) Do(req *http.Request) (*http.Response, error) {
	response, err := c.cb.Execute(func() (interface{}, error) {
		return c.doer.Do(req)
	})
	if err != nil {
		return nil, err
	}

	httpResponse, ok := response.(*http.Response)
	if !ok {
		return nil, fmt.Errorf("invalid http response object")
	}

	return httpResponse, nil
}

// DefaultReadyToTrip default function for circuit breaker
func DefaultReadyToTrip(
	requestThreshold uint32,
	consecutiveFailures uint32,
	failureRate float64,
) func(counts gobreaker.Counts) bool {
	return func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return (counts.Requests >= requestThreshold && failureRatio >= failureRate) ||
			(counts.ConsecutiveFailures >= consecutiveFailures)
	}
}

// NewCircuitBreakerClient function to instantiate http client.
func NewCircuitBreakerClient(options ...Options) httpclient.ClientDoer {
	client := circuitBreakerClient{
		doer: &http.Client{},
		cbConfig: gobreaker.Settings{
			Name:          "",
			Interval:      defaultInterval,
			Timeout:       defaultTimeout,
			MaxRequests:   defaultMaxRequests,
			ReadyToTrip:   DefaultReadyToTrip(defaultRequestsThreshold, defaultConsecutiveFailures, defaultFailureRatio),
			OnStateChange: func(name string, from, to gobreaker.State) {},
		},
	}

	for _, option := range options {
		option.Apply(&client)
	}

	client.cb = gobreaker.NewCircuitBreaker(client.cbConfig)

	return &client
}
