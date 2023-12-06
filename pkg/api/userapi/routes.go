package userapi

import (
	"chronos/config"
	"chronos/pkg/types"

	"github.com/labstack/echo/v4"
)

var AvailableRoutes []types.Route = []types.Route{
	{Path: "/api/user/add", Method: types.MethodPOST, Fn: createUser},
	{Path: "/api/user/:id", Method: types.MethodGET, Fn: getUser,
		Middlewares: []echo.MiddlewareFunc{config.JWTMiddleware()}},
	{Path: "/api/user/:id", Method: types.MethodPUT, Fn: updateUser,
		Middlewares: []echo.MiddlewareFunc{config.JWTMiddleware()}},
	{Path: "/api/user/:id", Method: types.MethodDELETE, Fn: deleteUser,
		Middlewares: []echo.MiddlewareFunc{config.JWTMiddleware()}},
}
