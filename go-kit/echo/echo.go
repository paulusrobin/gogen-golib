package echo

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	option interface {
		Apply(o *options)
	}
	options struct {
		encoder      func(c echo.Context, response interface{}) error
		decoder      func(c echo.Context) (interface{}, error)
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
)

// Handler is function to adapt go-kit to echo http framework.
func Handler(ep endpoint.Endpoint, opts ...option) func(c echo.Context) error {
	var settings = options{
		encoder:      defaultEncoder,
		decoder:      defaultDecoder,
		errorHandler: nil,
	}
	for _, opt := range opts {
		opt.Apply(&settings)
	}

	return func(c echo.Context) error {
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
}
