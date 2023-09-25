package api

import (
	"chronos/pkg/api/employeeapi"
	"chronos/pkg/api/other"
	"chronos/pkg/api/userapi"
	"chronos/pkg/types"
)

var AllAvailableRoutes []types.Route

func init() {
	AllAvailableRoutes = append(AllAvailableRoutes, other.AvailableRoutes...)
	AllAvailableRoutes = append(AllAvailableRoutes, userapi.AvailableRoutes...)
	AllAvailableRoutes = append(AllAvailableRoutes, employeeapi.AvailableRoutes...)
}
