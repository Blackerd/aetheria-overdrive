package store

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

func ensureMigrationsTable(ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version    VARCHAR(64) PRIMARY KEY,
			applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);
	`)
	if err != nil {
		return fmt.Errorf("ensure schema_migrations table: %w", err)
	}
	return nil
}

// MigrateUp applies all unapplied *.up.sql migrations in lexicographical order.
func MigrateUp(ctx context.Context, pool *pgxpool.Pool) (int, error) {
	if err := ensureMigrationsTable(ctx, pool); err != nil {
		return 0, err
	}

	entries, err := fs.ReadDir(migrationFiles, "migrations")
	if err != nil {
		return 0, fmt.Errorf("read embedded migrations: %w", err)
	}

	var versions []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasSuffix(name, ".up.sql") {
			versions = append(versions, strings.TrimSuffix(name, ".up.sql"))
		}
	}
	sort.Strings(versions)

	applied := 0
	for _, version := range versions {
		var exists bool
		err := pool.QueryRow(ctx, `SELECT EXISTS (SELECT 1 FROM schema_migrations WHERE version = $1)`, version).Scan(&exists)
		if err != nil {
			return applied, fmt.Errorf("check migration %s: %w", version, err)
		}
		if exists {
			continue
		}

		sqlBytes, err := migrationFiles.ReadFile("migrations/" + version + ".up.sql")
		if err != nil {
			return applied, fmt.Errorf("read up migration %s: %w", version, err)
		}

		tx, err := pool.Begin(ctx)
		if err != nil {
			return applied, fmt.Errorf("begin tx for %s: %w", version, err)
		}

		if _, err := tx.Exec(ctx, string(sqlBytes)); err != nil {
			_ = tx.Rollback(ctx)
			return applied, fmt.Errorf("exec up migration %s: %w", version, err)
		}

		if _, err := tx.Exec(ctx, `INSERT INTO schema_migrations (version) VALUES ($1)`, version); err != nil {
			_ = tx.Rollback(ctx)
			return applied, fmt.Errorf("record migration %s: %w", version, err)
		}

		if err := tx.Commit(ctx); err != nil {
			return applied, fmt.Errorf("commit migration %s: %w", version, err)
		}
		applied++
	}

	return applied, nil
}

// MigrateDown rolls back applied migrations in reverse order.
func MigrateDown(ctx context.Context, pool *pgxpool.Pool) (int, error) {
	if err := ensureMigrationsTable(ctx, pool); err != nil {
		return 0, err
	}

	rows, err := pool.Query(ctx, `SELECT version FROM schema_migrations ORDER BY version DESC`)
	if err != nil {
		return 0, fmt.Errorf("query applied migrations: %w", err)
	}
	defer rows.Close()

	var versions []string
	for rows.Next() {
		var v string
		if err := rows.Scan(&v); err != nil {
			return 0, fmt.Errorf("scan migration version: %w", err)
		}
		versions = append(versions, v)
	}
	if err := rows.Err(); err != nil {
		return 0, fmt.Errorf("iterate migration versions: %w", err)
	}

	rolledBack := 0
	for _, version := range versions {
		sqlBytes, err := migrationFiles.ReadFile("migrations/" + version + ".down.sql")
		if err != nil {
			return rolledBack, fmt.Errorf("read down migration %s: %w", version, err)
		}

		tx, err := pool.Begin(ctx)
		if err != nil {
			return rolledBack, fmt.Errorf("begin tx for down %s: %w", version, err)
		}

		if _, err := tx.Exec(ctx, string(sqlBytes)); err != nil {
			_ = tx.Rollback(ctx)
			return rolledBack, fmt.Errorf("exec down migration %s: %w", version, err)
		}

		if _, err := tx.Exec(ctx, `DELETE FROM schema_migrations WHERE version = $1`, version); err != nil {
			_ = tx.Rollback(ctx)
			return rolledBack, fmt.Errorf("delete migration record %s: %w", version, err)
		}

		if err := tx.Commit(ctx); err != nil {
			return rolledBack, fmt.Errorf("commit down migration %s: %w", version, err)
		}
		rolledBack++
	}

	return rolledBack, nil
}
