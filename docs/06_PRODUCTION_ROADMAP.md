# PRODUCTION ROADMAP & MILESTONES
## Project: Aetheria: Overdrive
**Document Version**: 2.1.0 (Backend & Server-First Strategy)  
**Classification**: Milestone Schedule & Acceptance Criteria  
**Project Lead**: Nguyễn Quý Dương (`Blackerd`)

---

## 1. Master Development Lifecycle & Milestones

The project is structured into **6 disciplined milestones**, prioritizing **Backend & Authoritative Server foundations first** before building the Unity client:

```mermaid
gantt
    title Aetheria: Overdrive Production Timeline (Backend-First)
    dateFormat  YYYY-MM-DD
    section Pre-Production
    M0: GDD, TDD & Environment Guide        :done, m0, 2026-11-22, 1d
    section Backend & Netcode
    M1: Infra (Docker) & Go Cloud Backend   :active, m1, after m0, 3d
    M2: Go Authoritative Server & WS Tests  :m2, after m1, 4d
    section Game Client & Content
    M3: Unity Core Client & Netcode Sync    :m3, after m2, 4d
    M4: 3 Classes, Enemies & Boss Fight     :m4, after m3, 4d
    section Launch & LiveOps
    M5: WebGL Polish & Portfolio Deploy     :m5, after m4, 2d
```

---

## 2. Milestone Breakdown & Definition of Done (DoD)

### Milestone 0: Pre-Production & Architecture Sign-Off (Completed ✅)
- [x] Complete Game Design Document (GDD) with full combat formulas and class design.
- [x] Complete Technical Design Document (TDD) covering Unity C# and Go (Golang) Game Server.
- [x] Complete Multiplayer Netcode & Schema Protocol Specification in Go.
- [x] Complete UI/UX Flow, HUD wireframes, and cross-platform input mappings.
- [x] Complete Art & Audio Asset Specification with color tokens and SFX manifest.
- [x] Complete `AGENTS.md` and `07_DEV_ENV_AND_IMPL_GUIDE.md` for portable AI & machine setup.
- **DoD Acceptance**: All specifications reviewed, approved, and versioned in Git.

---

### Milestone 1: Local Infra & Golang Cloud Backend (`backend/`)
- [ ] Setup `docker-compose.yml` for PostgreSQL 16 and Redis 7 with healthchecks.
- [ ] Initialize Go module (`go mod init aetheria-backend`) in `backend/` with config & `/healthz`.
- [ ] Implement SQL schema & migrations (`players`, `meta_talents`, `match_history`).
- [ ] Implement Guest Authentication (`POST /api/v1/auth/guest`) and JWT middleware.
- [ ] Implement Redis Sorted Sets (`ZSET`) global leaderboard & server-signed score submission.
- [ ] Implement transactional Core Relic Meta-Talent upgrade endpoints.
- **DoD Acceptance**: `go test -race ./...` passes 100% of unit, HTTP (`httptest`), DB transaction concurrency, and Redis leaderboard tests.

---

### Milestone 2: Golang Authoritative Game Server (`server/`)
- [ ] Initialize Go module (`go mod init aetheria-server`) in `server/` directory.
- [ ] Implement network protocol structs (`internal/protocol/`) with round-trip tests.
- [ ] Implement 2D Spatial Hash Grid & combat math (`internal/engine/`) with Go benchmarks (< 2ms per 1,000 entities).
- [ ] Implement `Room` Goroutine with 30 Hz `time.Ticker`, channel inbox, anti-cheat velocity clamping, wave spawner, perk draft pause, and revive logic.
- [ ] Implement WebSocket Hub, `ROOM-XXXX` Redis registry, 60s reconnect tokens, and backend webhook integration.
- [ ] Build lightweight `server/debug-client.html` 2D Canvas visualizer.
- **DoD Acceptance**: Multi-client Go WebSocket integration tests pass with `go test -race ./...`, and 2 browser tabs can play together in `debug-client.html` without desync.

---

### Milestone 3: Unity Core Client & Netcode Integration (`client/`)
- [ ] Initialize Unity 6 / 2022.3 LTS project in `client/` directory with clean MVP architecture.
- [ ] Implement `ServiceLocator`, `UniTask` async workflows, and `FastObjectPool<T>` for zero-GC combat.
- [ ] Implement `OperativeController` with WASD movement, mouse aiming, dash I-frames, and 60 FPS camera follow.
- [ ] Implement `GoSocketProtocol` WebSocket client, Client-Side Prediction, and 100ms Snapshot `Interpolator`.
- **DoD Acceptance**: 2 Unity instances join the same `ROOM-XXXX`, move, shoot, and reflect server-authoritative health changes with zero GC spikes.

---

### Milestone 4: Full Game Content, Classes & Boss Encounter
- [ ] Implement 3 Operative Classes: Vanguard, Shadowblade, Lumina (Active Skills & Ultimates).
- [ ] Implement 4 regular enemy archetypes: Swarmer, Brute, Sniper, Artillery Spitter.
- [ ] Implement Wave 10 Boss: **Apex Overlord** featuring all 3 combat phases.
- [ ] Implement 3-Card Roguelite Perk Draft modal and Co-op Downed/Revive UI.
- **DoD Acceptance**: Complete 10-wave run playable from start to finish in both Solo and Co-op mode, submitting final scores to `backend/`.

---

### Milestone 5: WebGL Polish & Portfolio Integration
- [ ] Unity WebGL export optimization (Brotli compression, 256MB initial memory cap).
- [ ] Responsive full-screen HTML5 canvas integration.
- [ ] Embedded "Play Game Now" section added to portfolio website ([blackerd.github.io/portfolio](https://blackerd.github.io/portfolio/)).
- [ ] Windows PC standalone `.exe` build packaged and uploaded as downloadable release.
- **DoD Acceptance**: Game loads under 5 seconds in Google Chrome, plays smoothly, and connects to live production Go multiplayer server.
