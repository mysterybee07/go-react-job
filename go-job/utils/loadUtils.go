package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
