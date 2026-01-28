package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseURL string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	appConfig := &Config{
		Port:        getEnv("PORT", ":8080"),
		JWTSecret:   getEnv("JWT_SECRET", "defaultsecret"),
		DatabaseURL: getEnv("database_url", "mysql:root@tcp(localhost:3306)/gopostdb"),
	}
	return appConfig
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
