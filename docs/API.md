# API 說明文件

## Auth API 說明文件 `/auth`

提供使用者登入、登出功能，採用 JSON Web Token（JWT）作為身分驗證方式。

---

### 📌 API 一覽

| 方法 | 路徑       | 說明     |
|------|------------|----------|
| POST | `/login`  | 使用者登入，取得 JWT |
| POST | `/logout` | 使用者登出 |

---

### 🟢 POST `/auth/login`

用戶登入並取得 JWT token。

#### 🔸 Request

- Header: `Content-Type: application/json`
- Body:
```json
{
  "email": "user@example.com",
  "password": "yourpassword"
}
```

#### 🔸 Response

- **Status 200 OK**
```json
{
  "message": "Login successful",
  "token": "your.jwt.token",
  "user": {
    "id": 1,
    "email": "user@example.com",
    "name": "John Doe"
  }
}
```

- **Status 400 Bad Request**
```json
{
  "error": "Invalid request format"
}
```

- **Status 401 Unauthorized**
```json
{
  "error": "email or password incorrect"
}
```

---

### 🔴 POST `/auth/logout`

模擬登出動作（前端只需刪除 JWT）。如採用 token blacklist，可額外實作伺服器端記錄失效 token。

#### 🔸 Request

- Header: (通常包含 `Authorization: Bearer <token>`)

#### 🔸 Response

- **Status 200 OK**
```json
{
  "message": "Logout successful"
}
```

---

#### 📘 補充說明

- JWT 可存於前端的 `localStorage` 或 `cookie`
- 每次 API 請求應在 Header 加上：

```http
Authorization: Bearer <your.jwt.token>
```

## Clock API 說明文件 `/clock`

用於員工打卡（進出紀錄）。該 API 會從 JWT Token 自動提取使用者身份，產生唯一 access_id，並寫入資料庫中。

---


### 🟢 POST `/clock`

需登入，並在 Authorization header 中提供有效的 JWT Token。

#### 🔸 Request

- Header: `Content-Type: application/json`
- Body:
```json
{
  "access_time": "2025-04-26T09:00:00Z",
  "direction": "in",         // "in" 或 "out"
  "gate_type": "main_gate",  // 例如：main_gate / side_gate（看 db 組）
  "gate_name": "北門"         // 打卡的門口名稱（看 db 組）
}

```

#### 🔸 Response

- **Status 200 OK**
```json
{
  "message": "Check-in successful"
}
```

- **Status 400 Bad Request**
```json
{
  "error": "Invalid request format"
}
```

- **Status 401 Unauthorized**
```json
{
  "error": "User ID not found in token"
}
```
