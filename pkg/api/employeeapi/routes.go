package employeeapi

import (
	"chronos/config"
	"chronos/pkg/types"

	"github.com/labstack/echo/v4"
)

var AvailableRoutes []types.Route = []types.Route{
	{Path: "/api/employee/add", Method: types.MethodPOST, Fn: createEmployee,
		Middlewares: []echo.MiddlewareFunc{config.JWTMiddleware()}},
	{Path: "/api/employee/:id", Method: types.MethodGET, Fn: getEmployee,
		Middlewares: []echo.MiddlewareFunc{config.JWTMiddleware()}},
	{Path: "/api/employee/:id", Method: types.MethodPUT, Fn: updateEmployee,
		Middlewares: []echo.MiddlewareFunc{config.JWTMiddleware()}},
	{Path: "/api/employee/:id", Method: types.MethodDELETE, Fn: deleteEmployee,
		Middlewares: []echo.MiddlewareFunc{config.JWTMiddleware()}},
}
