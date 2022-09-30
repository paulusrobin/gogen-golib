package http

import "github.com/labstack/echo/v4"

type registerRoute func(ec *echo.Echo)

// Apply function to implement Option.
func (r registerRoute) Apply(o *options) {
	o.routes = append(o.routes, r)
}

// RegisterRoute function to add route to server.
func RegisterRoute(fn func(*echo.Echo)) Option {
	return registerRoute(fn)
}
