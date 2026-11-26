# DEVELOPMENT ENVIRONMENT & IMPLEMENTATION GUIDE
## Project: Aetheria: Overdrive
**Document Version**: 1.0.0  
**Classification**: Onboarding, Toolchain & TDD Implementation Order  
**Lead Architect**: Nguyễn Quý Dương (`Blackerd`)

---

## 1. Required Software & Toolchain Matrix

To build **Aetheria: Overdrive** cleanly across any machine without manual database installation headaches, tools are split into two phases so you only install what you need right now:

### Phase 1 & 2 Tools (Required Now — Backend & Game Server)
| Software | Version | Purpose | Notes |
| :--- | :--- | :--- | :--- |
| **Git** | Latest | Version control & repo sync | Already installed on dev machine |
| **Go (Golang)** | `1.23+` | Compiles & tests `backend/` and `server/` | Includes `go test -race` & benchmarks |
| **Docker Desktop** | Latest | Runs PostgreSQL 16 & Redis 7 in containers | Avoids polluting Windows with raw DB installs |
| **VS Code** | Latest | Primary editor for Go & Markdown | Recommended ext: `golang.go` |
| **API Client (Optional)** | `curl` / Bruno / Postman | Manual REST endpoint smoke checks | Automated tests use Go `net/http/httptest` |

### Phase 3 & 4 Tools (Needed Later — Unity Game Client)
| Software | Version | Purpose | Notes |
| :--- | :--- | :--- | :--- |
| **Unity Hub** | Latest | Manages Unity Engine versions | Install when starting Milestone 3 |
| **Unity Editor** | `Unity 6` or `2022.3 LTS` | Builds `client/` C# project | Must include **WebGL Build Support** & **Windows Build Support (IL2CPP)** modules |
| **C# IDE** | VS Code (`C# Dev Kit`) or JetBrains Rider | C# IntelliSense & Unity debugging | Optional: `Cysharp/UniTask` package inside Unity |

---

## 2. Environment Audit & Installation Commands

### Windows (PowerShell + `winget`)
Check installed tools:
```powershell
foreach ($t in @("git", "go", "docker")) {
  $c = Get-Command $t -ErrorAction SilentlyContinue
  if ($c) { "$t : FOUND ($($c.Source))" } else { "$t : MISSING" }
}
Test-Path "C:\Program Files\Unity Hub\Unity Hub.exe"
```

Install missing Phase 1 & 2 tools:
```powershell
# Install Go 1.23+
winget install GoLang.Go --accept-package-agreements --accept-source-agreements

# Install Docker Desktop (for PostgreSQL 16 + Redis 7)
winget install Docker.DockerDesktop --accept-package-agreements --accept-source-agreements
```

Install Phase 3 & 4 tools (when ready for Unity):
```powershell
winget install Unity.UnityHub --accept-package-agreements --accept-source-agreements
```

> **Post-Install Reminder**: Restart your terminal after `winget` finishes so `go` and `docker` appear in `PATH`, and launch **Docker Desktop** so the Linux engine is active.

### macOS (`Homebrew`)
```bash
brew install git go
brew install --cask docker unity-hub
```

---

## 3. Exact Step-by-Step Implementation & Testing Sequence

This project uses a **Backend & Server-First, Test-Driven** build sequence. Never write client code for a feature until the server contract and tests for that feature are green.

### Step 1: Local Infrastructure (`docker-compose.yml`)
- **What to build**:
  - `docker-compose.yml` defining `postgres:16-alpine` (`5432`) and `redis:7-alpine` (`6379`) with healthchecks and persistent volumes.
  - `.env.example` with `POSTGRES_DSN`, `REDIS_ADDR`, `JWT_SECRET`, `SERVER_INTERNAL_SECRET`.
- **What to test first**:
  - Run `docker compose up -d` and `docker compose ps` -> both services must report `(healthy)`.
  - Verify connectivity: `docker compose exec redis redis-cli ping` returns `PONG`.

### Step 2: Golang Cloud Backend (`backend/`)
Build module-by-module, verifying each with unit and HTTP integration tests before proceeding:

1. **Step 2.1 — Module Init, Config & Healthcheck**
   - *Code*: `go mod init aetheria-backend`, config struct, Postgres/Redis connection pools, `GET /healthz`.
   - *Test*: Integration test checking `/healthz` returns `200 OK` when DB and Redis are up.
