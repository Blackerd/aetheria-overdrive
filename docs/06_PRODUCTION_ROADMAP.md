# PRODUCTION ROADMAP & MILESTONES
## Project: Aetheria: Overdrive
**Document Version**: 2.0.0 (Go & Unity Architecture)  
**Classification**: Milestone Schedule & Acceptance Criteria  
**Project Lead**: Nguyễn Quý Dương (`Blackerd`)

---

## 1. Master Development Lifecycle & Milestones

The project is structured into **6 disciplined milestones** following commercial game studio staging:

```mermaid
gantt
    title Aetheria: Overdrive Production Timeline
    dateFormat  YYYY-MM-DD
    section Pre-Production
    M0: GDD, TDD & Architecture Sign-Off    :done, m0, 2026-11-22, 1d
    section Production
    M1: Core Mechanics & Prototype           :active, m1, after m0, 3d
    M2: Golang Authoritative Netcode Server  :m2, after m1, 4d
    M3: Content, Classes & Boss Fight        :m3, after m2, 4d
    M4: Golang Cloud Gateway & Redis Leaderboard:m4, after m3, 3d
    section Launch & LiveOps
    M5: WebGL Polish & Portfolio Deploy      :m5, after m4, 2d
```

---

## 2. Milestone Breakdown & Definition of Done (DoD)

### Milestone 0: Pre-Production & Architecture Sign-Off (Current Milestone)
- [x] Complete Game Design Document (GDD) with full combat formulas and class design.
- [x] Complete Technical Design Document (TDD) covering Unity C# and Go (Golang) Game Server.
- [x] Complete Multiplayer Netcode & Schema Protocol Specification in Go.
- [x] Complete UI/UX Flow, HUD wireframes, and cross-platform input mappings.
- [x] Complete Art & Audio Asset Specification with color tokens and SFX manifest.
- [x] Project workspace initialized with Git repository and documentation suite.
- **DoD Acceptance**: All specifications reviewed, approved, and versioned in Git.

---

### Milestone 1: Core Mechanics & Solo Gameplay Prototype
- [ ] Initialize Unity 6 / 2022 LTS project in `client/` directory with clean architecture.
- [ ] Implement `CharacterController2D` with responsive WASD movement, mouse aiming, and 60 FPS camera follow.
- [ ] Implement `FastObjectPool` for bullets, hit effects, and floating damage numbers.
- [ ] Implement Primary Fire, Invulnerability Dash (0.2s I-frame window), and basic dummy targets.
- [ ] Basic single-player wave spawner delivering Swarmer Drones.
- **DoD Acceptance**: Operative can move, shoot, dash, eliminate 100 test enemies without GC allocation spikes, and maintain rock-solid 60 FPS.

---

### Milestone 2: Golang Authoritative Netcode Server (Go 1.23)
- [ ] Initialize Go module (`go mod init aetheria-server`) in `server/` directory.
- [ ] Implement `Room` with 30-60 Hz fixed simulation loop using `time.Ticker`.
- [ ] Implement Goroutine-per-room lifecycle and thread-safe channel inbox.
- [ ] Unity C# WebSocket client wrapper connecting to Go Game Server.
- [ ] Interpolation buffer smoothing remote player movements with zero jitter.
- [ ] Shareable Room Code mechanism (`ROOM-XXXX`) enabling 2-4 players to join the same instance.
- **DoD Acceptance**: 2 distinct browser tabs can join the same room, see each other move, shoot, and reflect health changes without desync.

---

### Milestone 3: Content Expansion, Operative Classes & Boss Encounter
- [ ] Implement 3 Operative Classes: Vanguard, Shadowblade, Lumina (Skills & Ultimates).
- [ ] Implement 4 regular enemy archetypes: Swarmer, Brute, Sniper, Artillery Spitter.
- [ ] Implement Wave 10 Boss: **Apex Overlord** featuring all 3 combat phases.
- [ ] Implement 3-Card Roguelite Perk Draft modal upon wave completion.
- [ ] Downed Stasis state and Co-op Teammate Revive channeling ring.
- **DoD Acceptance**: Complete 10-wave run playable from start to finish in both Solo and Co-op mode.

---

### Milestone 4: Golang Cloud Services & Meta-Progression (PostgreSQL + Redis)
- [ ] Implement REST API gateway in Go (`backend/` or unified service).
- [ ] Guest Authentication & JWT session generation.
- [ ] Redis Sorted Sets (`ZSET`) global high-score leaderboard tracking wave clears & kills.
- [ ] Match completion webhook with anti-fraud signature verification.
- [ ] Persistent Core Relic currency and Meta Talent Tree in PostgreSQL.
- **DoD Acceptance**: Match results securely post to Go backend and populate real-time global leaderboard.

---

### Milestone 5: WebGL Polish & Portfolio Integration
- [ ] Unity WebGL export optimization (Brotli compression, 256MB memory cap).
- [ ] Responsive full-screen HTML5 canvas integration.
- [ ] Embedded "Play Game Now" section added to portfolio website ([blackerd.github.io/portfolio](https://blackerd.github.io/portfolio/)).
- [ ] Windows PC standalone `.exe` build packaged and uploaded as downloadable release.
- **DoD Acceptance**: Game loads under 5 seconds in Google Chrome, plays smoothly, and connects to live production Go multiplayer server.
