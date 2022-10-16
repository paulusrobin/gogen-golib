package echo

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/labstack/echo/v4"
)

// Handler is function to adapt go-kit to echo http framework.
func Handler(ep endpoint.Endpoint, opts ...Option) func(c echo.Context) error {
	var settings = defaultOptions
	for _, opt := range opts {
		opt.Apply(&settings)
	}

	next := func(c echo.Context) error {
		errFunc := func(err error) error {
			if settings.errorHandler != nil {
				return settings.errorHandler(c, err)
			}
			return err
		}

		request, err := settings.decoder(c)
		if err != nil {
			return errFunc(err)
		}

		response, err := ep(c.Request().Context(), request)
		if err != nil {
			return errFunc(err)
		}

		return settings.encoder(c, response)
	}
	for i := len(settings.middlewares) - 1; i >= 0; i-- {
		next = settings.middlewares[i](next)
	}
	return next
}
