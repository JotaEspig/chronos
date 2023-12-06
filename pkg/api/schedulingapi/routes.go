package schedulingapi

import (
	"chronos/config"
	"chronos/pkg/types"

	"github.com/labstack/echo/v4"
)

var AvailableRoutes []types.Route = []types.Route{
	{Path: "/api/scheduling/add", Method: types.MethodPOST, Fn: createScheduling,
		Middlewares: []echo.MiddlewareFunc{config.JWTMiddleware()}},
	{Path: "/api/scheduling/:id", Method: types.MethodGET, Fn: getScheduling,
		Middlewares: []echo.MiddlewareFunc{config.JWTMiddleware()}},
	{Path: "/api/scheduling", Method: types.MethodGET, Fn: getSchedulingsByDate,
		Middlewares: []echo.MiddlewareFunc{config.JWTMiddleware()}},
	{Path: "/api/scheduling/:id", Method: types.MethodPUT, Fn: updateScheduling,
		Middlewares: []echo.MiddlewareFunc{config.JWTMiddleware()}},
	{Path: "/api/scheduling/:id", Method: types.MethodDELETE, Fn: deleteScheduling,
		Middlewares: []echo.MiddlewareFunc{config.JWTMiddleware()}},
}
