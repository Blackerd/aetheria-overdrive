package store

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func loadTestPostgresDSN() string {
	_ = godotenv.Load("../../../.env", "../../.env")
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		dsn = "postgres://aetheria:aetheria_dev_pass@localhost:5432/aetheria_overdrive?sslmode=disable"
	}
	return dsn
}

func TestMigrateUp_IdempotentAndConstraints(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping Docker Postgres migration test in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := NewPostgresPool(ctx, loadTestPostgresDSN())
	if err != nil {
		t.Fatalf("connect postgres: %v", err)
	}
	defer pool.Close()

	if _, err := MigrateUp(ctx, pool); err != nil {
		t.Fatalf("first MigrateUp failed: %v", err)
	}

	appliedSecond, err := MigrateUp(ctx, pool)
	if err != nil {
		t.Fatalf("second MigrateUp failed: %v", err)
	}
	if appliedSecond != 0 {
		t.Fatalf("expected 0 migrations on second idempotent run, got %d", appliedSecond)
	}

	// Test constraints and JSONB inside a rolled-back transaction so public data stays clean.
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatalf("begin test tx: %v", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	// 1. Verify CHECK (core_relics >= 0) rejects negative currency.
	_, err = tx.Exec(ctx, `SAVEPOINT sp_neg_relics`)
	if err != nil {
		t.Fatalf("savepoint: %v", err)
	}
	_, err = tx.Exec(ctx,
		`INSERT INTO players (guest_token, display_name, core_relics) VALUES ($1, $2, $3)`,
		"guest_neg_test", "BadPlayer", -5,
	)
	if err == nil {
		t.Fatal("expected CHECK constraint error when inserting negative core_relics, got nil")
	}
	if _, rbErr := tx.Exec(ctx, `ROLLBACK TO SAVEPOINT sp_neg_relics`); rbErr != nil {
		t.Fatalf("rollback to savepoint: %v", rbErr)
	}

	// 2. Insert valid player, talent, and match_history with JSONB summary.
	var playerID string
	err = tx.QueryRow(ctx,
		`INSERT INTO players (guest_token, display_name, core_relics) VALUES ($1, $2, $3) RETURNING id`,
		"guest_valid_test", "DuoTest", 120,
	).Scan(&playerID)
	if err != nil {
		t.Fatalf("insert valid player: %v", err)
	}

	if _, err := tx.Exec(ctx,
		`INSERT INTO meta_talents (player_id, talent_key, level) VALUES ($1, $2, $3)`,
		playerID, "max_shield_boost", 2,
	); err != nil {
		t.Fatalf("insert meta_talent: %v", err)
	}

	if _, err := tx.Exec(ctx,
		`INSERT INTO match_history (match_id, player_id, operative_class, mode, wave_cleared, total_kills, duration_seconds, relics_earned, victory, run_summary)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10::jsonb)`,
		"match_test_001", playerID, "VANGUARD", "COOP", 10, 420, 915, 50, true,
		`{"perks":["plasma_bounce","overclock_core"],"boss_killed":true}`,
	); err != nil {
		t.Fatalf("insert match_history: %v", err)
	}

	// 3. Query JSONB using containment operator @> (backed by GIN index).
	var foundMatchID string
	err = tx.QueryRow(ctx,
		`SELECT match_id FROM match_history WHERE player_id = $1 AND run_summary @> '{"boss_killed":true}'::jsonb`,
		playerID,
	).Scan(&foundMatchID)
	if err != nil {
		t.Fatalf("query JSONB boss_killed: %v", err)
	}
	if foundMatchID != "match_test_001" {
		t.Fatalf("expected match_test_001, got %q", foundMatchID)
	}
}

func TestMigrate_UpDownRoundTrip(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping Docker Postgres round-trip test in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dsn := loadTestPostgresDSN()
	adminPool, err := NewPostgresPool(ctx, dsn)
	if err != nil {
		t.Fatalf("connect admin postgres: %v", err)
	}
	defer adminPool.Close()

	tempSchema := fmt.Sprintf("mig_test_%d", time.Now().UnixNano())
	if _, err := adminPool.Exec(ctx, fmt.Sprintf(`CREATE SCHEMA %s`, tempSchema)); err != nil {
		t.Fatalf("create temp schema: %v", err)
	}
	defer func() {
		_, _ = adminPool.Exec(context.Background(), fmt.Sprintf(`DROP SCHEMA IF EXISTS %s CASCADE`, tempSchema))
	}()

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		t.Fatalf("parse dsn: %v", err)
	}
	if cfg.ConnConfig.RuntimeParams == nil {
		cfg.ConnConfig.RuntimeParams = map[string]string{}
	}
	cfg.ConnConfig.RuntimeParams["search_path"] = tempSchema

	scopedPool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		t.Fatalf("create scoped pool: %v", err)
	}
	defer scopedPool.Close()

	applied, err := MigrateUp(ctx, scopedPool)
	if err != nil {
		t.Fatalf("scoped MigrateUp failed: %v", err)
	}
	if applied < 1 {
		t.Fatalf("expected at least 1 applied migration, got %d", applied)
	}

	var playersTable *string
	if err := scopedPool.QueryRow(ctx, `SELECT to_regclass($1)::text`, tempSchema+".players").Scan(&playersTable); err != nil {
		t.Fatalf("check players table exists: %v", err)
	}
	if playersTable == nil {
		t.Fatal("expected players table to exist after MigrateUp")
	}

	rolledBack, err := MigrateDown(ctx, scopedPool)
	if err != nil {
		t.Fatalf("scoped MigrateDown failed: %v", err)
	}
	if rolledBack != applied {
		t.Fatalf("expected rolledBack == %d, got %d", applied, rolledBack)
	}

	var afterDown *string
	if err := scopedPool.QueryRow(ctx, `SELECT to_regclass($1)::text`, tempSchema+".players").Scan(&afterDown); err != nil {
		t.Fatalf("check players table removed: %v", err)
	}
	if afterDown != nil {
		t.Fatalf("expected players table to be dropped after MigrateDown, got %q", *afterDown)
	}

	reapplied, err := MigrateUp(ctx, scopedPool)
	if err != nil {
		t.Fatalf("re-apply MigrateUp failed: %v", err)
	}
	if reapplied != applied {
		t.Fatalf("expected re-applied == %d, got %d", applied, reapplied)
	}
}
