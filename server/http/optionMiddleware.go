package http

import "github.com/labstack/echo/v4"

type registerMiddleware echo.MiddlewareFunc

// Apply function to implement Option.
func (r registerMiddleware) Apply(o *options) {
	o.middlewares = append(o.middlewares, echo.MiddlewareFunc(r))
}

// RegisterMiddleware function to add middleware to server.
func RegisterMiddleware(fn echo.MiddlewareFunc) Option {
	return registerMiddleware(fn)
}
