package sarama

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/go-kit/kit/endpoint"
)

// Handler is function to adapt go-kit to echo http framework.
func Handler(ep endpoint.Endpoint, opts ...Option) HandlerFunc {
	var settings = defaultOptions
	for _, opt := range opts {
		opt.Apply(&settings)
	}

	next := func(ctx context.Context, message *sarama.ConsumerMessage) error {
		errFunc := func(err error) error {
			if settings.errorHandler != nil {
				return settings.errorHandler(ctx, err)
			}
			return err
		}

		request, err := settings.decoder(ctx, message)
		if err != nil {
			return errFunc(err)
		}

		response, err := ep(ctx, request)
		if err != nil {
			return errFunc(err)
		}

		return settings.encoder(ctx, response)
	}
	for i := len(settings.middlewares) - 1; i >= 0; i-- {
		next = settings.middlewares[i](next)
	}
	return next
}
