package http

import (
	"github.com/labstack/echo/v4"
	validator "github.com/paulusrobin/gogen-golib/validator/interface"
	"net/http"
)

type (
	Option interface {
		Apply(o *options)
	}
	options struct {
		errorHandler func(err error, c echo.Context)
		middlewares  []echo.MiddlewareFunc
		routes       []func(ec *echo.Echo)
	}
)

var (
	defaultErrorHandler = func(err error, c echo.Context) {
		if vErr, isValidationErr := validator.IsValidationError(err); isValidationErr {
			_ = c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": vErr.Error(),
				"details": vErr.Details,
			})
			return
		}
		_ = c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	defaultOption = options{
		errorHandler: defaultErrorHandler,
		middlewares:  make([]echo.MiddlewareFunc, 0),
		routes:       make([]func(*echo.Echo), 0),
	}
)
