package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               string
	DatabaseURL        string
	SupabaseURL        string
	SupabaseServiceKey string
	SupabaseJWTSecret  string
	FrontendURL        string
}

func LoadConfig() *Config {
	// Attempt to load .env from current directory or parent
	if err := godotenv.Load(); err != nil {
		_ = godotenv.Load("../../.env")
	}

	cfg := &Config{
		Port:               getEnv("PORT", "8080"),
		DatabaseURL:        getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/jayaabadisolusi?sslmode=disable"),
		SupabaseURL:        getEnv("SUPABASE_URL", "https://xyzcompany.supabase.co"),
		SupabaseServiceKey: getEnv("SUPABASE_SERVICE_ROLE_KEY", "dummy_service_role_key"),
		SupabaseJWTSecret:  getEnv("SUPABASE_JWT_SECRET", "dummy_jwt_secret_for_local_dev"),
		FrontendURL:        getEnv("FRONTEND_URL", "http://localhost:5173"),
	}

	log.Printf("[CONFIG] Running on port :%s | DB: configured | Supabase URL: %s", cfg.Port, cfg.SupabaseURL)
	return cfg
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
