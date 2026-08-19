package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	DatabaseURL  string
	JWTSecret    string
	GeminiAPIKey string

	ObjectStorageEndpoint  string
	ObjectStorageBucket    string
	ObjectStorageAccessKey string
	ObjectStorageSecretKey string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		Port:         getEnv("PORT", "8080"),
		DatabaseURL:  getEnv("DATABASE_URL", ""),
		JWTSecret:    getEnv("JWT_SECRET", ""),
		GeminiAPIKey: getEnv("GEMINI_API_KEY", ""),

		ObjectStorageEndpoint:  getEnv("OBJECT_STORAGE_ENDPOINT", ""),
		ObjectStorageBucket:    getEnv("OBJECT_STORAGE_BUCKET", ""),
		ObjectStorageAccessKey: getEnv("OBJECT_STORAGE_ACCESS_KEY", ""),
		ObjectStorageSecretKey: getEnv("OBJECT_STORAGE_SECRET_KEY", ""),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
