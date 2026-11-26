# HƯỚNG DẪN MÔI TRƯỜNG PHÁT TRIỂN & THỨ TỰ TRIỂN KHAI (TDD)
## Dự án: Aetheria: Overdrive
**Phiên bản tài liệu**: 1.0.0 (Bản Tiếng Việt)  
**Phân loại**: Thiết lập Môi trường & Quy trình Code/Test Chuẩn  
**Chủ nhiệm kỹ thuật**: Nguyễn Quý Dương (`Blackerd`)

---

## 1. Danh Sách Phần Mềm Cần Thiết Theo Từng Giai Đoạn

Để tránh cài đặt nặng máy không cần thiết ngay từ đầu và đảm bảo chuyển sang máy mới không bị thiếu thư viện, công cụ được chia làm 2 giai đoạn rõ ràng:

### Giai đoạn 1 & 2 (Cần ngay bây giờ — Làm Backend & Game Server)
| Phần mềm | Phiên bản | Vai trò trong dự án | Ghi chú |
| :--- | :--- | :--- | :--- |
| **Git** | Mới nhất | Quản lý mã nguồn & đồng bộ GitHub | Đã có sẵn trên máy |
| **Go (Golang)** | `1.23+` | Viết và chạy test cho `backend/` và `server/` | Có sẵn `go test -race` kiểm tra xung đột luồng |
| **Docker Desktop** | Mới nhất | Chạy **PostgreSQL 16** và **Redis 7** bằng container | Không phải cài DB trực tiếp làm rác Windows |
| **VS Code** | Mới nhất | Trình soạn thảo code Go & Markdown | Cài thêm extension: `Go (golang.go)` |
| **Công cụ test API (Tuỳ chọn)** | `curl` / Bruno / Postman | Test tay các API REST | Test tự động đã có `net/http/httptest` trong Go |

### Giai đoạn 3 & 4 (Cài sau — Khi bắt đầu làm Game Client Unity)
| Phần mềm | Phiên bản | Vai trò trong dự án | Ghi chú |
| :--- | :--- | :--- | :--- |
| **Unity Hub** | Mới nhất | Quản lý phiên bản Unity Engine | Cài khi bước sang Cột mốc 3 |
| **Unity Editor** | `Unity 6` hoặc `2022.3 LTS` | Mở và build dự án `client/` | Tick chọn thêm module **WebGL Build Support** và **Windows Build Support (IL2CPP)** |
| **IDE C#** | VS Code (`C# Dev Kit`) hoặc Rider | Gợi ý code C# và debug Unity | Có thể dùng thêm gói `UniTask` trong Unity |

---

## 2. Kiểm Tra Máy & Lệnh Cài Đặt Nhanh

### Trên Windows (PowerShell + `winget`)
Kiểm tra máy hiện tại đang có gì:
```powershell
foreach ($t in @("git", "go", "docker")) {
  $c = Get-Command $t -ErrorAction SilentlyContinue
  if ($c) { "$t : ĐÃ CÓ ($($c.Source))" } else { "$t : CHƯA CÓ" }
}
Test-Path "C:\Program Files\Unity Hub\Unity Hub.exe"
```

Lệnh cài nhanh phần mềm cho **Giai đoạn Backend & Server**:
```powershell
# 1. Cài Go (Golang)
winget install GoLang.Go --accept-package-agreements --accept-source-agreements

# 2. Cài Docker Desktop (để chạy PostgreSQL 16 + Redis 7)
winget install Docker.DockerDesktop --accept-package-agreements --accept-source-agreements
```

Lệnh cài **Unity Hub** (chỉ cần chạy khi chuyển sang làm Client):
```powershell
winget install Unity.UnityHub --accept-package-agreements --accept-source-agreements
```

> **Lưu ý quan trọng sau khi cài:** Sau khi chạy `winget`, hãy **khởi động lại Terminal / VS Code** để máy nhận lệnh `go` và `docker`, đồng thời mở app **Docker Desktop** lên chạy nền trước khi gõ `docker compose up -d`.

### Trên macOS (`Homebrew`)
```bash
brew install git go
brew install --cask docker unity-hub
```

---

## 3. Thứ Tự Chính Xác: "Code Gì Trước → Test Gì Trước → Làm Gì Tiếp"

Dự án đi theo hướng **Backend & Server First (Xây phần nền móng trước)** và **Test-Driven** (xong bước nào kiểm tra xanh bước đó rồi mới qua bước tiếp theo).

### Bước 1: Hạ tầng Database cục bộ (`docker-compose.yml`)
- **Code trước**: Tạo file `docker-compose.yml` chạy `postgres:16-alpine` (cổng `5432`) và `redis:7-alpine` (cổng `6379`) kèm healthcheck, và file `.env.example`.
- **Test trước khi đi tiếp**:
  - Chạy `docker compose up -d` và `docker compose ps` -> cả 2 container đều báo `healthy`.
  - Chạy `docker compose exec redis redis-cli ping` -> trả về `PONG`.

### Bước 2: Cloud Backend bằng Go (`backend/`)
Code theo từng module nhỏ, module nào xong viết test ngay cho module đó:

1. **Bước 2.1 — Khởi tạo Go Module, Config & Healthcheck**
   - *Code*: `go mod init aetheria-backend`, đọc biến môi trường `.env`, kết nối Pool Postgres & Redis, tạo route `GET /healthz`.
   - *Test*: Gọi `/healthz` trả về `200 OK` khi DB và Redis đang chạy.
