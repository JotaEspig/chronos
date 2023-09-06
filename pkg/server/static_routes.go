package server

import (
	"chronos/pkg/types"
	"fmt"
	"os"
)

// staticRoutes is a slice of staticRoute.
// WARNING: Root fields should always be "%s/<dir-you-want>" where %s will be
// replaced by chronos root dir. See comment in setStaticRoutes function at the
// end of this file
var staticRoutes []types.StaticRoute = []types.StaticRoute{
	{Path: "/", Root: "%s/html"},
	{Path: "/css", Root: "%s/css"},
	{Path: "/js", Root: "%s/js"},
}

// setStaticRoutes sets the static routes using staticRoutes variable
func (s *Server) setStaticRoutes() {
	chronosRootDir := os.Getenv("CHRONOS_ROOT_DIR")
	if chronosRootDir == "" {
		panic("CHRONOS_ROOT_DIR is not set")
	}

	// format the root field from staticRoutes using CHRONOS_ROOT_DIR
	// because this way allow the admin to choose where the application files
	// are located, so the chronos binary file (executable file) can be located
	// in "/bin" and the application files in "/var" for example
	for idx, elem := range staticRoutes {
		staticRoutes[idx].Root = fmt.Sprintf(elem.Root, chronosRootDir)
		elem = staticRoutes[idx]

		s.echo.Static(elem.Path, elem.Root)
	}
}
