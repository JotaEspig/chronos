package common

import (
	"fmt"
	"io"
	"os"
)

// ReadFile reads a whole file and returns its content. DO NOT USE IT IF THE
// FILE IS TOO BIG
func ReadFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		pwd, _ := os.Getwd()
		fmt.Fprintf(
			os.Stderr, "ReadFile: unable to find file: '%s' at '%s'\n",
			filename,
			pwd,
		)
		return ""
	}
	content, err := io.ReadAll(file)
	if err != nil {
		return ""
	}
	return string(content)
}
