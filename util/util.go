package util

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Helper method to env when run from ide
func LoadEnvFromFileForTests() {

	if os.Getenv("POSTGRES_HOST") == "" {

		file, _ := os.Open("../.env.test")

		defer file.Close()

		// Load the env vars from .env
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			envvars := strings.Split(line, "=")
			os.Setenv(envvars[0], envvars[1])
			log.Println(envvars)
		}
	}
}
