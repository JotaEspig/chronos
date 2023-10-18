package types

import "github.com/labstack/echo/v4"

type Route struct {
	Path        string
	Method      string
	Fn          func(echo.Context) error
	Middlewares []echo.MiddlewareFunc
}
