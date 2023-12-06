package server

import (
	"chronos/pkg/api"
	"chronos/pkg/types"
)

func (s *Server) setRoute(route types.Route) {
	gp := s.echo.Group(route.Path, route.Middlewares...)
	switch route.Method {
	case types.MethodGET:
		gp.GET("", route.Fn)
	case types.MethodPOST:
		gp.POST("", route.Fn)
	case types.MethodPUT:
		gp.PUT("", route.Fn)
	case types.MethodDELETE:
		gp.DELETE("", route.Fn)
	case types.MethodOPTIONS:
		gp.OPTIONS("", route.Fn)
	}
}

func (s *Server) setRoutes() {
	for _, route := range api.AllAvailableRoutes {
		s.setRoute(route)
	}
}
