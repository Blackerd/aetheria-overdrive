# Server Directory — Authoritative Netcode Server

This directory contains the **Colyseus 0.16** authoritative multiplayer game server built on Node.js / TypeScript.

## Key Capabilities
- **Tick Rate**: Fixed 30 Hz (33.33ms) simulation loop.
- **Rooms**: `ArenaRoom` handling 1-4 operative co-op instances.
- **State Synchronization**: Colyseus binary schema delta patches.
- **Anti-Cheat**: Server-side movement velocity bounds, projectile casting, and damage mitigation.

See full specifications in [`docs/03_MULTIPLAYER_NETCODE_SPEC.md`](../docs/03_MULTIPLAYER_NETCODE_SPEC.md).
