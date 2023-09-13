package tests

import (
	"os"
)

func init() {
	os.Setenv("IS_TEST_RUN", "true")
}
