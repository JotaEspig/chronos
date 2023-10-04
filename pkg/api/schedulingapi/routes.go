package schedulingapi

import "chronos/pkg/types"

var AvailableRoutes []types.Route = []types.Route{
	{Path: "/api/scheduling/add", Method: types.MethodPOST, Fn: createScheduling},
	{Path: "/api/scheduling/:id", Method: types.MethodGET, Fn: getScheduling},
	{Path: "/api/scheduling", Method: types.MethodGET, Fn: getSchedulingsByDate},
	{Path: "/api/scheduling/:id", Method: types.MethodPUT, Fn: updateScheduling},
	{Path: "/api/scheduling/:id", Method: types.MethodDELETE, Fn: deleteScheduling},
}
