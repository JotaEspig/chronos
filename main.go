package main

import (
	//"chronos/config"
	"chronos/pkg/server"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	/*
		defer config.DB.Close()

		db := config.DB
		stmt, err := db.Prepare("INSERT INTO test(id) VALUES (?);")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()
		stmt.Exec("-- DROP TABLE test; --")

		os.Exit(0)
	*/
	godotenv.Load(".env")

	portStr := os.Getenv("CHRONOS_PORT")
	port, ok := strconv.Atoi(portStr)
	if ok != nil {
		panic("CHRONOS_PORT is not set")
	}

	s := server.NewServer(port)
	s.Start()
}
