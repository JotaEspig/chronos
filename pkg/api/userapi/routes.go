package userapi

import "chronos/pkg/types"

var AvailableRoutes []types.Route = []types.Route{
	{Path: "/api/user/", Method: types.MethodPOST, Fn: createUser},
	{Path: "/api/user/:id", Method: types.MethodGET, Fn: getUser},
	{Path: "/api/user/:id", Method: types.MethodDELETE, Fn: deleteUser},
}
