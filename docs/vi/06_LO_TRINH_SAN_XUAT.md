# LỘ TRÌNH SẢN XUẤT & CÁC CỘT MỐC (ROADMAP)
## Dự án: Aetheria: Overdrive
**Phiên bản tài liệu**: 2.1.0 (Chiến lược Backend & Server Trước - TDD)  
**Phân loại**: Kế hoạch Tiến độ & Tiêu chuẩn Nghiệm thu  
**Chủ nhiệm dự án**: Nguyễn Quý Dương (`Blackerd`)

---

## 1. Sơ Đồ 6 Cột Mốc Sản Xuất (Milestone Schedule)

Dự án đi theo chiến lược **Backend & Game Server First** (xây dựng nền móng hệ thống, logic vật lý và bộ test chuẩn xác trước khi làm đồ họa Unity Client):

- **Cột mốc 0 (M0)**: Tiền sản xuất & Hoàn thiện tài liệu kiến trúc + `AGENTS.md` (Đã hoàn thành ✅).
- **Cột mốc 1 (M1)**: Hạ tầng Docker (PostgreSQL + Redis) & Cloud Backend API bằng Go (`backend/`).
- **Cột mốc 2 (M2)**: Dedicated Game Server bằng Go (`server/`) + Test đa luồng WS + `debug-client.html`.
- **Cột mốc 3 (M3)**: Khởi tạo Unity Client (`client/`) & Kết nối đồng bộ Netcode 2 người chơi.
- **Cột mốc 4 (M4)**: Mở rộng nội dung Game (3 Lớp nhân vật, 4 loại quái, Trùm 3 phase *Apex Overlord*, Thẻ Perk).
- **Cột mốc 5 (M5)**: Tối ưu WebGL, Nhúng trực tiếp lên Portfolio & Phát hành bản PC Windows `.exe`.

---

## 2. Chi Tiết Từng Cột Mốc & Tiêu Chuẩn Nghiệm Thu (Definition of Done - DoD)

### Cột mốc 0: Tiền sản xuất & Duyệt tài liệu kiến trúc (ĐÃ HOÀN THÀNH ✅)
- [x] Soạn thảo trọn bộ GDD, TDD, Netcode Spec, UI/UX Wireframe, Art/Audio song ngữ Anh - Việt.
- [x] Tạo `AGENTS.md` và `07_HUONG_DAN_MOI_TRUONG_VA_TRIEN_KHAI.md` giúp AI ở bất kỳ máy nào cũng hiểu đúng stack & quy trình.
- [x] Khởi tạo kho lưu trữ Git và đẩy toàn bộ bộ khung dự án lên.
- **Tiêu chuẩn nghiệm thu (DoD)**: Bộ tài liệu đầy đủ, thống nhất thứ tự triển khai Backend-first.

---

### Cột mốc 1: Hạ tầng cục bộ & Cloud Backend bằng Go (`backend/`)
- [x] Tạo `docker-compose.yml` chạy `postgres:16-alpine` và `redis:7-alpine` kèm healthcheck.
- [x] Khởi tạo Go Module (`go mod init aetheria-backend`), config `.env` và `GET /healthz`.
- [x] Viết SQL schema & migration cho bảng `players`, `meta_talents`, `match_history`.
- [ ] Lập trình đăng nhập nhanh Guest (`POST /api/v1/auth/guest`) và JWT Middleware.
- [ ] Lập trình Bảng xếp hạng Redis Sorted Sets (`ZSET`) & API nộp điểm có chữ ký bảo mật từ Server.
- [ ] Lập trình API nâng cấp Cây Kỹ Năng (Meta-Talents) dùng transaction chống race condition.
- **Tiêu chuẩn nghiệm thu (DoD)**: `go test -race ./...` vượt qua 100% test đơn vị, HTTP `httptest`, test tranh chấp đồng thời DB và test Redis.

---

### Cột mốc 2: Game Server Phân Quyền bằng Go (`server/`)
- [ ] Khởi tạo Go Module (`go mod init aetheria-server`) trong thư mục `server/`.
- [ ] Định nghĩa struct gói tin (`internal/protocol/`) kèm unit test mã hóa/giải mã.
- [ ] Xây dựng Spatial Hash Grid 2D & công thức chiến đấu (`internal/engine/`) kèm benchmark `< 2ms` cho 1.000 thực thể.
- [ ] Xây dựng `Room` Goroutine chạy vòng lặp cố định 30 Hz (`time.Ticker`), kênh `inbox`, kẹp tốc độ chống hack, wave spawner, tạm dừng chọn Perk, hồi sinh đồng đội.
- [ ] Xây dựng WebSocket Hub, mã phòng `ROOM-XXXX` trên Redis, Reconnect Token 60s và gửi điểm về `backend/`.
- [ ] Tạo file `server/debug-client.html` (Canvas 2D tối giản) để test trực quan trên trình duyệt.
- **Tiêu chuẩn nghiệm thu (DoD)**: Integration test đa client WebSocket bằng Go chạy xanh với `go test -race ./...`, và 2 tab trình duyệt cùng chơi thử mượt mà trên `debug-client.html`.

---

### Cột mốc 3: Khởi Tạo Unity Client & Kết Nối Netcode (`client/`)
- [ ] Khởi tạo project Unity 6 / 2022.3 LTS trong `client/` theo Clean Architecture MVP.
- [ ] Dựng `ServiceLocator`, `UniTask`, và `FastObjectPool<T>` (0 cấp phát rác GC).
- [ ] Lập trình `OperativeController` (WASD, xoay chuột, lướt Dash I-frame 0.2s).
- [ ] Tích hợp `GoSocketProtocol`, Client-Side Prediction và bộ đệm nội suy `Interpolator` 100ms.
- **Tiêu chuẩn nghiệm thu (DoD)**: 2 bản Unity cùng vào 1 mã phòng `ROOM-XXXX`, di chuyển, bắn và đồng bộ máu từ Go Server mượt mà ở 60 FPS.

---

### Cột mốc 4: Hoàn Thiện 3 Nhân Vật, Hệ Thống Quái & Trùm Cuối
- [ ] Hoàn thiện chiêu thức và nội tại của 3 nhân vật: Vanguard, Shadowblade, Lumina.
- [ ] Hoàn thiện 4 loại quái thường: Swarmer, Brute, Phantom, Spitter.
- [ ] Lập trình Trùm cuối **Apex Overlord** với đủ 3 giai đoạn chuyển đổi biến hóa.
- [ ] Giao diện rút 3 thẻ Perk Roguelite và vòng tròn cứu đồng đội (Revive UI).
- **Tiêu chuẩn nghiệm thu (DoD)**: Hoàn thành trọn vẹn 1 trận 10 wave ở cả Solo lẫn Co-op, điểm cuối trận đẩy thẳng lên bảng xếp hạng `backend/`.

---

### Cột mốc 5: Tối Ưu WebGL & Tích Hợp Vào Web Portfolio
- [ ] Xuất bản Unity WebGL nén Brotli, giới hạn RAM 256MB.
- [ ] Nhúng trực tiếp khung chơi game vào trang Portfolio ([blackerd.github.io/portfolio](https://blackerd.github.io/portfolio/)).
- [ ] Đóng gói thêm bản Windows PC `.exe` cho nhà tuyển dụng tải về trải nghiệm.
- **Tiêu chuẩn nghiệm thu (DoD)**: Nhà tuyển dụng mở link portfolio là chơi được ngay trong 3 giây và kết nối được server multiplayer thật.
