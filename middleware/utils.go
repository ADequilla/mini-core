// Package utils provides ...
package middleware

import (
	"encoding/base64"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func ConvertToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
