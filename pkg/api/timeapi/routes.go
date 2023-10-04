package timeapi

import "chronos/pkg/types"

var AvailableRoutes []types.Route = []types.Route{
	{Path: "/api/time/add", Method: types.MethodPOST, Fn: createTime},
	{Path: "/api/time/:id", Method: types.MethodGET, Fn: getTime},
	{Path: "/api/time", Method: types.MethodGET, Fn: getTimesByDate},
	{Path: "/api/time/:id", Method: types.MethodPUT, Fn: updateTime},
	{Path: "/api/time/:id", Method: types.MethodDELETE, Fn: deleteTime},
}
