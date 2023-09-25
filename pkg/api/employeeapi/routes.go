package employeeapi

import "chronos/pkg/types"

var AvailableRoutes []types.Route = []types.Route{
	{Path: "/api/employee/add", Method: types.MethodPOST, Fn: createEmployee},
	{Path: "/api/employee/:id", Method: types.MethodGET, Fn: getEmployee},
}
