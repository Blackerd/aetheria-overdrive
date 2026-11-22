# ĐẶC TẢ GIAO THỨC ĐỒNG BỘ MULTIPLAYER (NETCODE)
## Dự án: Aetheria: Overdrive
**Phiên bản tài liệu**: 2.0.0 (Netcode Server Chuyên Biệt Bằng Golang)  
**Phân loại**: Đặc tả Giao thức Mạng & Netcode  
**Chủ nhiệm mạng**: Nguyễn Quý Dương (`Blackerd`)

---

## 1. Mô Hình Mạng & Cơ Chế Hoạt Động

Game áp dụng mô hình **Server Phân Quyền Tuyệt Đối (Dedicated Authoritative Server)** được xây dựng hoàn toàn bằng ngôn ngữ **Go (Golang 1.23)** trên nền tảng WebSocket nhị phân an toàn (`wss://`):

- **Client (Unity C#)**: Đóng vai trò hiển thị và nhận lệnh từ bàn phím/chuột gửi lên server.
- **Server Game (Go)**: Mỗi phòng đấu chạy trên 1 **Goroutine** độc lập với vòng lặp 30 - 60 Hz. Toàn bộ va chạm, đạn bay, sinh quái, trừ máu đều do Go tính toán và đồng bộ xuống cho các máy khách.

---

## 2. Cấu Trúc Dữ Liệu Đồng Bộ Trong Go (Structs)

### 2.1 Trạng thái của từng Người chơi (PlayerState trong Go)
```go
package protocol

type PlayerState struct {
    ID              string  `json:"id"`             // ID người chơi
    Name            string  `json:"name"`           // Tên hiển thị
    OperativeClass  string  `json:"class"`          // VANGUARD, BLADE, LUMINA

    // Toạ độ & Chuyển động
    X               float32 `json:"x"`
    Y               float32 `json:"y"`
    AimAngle        float32 `json:"angle"`          // Góc ngắm súng
    IsDashing       bool    `json:"dashing"`        // Đang lướt bất tử

    // Chỉ số sinh tồn
    Health          int32   `json:"hp"`             // Máu hiện tại
    MaxHealth       int32   `json:"max_hp"`
    Shield          int32   `json:"shield"`         // Giáp năng lượng
    MaxShield       int32   `json:"max_shield"`
    UltimateEnergy  float32 `json:"ult"`            // Năng lượng chiêu cuối

    // Trạng thái tổ đội
    IsDowned        bool    `json:"downed"`         // Bị gục ngã
    ReviveProgress  float32 `json:"revive_pct"`     // % Được cứu
    Kills           int32   `json:"kills"`          // Số quái đã hạ
}
```

### 2.2 Trạng thái của toàn phòng đấu (ArenaRoomState trong Go)
```go
type ArenaRoomState struct {
    RoomID             string                  `json:"room_id"`
    RoomCode           string                  `json:"code"`       // Ví dụ: ROOM-8822
    MatchState         string                  `json:"state"`      // LOBBY, COMBAT, DRAFT, WIPED, VICTORY
    CurrentWave        int32                   `json:"wave"`       // Wave 1 - 10
    WaveTimeRemaining  float32                 `json:"time_rem"`   // Đếm ngược giây
    Players            map[string]*PlayerState `json:"players"`
    Enemies            map[string]*EnemyState  `json:"enemies"`
}
```

---

## 3. Danh Mục Các Gói Tin Mạng (Network Messages)

### 3.1 Client gửi lên Server Go (C2S)
- `c2s_input`: Gửi toạ độ di chuyển và góc ngắm bắn ở tần số 30 Hz.
- `c2s_fire`: Yêu cầu bắn súng (gửi kèm loại súng và góc bắn).
- `c2s_dash`: Yêu cầu kích hoạt lướt né bất tử theo hướng chỉ định.
- `c2s_skill`: Yêu cầu bật chiêu kích hoạt (Q) hoặc chiêu cuối (R).
- `c2s_select_perk`: Chọn 1 trong 3 thẻ kỹ năng sau khi qua màn.
- `c2s_revive_target`: Bắt đầu quá trình đứng cứu đồng đội đang bị gục ngã.

### 3.2 Server Go thông báo xuống Client (S2C)
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
   - Khi server Go gửi vị trí chuẩn về, nếu sai lệch nhỏ hơn 5cm thì client tự động sửa sai mượt mà trong 3 frame tiếp theo.
2. **Bộ đệm nội suy 100ms (Snapshot Interpolation cho đồng đội & quái)**:
   - Chuyển động của đồng đội và quái vật được lưu vào một bộ đệm trễ 100ms (tương đương 3 tick server).
   - Client liên tục dùng thuật toán nội suy Lerp giữa vị trí cũ và vị trí mới, loại bỏ hoàn toàn hiện tượng dịch chuyển giật cục (teleport).

---

## 5. Cơ Chế Tự Động Kết Nối Lại Khi Mất Mạng

- **Giữ chỗ 60 giây trong Goroutine**: Khi người chơi bị rớt mạng hoặc vô tình đổi tab trình duyệt, server Go sẽ giữ nguyên nhân vật trong phòng ở trạng thái ngủ (Dormant) trong 60 giây.
- **Token khôi phục (Reconnection Token)**: Khi vào phòng, mỗi client nhận một mã token bí mật. Khi mạng có lại, client chỉ cần gửi token này lên là lập tức vào lại đúng trận đấu mà đồng đội không bị gián đoạn.
