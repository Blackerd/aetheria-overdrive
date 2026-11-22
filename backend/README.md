# Backend Directory — Cloud Game Services & API

This directory contains the **NestJS + Fastify** cloud backend engine managing identity, matchmaking, and persistence.

## Key Services
- **Auth Module**: Guest sessions & JWT player tokens.
- **Leaderboard Module**: Real-time Redis 7 Sorted Sets (`ZSET`) high scores.
- **Meta-Progression**: Persistent player skill tree & Core Relic storage.
- **Telemetry**: Prometheus metrics & CCU tracking.

See full specifications in [`docs/02_TDD_TECHNICAL_ARCHITECTURE.md`](../docs/02_TDD_TECHNICAL_ARCHITECTURE.md).