2. **Bước 2.2 — Bảng Database & Migration**
   - *Code*: Viết file SQL tạo bảng `players`, `meta_talents`, `match_history`.
   - *Test*: Chạy test migration `up` lên Postgres thật và kiểm tra chạy lại không lỗi (idempotent).
3. **Bước 2.3 — Đăng nhập nhanh Guest & JWT Token**
   - *Code*: `POST /api/v1/auth/guest` tạo hoặc lấy lại hồ sơ người chơi + sinh mã JWT.
   - *Test*: Dùng `httptest` test các case: tạo guest mới thành công, token sai/hết hạn trả về `401 Unauthorized`.
4. **Bước 2.4 — Bảng xếp hạng Redis & Nộp điểm chống gian lận**
   - *Code*: `GET /api/v1/leaderboard/global` (`ZREVRANGE`) và `POST /api/v1/leaderboard/submit` (bắt buộc có `SERVER_INTERNAL_SECRET` / chữ ký từ Game Server).
   - *Test*: Test công thức sắp xếp điểm (`wave * 1_000_000 + kills * 10 - seconds`) và đảm bảo người ngoài gửi điểm lậu bị chặn `403 Forbidden`.
5. **Bước 2.5 — Hệ thống Nâng cấp Cây Kỹ Năng (Meta-Talents)**
   - *Code*: `GET /api/v1/profile/meta-talents` và `POST /api/v1/profile/upgrade-talent` dùng transaction SQL (`SELECT ... FOR UPDATE`).
   - *Test*: Chạy `go test -race ./...` giả lập 2 request bấm nâng cấp cùng lúc để chứng minh số tiền Core Relic không bao giờ bị trừ âm (chống bug double-spend).

### Bước 3: Dedicated Game Server bằng Go (`server/`)
Xây dựng toàn bộ bộ não game và test chạy ổn định trước khi mở Unity:

1. **Bước 3.1 — Struct Dữ liệu & Gói tin Mạng (`internal/protocol/`)**
   - *Code*: Định nghĩa `PlayerState`, `EnemyState`, `ArenaRoomState`, các gói tin C2S/S2C.
   - *Test*: Unit test mã hóa / giải mã gói tin đảm bảo không rơi mất trường dữ liệu.
2. **Bước 3.2 — Lưới va chạm không gian & Công thức chiến đấu (`internal/engine/`)**
   - *Code*: Spatial Hash Grid 2D, kiểm tra trúng đạn, tính sát thương, trừ Giáp, tỉ lệ Chí mạng.
   - *Test*: Unit test va chạm các góc cạnh + Benchmark (`go test -bench=.`) chứng minh xử lý 1.000 thực thể chỉ mất dưới `2ms` (rất dư so với ngân sách `33.3ms` của 30 Hz).
3. **Bước 3.3 — Vòng lặp Phòng đấu 30 Hz (`internal/room/`)**
   - *Code*: Mô hình Goroutine-per-Room, `time.Ticker` 30 Hz, kênh `inbox`, kẹp tốc độ chống hack speed, sinh đợt quái, tạm dừng rút 3 thẻ Perk, gục ngã & hồi sinh đồng đội.
   - *Test*: Test từng tick giả lập (deterministic step test) + `go test -race ./...` đảm bảo không có Data Race khi nhiều người gửi input cùng lúc.
4. **Bước 3.4 — WebSocket Hub, Mã Phòng & Kết Nối Lại (`internal/ws/`)**
   - *Code*: Sinh mã phòng `ROOM-XXXX` lưu trên Redis, giữ chỗ 60s bằng Reconnect Token, kết thúc trận tự gọi `backend/` lưu điểm.
   - *Test*: Viết Integration Test tự động bằng Go mở 2–4 client ảo kết nối WebSocket vào cùng 1 mã phòng, di chuyển, bắn quái, cứu đồng đội, ngắt kết nối rồi nối lại thành công.
5. **Bước 3.5 — Công cụ Nhìn Trực Quan Trên Trình Duyệt (`server/debug-client.html`)**
   - *Code*: Một file HTML + Canvas 2D siêu nhẹ (chỉ để debug) nối thẳng vào WebSocket của `server/`.
   - *Test*: Mở 2 tab trình duyệt, Tab 1 tạo phòng, Tab 2 nhập Room Code -> tự tay điều khiển 2 chấm tròn chạy bắn quái mượt mà trước khi đụng vào Unity.

### Bước 4: Làm Game Client trên Unity (`client/`)
Chỉ bắt đầu khi Bước 1–3 đã chạy test xanh 100%:
1. Tạo project Unity, dựng `ServiceLocator`, `FastObjectPool<T>`, `UniTask`.
2. Làm `OperativeController` (WASD, xoay súng theo chuột, lướt Dash bất tử 0.2s) chạy test offline ổn định.
3. Nối `GoSocketProtocol` + Dự đoán phía Client + Bộ đệm nội suy `Interpolator` 100ms, mở 2 bản Unity kiểm tra đồng bộ qua Go Server.
4. Hoàn thiện 3 Class (`Vanguard`, `Shadowblade`, `Lumina`), 4 loại quái, Trùm 3 phase `Apex Overlord`, UI rút thẻ Perk & hồi sinh.
5. Build **WebGL (Brotli, 256MB RAM)** + **Windows `.exe`** và nhúng vào Portfolio.
