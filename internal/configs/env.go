package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Load func to get env value
func Load(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
