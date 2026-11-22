# THIẾT KẾ GIAO DIỆN & TRẢI NGHIỆM NGƯỜI DÙNG (UI/UX)
## Dự án: Aetheria: Overdrive
**Phiên bản tài liệu**: 1.0.0 (Bản Tiếng Việt)  
**Phân loại**: Đặc tả Luồng Giao Diện & Trải Nghiệm  
**Chủ nhiệm UI/UX**: Nguyễn Quý Dương (`Blackerd`)

---

## 1. Sơ Đồ Luồng Màn Hình Toàn Diện (Master Flow)

Hành trình người chơi được thiết kế tối giản hoá số lần click chuột để người xem / nhà tuyển dụng có thể vào trận đấu chỉ sau 3 giây:

```
[01. Màn hình Bắt Đầu]
        │
        ├─► [Bấm Chơi Đơn (Solo)] ───────────────┐
        │                                        │
        └─► [Bấm Tạo / Vào Phòng Co-op]          ▼
                 │                      [03. Chọn Nhân Vật]
                 ▼                      (Vanguard/Blade/Lumina)
        [02. Phòng Đợi Nhóm]                     │
        (Chia sẻ Mã Code: ROOM-8822)             │
                 │                               │
                 └───────────────────────────────┼────────► [04. Sàn Đấu Chiến Trận (HUD)]
                                                                    │
                                                                    ├─► Qua Wave ──► [05. Bốc 3 Thẻ Kỹ Năng]
                                                                    │
                                                                    ├─► Hết Máu ──► [06. Màn Hình Thất Bại]
                                                                    │
                                                                    └─► Diệt Boss ──► [07. Chiến Thắng & Bảng Xếp Hạng]
```

---

## 2. Bản Vẽ Giao Diện Sàn Đấu (In-Combat HUD Wireframe)

Giao diện được thiết kế theo phong cách khí tài công nghệ quân sự tương lai, đặt ở các góc viền để không che tầm nhìn chiến trận:

```
┌────────────────────────────────────────────────────────────────────────┐
│ [WAVE 04/10] ── ⏱️ CÒN LẠI: 01:42                   ĐIỂM: 142,500 ── 🏆  │
│ ────────────────────────────────────────────────────────────────────── │
│                                                                        │
│ ┌─ DANH SÁCH ĐỒNG ĐỘI ────┐                         ┌─ BẢN ĐỒ CON ───┐ │
│ │ P1: Bạn (Vanguard)      │                         │  ● Vị trí Boss │ │
│ │ HP [██████████] 180/200 │                         │  ▲ P1   ▲ P2   │ │
│ │ SH [████████░░]  80/100 │                         │  ■ Đàn Quái    │ │
│ │                         │                         └────────────────┘ │
│ │ P2: Alex (Shadowblade)  │                                            │
│ │ HP [██████░░░░] 60/110  │                                            │
│ │ SH [░░░░░░░░░░]  0/50   │                                            │
│ └─────────────────────────┘                                            │
│                                                                        │
│                                                                        │
│                                Nhân Vật                                │
│                                   ▲                                    │
│                                                                        │
│ ────────────────────────────────────────────────────────────────────── │
│  MÁU: [█████████████████░░] 180/200       GIÁP: [████████░░] 80/100    │
│                                                                        │
│  [LƯỚT (SPACE)]     [CHIÊU Q: HÚT QUÁI]       [CHIÊU CUỐI R: KHIÊN]    │
│    SẴN SÀNG          ██████░░░░ (3.2s)          [████████████] SẴN SÀNG│
└────────────────────────────────────────────────────────────────────────┘
```

---

## 3. Bảng Chọn 3 Thẻ Nâng Cấp Roguelite (Draft Modal)

Mỗi khi dọn xong 1 đợt wave, màn hình dừng lại 15 giây cho phép từng thành viên tự chọn lối đột biến sức mạnh cho riêng mình:

