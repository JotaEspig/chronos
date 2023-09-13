package main

import (
	"chronos/config"
	"chronos/pkg/server"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	config.InitDB()
	defer config.DB.Close()

	portStr := os.Getenv("CHRONOS_PORT")
	port, ok := strconv.Atoi(portStr)
	if ok != nil {
		panic("CHRONOS_PORT is not set")
	}

	s := server.NewServer(port)
	s.Start()
}
