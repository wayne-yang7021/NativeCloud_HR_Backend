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
cd /opt/shared/NativeCloud_HR
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

---

### 6. 自動部署流程

當你將程式碼推送到 `main` 分支，並且 **加上 Git Tag（例如 `v1.0.0`）** 時，GitHub Actions 會自動進行以下動作：

#### ✅ 自動觸發的步驟如下：

1. **構建並推送 Docker Image 至 Docker Hub**：

   * 自動使用 Git tag 作為版本號，例如：`yourdockerhub/native-cloud-hr:v1.0.0`
   * 同時也會推送一份 `latest` tag 的 image，方便 GCP VM 使用

2. **SSH 連接 GCP VM**：

   * GitHub Actions 會使用你提供的 SSH 金鑰連線到 `native-cloud-hr` VM

3. **拉取最新 Docker Image 並重新部署應用**：

   * 使用 `docker pull` 取得最新 image（版本 tag 或 latest）
   * 透過 `sudo docker-compose down` 停止舊容器
   * 使用 `sudo docker-compose up -d` 背景啟動新容器

#### 🛠 需要你準備好的條件：

* 已設定以下 GitHub Secrets：

  * `DOCKER_USERNAME`, `DOCKER_PASSWORD`
  * `GCP_SSH_USER`, `GCP_SSH_KEY`, `GCP_VM_IP`
* VM 上已經配置好對應的 `docker-compose.yml`，可拉取正確的 image
* GitHub Actions workflow 已配置好正確的自動化流程（如使用 `appleboy/ssh-action`）

#### ⏩ 例子：自動化流程觸發方式

```bash
# 修改完程式碼後，提交變更
git add .
git commit -m "新增功能"
git push

# 建立並推送 Git tag（這會觸發 GitHub Actions）
git tag v1.0.0
git push origin v1.0.0
```

---

這樣你每次標記一個新版本，只要 Push Tag，就會：

* 自動打包 image
* 上傳 Docker Hub
* 在 GCP VM 上自動重啟部署應用 🎉
