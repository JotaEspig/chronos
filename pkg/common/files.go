package common

import (
	"io"
	"os"
)

// ReadFile reads a whole file and returns its content. DO NOT USE IT IF THE
// FILE IS TOO BIG
func ReadFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		return ""
	}
	content, err := io.ReadAll(file)
	if err != nil {
		return ""
	}
	return string(content)
}
