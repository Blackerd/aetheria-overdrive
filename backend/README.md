# Backend Directory — Golang Cloud Game Services & API

This directory contains the **Go (Golang)** cloud services engine managing player authentication, matchmaking, and rankings.

## Key Services
- **Auth Module**: Guest sessions & JWT token issuance.
- **Leaderboard Module**: Real-time Redis 7 Sorted Sets (`ZSET`) high scores.
- **Persistence Layer**: PostgreSQL 16 schema storing match audit history and persistent Core Relic meta-talents.
- **Anti-Fraud Validator**: Server-side match signature verification.

See full specifications in [`docs/02_TDD_TECHNICAL_ARCHITECTURE.md`](../docs/02_TDD_TECHNICAL_ARCHITECTURE.md).
