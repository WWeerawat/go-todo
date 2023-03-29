package config

import (
	"github.com/joho/godotenv"
	"os"
)

// EnvVar function is for read .env file
func EnvVar(key string, defaultVal string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultVal
	}
	return value
}
