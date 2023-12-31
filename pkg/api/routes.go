// package api provides api endpoints
package api

import (
	"chronos/pkg/api/employeeapi"
	"chronos/pkg/api/login"
	"chronos/pkg/api/other"
	"chronos/pkg/api/schedulingapi"
	"chronos/pkg/api/timeapi"
	"chronos/pkg/api/userapi"
	"chronos/pkg/types"
)

var AllAvailableRoutes []types.Route

func init() {
	AllAvailableRoutes = append(AllAvailableRoutes, other.AvailableRoutes...)
	AllAvailableRoutes = append(AllAvailableRoutes, userapi.AvailableRoutes...)
	AllAvailableRoutes = append(AllAvailableRoutes, employeeapi.AvailableRoutes...)
	AllAvailableRoutes = append(AllAvailableRoutes, timeapi.AvailableRoutes...)
	AllAvailableRoutes = append(AllAvailableRoutes, schedulingapi.AvailableRoutes...)
	AllAvailableRoutes = append(AllAvailableRoutes, login.AvailableRoutes...)
}
