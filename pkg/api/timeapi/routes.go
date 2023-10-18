package timeapi

import (
	"chronos/config"
	"chronos/pkg/types"

	"github.com/labstack/echo/v4"
)

var AvailableRoutes []types.Route = []types.Route{
	{Path: "/api/time/add", Method: types.MethodPOST, Fn: createTime,
		Middlewares: []echo.MiddlewareFunc{config.JWTMiddleware()}},
	{Path: "/api/time/:id", Method: types.MethodGET, Fn: getTime,
		Middlewares: []echo.MiddlewareFunc{config.JWTMiddleware()}},
	{Path: "/api/time", Method: types.MethodGET, Fn: getTimesByDate,
		Middlewares: []echo.MiddlewareFunc{config.JWTMiddleware()}},
	{Path: "/api/time/:id", Method: types.MethodPUT, Fn: updateTime,
		Middlewares: []echo.MiddlewareFunc{config.JWTMiddleware()}},
	{Path: "/api/time/:id", Method: types.MethodDELETE, Fn: deleteTime,
		Middlewares: []echo.MiddlewareFunc{config.JWTMiddleware()}},
}
