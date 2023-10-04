package tests

import (
	"chronos/config"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	os.Chdir("..")
	godotenv.Load("./.env")
	config.InitDB()
}
