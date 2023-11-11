package utils

import (
	dotenv "github.com/joho/godotenv"
	"path"
)

func LoadEnv() {
	envPath := path.Join(GetCurrentDir(), "..", ".env")
	if err := dotenv.Load(envPath); err != nil {
		panic("Error loading .env file.")
	}
}
