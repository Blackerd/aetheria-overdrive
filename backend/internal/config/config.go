package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const (
	defaultPort      = "8080"
	defaultRedisAddr = "localhost:6379"
)

// Config holds runtime configuration for the backend cloud services daemon.
type Config struct {
	Port                 string
	PostgresDSN          string
	RedisAddr            string
	JWTSecret            string
	ServerInternalSecret string
}

// Load reads configuration from environment variables, falling back to .env files if present.
func Load() (Config, error) {
	// Best-effort load for local development; OS env vars still take precedence.
	_ = godotenv.Load(".env", "../.env")

	cfg := Config{
		Port:                 getEnvOrDefault("BACKEND_PORT", defaultPort),
		PostgresDSN:          strings.TrimSpace(os.Getenv("POSTGRES_DSN")),
		RedisAddr:            getEnvOrDefault("REDIS_ADDR", defaultRedisAddr),
		JWTSecret:            strings.TrimSpace(os.Getenv("JWT_SECRET")),
		ServerInternalSecret: strings.TrimSpace(os.Getenv("SERVER_INTERNAL_SECRET")),
	}

	var missing []string
	if cfg.PostgresDSN == "" {
		missing = append(missing, "POSTGRES_DSN")
	}
	if cfg.JWTSecret == "" {
		missing = append(missing, "JWT_SECRET")
	}
	if cfg.ServerInternalSecret == "" {
		missing = append(missing, "SERVER_INTERNAL_SECRET")
	}

	if len(missing) > 0 {
		return Config{}, fmt.Errorf("missing required configuration: %s", strings.Join(missing, ", "))
	}

	return cfg, nil
}

func getEnvOrDefault(key, fallback string) string {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		return fallback
	}
	return val
}
