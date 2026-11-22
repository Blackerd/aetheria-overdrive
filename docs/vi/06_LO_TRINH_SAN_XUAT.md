# LỘ TRÌNH SẢN XUẤT & CÁC CỘT MỐC (ROADMAP)
## Dự án: Aetheria: Overdrive
**Phiên bản tài liệu**: 1.0.0 (Bản Tiếng Việt)  
**Phân loại**: Kế hoạch Tiến độ & Tiêu chuẩn Nghiệm thu  
**Chủ nhiệm dự án**: Nguyễn Quý Dương (`Blackerd`)

---

## 1. Sơ Đồ 6 Cột Mốc Sản Xuất (Milestone Schedule)

Dự án được chia thành 6 giai đoạn phát triển có kiểm soát chặt chẽ theo mô hình Game Studio chuyên nghiệp:

- **Cột mốc 0 (M0)**: Tiền sản xuất & Hoàn thiện toàn bộ tài liệu thiết kế (Đã hoàn thành ✅).
- **Cột mốc 1 (M1)**: Bản chơi thử cơ chế cốt lõi (Core Mechanics & Solo Prototype).
- **Cột mốc 2 (M2)**: Ghép nối Netcode đồng bộ phòng đấu nhiều người chơi (Colyseus Multiplayer).
- **Cột mốc 3 (M3)**: Mở rộng nội dung (3 Lớp nhân vật, 4 loại quái & Trùm cuối Apex Overlord).
- **Cột mốc 4 (M4)**: Dịch vụ đám mây, Bảng xếp hạng Redis & Cây kỹ năng vĩnh viễn (NestJS Backend).
- **Cột mốc 5 (M5)**: Tối ưu WebGL, Nhúng trực tiếp lên Portfolio & Phát hành bản PC Windows `.exe`.

---

## 2. Chi Tiết Từng Cột Mốc & Tiêu Chuẩn Nghiệm Thu (Definition of Done - DoD)

### Cột mốc 0: Tiền sản xuất & Duyệt tài liệu kiến trúc (ĐÃ HOÀN THÀNH ✅)
- [x] Soạn thảo toàn bộ Tài liệu thiết kế Game (GDD) song ngữ Anh - Việt.
- [x] Soạn thảo Tài liệu kiến trúc kỹ thuật (TDD) mô tả hệ thống 3 tầng.
- [x] Soạn thảo Đặc tả giao thức mạng Netcode Colyseus và cơ chế chống lag.
- [x] Soạn thảo Luồng giao diện UI/UX và bản vẽ wireframe HUD.
- [x] Soạn thảo Quy chuẩn mỹ thuật, mã màu và âm thanh.
- [x] Khởi tạo kho lưu trữ Git và đẩy toàn bộ bộ khung dự án lên.
- **Tiêu chuẩn nghiệm thu (DoD)**: Bộ tài liệu đầy đủ, không thiếu sót công đoạn nào, sẵn sàng cho việc triển khai code.

---

### Cột mốc 1: Bản chơi thử cơ chế di chuyển & bắn súng (M1)
- [ ] Khởi tạo project Unity 6 / 2022 LTS trong thư mục `client/`.
- [ ] Lập trình di chuyển mượt mà 8 hướng qua cụm phím WASD, xoay súng theo chuột.
- [ ] Lập trình cơ chế Lướt né bất tử (Dash I-frame 0.2s).
- [ ] Xây dựng bộ đệm Object Pool quản lý đạn bắn và hiệu ứng nổ.
- [ ] Bộ tạo quái cơ bản sinh ra đàn quái Swarmer Drone đuổi theo người chơi.
- **Tiêu chuẩn nghiệm thu (DoD)**: Người chơi có thể di chuyển, bắn đạn, lướt né và tiêu diệt 100 quái thử nghiệm mà game giữ vững 60 FPS, không bị drop khung hình do rác bộ nhớ.

---

### Cột mốc 2: Ghép nối Netcode Colyseus nhiều người chơi (M2)
- [ ] Khởi tạo Server Colyseus TypeScript trong thư mục `server/`.
- [ ] Xây dựng vòng lặp mô phỏng vật lý cố định 30 Hz trên server.
- [ ] Tích hợp thư viện Colyseus Client vào Unity qua WebSocket.
- [ ] Lập trình cơ chế tạo phòng và vào phòng bằng Mã Code (`ROOM-XXXX`).
- [ ] Thuật toán nội suy Interpolation làm mượt chuyển động của đồng đội.
- **Tiêu chuẩn nghiệm thu (DoD)**: 2 trình duyệt web mở 2 tab khác nhau cùng vào 1 phòng, nhìn thấy nhau di chuyển, bắn súng và trừ máu mượt mà không bị lệch đồng bộ.

---

### Cột mốc 3: Hoàn thiện 3 Nhân Vật, Hệ Thống Quái & Trùm Cuối (M3)
- [ ] Hoàn thiện chiêu thức và nội tại của 3 nhân vật: Vanguard, Shadowblade, Lumina.
- [ ] Hoàn thiện 4 loại quái thường: Swarmer, Brute, Phantom, Spitter.
- [ ] Lập trình Trùm cuối **Apex Overlord** với đủ 3 giai đoạn chuyển đổi biến hóa.
- [ ] Hệ thống rút 3 thẻ Perk Roguelite sau mỗi wave.
- [ ] Cơ chế gục ngã và đồng đội đứng cứu hồi sinh.
- **Tiêu chuẩn nghiệm thu (DoD)**: Hoàn thành trọn vẹn một trận đấu 10 wave từ đầu đến cuối ở cả chế độ Chơi Đơn lẫn Chơi Chung Đội.

---

### Cột mốc 4: Dịch Vụ Đám Mây & Bảng Xếp Hạng (NestJS + Redis) (M4)
- [ ] Khởi tạo backend NestJS + Fastify trong thư mục `backend/`.
- [ ] Hệ thống cấp phiên chơi nhanh (Guest Login) và mã Token bảo mật.
- [ ] Bảng xếp hạng toàn cầu thời gian thực lưu trên Redis Sorted Sets (`ZSET`).
- [ ] Cơ chế kiểm tra tính hợp lệ của điểm số chống hack gửi điểm giả.
- [ ] Cây nâng cấp kỹ năng vĩnh viễn (Meta-Talent Tree).
- **Tiêu chuẩn nghiệm thu (DoD)**: Điểm số sau trận đấu tự động gửi lên đám mây và người chơi thấy tên mình xuất hiện ngay trên Bảng Xếp Hạng Top 100.

---

### Cột mốc 5: Tối Ưu WebGL & Tích Hợp Vào Web Portfolio (M5)
- [ ] Xuất bản Unity WebGL nén Brotli, giới hạn bộ nhớ RAM 256MB.
- [ ] Nhúng trực tiếp khung chơi game vào trang Portfolio cá nhân ([blackerd.github.io/portfolio](https://blackerd.github.io/portfolio/)).
- [ ] Đóng gói thêm bản cài đặt Windows PC `.exe` để người tuyển dụng có thể tải về nếu muốn trải nghiệm tối đa đồ họa.
- **Tiêu chuẩn nghiệm thu (DoD)**: Nhà tuyển dụng truy cập link portfolio là có thể bấm chơi trực tiếp trên web trong 3 giây.
