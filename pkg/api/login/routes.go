package login

import "chronos/pkg/types"

var AvailableRoutes []types.Route = []types.Route{
	{Path: "/api/login", Method: types.MethodPOST, Fn: login},
}
