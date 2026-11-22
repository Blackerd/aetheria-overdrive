# ĐẶC TẢ MỸ THUẬT & ÂM THANH (ART & AUDIO)
## Dự án: Aetheria: Overdrive
**Phiên bản tài liệu**: 1.0.0 (Bản Tiếng Việt)  
**Phân loại**: Đặc tả Thiết kế Hình ảnh & Âm thanh  
**Chủ nhiệm mỹ thuật**: Nguyễn Quý Dương (`Blackerd`)

---

## 1. Định Hướng Phong Cách Mỹ Thuật (Art Direction)

Game lựa chọn phong cách **Neo-Cyberpunk Stylized Sci-Fi**: kết hợp giữa bối cảnh sàn đấu công nghiệp tối màu với các luồng sáng dạ quang (neon emissive) rực rỡ, giúp mắt người chơi không bị mỏi và dễ dàng phân biệt được vị trí của đạn, quái và đồng đội trong lúc chiến đấu căng thẳng.

### Quy tắc phân cấp thị giác (Visual Hierarchy):
1. **Mặt sàn đấu trường**: Tông màu tối lạnh (`#050811`), có các đường vân kim loại hình tổ ong chìm phía dưới.
2. **Nhân vật người chơi**: Màu sắc nổi bật, viền sáng rõ nét, có bóng đổ tiếp đất rõ ràng.
3. **Đạn và vùng nguy hiểm**: Luôn có độ phát sáng cao nhất (Bloom); đạn của quái luôn mang màu Đỏ Rực / Cam Cháy (`#EF4444`) để người chơi lập tức phản xạ né tránh.

---

## 2. Bảng Mã Màu Chủ Đạo (Color Palette Tokens)

| Tên mã màu | Mã Hex | Ý nghĩa & Vị trí áp dụng |
| :--- | :--- | :--- |
| **Xanh Lam Aether** | `#00F5FF` | Màu giáp năng lượng Vanguard, đạn plasma của phe người chơi, viền HUD |
| **Tím Hư Không** | `#D946EF` | Màu chiêu thức Shadowblade, các vết rách không gian, thẻ kỹ năng Sử Thi |
| **Xanh Lá Tái Sinh**| `#22C55E` | Hào quang hồi máu của Lumina, thanh máu của người chơi |
| **Đỏ Báo Động** | `#EF4444` | Đạn quái vật, tia laser của trùm Apex Overlord, trạng thái sắp chết |
| **Vàng Quá Tải** | `#F59E0B` | Số nhảy sát thương chí mạng, ngọc kinh nghiệm, thẻ kỹ năng Huyền Thoại |
| **Đen Kim Loại** | `#0B1120` | Giáp thân vỏ của nhân vật, vỏ ngoài của robot an ninh |

---

## 3. Danh Mục Hiệu Ứng Hạt (Particle VFX Manifest)

Tất cả các hiệu ứng đều được khởi tạo sẵn trong bộ đệm đối tượng (Object Pool) để không làm giảm FPS:

- **Tia lửa đầu nòng (Muzzle Flash)**: Chớp sáng hình nón kèm tia lửa bắn ra trong 60ms khi người chơi nhấn bắn.
- **Vệt đạn bay (Bullet Tracer)**: Dải ruy băng năng lượng phát sáng kéo dài phía sau viên đạn.
- **Chạm đích (Hit Impact)**: Văng ra 6-8 mảnh vụn năng lượng kèm hiệu ứng nháy trắng của quái.
- **Vỡ giáp (Shield Break)**: Nổ tung thành các mảnh kính lục giác màu xanh lam khi giáp về 0.
- **Bóng mờ lướt (Dash Trail)**: Để lại 4 bóng ma mờ phía sau nhân vật khi lướt né đòn.
- **Xoáy hút (Gravity Vortex)**: Luồng hạt xoay tròn cuốn bụi sàn và quái vật vào tâm hố đen trong 2 giây.

---

## 4. Kiến Trúc Âm Thanh & Nhạc Nền (Sound Design)

### 4.1 Danh mục âm thanh hành động (SFX)
- `Tiếng Bắn Súng`: Âm bass đanh thép kết hợp tiếng lách cách cơ khí đẩy vỏ đạn.
- `Tiếng Lướt Né`: Âm thanh rít gió chân không kèm độ trầm kéo tụt nhanh.
- `Tiếng Trúng Giáp`: Âm thanh kim loại va chạm leng keng vang vọng.
- `Tiếng Chọn Thẻ Kỹ Năng`: Hợp âm chuông ngân vang (Đô Trưởng) báo hiệu sức mạnh mới được kích hoạt.
- `Tiếng Nhịp Tim Sắp Chết`: Khi máu dưới 25%, âm thanh xung quanh bị lọc nghẹt (Low-pass Filter) và nhịp tim đập dồn dập vang lên.

### 4.2 Nhạc nền thích ứng theo nhịp đấu (Adaptive BGM)
- **Trong sảnh chờ & Rút thẻ**: Nhạc ambient điện tử không lời thư giãn (75 BPM).
- **Trong các đợt wave quái**: Nhạc Synthwave / Drum & Bass tốc độ cao (135 BPM) thúc giục nhịp độ chiến đấu.
- **Trận đấu Trùm Apex Overlord (Wave 10)**: Kết hợp giữa guitar điện bốc lửa, dàn hợp xướng hùng tráng và trống điện tử dồn dập (150 BPM) mang lại cảm giác hoành tráng tột cùng.
