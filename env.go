package sumaregi

import (
	"fmt"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from .env file.
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
