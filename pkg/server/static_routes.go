package server

import (
	"fmt"
	"os"
)

// staticRoute represents a static route with ´path´ being the URL that
// the user access the route (e.g. "/html") and ´root´ being the path to dir
// where the files you want to route are located
// (e.g. "/home/john/html-files").
type staticRoute struct {
	path, root string
}

// staticRoutes is a slice of staticRoute.
// WARNING: Root fields should always be "%s/<dir-you-want>" where %s will be
// replaced by chronos root dir. See the documentation of init function at the
// end of this file
var staticRoutes []staticRoute = []staticRoute{
	{"/", "%s/html"},
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
		staticRoutes[idx].root = fmt.Sprintf(elem.root, chronosRootDir)
		elem = staticRoutes[idx]

		s.echo.Static(elem.path, elem.root)
	}
}
