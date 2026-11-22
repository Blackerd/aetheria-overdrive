# TÀI LIỆU THIẾT KẾ KIẾN TRÚC KỸ THUẬT (TDD)
## Dự án: Aetheria: Overdrive
**Phiên bản tài liệu**: 1.0.0 (Bản Tiếng Việt)  
**Phân loại**: Đặc tả Hệ thống & Kỹ thuật Phần mềm  
**Chủ nhiệm kỹ thuật**: Nguyễn Quý Dương (`Blackerd`)

---

## 1. Kiến Trúc Hệ Thống Tổng Thể 3 Tầng

**Aetheria: Overdrive** tuân thủ mô hình kiến trúc phân tách 3 tầng độc lập chuẩn công nghiệp:

```
┌────────────────────────────────────────────────────────────────────────┐
│ 1. TẦNG CLIENT (Unity 6 / C# 12)                                       │
│    • Render đồ họa 60 FPS, Hệ thống Camera Follow mượt mà             │
│    • Xử lý Input (Chuột + Phím trên PC / Phím ảo Touch trên Web/Mobile)│
│    • Quản lý Animation State Machine, Particle System VFX              │
│    • Bộ đệm Object Pool quản lý hàng nghìn viên đạn không gây giật rác │
├────────────────────────────────────────────────────────────────────────┤
│ 2. TẦNG MẠNG ĐỘC LẬP / NETCODE (Colyseus 0.16 + Node.js 22)            │
│    • Server Authoritative: Mô phỏng toàn bộ logic game trên server     │
│    • Vòng lặp vật lý cố định 30 Hz (33.33ms mỗi tick)                  │
│    • Quản lý vòng đời phòng đấu (ArenaRoom Lifecycle: 1 - 4 người)     │
│    • Đồng bộ trạng thái chênh lệch (Schema Delta Patch Sync)           │
├────────────────────────────────────────────────────────────────────────┤
│ 3. TẦNG DỊCH VỤ ĐÁM MÂY (NestJS + Fastify Engine)                      │
│    • Cổng API RESTful xác thực người chơi và phiên Guest               │
│    • Bảng xếp hạng thời gian thực siêu tốc qua Redis 7 (Sorted Sets)   │
│    • Lưu trữ dữ liệu nâng cấp cây kỹ năng (Meta-Progression)           │
│    • Thu thập số liệu giám sát hệ thống (Telemetry / CCU Metrics)      │
└────────────────────────────────────────────────────────────────────────┘
```

---

## 2. Kiến Trúc Chi Tiết Client Unity (C#)

### 2.1 Các mẫu thiết kế (Design Patterns) áp dụng
1. **Model-View-Presenter (MVP)**: Giao diện UI hoàn toàn tách rời logic gameplay. UI View chỉ hiển thị dữ liệu và gửi sự kiện người dùng đến Presenter.
2. **Service Locator**: Các hệ thống toàn cục (Âm thanh, Mạng, Tải tài nguyên, Input) được đăng ký và truy xuất tập trung thông qua `ServiceLocator`.
3. **UniTask (Xử lý bất đồng bộ)**: Loại bỏ hoàn toàn Coroutine truyền thống của Unity để tránh việc cấp phát rác bộ nhớ (GC Allocation); sử dụng `UniTask` với `CancellationToken` an toàn.
4. **ScriptableObject Event Channels**: Các module độc lập phát và nhận sự kiện qua các file asset ScriptableObject, giúp code không bị phụ thuộc lồng chéo (loose coupling).
5. **High-Performance Object Pooling**: Tạo sẵn các hồ đệm đối tượng cho đạn, hiệu ứng nổ và chữ số sát thương trước khi trận đấu bắt đầu.

### 2.2 Cấu trúc thư mục mã nguồn Unity
```
Assets/
├── _Scripts/
│   ├── Core/
│   │   ├── ServiceLocator.cs       # Quản lý dependency toàn cục
│   │   ├── GameManager.cs          # Quản lý vòng đời trận đấu
│   │   └── ObjectPool/             # Hệ thống hồ chứa đối tượng
│   ├── Gameplay/
│   │   ├── Characters/             # Controller di chuyển và bắn
│   │   ├── Combat/                 # Logic đạn, hitbox, sát thương
│   │   ├── Enemies/                # AI quái và trùm Apex Overlord
│   │   └── WaveSystem/             # Bộ điều phối đợt quái
│   ├── Netcode/
│   │   ├── NetworkManager.cs       # Khởi tạo kết nối Colyseus
│   │   └── Interpolator.cs         # Làm mượt chuyển động người chơi khác
│   └── UI/
│       ├── HUD/                    # Thanh máu, năng lượng, đồng đội
│       └── Draft/                  # Giao diện chọn 3 thẻ kỹ năng
```

