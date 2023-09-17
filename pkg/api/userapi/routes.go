package userapi

import "chronos/pkg/types"

var AvailableRoutes []types.Route = []types.Route{
	{Path: "/api/user/add", Method: types.MethodPOST, Fn: createUser},
	{Path: "/api/user/:id", Method: types.MethodGET, Fn: getUser},
	{Path: "/api/user/:id", Method: types.MethodPUT, Fn: updateUser},
	{Path: "/api/user/:id", Method: types.MethodDELETE, Fn: deleteUser},
}
