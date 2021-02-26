package util

import (
	"bufio"
	"os"
	"strings"
)

// Helper method to create mock data for tests
func LoadEnvFromFileForTests() {

	file, _ := os.Open("../.env.test")

	defer file.Close()

	// Load the env vars from .env
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		envvars := strings.Split(line, "=")
		os.Setenv(envvars[0], envvars[1])
	}
}
