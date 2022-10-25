package sarama

import (
	"context"
	"github.com/Shopify/sarama"
)

type (
	EncoderFunc    func(ctx context.Context, response interface{}) error
	DecoderFunc    func(ctx context.Context, message *sarama.ConsumerMessage) (interface{}, error)
	HandlerFunc    func(ctx context.Context, message *sarama.ConsumerMessage) error
	MiddlewareFunc func(next HandlerFunc) HandlerFunc

	Option interface {
		Apply(o *options)
	}
	options struct {
		decoder      func(ctx context.Context, message *sarama.ConsumerMessage) (interface{}, error)
		encoder      func(ctx context.Context, response interface{}) error
		middlewares  []MiddlewareFunc
		errorHandler func(ctx context.Context, err error) error
	}
)

var (
	defaultDecoder = func(c context.Context, message *sarama.ConsumerMessage) (interface{}, error) {
		return nil, nil
	}
	defaultEncoder = func(c context.Context, response interface{}) error {
		return nil
	}
	defaultOptions = options{
		encoder:      defaultEncoder,
		decoder:      defaultDecoder,
		middlewares:  make([]MiddlewareFunc, 0),
		errorHandler: nil,
	}
)

// Decoder Option
type withDecoder func(ctx context.Context, message *sarama.ConsumerMessage) (interface{}, error)

func (w withDecoder) Apply(o *options) {
	o.decoder = w
}

func WithDecoder(decoder func(c context.Context, message *sarama.ConsumerMessage) (interface{}, error)) Option {
	return withDecoder(decoder)
}

// Encoder Option
type withEncoder func(c context.Context, response interface{}) error

func (w withEncoder) Apply(o *options) {
	o.encoder = w
}

func WithEncoder(encoder func(c context.Context, response interface{}) error) Option {
	return withEncoder(encoder)
}

// ErrorHandler Option
type withErrorHandler func(c context.Context, err error) error

func (w withErrorHandler) Apply(o *options) {
	o.errorHandler = w
}

func WithErrorHandler(errorHandler func(c context.Context, err error) error) Option {
	return withErrorHandler(errorHandler)
}

// Middleware Option
type withMiddleware []MiddlewareFunc

func (w withMiddleware) Apply(o *options) {
	o.middlewares = append(o.middlewares, w...)
}

func WithMiddleware(middlewares ...MiddlewareFunc) Option {
	return withMiddleware(middlewares)
}
