package server

import (
	"chronos/pkg/api"
	"chronos/pkg/types"
)

func (s *Server) setRoute(route types.Route) {
	switch route.Method {
	case types.MethodGET:
		s.echo.GET(route.Path, route.Fn)
	case types.MethodPOST:
		s.echo.POST(route.Path, route.Fn)
	case types.MethodPUT:
		s.echo.PUT(route.Path, route.Fn)
	case types.MethodDELETE:
		s.echo.DELETE(route.Path, route.Fn)
	case types.MethodOPTIONS:
		s.echo.OPTIONS(route.Path, route.Fn)
	}
}

func (s *Server) setRoutes() {
	for _, route := range api.AllAvailableRoutes {
		s.setRoute(route)
	}
}
