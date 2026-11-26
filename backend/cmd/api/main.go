package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"aetheria-backend/internal/config"
	"aetheria-backend/internal/httpapi"
	"aetheria-backend/internal/store"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	startupCtx, cancelStartup := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelStartup()

	pgPool, err := store.NewPostgresPool(startupCtx, cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("init postgres: %v", err)
	}
	defer pgPool.Close()

	rdb, err := store.NewRedisClient(startupCtx, cfg.RedisAddr)
	if err != nil {
		log.Fatalf("init redis: %v", err)
	}
	defer rdb.Close()

	handler := httpapi.NewServer(pgPool, httpapi.RedisPinger{Client: rdb})
	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	sigCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("aetheria-backend listening on :%s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("http server error: %v", err)
		}
	}()

	<-sigCtx.Done()
	log.Println("shutting down aetheria-backend...")

	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("graceful shutdown error: %v", err)
	}
}
