package main

import (
	"chronos/pkg/server"
	"os"
	"strconv"
)

func main() {
	portStr := os.Getenv("CHRONOS_PORT")
	port, ok := strconv.Atoi(portStr)
	if ok != nil {
		panic("CHRONOS_PORT is not set")
	}

	s := server.NewServer(port)
	s.Start()
}
