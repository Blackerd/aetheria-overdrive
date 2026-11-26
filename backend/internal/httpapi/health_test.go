package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"aetheria-backend/internal/store"
	"github.com/joho/godotenv"
)

type mockPinger struct {
	err error
}

func (m mockPinger) Ping(context.Context) error {
	return m.err
}

func TestHealthz_AllHealthy(t *testing.T) {
	srv := NewServer(mockPinger{}, mockPinger{})
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	srv.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var resp healthResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if resp.Status != "ok" {
		t.Errorf("expected status ok, got %q", resp.Status)
	}
	if resp.Services["postgres"] != "ok" {
		t.Errorf("expected postgres ok, got %q", resp.Services["postgres"])
	}
	if resp.Services["redis"] != "ok" {
		t.Errorf("expected redis ok, got %q", resp.Services["redis"])
	}
}

func TestHealthz_PostgresDown(t *testing.T) {
	srv := NewServer(mockPinger{err: errors.New("connection refused")}, mockPinger{})
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	srv.ServeHTTP(rec, req)

	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected status 503, got %d", rec.Code)
	}

	var resp healthResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if resp.Status != "degraded" {
		t.Errorf("expected status degraded, got %q", resp.Status)
	}
	if resp.Services["postgres"] != "down" {
		t.Errorf("expected postgres down, got %q", resp.Services["postgres"])
	}
	if resp.Services["redis"] != "ok" {
		t.Errorf("expected redis ok, got %q", resp.Services["redis"])
	}
}

func TestHealthz_RedisDown(t *testing.T) {
	srv := NewServer(mockPinger{}, mockPinger{err: errors.New("timeout")})
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	srv.ServeHTTP(rec, req)

	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected status 503, got %d", rec.Code)
	}

	var resp healthResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if resp.Status != "degraded" {
		t.Errorf("expected status degraded, got %q", resp.Status)
	}
	if resp.Services["postgres"] != "ok" {
		t.Errorf("expected postgres ok, got %q", resp.Services["postgres"])
	}
	if resp.Services["redis"] != "down" {
		t.Errorf("expected redis down, got %q", resp.Services["redis"])
	}
}

func TestHealthz_IntegrationDocker(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping Docker integration test in short mode")
	}

	_ = godotenv.Load("../../../.env", "../../.env")

	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		dsn = "postgres://aetheria:aetheria_dev_pass@localhost:5432/aetheria_overdrive?sslmode=disable"
	}
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pgPool, err := store.NewPostgresPool(ctx, dsn)
	if err != nil {
		t.Fatalf("connect docker postgres: %v", err)
	}
	defer pgPool.Close()

	rdb, err := store.NewRedisClient(ctx, redisAddr)
	if err != nil {
		t.Fatalf("connect docker redis: %v", err)
	}
	defer rdb.Close()

	srv := NewServer(pgPool, RedisPinger{Client: rdb})
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	srv.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d (body: %s)", rec.Code, rec.Body.String())
	}

	var resp healthResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if resp.Status != "ok" || resp.Services["postgres"] != "ok" || resp.Services["redis"] != "ok" {
		t.Fatalf("unexpected health response: %+v", resp)
	}
}
