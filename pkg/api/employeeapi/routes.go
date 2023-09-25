package employeeapi

import "chronos/pkg/types"

var AvailableRoutes []types.Route = []types.Route{
	{Path: "/api/employee/add", Method: types.MethodPOST, Fn: createEmployee},
}
