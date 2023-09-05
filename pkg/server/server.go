package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
	port int
}

// NewServer creates a new Server
func NewServer(port int) *Server {
	s := &Server{}

	if port <= 0 {
		panic("PORT is a invalid number (port <= 0)")
	}

	e := echo.New()
	s.echo = e

	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.CORS())
	s.setStaticRoutes()

	return s
}

// Start starts the server using ´port´
func (s *Server) Start() {
	addr := fmt.Sprintf(":%d", s.port)
	s.echo.Logger.Fatal(s.echo.Start(addr))
}
