CREATE TABLE IF NOT EXISTS players (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    guest_token  VARCHAR(128) NOT NULL UNIQUE,
    display_name VARCHAR(32) NOT NULL,
    core_relics  INTEGER NOT NULL DEFAULT 0 CHECK (core_relics >= 0),
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS meta_talents (
    player_id  UUID NOT NULL REFERENCES players(id) ON DELETE CASCADE,
    talent_key VARCHAR(64) NOT NULL,
    level      INTEGER NOT NULL DEFAULT 1 CHECK (level >= 1),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (player_id, talent_key)
);

CREATE TABLE IF NOT EXISTS match_history (
    id               BIGSERIAL PRIMARY KEY,
    match_id         VARCHAR(64) NOT NULL UNIQUE,
    player_id        UUID NOT NULL REFERENCES players(id) ON DELETE CASCADE,
    operative_class  VARCHAR(32) NOT NULL,
    mode             VARCHAR(16) NOT NULL CHECK (mode IN ('SOLO', 'COOP')),
    wave_cleared     INTEGER NOT NULL CHECK (wave_cleared >= 0),
    total_kills      INTEGER NOT NULL CHECK (total_kills >= 0),
    duration_seconds INTEGER NOT NULL CHECK (duration_seconds >= 0),
    relics_earned    INTEGER NOT NULL DEFAULT 0 CHECK (relics_earned >= 0),
    victory          BOOLEAN NOT NULL DEFAULT FALSE,
    run_summary      JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_match_history_player_created
    ON match_history (player_id, created_at DESC);

CREATE INDEX IF NOT EXISTS idx_match_history_run_summary_gin
    ON match_history USING GIN (run_summary);
