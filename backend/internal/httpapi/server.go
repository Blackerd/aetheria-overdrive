package httpapi

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

// Pinger abstracts health check connectivity for databases and caches.
type Pinger interface {
	Ping(ctx context.Context) error
}

// RedisPinger adapts *redis.Client to the Pinger interface.
type RedisPinger struct {
	Client *redis.Client
}

// Ping executes a Redis PING command and returns any error.
func (r RedisPinger) Ping(ctx context.Context) error {
	return r.Client.Ping(ctx).Err()
}

// Server exposes the REST HTTP API using the Go standard library ServeMux.
type Server struct {
	mux *http.ServeMux
	pg  Pinger
	rdb Pinger
}

// NewServer constructs an HTTP API server with all registered routes.
func NewServer(pg Pinger, rdb Pinger) *Server {
	s := &Server{
		mux: http.NewServeMux(),
		pg:  pg,
		rdb: rdb,
	}

	s.mux.HandleFunc("GET /healthz", s.handleHealthz)
	return s
}

// ServeHTTP implements http.Handler.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

type healthResponse struct {
	Status   string            `json:"status"`
	Services map[string]string `json:"services"`
}

func (s *Server) handleHealthz(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	pgStatus := "ok"
	if s.pg == nil || s.pg.Ping(ctx) != nil {
		pgStatus = "down"
	}

	redisStatus := "ok"
	if s.rdb == nil || s.rdb.Ping(ctx) != nil {
		redisStatus = "down"
	}

	overallStatus := "ok"
	statusCode := http.StatusOK
	if pgStatus != "ok" || redisStatus != "ok" {
		overallStatus = "degraded"
		statusCode = http.StatusServiceUnavailable
	}

	writeJSON(w, statusCode, healthResponse{
		Status: overallStatus,
		Services: map[string]string{
			"postgres": pgStatus,
			"redis":    redisStatus,
		},
	})
}

func writeJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(payload)
}
