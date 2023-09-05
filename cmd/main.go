package main

import (
	"chronos/pkg/server"
	"os"
	"strconv"
)

func main() {
	// go to root dir
	os.Chdir("..")

	portStr := os.Getenv("PORT")
	port, ok := strconv.Atoi(portStr)
	if ok != nil {
		panic("PORT is not set")
	}

	s := server.NewServer(port)
	s.Start()
}