### 2.3 Bộ đệm đối tượng không rác bộ nhớ (Zero-Allocation Pool)
```csharp
// Bộ đệm đối tượng tốc độ cao không cấp phát heap trong khi bắn
public interface IPoolable {
    void OnSpawnFromPool();
    void OnReturnToPool();
}

public class FastObjectPool<T> where T : Component, IPoolable {
    private readonly Queue<T> _pool = new();
    private readonly T _prefab;
    private readonly Transform _parent;

    public FastObjectPool(T prefab, int initialCapacity, Transform parent = null) {
        _prefab = prefab;
        _parent = parent;
        for (int i = 0; i < initialCapacity; i++) {
            T obj = Object.Instantiate(_prefab, _parent);
            obj.gameObject.SetActive(false);
            _pool.Enqueue(obj);
        }
    }

    public T Rent(Vector3 position, Quaternion rotation) {
        T obj = _pool.Count > 0 ? _pool.Dequeue() : Object.Instantiate(_prefab, _parent);
        Transform t = obj.transform;
        t.position = position;
        t.rotation = rotation;
        obj.gameObject.SetActive(true);
        obj.OnSpawnFromPool();
        return obj;
    }

    public void Return(T obj) {
        obj.OnReturnToPool();
        obj.gameObject.SetActive(false);
        _pool.Enqueue(obj);
    }
}
```

---

## 3. Server Netcode Phân Quyền (Colyseus Server)

### 3.1 Vòng lặp mô phỏng 30 Hz (33.33ms)
Server chạy một vòng lặp logic cố định độc lập với tốc độ khung hình client:
1. Tiếp nhận và gom các gói tin di chuyển / bắn súng từ người chơi gửi lên.
2. Kiểm tra tính hợp lệ của vị trí (chống hack tốc độ di chuyển và xuyên tường).
3. Di chuyển vị trí của đạn và tính toán va chạm theo thuật toán lưới ô không gian (Spatial Hash Grid).
4. Cập nhật AI quái vật đuổi theo người chơi gần nhất.
5. Kiểm tra điều kiện hết quái để kích hoạt sự kiện rút thẻ Perk.
6. Đóng gói các trường dữ liệu có thay đổi thành dạng nhị phân gửi về cho client.

### 3.2 Quy tắc chống gian lận (Anti-Cheat Rules)
- **Kiểm soát vị trí**: Tốc độ di chuyển của người chơi được giới hạn bởi `TốcĐộTốiĐa * ThờiGianTick * 1.15`. Nếu người chơi cố tình chỉnh sửa memory gửi vị trí vượt quá khoảng này, server tự động kéo giật người chơi về vị trí hợp lệ.
- **Quyền sinh đạn**: Client chỉ được gửi yêu cầu *"Tôi muốn bắn theo góc X"*. Server tự kiểm tra thời gian hồi chiêu rồi mới sinh ra viên đạn thật.
- **Quyền tính máu**: Client tuyệt đối không được tự gửi thông điệp *"Tôi đã trừ quái 100 máu"*. Toàn bộ việc trừ máu do Server phán quyết.

---

## 4. Dịch Vụ Đám Mây & Bảng Xếp Hạng (NestJS + Fastify)

### 4.1 Danh sách API chính
- `POST /api/v1/auth/guest`: Đăng nhập nhanh cho người chơi mới vào thẳng trận mà không cần đăng ký tài khoản rườm rà.
- `GET /api/v1/leaderboard/global`: Lấy danh sách Top 100 người chơi có số wave vượt qua cao nhất.
- `POST /api/v1/leaderboard/submit`: Server Colyseus gửi kết quả điểm số trận đấu lên (được ký bằng Secret Key nội bộ chống người ngoài gửi điểm ảo).
- `GET /api/v1/profile/meta-talents`: Lấy danh sách kỹ năng vĩnh viễn đã mở khóa.

### 4.2 Bảng xếp hạng Redis Sorted Sets siêu tốc
Sử dụng cấu trúc `ZSET` của Redis giúp việc đọc và cập nhật bảng xếp hạng chỉ mất dưới 1 miligiây với độ phức tạp $O(\log N)$:
```typescript
// Tính điểm tổng hợp: Điểm Wave ưu tiên hàng đầu, sau đó đến số quái diệt được
const scoreKey = 'leaderboard:global:wave_score';
const compositeScore = waveCleared * 1_000_000 + totalKills * 10 - elapsedSeconds;
await redisClient.zadd(scoreKey, compositeScore, playerId);
```

---

## 5. Chiến Lược Tối Ưu Hóa Unity WebGL

Để game có thể tải trong vòng 3-5 giây trên trình duyệt web của nhà tuyển dụng mà không bị đơ giật:
1. **Nén Brotli cấp cao**: Toàn bộ dữ liệu game được nén bằng chuẩn Brotli (`.br`), giảm kích thước file tải về hơn 60% so với Gzip thông thường.
2. **Khống chế dung lượng RAM 256MB**: Cấu hình bộ nhớ WebGL Heap ban đầu ở mức 256MB, đảm bảo không bao giờ bị trình duyệt Chrome/Safari trên điện thoại tắt tab vì tràn bộ nhớ.
3. **Cắt tỉa code (Managed Stripping Level = High)**: Trình biên dịch Unity IL2CPP tự động loại bỏ toàn bộ các module và class không được dùng đến trong game, giữ bundle siêu nhỏ gọn.
