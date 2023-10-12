package other

import (
	"chronos/config"
	"chronos/pkg/types"

	"github.com/labstack/echo/v4"
)

var AvailableRoutes []types.Route = []types.Route{
	{Path: "/api/hello", Method: types.MethodGET, Fn: hello},
	{Path: "/api/jwt", Method: types.MethodGET, Fn: needsJWT,
		Middlewares: []echo.MiddlewareFunc{config.JWTMiddleware()}},
}
