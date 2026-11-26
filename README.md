# AETHERIA: OVERDRIVE 🎮⚡
### Cross-Platform Co-op Roguelite Arena Game • Unity & Golang Full-Stack Architecture
*(Tựa Game Hành Động Roguelite Co-op Đa Nền Tảng • Kiến Trúc Game Chuẩn Studio Unity & Golang)*

[![Unity](https://img.shields.io/badge/Client-Unity_6_/_2022_LTS-000000?style=for-the-badge&logo=unity&logoColor=white)](https://unity.com/)
[![C#](https://img.shields.io/badge/Language-C%23_12-239120?style=for-the-badge&logo=c-sharp&logoColor=white)](https://learn.microsoft.com/en-us/dotnet/csharp/)
[![Go](https://img.shields.io/badge/Server-Go_(Golang)_1.23-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![PostgreSQL](https://img.shields.io/badge/Database-PostgreSQL_16-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Redis](https://img.shields.io/badge/Caching_%26_Rankings-Redis_7-DC382D?style=for-the-badge&logo=redis&logoColor=white)](https://redis.io/)
[![WebGL](https://img.shields.io/badge/Platform-WebGL_+_Windows_PC-990000?style=for-the-badge&logo=webgl&logoColor=white)](https://webgl.org/)
[![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)](LICENSE)

> **AETHERIA: OVERDRIVE** là dự án game hành động roguelite góc nhìn từ trên xuống (top-down), hỗ trợ chơi Solo và Co-op nhiều người chơi thời gian thực (1-4 người). Dự án được thiết kế hoàn chỉnh theo tiêu chuẩn của một **Full-Stack Game Developer**: kết hợp giữa Client **Unity C#**, Server mạng phân quyền viết bằng **Go (Golang)** siêu tốc độ, và cơ sở dữ liệu **PostgreSQL + Redis**.

> 🤖 **Dành cho AI Agent / Máy mới:** Hãy đọc [`AGENTS.md`](file:///c:/cty/aetheria-overdrive/AGENTS.md) và [`docs/vi/07_HUONG_DAN_MOI_TRUONG_VA_TRIEN_KHAI.md`](file:///c:/cty/aetheria-overdrive/docs/vi/07_HUONG_DAN_MOI_TRUONG_VA_TRIEN_KHAI.md) trước tiên để nắm chính xác thứ tự cài đặt phần mềm, code & test (Backend/Server trước -> Unity Client sau).

---

## 🧭 Kiến Trúc Hệ Thống (System Topology)

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                      AETHERIA: OVERDRIVE SYSTEM TOPOLOGY                     │
└─────────────────────────────────────────────────────────────────────────────┘

       [ Unity Client (WebGL / Browser) ]       [ Unity Client (Friend / Remote) ]
                     │                                         │
                     ├──────── WebSocket Binary (30-60 Hz) ────┤
                     │                                         │
                     ▼                                         ▼
   ┌───────────────────────────────────────────────────────────────────────┐
   │            GOLANG DEDICATED GAME SERVER (Go 1.23 Concurrent Daemon)   │
   │  • Goroutine-per-Room Concurrency • Fixed 30/60 TPS Game Loop         │
   │  • Spatial Hash Grid 2D Collision • Binary Delta State Sync           │
   │  • Server Authoritative Damage    • Zero-GC Channel Message Queues    │
   └───────────────────────────────────┬───────────────────────────────────┘
                                       │ Internal gRPC / IPC
                                       ▼
   ┌───────────────────────────────────────────────────────────────────────┐
   │            GOLANG CLOUD SERVICES & API (REST Gateway)                 │
   │  • Matchmaking & Room Code Hub    • Redis Global ZSET Leaderboards    │
   │  • Guest Session & JWT Identity   • Meta-Progression Cloud Storage    │
   │  • Prometheus Telemetry & CCU     • Anti-Cheat Match Audit Verifier   │
   └───────────────────────────────────┬───────────────────────────────────┘
                                       │
                      ┌────────────────┴────────────────┐
                      ▼                                 ▼
             [ Redis 7 (In-Memory) ]        [ PostgreSQL 16 Database ]
             • Real-time Leaderboards       • User Profiles & Save Data
             • Active Room Registry         • Match History & Replay Audit
```

---

## 📚 Bộ Tài Liệu Thiết Kế Đầy Đủ Song Ngữ (Documentation Suite)

Dự án cung cấp đầy đủ cả hai phiên bản **Tiếng Anh (EN)** và **Tiếng Việt (VI)**:

| STT | Tài liệu | Bản Tiếng Anh (English) | Bản Tiếng Việt (Vietnamese) | Phạm vi & Nội dung |
| :---: | :--- | :---: | :---: | :--- |
| **00** | **AI Agent Handbook** | [`AGENTS.md`](file:///c:/cty/aetheria-overdrive/AGENTS.md) | — | Cẩm nang chuẩn cho AI khi chuyển máy mới: Hard rules, Stack, thứ tự code & kiểm tra công cụ. |
| **01** | **Game Design (GDD)** | [English Doc](file:///c:/cty/aetheria-overdrive/docs/01_GDD_GAME_DESIGN_DOCUMENT.md) | [Bản Tiếng Việt](file:///c:/cty/aetheria-overdrive/docs/vi/01_GDD_TAI_LIEU_THIET_KE_GAME.md) | Cốt truyện, 3 lớp nhân vật, công thức sát thương, quái, Boss 3 phase, thẻ Perk Roguelite. |
| **02** | **Technical Arch (TDD)** | [English Doc](file:///c:/cty/aetheria-overdrive/docs/02_TDD_TECHNICAL_ARCHITECTURE.md) | [Bản Tiếng Việt](file:///c:/cty/aetheria-overdrive/docs/vi/02_TDD_KIEN_TRUC_KY_THUAT.md) | Clean Architecture Unity C#, Go Server Goroutines, Spatial Grid, PostgreSQL + Redis, tối ưu WebGL. |
| **03** | **Multiplayer Netcode** | [English Doc](file:///c:/cty/aetheria-overdrive/docs/03_MULTIPLAYER_NETCODE_SPEC.md) | [Bản Tiếng Việt](file:///c:/cty/aetheria-overdrive/docs/vi/03_NETCODE_DONG_BO_MULTIPLAYER.md) | Giao thức đồng bộ Go WebSocket Binary, nội suy Interpolation 100ms chống lag, chống hack vị trí. |
| **04** | **UI/UX & Wireframes** | [English Doc](file:///c:/cty/aetheria-overdrive/docs/04_UI_UX_FLOW_AND_WIREFRAMES.md) | [Bản Tiếng Việt](file:///c:/cty/aetheria-overdrive/docs/vi/04_UI_UX_LUONG_MAN_HINH_VA_WIREFRAME.md) | Sơ đồ luồng màn hình, bản vẽ wireframe HUD chiến đấu, mã phòng Co-op, nút bấm PC/Mobile. |
| **05** | **Art, Audio & Style** | [English Doc](file:///c:/cty/aetheria-overdrive/docs/05_ART_AND_AUDIO_SPEC.md) | [Bản Tiếng Việt](file:///c:/cty/aetheria-overdrive/docs/vi/05_MY_THUAT_VA_AM_THANH.md) | Phong cách Neo-Cyberpunk, bảng mã màu Token, danh mục hiệu ứng hạt VFX và âm thanh SFX/BGM. |
| **06** | **Production Roadmap** | [English Doc](file:///c:/cty/aetheria-overdrive/docs/06_PRODUCTION_ROADMAP.md) | [Bản Tiếng Việt](file:///c:/cty/aetheria-overdrive/docs/vi/06_LO_TRINH_SAN_XUAT.md) | 6 Cột mốc sản xuất (M0 -> M5) theo hướng Backend/Server-First và Definition of Done. |
| **07** | **Dev Env & Impl Order**| [English Doc](file:///c:/cty/aetheria-overdrive/docs/07_DEV_ENV_AND_IMPL_GUIDE.md) | [Bản Tiếng Việt](file:///c:/cty/aetheria-overdrive/docs/vi/07_HUONG_DAN_MOI_TRUONG_VA_TRIEN_KHAI.md) | Phần mềm cần cài theo từng giai đoạn, lệnh cài trên máy mới, thứ tự Code trước -> Test trước. |

---

## 🗂️ Cấu Trúc Thư Mục Dự Án (Project Structure)

```
c:\cty\aetheria-overdrive/
├── AGENTS.md                          # 🤖 Cẩm nang tự động cho AI khi mở repo ở bất kỳ máy nào
├── docs/                              # Toàn bộ tài liệu thiết kế & kỹ thuật
│   ├── vi/                            # 🇻🇳 Bộ tài liệu chi tiết Tiếng Việt (01 -> 07)
│   ├── 01_GDD_GAME_DESIGN_DOCUMENT.md # 🇺🇸 Bộ tài liệu chuẩn quốc tế Tiếng Anh (01 -> 07)
│   └── ...
├── client/                            # Mã nguồn Unity C# Client Project (WebGL / Windows PC)
├── server/                            # Mã nguồn Golang Dedicated Game Server (Go 1.23)
└── backend/                           # Mã nguồn Golang Cloud API & Services (PostgreSQL + Redis)
```

---

## 👨‍💻 Tác giả & Thông tin liên hệ
- **Lead Developer**: Nguyễn Quý Dương (`Blackerd`)
- **Vai trò**: Full-Stack Game Developer
- **Email**: nguyenquyduong0802@gmail.com
- **GitHub**: [github.com/Blackerd](https://github.com/Blackerd)
- **Portfolio Trực Tuyến**: [blackerd.github.io/portfolio](https://blackerd.github.io/portfolio/)
