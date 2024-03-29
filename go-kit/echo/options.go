package echo

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	Option interface {
		Apply(o *options)
	}
	options struct {
		encoder      func(c echo.Context, response interface{}) error
		decoder      func(c echo.Context) (interface{}, error)
		middlewares  []echo.MiddlewareFunc
		errorHandler func(c echo.Context, err error) error
	}
)

var (
	defaultDecoder = func(c echo.Context) (interface{}, error) {
		return nil, nil
	}
	defaultEncoder = func(c echo.Context, response interface{}) error {
		return c.JSON(http.StatusOK, response)
	}
	defaultOptions = options{
		encoder:      defaultEncoder,
		decoder:      defaultDecoder,
		middlewares:  make([]echo.MiddlewareFunc, 0),
		errorHandler: nil,
	}
)

// Decoder Option
type withDecoder func(c echo.Context) (interface{}, error)

func (w withDecoder) Apply(o *options) {
	o.decoder = w
}

func WithDecoder(decoder func(c echo.Context) (interface{}, error)) Option {
	return withDecoder(decoder)
}

// Encoder Option
type withEncoder func(c echo.Context, response interface{}) error

func (w withEncoder) Apply(o *options) {
	o.encoder = w
}

func WithEncoder(encoder func(c echo.Context, response interface{}) error) Option {
	return withEncoder(encoder)
}

// ErrorHandler Option
type withErrorHandler func(c echo.Context, err error) error

func (w withErrorHandler) Apply(o *options) {
	o.errorHandler = w
}

func WithErrorHandler(errorHandler func(c echo.Context, err error) error) Option {
	return withErrorHandler(errorHandler)
}

// Middleware Option
type withMiddleware []echo.MiddlewareFunc

func (w withMiddleware) Apply(o *options) {
	o.middlewares = append(o.middlewares, w...)
}

func WithMiddleware(middlewares ...echo.MiddlewareFunc) Option {
	return withMiddleware(middlewares)
}