```
┌────────────────────────────────────────────────────────────────────────┐
│                   === HOÀN THÀNH WAVE // CHỌN THẺ PERK ===             │
│                          THỜI GIAN CÒN LẠI: 12 GIÂY                    │
│                                                                        │
│  ┌───────────────────────┐ ┌───────────────────────┐ ┌───────────────┐ │
│  │   [THẺ PHỔ THÔNG]     │ │     [THẺ HIẾM]        │ │ [THẺ SỬ THI]  │ │
│  │                       │ │                       │ │               │ │
│  │   ⚡ TỤ ĐIỆN QUÁ TẢI   │ │   🌀 ĐẠN PLASMA NẢY   │ │ 🌌 HỐ ĐEN     │ │
│  │                       │ │                       │ │               │ │
│  │  +20% Tốc độ bắn      │ │  Đạn bắn trúng nảy    │ │ Lướt né tạo   │ │
│  │  +15% Tốc độ bay đạn  │ │  sang 2 mục tiêu lân  │ │ ra hố đen hút │ │
│  │                       │ │  cận gây 60% dame.    │ │ quái lại.     │ │
│  │                       │ │                       │ │               │ │
│  │      [ LỰA CHỌN ]     │ │      [ LỰA CHỌN ]     │ │  [ LỰA CHỌN ] │ │
│  └───────────────────────┘ └───────────────────────┘ └───────────────┘ │
│                                                                        │
│  [Đồng đội Alex đã chọn: ĐẠN PLASMA NẢY]                               │
└────────────────────────────────────────────────────────────────────────┘
```

---

## 4. Bảng Phím Bấm Đa Nền Tảng (Input Mapping)

Game tự động nhận diện thiết bị của người chơi:

| Thao tác trong game | Chơi trên Máy tính (PC / Laptop) | Chơi trên Web Cảm ứng (Mobile / Tablet) | Tay cầm (Xbox / PS) |
| :--- | :--- | :--- | :--- |
| **Di chuyển nhân vật** | Phím `W`, `A`, `S`, `D` | Cần điều khiển ảo bên Trái (Left Stick) | Cần gạt Analog Trái |
| **Ngắm hướng bắn** | Di chuột theo góc 360° | Cần điều khiển ảo bên Phải (Right Stick)| Cần gạt Analog Phải |
| **Bắn súng** | Nhấn giữ Chuột Trái | Tự động bắn khi kéo cần ngắm / Nút bấm | Cò bấm Phải (`RT` / `R2`) |
| **Lướt né đòn** | Phím cách (`Spacebar`) | Nút lướt tròn bên phải | Nút `A` / `L1` |
| **Tung chiêu (Q)** | Phím `Q` | Nút kỹ năng 1 | Nút `LB` / `L2` |
| **Chiêu cuối (R)** | Phím `R` | Nút kỹ năng Ultimate phát sáng | Nút `Y` / `R1` |
| **Cứu đồng đội** | Phím `E` (Đứng trong vòng) | Tự động kích hoạt khi lại gần | Nút `X` / `Vuông` |

---

## 5. Phản Hồi Lực & Cảm Giác Đánh (Juice & Game Feel)

1. **Hiệu ứng Khựng Đòn (Hit Freeze Frame)**: Khi trúng đòn chí mạng hoặc đánh vỡ giáp quái, khung hình dừng lại 30 miligiây để tạo cảm giác cú đánh cực kỳ uy lực.
2. **Rung màn hình (Screen Shake)**: Rung mạnh theo tỉ lệ sát thương nhận vào; nếu đồng đội ở xa bắn đạn lớn thì rung nhẹ dần theo khoảng cách.
3. **Màu số nhảy (Damage Popup Numbers)**:
   - **Số Trắng**: Sát thương thường.
   - **Số Vàng To Nảy Lên**: Đòn chí mạng (Critical Strike).
   - **Số Xanh Lam**: Sát thương trừ vào giáp năng lượng.
   - **Số Xanh Lá**: Hồi phục máu.
4. **Nháy sáng khi trúng đòn (Hit Flash)**: Quái vật nhấp nháy màu trắng tuyền trong 60 miligiây khi trúng đạn.
