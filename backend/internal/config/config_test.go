package config

import (
	"strings"
	"testing"
)

func TestLoad_SuccessAndDefaults(t *testing.T) {
	t.Setenv("BACKEND_PORT", "")
	t.Setenv("REDIS_ADDR", "")
	t.Setenv("POSTGRES_DSN", "postgres://aetheria:pass@localhost:5432/aetheria_overdrive?sslmode=disable")
	t.Setenv("JWT_SECRET", "test_jwt_secret_1234567890")
	t.Setenv("SERVER_INTERNAL_SECRET", "test_internal_secret_1234567890")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if cfg.Port != "8080" {
		t.Errorf("expected default Port 8080, got %q", cfg.Port)
	}
	if cfg.RedisAddr != "localhost:6379" {
		t.Errorf("expected default RedisAddr localhost:6379, got %q", cfg.RedisAddr)
	}
	if cfg.PostgresDSN == "" || cfg.JWTSecret == "" || cfg.ServerInternalSecret == "" {
		t.Errorf("expected required fields to be populated, got %+v", cfg)
	}
}

func TestLoad_CustomValues(t *testing.T) {
	t.Setenv("BACKEND_PORT", "9090")
	t.Setenv("REDIS_ADDR", "redis:6380")
	t.Setenv("POSTGRES_DSN", "postgres://u:p@db:5432/app")
	t.Setenv("JWT_SECRET", "custom_jwt")
	t.Setenv("SERVER_INTERNAL_SECRET", "custom_internal")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if cfg.Port != "9090" {
		t.Errorf("expected Port 9090, got %q", cfg.Port)
	}
	if cfg.RedisAddr != "redis:6380" {
		t.Errorf("expected RedisAddr redis:6380, got %q", cfg.RedisAddr)
	}
}

func TestLoad_MissingRequired(t *testing.T) {
	t.Setenv("POSTGRES_DSN", "")
	t.Setenv("JWT_SECRET", "")
	t.Setenv("SERVER_INTERNAL_SECRET", "")

	_, err := Load()
	if err == nil {
		t.Fatal("expected error when required environment variables are missing, got nil")
	}

	msg := err.Error()
	for _, key := range []string{"POSTGRES_DSN", "JWT_SECRET", "SERVER_INTERNAL_SECRET"} {
		if !strings.Contains(msg, key) {
			t.Errorf("expected error message %q to mention %s", msg, key)
		}
	}
}
