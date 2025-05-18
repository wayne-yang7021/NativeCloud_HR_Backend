# 使用 GCP 虛擬機連接與管理 NativeCloud\_HR 專案

## 前置條件

1. 擁有一個 Google Cloud 帳戶。
2. 已將你加入專案。

---

## 步驟

### 1. 使用 SSH 連接到 VM

請使用以下指令進入你已建立的虛擬機：

```bash
gcloud compute ssh native-cloud-hr --zone=asia-east1-c
```

這條指令會透過 `gcloud` 工具使用 SSH 連接到名為 `native-cloud-hr` 的虛擬機，zone 為 `asia-east1-c`。

---

### 2. 在 VM 上使用 Git 管理專案

#### 2.1. Clone 專案程式碼

若是第一次部署，可以使用 Git 下載專案：

```bash
git clone https://github.com/your-username/NativeCloud_HR.git
cd NativeCloud_HR
```

#### 2.2. 更新專案程式碼

若已經 clone 過，更新程式碼：

```bash
cd NativeCloud_HR
git pull
```

---

### 3. 使用 Docker Compose 管理應用

#### 3.1. 建立與啟動應用（含 build）

```bash
sudo docker-compose up --build
```

此指令會依照 `docker-compose.yml` 重新建構所有服務，並啟動應用。

#### 3.2. 以背景模式執行應用

```bash
sudo docker-compose up -d
```

加上 `-d` 會讓應用在背景執行，不會鎖住終端機。使用背景執行後，如果想查看 log 可以輸入

```bash
sudo docker-compose logs -f
```

#### 3.3. 停止並移除容器

```bash
sudo docker-compose down
```

這條指令會停止所有服務並清除相關資源。

---

### 4. 常見錯誤處理

#### ❌ 如果遇到 Image 或 Container 問題，例如：

```bash
ERROR: for <service> Container "xxxx" is unhealthy.
ERROR: for <service> 'ContainerConfig' KeyError
```

請嘗試清除並重新建構：

```bash
sudo docker-compose down -v --rmi all
sudo docker system prune -af
sudo docker-compose up --build
```

---

### 5. 配置 GCP 防火牆開放端口（例如 8080）

```bash
gcloud compute firewall-rules create allow-http --allow tcp:8080
```