2. **Step 2.2 — SQL Schema & Migrations**
   - *Code*: SQL tables `players`, `meta_talents`, `match_history`.
   - *Test*: Migration test verifying clean `up` and idempotent re-runs against PostgreSQL.
3. **Step 2.3 — Guest Auth & JWT Middleware**
   - *Code*: `POST /api/v1/auth/guest` + JWT signing & auth middleware.
   - *Test*: `httptest` cases for guest creation, returning existing player, and rejecting expired/tampered JWT (`401`).
4. **Step 2.4 — Redis Leaderboard & Signed Score Submission**
   - *Code*: `GET /api/v1/leaderboard/global` (`ZREVRANGE`), `POST /api/v1/leaderboard/submit` (protected by `SERVER_INTERNAL_SECRET` / HMAC).
   - *Test*: Unit & Redis integration tests checking composite score ordering (`wave * 1_000_000 + kills * 10 - seconds`) and rejecting unsigned requests (`403`).
5. **Step 2.5 — Meta-Progression & Core Relic Talents**
   - *Code*: `GET /api/v1/profile/meta-talents`, `POST /api/v1/profile/upgrade-talent` using SQL transactions (`SELECT ... FOR UPDATE`).
   - *Test*: Concurrent race test (`go test -race ./...`) firing parallel upgrade requests to verify Core Relic balance never goes negative.

### Step 3: Golang Dedicated Game Server (`server/`)
Build the authoritative simulation engine and verify it headless (without Unity) first:

1. **Step 3.1 — Network Protocol & State Structs**
   - *Code*: `internal/protocol/` (`PlayerState`, `EnemyState`, `ArenaRoomState`, C2S/S2C JSON/Binary packet structs).
   - *Test*: Round-trip encode/decode unit tests ensuring zero field loss.
2. **Step 3.2 — Spatial Hash Grid 2D & Combat Math**
   - *Code*: `internal/engine/` (2D spatial grid, circle/ray hit detection, damage + armor mitigation + critical hit math).
   - *Test*: Unit tests for edge collisions and Go benchmarks (`go test -bench=.`) confirming 1,000 entities simulate well within the 33.3ms tick budget (target < 2ms).
3. **Step 3.3 — Authoritative Room Loop (30 Hz Ticker)**
   - *Code*: `internal/room/` (Goroutine-per-Room, `time.Ticker` 30 Hz, channel `inbox`, anti-cheat velocity clamping, wave spawner, 3-card perk draft pause, downed/revive channeling).
   - *Test*: Deterministic step-by-step simulation tests + `go test -race ./...` ensuring zero data races across concurrent rooms.
4. **Step 3.4 — WebSocket Hub, Room Codes & Reconnect**
   - *Code*: `internal/ws/` (`ROOM-XXXX` generator + Redis room lookup, 60s dormant reconnect token, end-of-match webhook call to `backend/`).
   - *Test*: Multi-client Go WebSocket integration test spinning up 2–4 virtual clients connecting via room code, moving, shooting, taking damage, reviving, and reconnecting cleanly.
5. **Step 3.5 — Lightweight HTML5 Canvas Debug Visualizer (`server/debug-client.html`)**
   - *Code*: Single static HTML file drawing circles/rectangles on a 2D Canvas from WebSocket packets.
   - *Test*: Open 2 browser tabs, create a room in Tab 1, join with Room Code in Tab 2, and visually verify smooth co-op combat & enemy waves before touching Unity.

### Step 4: Unity Client (`client/`)
Only start once Stages 1–3 are 100% tested and stable:
1. Initialize Unity project + `ServiceLocator` + `FastObjectPool<T>` + `UniTask`.
2. Build local `OperativeController` (WASD, mouse aim, 0.2s Dash I-frames) and test offline solo feel.
3. Wire `GoSocketProtocol` + Client-Side Prediction + 100ms Snapshot `Interpolator` and verify 2 Unity instances sync cleanly against the running Go `server/`.
4. Implement 3 Classes (`Vanguard`, `Shadowblade`, `Lumina`), 4 Enemy types, `Apex Overlord` 3-phase Boss, 3-Card Draft modal, and Revive UI.
5. Export **WebGL (Brotli, 256MB initial memory)** + **Windows `.exe`** and embed into the Portfolio.
