# AETHERIA: OVERDRIVE 🎮⚡
### Cross-Platform Co-op Roguelite Arena Game • Full-Stack Game Architecture
*(Tựa Game Hành Động Roguelite Co-op Đa Nền Tảng • Kiến Trúc Game Full-Stack)*

[![Unity](https://img.shields.io/badge/Client-Unity_6_/_2022_LTS-000000?style=for-the-badge&logo=unity&logoColor=white)](https://unity.com/)
[![C#](https://img.shields.io/badge/Language-C%23_12-239120?style=for-the-badge&logo=c-sharp&logoColor=white)](https://learn.microsoft.com/en-us/dotnet/csharp/)
[![Colyseus](https://img.shields.io/badge/Netcode-Colyseus_0.16-FF4081?style=for-the-badge&logo=socketdotio&logoColor=white)](https://colyseus.io/)
[![NestJS](https://img.shields.io/badge/Backend-NestJS_+_Fastify-E0234E?style=for-the-badge&logo=nestjs&logoColor=white)](https://nestjs.com/)
[![Redis](https://img.shields.io/badge/Caching_%26_Rankings-Redis_7-DC382D?style=for-the-badge&logo=redis&logoColor=white)](https://redis.io/)
[![WebGL](https://img.shields.io/badge/Platform-WebGL_+_Windows_PC-990000?style=for-the-badge&logo=webgl&logoColor=white)](https://webgl.org/)
[![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)](LICENSE)

> **AETHERIA: OVERDRIVE** là dự án game hành động roguelite góc nhìn từ trên xuống (top-down), hỗ trợ chơi Solo và Co-op nhiều người chơi thời gian thực (1-4 người). Dự án được thiết kế hoàn chỉnh theo tiêu chuẩn của một **Full-Stack Game Developer**: kết hợp giữa Client **Unity C#**, Server mạng phân quyền **Colyseus WebSocket**, và hệ thống đám mây **NestJS/Fastify + Redis**.

---

## 🧭 Kiến Trúc Hệ Thống (System Topology)

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

## 📚 Bộ Tài Liệu Thiết Kế Đầy Đủ Song Ngữ (Documentation Suite)

Dự án cung cấp đầy đủ cả hai phiên bản **Tiếng Anh (EN)** và **Tiếng Việt (VI)**:

| STT | Tài liệu | Bản Tiếng Anh (English) | Bản Tiếng Việt (Vietnamese) | Phạm vi & Nội dung |
| :---: | :--- | :---: | :---: | :--- |
| **01** | **Game Design (GDD)** | [English Doc](file:///c:/cty/aetheria-overdrive/docs/01_GDD_GAME_DESIGN_DOCUMENT.md) | [Bản Tiếng Việt](file:///c:/cty/aetheria-overdrive/docs/vi/01_GDD_TAI_LIEU_THIET_KE_GAME.md) | Cốt truyện, 3 lớp nhân vật, công thức sát thương, quái, Boss 3 phase, thẻ Perk Roguelite. |
| **02** | **Technical Arch (TDD)** | [English Doc](file:///c:/cty/aetheria-overdrive/docs/02_TDD_TECHNICAL_ARCHITECTURE.md) | [Bản Tiếng Việt](file:///c:/cty/aetheria-overdrive/docs/vi/02_TDD_KIEN_TRUC_KY_THUAT.md) | Clean Architecture Unity C#, Object Pool, Colyseus loop 30 TPS, API NestJS, tối ưu WebGL. |
| **03** | **Multiplayer Netcode** | [English Doc](file:///c:/cty/aetheria-overdrive/docs/03_MULTIPLAYER_NETCODE_SPEC.md) | [Bản Tiếng Việt](file:///c:/cty/aetheria-overdrive/docs/vi/03_NETCODE_DONG_BO_MULTIPLAYER.md) | Giao thức đồng bộ Colyseus Schema, nội suy Interpolation 100ms chống lag, chống hack vị trí. |
| **04** | **UI/UX & Wireframes** | [English Doc](file:///c:/cty/aetheria-overdrive/docs/04_UI_UX_FLOW_AND_WIREFRAMES.md) | [Bản Tiếng Việt](file:///c:/cty/aetheria-overdrive/docs/vi/04_UI_UX_LUONG_MAN_HINH_VA_WIREFRAME.md) | Sơ đồ luồng màn hình, bản vẽ wireframe HUD chiến đấu, mã phòng Co-op, nút bấm PC/Mobile. |
| **05** | **Art, Audio & Style** | [English Doc](file:///c:/cty/aetheria-overdrive/docs/05_ART_AND_AUDIO_SPEC.md) | [Bản Tiếng Việt](file:///c:/cty/aetheria-overdrive/docs/vi/05_MY_THUAT_VA_AM_THANH.md) | Phong cách Neo-Cyberpunk, bảng mã màu Token, danh mục hiệu ứng hạt VFX và âm thanh SFX/BGM. |
| **06** | **Production Roadmap** | [English Doc](file:///c:/cty/aetheria-overdrive/docs/06_PRODUCTION_ROADMAP.md) | [Bản Tiếng Việt](file:///c:/cty/aetheria-overdrive/docs/vi/06_LO_TRINH_SAN_XUAT.md) | 6 Cột mốc sản xuất (M0 -> M5), kế hoạch Sprint và tiêu chuẩn nghiệm thu Definition of Done. |

---

## 🗂️ Cấu Trúc Thư Mục Dự Án (Project Structure)

```
c:\cty\aetheria-overdrive/
├── docs/                              # Toàn bộ tài liệu thiết kế & kỹ thuật
│   ├── vi/                            # 🇻🇳 Bộ tài liệu chi tiết Tiếng Việt
│   │   ├── 01_GDD_TAI_LIEU_THIET_KE_GAME.md
│   │   ├── 02_TDD_KIEN_TRUC_KY_THUAT.md
│   │   ├── 03_NETCODE_DONG_BO_MULTIPLAYER.md
│   │   ├── 04_UI_UX_LUONG_MAN_HINH_VA_WIREFRAME.md
│   │   ├── 05_MY_THUAT_VA_AM_THANH.md
│   │   └── 06_LO_TRINH_SAN_XUAT.md
│   ├── 01_GDD_GAME_DESIGN_DOCUMENT.md # 🇺🇸 Bộ tài liệu chuẩn quốc tế Tiếng Anh
│   ├── 02_TDD_TECHNICAL_ARCHITECTURE.md
│   ├── 03_MULTIPLAYER_NETCODE_SPEC.md
│   ├── 04_UI_UX_FLOW_AND_WIREFRAMES.md
│   ├── 05_ART_AND_AUDIO_SPEC.md
│   └── 06_PRODUCTION_ROADMAP.md
├── client/                            # Mã nguồn Unity C# Client Project
├── server/                            # Mã nguồn Colyseus Netcode Server
└── backend/                           # Mã nguồn NestJS + Fastify Cloud Services
```

---

## 👨‍💻 Tác giả & Thông tin liên hệ
- **Lead Developer**: Nguyễn Quý Dương (`Blackerd`)
- **Vai trò**: Full-Stack Game Developer
- **Email**: nguyenquyduong0802@gmail.com
- **GitHub**: [github.com/Blackerd](https://github.com/Blackerd)
- **Portfolio Trực Tuyến**: [blackerd.github.io/portfolio](https://blackerd.github.io/portfolio/)
