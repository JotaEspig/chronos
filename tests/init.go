package tests

import (
	"chronos/config"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../.env")
	os.Setenv("IS_TEST_RUN", "true")
	config.InitDB()
}
