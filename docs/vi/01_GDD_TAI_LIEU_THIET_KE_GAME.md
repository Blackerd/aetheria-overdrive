# TÀI LIỆU THIẾT KẾ GAME (GDD)
## Dự án: Aetheria: Overdrive
**Phiên bản tài liệu**: 1.0.0 (Bản Tiếng Việt)  
**Phân loại**: Đặc tả Thiết kế Sản xuất  
**Chủ nhiệm thiết kế**: Nguyễn Quý Dương (`Blackerd`)

---

## 1. Tóm Tắt Ý Tưởng & Khái Niệm Cốt Lõi (High Concept)

### 1.1 Khái niệm một câu (One-Sentence Pitch)
**Aetheria: Overdrive** là tựa game hành động roguelite góc nhìn từ trên xuống (top-down isometric), hỗ trợ chơi đơn và tổ đội 1 đến 4 người, nơi các chiến binh tăng cường công nghệ sinh học chiến đấu chống lại các đợt quái máy ngoài hành tinh, kết hợp chọn thẻ nâng cấp kỹ năng đột biến để kích hoạt quá tải lõi năng lượng Aetheria Core trước khi toàn bộ trạm sụp đổ.

### 1.2 Nền tảng & Đối tượng mục tiêu
- **Nền tảng chính**: Unity WebGL (chơi trực tiếp trên trình duyệt qua Web Portfolio, 0 giây cài đặt).
- **Nền tảng phụ**: Windows PC Standalone (bản cài `.exe` tối đa hóa đồ họa và FPS).
- **Đối tượng người chơi**: Fan của các tựa game *Enter the Gungeon*, *Soul Knight*, *Risk of Rain*, và *Brawl Stars*.
- **Game Engine**: Unity 6 / Unity 2022 LTS (C#).

### 1.3 Ba trụ cột thiết kế chính (Core Pillars)
1. **Chiến đấu tốc độ cao, phản hồi lực lập tức**: Điều khiển nhạy bén, đòn đánh có trọng lượng (hiệu ứng rung màn hình, số sát thương nảy lên, ngưng đọng khung hình khi chí mạng - hit freeze frame), lướt né bất tử đòi hỏi kỹ năng cao.
2. **Khả năng phối hợp đồng đội (Co-op Synergy)**: Chiêu thức của các lớp nhân vật được thiết kế để kết hợp đòn combo (Ví dụ: Vanguard hút quái lại một điểm, Shadowblade tung bão kiếm dọn sạch, Lumina kích hoạt quá tải tăng tốc bắn toàn đội).
3. **Giá trị chơi lại cao (Roguelite Draft)**: Mỗi khi dọn xong một đợt quái, người chơi được rút 1 trong 3 thẻ kỹ năng ngẫu nhiên làm biến đổi hoàn toàn phong cách chiến đấu trong màn chơi đó.

---

## 2. Bối Cảnh & Cốt Truyện

### 2.1 Căn cứ nghiên cứu Aetheria 09
Sâu dưới đáy đại dương tại trạm nghiên cứu *Aetheria Station 09*, nhân loại nỗ lực khai thác nguồn năng lượng phản vật chất bí ẩn mang tên **The Void Flux**. Một sự cố rạn nứt không gian bất ngờ đã làm mất ổn định lõi lò phản ứng, biến toàn bộ hệ thống robot an ninh thành những cỗ máy sát thủ bị tha hóa.

### 2.2 Giao thức Overdrive (The Overdrive Protocol)
Khi các cửa ngăn khẩn cấp tự động đóng chặt, AI trung tâm kích hoạt **Giao thức Overdrive**: triển khai một tiểu đội chiến binh đột kích tinh nhuệ vào buồng phản ứng. Mục tiêu duy nhất: sống sót qua 10 đợt phong tỏa liên tiếp, tiêu diệt trùm AI tha hóa **Apex Overlord** và kích hoạt thiết bị ổn định lõi bằng tay.

---

## 3. Ba Vòng Lặp Gameplay (Core Loops)

### 3.1 Vòng lặp tức thời (Micro Loop: 5 – 10 giây)
```
Đọc quỹ đạo đạn địch ──► Di chuyển & Ngắm bắn ──► Xả đạn / Tung chiêu ──► Lướt né đòn bất tử ──► Nhặt mảnh Aether
```
- Người chơi liên tục di chuyển né đạn, tận dụng 0.2 giây bất tử khi lướt (Dash I-frame) và tiêu diệt các mục tiêu nguy hiểm trước.

### 3.2 Vòng lặp từng đợt (Meso Loop: 2 – 3 phút / đợt wave)
```
Quái xuất hiện ──► Chống chọi đợt quái đông ──► Hạ gục Chỉ huy đợt ──► Dọn sạch sàn đấu ──► Bốc 1 trong 3 thẻ Perk
```
- Độ khó tăng dần sau mỗi wave: số lượng quái đông hơn, quái có thêm giáp và cạm bẫy laser trong sàn đấu.

### 3.3 Vòng lặp toàn trận (Macro Loop: 15 – 20 phút / lượt chơi)
```
Chọn Nhân vật ──► Vào Đấu trường (Solo hoặc Ghép bạn) ──► Vượt qua 10 Waves ──► Diệt Trùm Apex Overlord ──► 
Nhận Cổ vật Core Relic ──► Nâng cấp Cây Kỹ Năng vĩnh viễn ──► Leo Bảng Xếp Hạng Toàn Cầu
```

---

## 4. Ba Lớp Nhân Vật Chiến Binh (Operative Classes)

Người chơi có thể chọn 1 trong 3 nhân vật với vai trò bổ trợ chiến thuật rõ rệt:

| Chỉ số / Đặc tính | Vanguard (Đấu sĩ Thiết Giáp) | Shadowblade (Sát Thủ Cơ Động) | Lumina (Kỹ Sư Công Nghệ - Hỗ Trợ) |
| :--- | :--- | :--- | :--- |
| **Vai trò** | Đỡ đòn tuyến đầu / Khống chế quái | Sát thương dồn cực nhanh / Né tránh | Hỗ trợ hồi máu / Tăng sức mạnh đội |
| **Máu cơ bản (HP)** | 200 | 110 | 140 |
| **Giáp năng lượng (Shield)** | 100 | 50 | 80 |
| **Tốc độ chạy** | 5.5 m/s | 7.2 m/s | 6.2 m/s |
| **Vũ khí chính** | Súng săn hạng nặng (Scatter Shotgun) | Song kiếm Plasma & Phi tiêu năng lượng | Súng trường chùm tia phân rã (Beam Rifle) |
| **Chiêu kích hoạt (Q)** | **Xoáy Hấp Dẫn (Gravity Vortex)**: Hút toàn bộ quái trong phạm vi 6m vào tâm trong 2s. | **Tốc Biến Bóng Đêm (Shadow Blink)**: Dịch chuyển tức thời 8m về phía trước, chém mọi kẻ địch trên đường đi. | **Trường Hồi Phục Nanite**: Tạo vòng hào quang bán kính 5m hồi 15 HP/giây cho bản thân và đồng đội trong 6s. |
| **Chiêu cuối (R)** | **Khiên Bất Hoại (Aegis Fortress)**: Tạo khiên bất tử cho cả đội + phản lại 50% sát thương nhận vào trong 4s. | **Bão Kiếm Tử Thần (Blade Storm)**: Phóng ra 24 phi đao năng lượng tự động tìm mục tiêu xung quanh 360 độ. | **Ma Trận Quá Tải (Overcharge Matrix)**: Tăng 40% tốc độ bắn và 30% tỉ lệ chí mạng cho cả đội trong 8s. |
| **Nội tại (Passive)** | **Vỏ Giáp Động Năng**: Mỗi khi nhận sát thương giải phóng một luồng sốc điện (20 DMG, hồi chiêu 3s). | **Hành Quyết**: Gây thêm 35% sát thương lên kẻ địch còn dưới 30% máu. | **Bác Sĩ Chiến Trường**: Cứu đồng đội bị gục ngã chỉ mất 1.5s (thay vì 4.0s như bình thường). |

---

## 5. Hệ Thống Chiến Đấu & Công Thức Tính Toán

### 5.1 Công thức sát thương
$$\text{Sát Thương Thực Tế} = (\text{Sát Thương Vũ Khí} \times \text{Hệ Số Đòn Đánh}) \times (1 + \text{Tổng Buff}) \times \text{Hệ Số Chí Mạng}$$

- **Đòn đánh chí mạng cơ bản**: 5% cơ hội, gây 150% sát thương.
- **Công thức giảm trừ qua Giáp (Armor)**:
$$\text{Sát Thương Nhận Vào} = \text{Sát Thương Gốc} \times \frac{100}{100 + \text{Chỉ Số Giáp}}$$

### 5.2 Cơ chế Giáp năng lượng (Shield) và Máu (Health)
1. Đòn đánh của địch trừ vào **Giáp năng lượng** trước.
2. Khi Giáp về 0, phần sát thương dư mới trừ trực tiếp vào **Máu**.
3. Nếu không nhận sát thương trong **3.0 giây**, Giáp tự động hồi phục với tốc độ **25 điểm/giây**.
4. Máu không tự hồi phục (cần nhặt túi máu rơi ra từ quái hoặc nhờ chiêu hồi máu của Lumina).

### 5.3 Kỹ năng lướt né đòn (Invulnerability Dash)
- **Thời gian hồi**: 1.8 giây.
- **Thời gian lướt**: 0.35 giây.
- **Cửa sổ bất tử (I-frame)**: 0.20 giây ở giữa nhịp lướt. Trong lúc này nhân vật có thể xuyên qua đạn và quái nhỏ mà không bị mất máu.

---

## 6. Danh Mục Quái Vật & Trùm Cuối (Bestiary & Boss)

### 6.1 Quái vật thông thường
| Tên quái vật | Loại quái | Tốc độ | Máu | Kiểu tấn công | Mức độ nguy hiểm |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **Fly-Bot Tự Sát (Swarmer)** | Cận chiến | 6.8 m/s | 40 | Lao thành đàn 5-10 con về phía người chơi; phát nổ khi chạm (25 DMG nổ diện rộng). | Thấp |
| **Thiết Giáp Hư Không (Brute)** | Đỡ đòn | 3.5 m/s | 280 | Cầm khiên chắn đạn phía trước; húc thẳng về phía trước gây choáng và đẩy lùi. | Trung bình |
| **Xạ Thủ Bắn Tỉa (Phantom)** | Đánh xa | 4.0 m/s | 80 | Khóa mục tiêu bằng tia laser đỏ; bắn ra tia laser xuyên thấu sau 1.5s nhắm (65 DMG). | Cao |
| **Pháo Tháp Hư Không (Spitter)**| Pháo cẩu | 0 m/s (Cố định)| 180 | Bắn đạn pháo cầu vồng tạo ra vùng lửa cháy trên mặt đất trong 5s. | Cao |

### 6.2 Trùm cuối: Apex Overlord (Xuất hiện ở Wave 10)
Cỗ máy an ninh tối thượng bị tha hóa, trang bị pháo hạt kép và các tấm khiên quỹ đạo xoay tròn.

- **Giai đoạn 1 (100% - 65% Máu)**: Quét pháo súng máy 2 nòng khắp đấu trường + bắn 8 quả tên lửa mini tự đuổi người chơi gần nhất.
- **Giai đoạn 2 (65% - 30% Máu)**: Kích hoạt 3 tấm khiên xoay quanh người chặn 100% đạn trực diện; kích hoạt 4 tia laser chia sàn đấu thành 4 góc quay tròn; triệu hồi bầy Swarmer mỗi 10 giây.
- **Giai đoạn 3 (Dưới 30% Máu - Cuồng Nộ)**: Tốc độ đánh tăng 50%; viền đấu trường thu hẹp dần do khí độc; trùm gồng chiêu diệt chủng đòi hỏi cả đội phải dồn sát thương làm choáng trùm trong vòng 6 giây.

---

## 7. Hệ Thống Rút Thẻ Kỹ Năng Roguelite (Draft Cards)

Sau mỗi wave, người chơi được chọn 1 trong 3 thẻ bài ngẫu nhiên để biến đổi năng lực:
1. **Tụ Điện Quá Tải (Thường)**: +15% tốc độ bắn, +10% tốc độ bay của đạn.
2. **Nanite Hút Máu (Hiếm)**: Tiêu diệt mỗi con quái hồi 4 HP và 8 Shield.
3. **Đạn Plasma Nảy (Cực Hiếm)**: Đạn bắn trúng kẻ địch sẽ nảy sang 2 mục tiêu lân cận gây 60% sát thương.
4. **Hố Đen Sau Khi Lướt (Sử Thi)**: Lướt né để lại một hố đen mini phát nổ và hút quái lại một chỗ.
5. **Khai Mở Tiềm Năng (Huyền Thoại)**: Giảm 40% hồi chiêu cuối; khi bật chiêu cuối cả đội nhận 100% tỉ lệ chí mạng trong 5 giây.

---

## 8. Cơ Chế Co-op & Gục Ngã (Downed State)

- **Cơ chế Gục Ngã (Downed Stasis)**: Khi hết máu, nhân vật không chết hẳn mà chuyển sang trạng thái gục ngã (bò chậm với tốc độ 1.5 m/s và bắn súng lục yếu).
- **Hồi sinh đồng đội (Revive)**: Đồng đội còn sống chỉ cần đứng trong vòng tròn quanh người bị gục trong **4.0 giây** (1.5s với Lumina) để hồi sinh bạn với 50% máu.
- **Thua cuộc (Wipe)**: Nếu toàn bộ người chơi trong phòng đều bị gục ngã cùng lúc, trận đấu kết thúc thất bại.
