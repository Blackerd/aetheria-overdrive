# Client Directory — Unity Engine Project

This directory contains the Unity 6 / Unity 2022 LTS C# game client project for **Aetheria: Overdrive**.

## Architecture & Conventions
- **Clean Architecture**: Decoupled Model-View-Presenter (MVP).
- **Asynchronous Workflows**: `UniTask` (Zero-allocation async/await).
- **Object Pooling**: `FastObjectPool<T>` for zero-GC combat loops.
- **Network Interface**: Colyseus WebSocket Client.
- **Build Targets**: Unity WebGL (Browser) and Windows PC Standalone (`.exe`).

See full specifications in [`docs/02_TDD_TECHNICAL_ARCHITECTURE.md`](../docs/02_TDD_TECHNICAL_ARCHITECTURE.md).
