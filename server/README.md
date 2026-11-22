# Server Directory — Golang Dedicated Game Server

This directory contains the **Golang (Go 1.23+)** authoritative multiplayer game server for **Aetheria: Overdrive**.

## Key Architecture
- **Goroutine-per-Room Concurrency**: Every battle instance runs inside an isolated, lightweight Goroutine.
- **Fixed Simulation Tick**: 30 - 60 Hz physics & hit-registration loop using Go's `time.Ticker`.
- **Low-Overhead WebSockets**: Native binary state synchronization with minimal packet overhead.
- **Spatial Hash Grid (2D)**: Ultra-fast projectile and enemy collision checking in pure Go.

See full specifications in [`docs/02_TDD_TECHNICAL_ARCHITECTURE.md`](../docs/02_TDD_TECHNICAL_ARCHITECTURE.md) and [`docs/03_MULTIPLAYER_NETCODE_SPEC.md`](../docs/03_MULTIPLAYER_NETCODE_SPEC.md).
