# AETHERIA: OVERDRIVE 🎮⚡
### Cross-Platform Co-op Roguelite Arena Game • Full-Stack Game Architecture

[![Unity](https://img.shields.io/badge/Client-Unity_6_/_2022_LTS-000000?style=for-the-badge&logo=unity&logoColor=white)](https://unity.com/)
[![C#](https://img.shields.io/badge/Language-C%23_12-239120?style=for-the-badge&logo=c-sharp&logoColor=white)](https://learn.microsoft.com/en-us/dotnet/csharp/)
[![Colyseus](https://img.shields.io/badge/Netcode-Colyseus_0.16-FF4081?style=for-the-badge&logo=socketdotio&logoColor=white)](https://colyseus.io/)
[![NestJS](https://img.shields.io/badge/Backend-NestJS_+_Fastify-E0234E?style=for-the-badge&logo=nestjs&logoColor=white)](https://nestjs.com/)
[![Redis](https://img.shields.io/badge/Caching_%26_Rankings-Redis_7-DC382D?style=for-the-badge&logo=redis&logoColor=white)](https://redis.io/)
[![WebGL](https://img.shields.io/badge/Platform-WebGL_+_Windows_PC-990000?style=for-the-badge&logo=webgl&logoColor=white)](https://webgl.org/)
[![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)](LICENSE)

> **AETHERIA: OVERDRIVE** is a fast-paced, top-down isometric co-op roguelite arena game engineered from the ground up to demonstrate production-grade **Full-Stack Game Development**. It combines an expressive **Unity C#** game client, an **Authoritative Colyseus WebSocket Netcode Server**, and a high-concurrency **NestJS/Fastify** cloud services backend.

---

## 🧭 Project Architecture Overview

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                      AETHERIA: OVERDRIVE SYSTEM TOPOLOGY                     │
└─────────────────────────────────────────────────────────────────────────────┘

       [ Unity Client (WebGL / Windows) ]       [ Unity Client (Friend / Remote) ]
                     │                                         │
                     ├────────────── WebSocket (30 Hz) ────────┤
                     │                                         │
                     ▼                                         ▼
   ┌───────────────────────────────────────────────────────────────────────┐
   │             AUTHORITATIVE NETCODE SERVER (Colyseus / Node.js)         │
   │  • Game Loop Simulation (30 TPS)  • Collision & Hit Registration      │
   │  • State Patch Delta Sync         • Hostless Room Lifecycle           │
   │  • Anti-Cheat Validation          • Reconnection Buffer               │
   └───────────────────────────────────┬───────────────────────────────────┘
                                       │ Internal IPC / HTTP
                                       ▼
   ┌───────────────────────────────────────────────────────────────────────┐
   │             CLOUD BACKEND & SERVICES (NestJS / Fastify Daemon)        │
   │  • Matchmaking Queue Orchestrator • Global Redis Leaderboards         │
   │  • Player Session & Guest Profiles• Meta-Progression Cloud Storage    │
   │  • Telemetry & Server Diagnostics • Anti-Fraud Score Auditing         │
   └───────────────────────────────────┬───────────────────────────────────┘
                                       │
                      ┌────────────────┴────────────────┐
                      ▼                                 ▼
             [ Redis 7 (In-Memory) ]       [ PostgreSQL / SQLite WAL ]
             • Real-time Leaderboards      • User Profiles & Save Data
             • Room Registry & CCU Stats   • Match Audit Replays
```

---

## 📚 Master Documentation Suite

Every stage of development is documented according to commercial game studio standards:

| Document | File Path | Scope & Purpose |
| :--- | :--- | :--- |
| **01. Game Design Document (GDD)** | [`docs/01_GDD_GAME_DESIGN_DOCUMENT.md`](file:///c:/cty/aetheria-overdrive/docs/01_GDD_GAME_DESIGN_DOCUMENT.md) | World lore, character classes, combat formulas, wave mechanics, enemies, and meta-economy. |
| **02. Technical Architecture (TDD)** | [`docs/02_TDD_TECHNICAL_ARCHITECTURE.md`](file:///c:/cty/aetheria-overdrive/docs/02_TDD_TECHNICAL_ARCHITECTURE.md) | Unity Clean Architecture, Service Locator, UniTask, Object Pools, Colyseus loop, NestJS services. |
| **03. Multiplayer Netcode Spec** | [`docs/03_MULTIPLAYER_NETCODE_SPEC.md`](file:///c:/cty/aetheria-overdrive/docs/03_MULTIPLAYER_NETCODE_SPEC.md) | Client-Server contracts, RoomState schema, 30 TPS tick loop, interpolation buffer, anti-cheat rules. |
| **04. UI/UX Flow & Wireframes** | [`docs/04_UI_UX_FLOW_AND_WIREFRAMES.md`](file:///c:/cty/aetheria-overdrive/docs/04_UI_UX_FLOW_AND_WIREFRAMES.md) | Complete screen journey, in-combat HUD wireframes, cross-platform input mapping (PC & Touch). |
| **05. Art, Audio & Style Guide** | [`docs/05_ART_AND_AUDIO_SPEC.md`](file:///c:/cty/aetheria-overdrive/docs/05_ART_AND_AUDIO_SPEC.md) | Neo-cyberpunk visual identity, color palettes, particle VFX specs, audio chimes and SFX manifest. |
| **06. Production Roadmap** | [`docs/06_PRODUCTION_ROADMAP.md`](file:///c:/cty/aetheria-overdrive/docs/06_PRODUCTION_ROADMAP.md) | 6 Development Milestones (M0 to M5), sprint breakdown, and strict Definition of Done (DoD). |

---

## 🗂️ Repository Directory Structure

```
c:\cty\aetheria-overdrive/
├── .github/                       # CI/CD Workflows for automated build & deploy
├── docs/                          # Comprehensive Technical & Design Documentation
│   ├── 01_GDD_GAME_DESIGN_DOCUMENT.md
│   ├── 02_TDD_TECHNICAL_ARCHITECTURE.md
│   ├── 03_MULTIPLAYER_NETCODE_SPEC.md
│   ├── 04_UI_UX_FLOW_AND_WIREFRAMES.md
│   ├── 05_ART_AND_AUDIO_SPEC.md
│   └── 06_PRODUCTION_ROADMAP.md
├── client/                        # Unity C# Client Project (WebGL / Windows PC)
│   ├── Assets/
│   │   ├── _Scripts/              # Core Game Code (Core, Gameplay, Netcode, UI)
│   │   ├── Prefabs/               # Characters, Projectiles, Monsters, VFX
│   │   └── Settings/              # ScriptableObject Data Assets
│   └── ProjectSettings/
├── server/                        # Colyseus Authoritative Game Server
│   ├── src/
│   │   ├── rooms/                 # ArenaRoom, LobbyRoom logic
│   │   ├── schema/                # ArenaRoomState, PlayerState, EnemyState
│   │   └── index.ts
│   └── package.json
└── backend/                       # NestJS + Fastify Cloud Services Daemon
    ├── src/
    │   ├── modules/               # auth, leaderboard, matchmaking, telemetry
    │   └── main.ts
    └── package.json
```

---

## 🎮 Key Features for Portfolio Showcase

1. **Zero-Friction Web Playable**: Instant WebGL execution embedded directly into your developer portfolio website. Recruiter clicks *Play* and is in battle within 3 seconds.
2. **True Co-op Teamplay (2-4 Players)**: Shareable Room Codes (e.g., `ROOM-8822`) allow playtesting with friends or recruiters across remote machines.
3. **Dedicated Solo Experience**: Full single-player offline-capable game loop when network is unavailable.
4. **Hardcore Game Engineering**: High-density object pooling (1000+ bullets at 60 FPS), server-authoritative hit registration, lag-compensated interpolation buffer, and real-time Redis leaderboards.

---

## 👨‍💻 Author
- **Lead Developer**: Nguyễn Quý Dương (`Blackerd`)
- **Role**: Full-Stack Game Developer
- **Email**: nguyenquyduong0802@gmail.com
- **GitHub**: [github.com/Blackerd](https://github.com/Blackerd)
