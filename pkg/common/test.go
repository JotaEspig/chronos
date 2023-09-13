package common

import (
	"os"
)

// IsTestRun checks if it's running as a test or not
func IsTestRun() bool {
	env := os.Getenv("IS_TEST_RUN")
	return env == "true"
}
