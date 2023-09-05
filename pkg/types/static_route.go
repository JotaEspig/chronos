package types

// StaticRoute represents a static route with ´path´ being the URL that
// the user access the route (e.g. "/html") and ´root´ being the path to dir
// where the files you want to route are located
// (e.g. "/home/john/html-files").
type StaticRoute struct {
	Path, Root string
}
