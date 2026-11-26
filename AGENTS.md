# AGENTS.md — Portable AI Agent Operational Manual
## Project: Aetheria: Overdrive

> **READ THIS FIRST.** Every AI coding assistant (on any machine, OS, or IDE) MUST read this file before writing code, installing dependencies, or suggesting architecture changes.

---

## 1. Project Identity & Hard Rules

**Aetheria: Overdrive** is an independent, personal Full-Stack Game Engineering portfolio project by **Nguyễn Quý Dương (`Blackerd`)**.

### Approved Tech Stack (NON-NEGOTIABLE)
| Layer | Technology | Target Directory |
| :--- | :--- | :--- |
| **Game Client** | **Unity 6 / 2022.3 LTS (C# 12)** -> WebGL & Windows x64 | `client/` |
| **Game Server (Netcode)** | **Go (Golang 1.23+)** Dedicated Authoritative Server (WebSockets, Goroutine-per-Room, 30 Hz Ticker) | `server/` |
| **Cloud Backend & API** | **Go (Golang 1.23+)** REST Gateway (Guest Auth/JWT, Matchmaking, Leaderboard, Meta-Talents) | `backend/` |
| **Database & Cache** | **PostgreSQL 16** (Persistent Profiles/Talents) + **Redis 7** (Real-time ZSET Leaderboards & Room Registry) | Docker Compose (`infra/` or root `docker-compose.yml`) |

### Forbidden Technologies (HARD BAN)
- **NEVER** use or import **Colyseus**, **NestJS**, **Fastify**, or any Node.js/TypeScript game server frameworks.
- **NEVER** copy, reference, or reuse any source code from company projects in `c:\cty\sena_*` or similar external workspaces. This project must remain 100% original and independent.

---

## 2. Strict Build Order: Backend & Server First (TDD)

Follow this exact implementation order. **Do NOT skip ahead to Unity (`client/`) until Stages 1, 2, and 3 pass all automated tests.**

### Stage 1: Local Infrastructure (`docker-compose.yml`)
- **Code**: `docker-compose.yml` running `postgres:16-alpine` (port `5432`) and `redis:7-alpine` (port `6379`) with healthchecks and persistent volumes, plus `.env.example`.
- **Verify**: `docker compose up -d` -> both containers report `healthy`. Verify connection using `pg_isready` / `redis-cli ping` inside containers.

### Stage 2: Golang Cloud Backend (`backend/`)
1. **Scaffold & Config**: `go mod init aetheria-backend`, config loader (`.env`), PostgreSQL & Redis client pools.
   - *Test*: DB/Redis ping healthcheck endpoint `GET /healthz` returns `200 OK`.
2. **Database Schema & Migrations**: SQL migrations for `players`, `meta_talents`, `match_history`.
   - *Test*: Migration up/down idempotency test against local Postgres.
3. **Guest Auth & JWT**: `POST /api/v1/auth/guest` (creates/resumes guest profile + returns JWT).
   - *Test*: Unit + `httptest` suite verifying valid JWT, invalid token rejection (`401`), and persistence in Postgres.
4. **Redis Global Leaderboard & Anti-Fraud Score Submit**: `GET /api/v1/leaderboard/global` and `POST /api/v1/leaderboard/submit` (HMAC/Server Secret signed).
   - *Test*: `go test -race ./...` verifying composite score ranking (`ZADD`/`ZREVRANGE`), tie-breakers, and `403 Forbidden` on missing/invalid Server Secret.
5. **Meta-Progression & Talents**: `GET /api/v1/profile/meta-talents` and `POST /api/v1/profile/upgrade-talent`.
   - *Test*: Transactional concurrency test (`-race`) ensuring Core Relics cannot go negative on double-spend attempts.

### Stage 3: Golang Authoritative Game Server (`server/`)
1. **Protocol & State Structs**: `internal/protocol/` (`PlayerState`, `EnemyState`, `ArenaRoomState`, C2S/S2C JSON/Binary envelopes).
   - *Test*: Serialization/deserialization round-trip unit tests.
2. **Spatial Hash Grid 2D & Combat Math**: `internal/engine/` (grid insertion/query, circle/ray hit detection, armor mitigation & crit formula).
   - *Test*: Unit tests & Go benchmarks (`go test -bench=.`) confirming 1,000 entities query well under 1ms.
3. **Room Loop & State Machine**: `internal/room/` (Goroutine-per-Room, `time.Ticker` 30 Hz = 33.3ms, channel `inbox`, wave spawner, 3-card draft pause, downed/revive channeling, anti-cheat velocity clamping).
   - *Test*: Deterministic tick-stepping unit tests (no real wall-clock waits) + `go test -race ./...` stress test.
4. **WebSocket Hub, Room Code & Reconnect**: `internal/ws/` (`ROOM-XXXX` registry in Redis, 60s dormant reconnect token, integration with `backend/` score submit).
   - *Test*: Automated multi-client Go WebSocket integration test simulating 2–4 concurrent players joining by room code, moving, firing, taking damage, reviving, and reconnecting after disconnect.
5. **Browser Debug Visualizer (`server/debug-client.html`)**: Lightweight standalone HTML5 Canvas 2D page connecting over WebSocket to `server/` so developers can visually test multiplayer movement, shooting, enemies, and state sync in 2 browser tabs BEFORE opening Unity.

### Stage 4: Unity Client (`client/`)
- Only start after `backend/` and `server/` pass `go test -race ./...` and multi-client WebSocket tests.
- Build order inside Unity:
  1. `ServiceLocator`, `FastObjectPool<T>`, `UniTask` setup.
  2. Local `OperativeController` (WASD + Mouse Aim + Dash I-frames).
  3. `GoSocketProtocol` + Client-Side Prediction + 100ms Snapshot `Interpolator`.
  4. 3 Classes (`Vanguard`, `Shadowblade`, `Lumina`), 4 Enemy archetypes, `Apex Overlord` 3-phase Boss, 3-Card Draft UI, Revive UI.
  5. WebGL Brotli build (256MB initial memory) + Windows Standalone `.exe`.

---

## 3. Environment Verification & Auto-Remediation Protocol

When starting a session on a new machine, ALWAYS check required tools first and report what is missing before running build commands.

### Quick Check Command (Windows PowerShell)
```powershell
foreach ($t in @("git", "go", "docker", "node")) {
  $c = Get-Command $t -ErrorAction SilentlyContinue
  if ($c) { "$t : OK ($($c.Source))" } else { "$t : MISSING" }
}
Test-Path "C:\Program Files\Unity Hub\Unity Hub.exe"
```

### Quick Check Command (macOS / Linux bash)
```bash
for t in git go docker; do
  command -v "$t" >/dev/null 2>&1 && echo "$t : OK ($(command -v $t))" || echo "$t : MISSING"
done
```

### Auto-Remediation Install Matrix
| Tool | Purpose | Windows (`winget`) | macOS (`brew`) | Linux (Ubuntu/Debian) |
| :--- | :--- | :--- | :--- | :--- |
| **Git** | Source control | `winget install Git.Git` | `brew install git` | `sudo apt install -y git` |
| **Go 1.23+** | `server/` & `backend/` | `winget install GoLang.Go` | `brew install go` | Download tarball from `go.dev/dl/` |
| **Docker Desktop** | Local PostgreSQL 16 & Redis 7 | `winget install Docker.DockerDesktop` | `brew install --cask docker` | `curl -fsSL https://get.docker.com \| sh` |
| **Unity Hub** | `client/` (Stage 4 only) | `winget install Unity.UnityHub` | `brew install --cask unity-hub` | Install Unity Hub AppImage / apt repo |

> **Note for AI Agents**: After installing `GoLang.Go` or `Docker.DockerDesktop` on Windows, remind the user to restart their terminal (or refresh `$env:Path`) and launch Docker Desktop once so the Docker daemon is running before executing `docker compose up -d`.

---

## 4. Master Documentation Index

- **GDD (Game Design)**: `docs/01_GDD_GAME_DESIGN_DOCUMENT.md` | `docs/vi/01_GDD_TAI_LIEU_THIET_KE_GAME.md`
- **TDD (Technical Architecture)**: `docs/02_TDD_TECHNICAL_ARCHITECTURE.md` | `docs/vi/02_TDD_KIEN_TRUC_KY_THUAT.md`
- **Netcode Spec**: `docs/03_MULTIPLAYER_NETCODE_SPEC.md` | `docs/vi/03_NETCODE_DONG_BO_MULTIPLAYER.md`
- **UI/UX & Wireframes**: `docs/04_UI_UX_FLOW_AND_WIREFRAMES.md` | `docs/vi/04_UI_UX_LUONG_MAN_HINH_VA_WIREFRAME.md`
- **Art & Audio Guide**: `docs/05_ART_AND_AUDIO_SPEC.md` | `docs/vi/05_MY_THUAT_VA_AM_THANH.md`
- **Production Roadmap**: `docs/06_PRODUCTION_ROADMAP.md` | `docs/vi/06_LO_TRINH_SAN_XUAT.md`
- **Dev Environment & Implementation Order**: `docs/07_DEV_ENV_AND_IMPL_GUIDE.md` | `docs/vi/07_HUONG_DAN_MOI_TRUONG_VA_TRIEN_KHAI.md`
