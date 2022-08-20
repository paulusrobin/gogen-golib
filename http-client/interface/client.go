package httpclient

import (
	"context"
	"github.com/paulusrobin/gogen-golib/encoding/json"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

type (
	client struct {
		doer ClientDoer
	}
	Client interface {
		Do(ctx context.Context, request *http.Request) (*http.Response, error)
	}
	ClientDoer interface {
		Do(request *http.Request) (*http.Response, error)
	}
)

func loggerWith(logger zerolog.Logger, parameter interface{}) zerolog.Logger {
	if parameter == nil {
		return logger
	}

	requestByte, err := json.Marshal(parameter)
	if err != nil {
		return logger
	}

	return logger.With().
		RawJSON("request", requestByte).
		Logger()
}

// Do function to implement Client
func (c client) Do(ctx context.Context, request *http.Request) (*http.Response, error) {
	logger := loggerWith(log.Ctx(ctx).With().Logger(), request)
	logger.Info().Msg("http-client pre-request")

	response, err := c.doer.Do(request)

	logger = loggerWith(logger.With().Logger(), response)
	event := logger.Info()
	if err != nil {
		event = logger.Error().Err(err)
	}
	event.Msg("http-client post-request")
	return response, err
}

// NewClient function to initialize http client
func NewClient(doer ClientDoer) Client {
	return &client{doer: doer}
}
