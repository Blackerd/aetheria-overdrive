# ĐẶC TẢ GIAO THỨC ĐỒNG BỘ MULTIPLAYER (NETCODE)
## Dự án: Aetheria: Overdrive
**Phiên bản tài liệu**: 1.0.0 (Bản Tiếng Việt)  
**Phân loại**: Đặc tả Giao thức Mạng & Netcode  
**Chủ nhiệm mạng**: Nguyễn Quý Dương (`Blackerd`)

---

## 1. Mô Hình Mạng & Cơ Chế Hoạt Động

Game áp dụng mô hình **Server Phân Quyền Tuyệt Đối (Dedicated Authoritative Server)** chạy bằng nền tảng **Colyseus 0.16** trên giao thức WebSocket an toàn (`wss://`):

- **Client**: Đóng vai trò là thiết bị hiển thị hình ảnh và thu thập phím bấm của người chơi để gửi lên server.
- **Server**: Chịu trách nhiệm hoàn toàn về việc tính toán vị trí người chơi, đạn bay, sinh quái vật, tính lượng máu bị trừ và chuyển đổi giữa các wave.

---

## 2. Cấu Trúc Dữ Liệu Đồng Bộ (Colyseus Schema)

### 2.1 Trạng thái của từng Người chơi (PlayerState)
```typescript
import { Schema, type } from "@colyseus/schema";

export class PlayerState extends Schema {
  @type("string") sessionId: string = "";           // ID phiên kết nối
  @type("string") name: string = "";                // Tên người chơi
  @type("string") operativeClass: string = "VANGUARD"; // Lớp nhân vật

  // Toạ độ & Chuyển động
  @type("number") x: number = 0;
  @type("number") y: number = 0;
  @type("number") aimAngle: number = 0;             // Góc xoay súng
  @type("boolean") isDashing: boolean = false;      // Đang lướt né đòn

  // Chỉ số sinh tồn
  @type("number") health: number = 200;             // Máu hiện tại
  @type("number") maxHealth: number = 200;
  @type("number") shield: number = 100;             // Giáp năng lượng
  @type("number") maxShield: number = 100;
  @type("number") ultimateEnergy: number = 0;       // Năng lượng chiêu cuối

  // Trạng thái tổ đội
  @type("boolean") isDowned: boolean = false;       // Bị gục ngã
  @type("number") reviveProgress: number = 0.0;     // Tiến độ được đồng đội cứu
  @type("number") kills: number = 0;                // Số quái đã hạ
}
```

### 2.2 Trạng thái của toàn phòng đấu (ArenaRoomState)
```typescript
import { MapSchema } from "@colyseus/schema";

export class ArenaRoomState extends Schema {
  @type({ map: PlayerState }) players = new MapSchema<PlayerState>();
  @type("string") matchState: string = "LOBBY";     // LOBBY, COMBAT, DRAFT, VICTORY, WIPED
  @type("number") currentWave: number = 1;          // Wave hiện tại (1 - 10)
  @type("number") waveTimeRemaining: number = 90;   // Thời gian đếm ngược của wave
}
```

---

## 3. Danh Mục Các Gói Tin Mạng (Network Messages)

### 3.1 Client gửi lên Server (C2S)
- `c2s_input`: Gửi toạ độ di chuyển và góc ngắm bắn ở tần số 30 Hz.
- `c2s_fire`: Yêu cầu bắn súng (gửi kèm loại súng và góc bắn).
- `c2s_dash`: Yêu cầu kích hoạt lướt né bất tử theo hướng chỉ định.
- `c2s_skill`: Yêu cầu bật chiêu kích hoạt (Q) hoặc chiêu cuối (R).
- `c2s_select_perk`: Chọn 1 trong 3 thẻ kỹ năng sau khi qua màn.
- `c2s_revive_target`: Bắt đầu quá trình đứng cứu đồng đội đang bị gục ngã.

### 3.2 Server thông báo xuống Client (S2C)
- `s2c_bullet_spawn`: Thông báo vị trí viên đạn mới được sinh ra để Client vẽ hiệu ứng.
- `s2c_combat_event`: Bắn tín hiệu trúng đạn / chí mạng / vỡ giáp để Client hiện số sát thương và âm thanh.
- `s2c_draft_opened`: Mở bảng chọn 3 thẻ kỹ năng cùng lúc trên màn hình của tất cả người chơi.
- `s2c_wave_complete`: Thông báo hoàn thành wave và chuẩn bị đợt quái tiếp theo.
- `s2c_mission_wiped`: Toàn bộ người chơi đều gục ngã; kích hoạt màn hình thất bại.

---

## 4. Xử Lý Giật Lag & Nội Suy Chuyển Động (Interpolation)

Để đảm bảo người chơi nhìn thấy bạn bè và quái vật di chuyển mượt mà ngay cả khi mạng có ping từ 40ms đến 150ms:

1. **Dự đoán phía Client (Client-Side Prediction cho bản thân)**: 
   - Nhân vật của chính bạn sẽ di chuyển ngay lập tức khi bạn bấm nút mà không cần đợi server trả lời.
   - Khi server gửi vị trí chuẩn về, nếu sai lệch nhỏ hơn 5cm thì client tự động sửa sai mượt mà trong 3 frame tiếp theo.
2. **Bộ đệm nội suy 100ms (Snapshot Interpolation cho đồng đội & quái)**:
   - Chuyển động của đồng đội và quái vật được lưu vào một bộ đệm trễ 100ms (tương đương 3 tick server).
   - Client liên tục dùng thuật toán nội suy Lerp giữa vị trí cũ và vị trí mới, loại bỏ hoàn toàn hiện tượng dịch chuyển giật cục (teleport).

---

## 5. Cơ Chế Tự Động Kết Nối Lại Khi Mất Mạng

- **Giữ chỗ 60 giây**: Khi người chơi bị rớt mạng hoặc vô tình đổi tab trình duyệt, server sẽ giữ nguyên nhân vật trong phòng ở trạng thái ngủ (Dormant) trong 60 giây.
- **Token khôi phục (Reconnection Token)**: Khi vào phòng, mỗi client nhận một mã token bí mật. Khi mạng có lại, client chỉ cần gửi token này lên là lập tức vào lại đúng trận đấu mà đồng đội không bị gián đoạn.
