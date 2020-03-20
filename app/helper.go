package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// SetURL for dynamic url
func SetURL(path string) string {
	url := os.Getenv("url") + path

	return url
}
