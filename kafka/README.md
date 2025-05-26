# 🧭 NativeCloud_HR Monitoring & Kafka Stack

本專案建置一組以 Kafka 為核心、結合 Prometheus 與 Grafana 監控的容器化系統。適合用於開發與測試以 Kafka 為訊息中介的應用程式，並可即時觀察 Kafka Lag 及系統健康狀態。

---

## 📦 使用的服務

| Service         | Image                                 | 作用說明                                                                 |
|----------------|----------------------------------------|--------------------------------------------------------------------------|
| `kafka`         | `bitnami/kafka:3.5`                   | Kafka Broker，提供訊息佇列服務。                                         |
| `app`           | 自行構建 Dockerfile                    | 你的應用程式，會發送或接收 Kafka 訊息。                                   |
| `kafka-exporter`| `danielqsj/kafka-exporter:latest`     | 將 Kafka Lag 指標轉換為 Prometheus 能理解的格式。                        |
| `prometheus`    | `prom/prometheus`                     | 監控系統，定期拉取 exporter 的 metrics。                                 |
| `grafana`       | `grafana/grafana`                     | 圖表化工具，可視化 Kafka metrics 與其他系統狀態。                         |

---

## 🧵 Kafka 預設設定與運作方式

使用的是 Bitnami Kafka 3.5 單節點架構，主要設定如下：

| 參數                        | 預設值說明                              |
|-----------------------------|------------------------------------------|
| Partition 預設數量          | **1**（預設主題建立時）                  |
| Replication Factor          | **1**（因為是單節點，無副本）             |
| Listener 設定               | `INTERNAL://kafka:9092` 使用 PLAINTEXT 傳輸 |
| Broker ID                   | `1`                                     |
| Inter-Broker 通訊           | 使用 `INTERNAL` listener                 |
| Controller/Broker 整合     | Kafka 同時作為 Controller 與 Broker     |

📌 訊息流程：

1. `app` 發送訊息至 Kafka topic。
2. Kafka 存儲訊息（預設單 partition）。
3. 消費者接收訊息，若沒有及時消費，則會在 Kafka Exporter 中觀察到 lag。

---

## 📊 Grafana 操作指南

<<<<<<< HEAD
1. 啟動服務後，瀏覽器前往：[http://localhost:3000](http://localhost:3000)
2. 預設帳密：
   - **帳號**：`admin`
   - **密碼**：`admin`
3. 登入後，點選左側側邊欄 → "Data Sources" → 選擇或新增 Prometheus（URL: `http://prometheus:9090`）
4. 匯入 Kafka Lag Dashboard：
   - 點選 "+" → "Import"
   - 可使用 dashboard id：`7589`（Kafka Exporter Overview）或匯入自訂 JSON。

---

## 🔧 如何操作與測試 Kafka

### 進入 Kafka container：

```bash
docker exec -it kafka bash
以下是為你這份 `docker-compose.yml` 撰寫的完整 `README.md` 範例，說明 Kafka 預設行為、Grafana 操作、每個服務用途與彼此如何互動，非常適合放在專案根目錄中：

---

````markdown
# 🧭 NativeCloud_HR Monitoring & Kafka Stack

本專案建置一組以 Kafka 為核心、結合 Prometheus 與 Grafana 監控的容器化系統。適合用於開發與測試以 Kafka 為訊息中介的應用程式，並可即時觀察 Kafka Lag 及系統健康狀態。

---

## 📦 使用的服務

| Service         | Image                                 | 作用說明                                                                 |
|----------------|----------------------------------------|--------------------------------------------------------------------------|
| `kafka`         | `bitnami/kafka:3.5`                   | Kafka Broker，提供訊息佇列服務。                                         |
| `app`           | 自行構建 Dockerfile                    | 你的應用程式，會發送或接收 Kafka 訊息。                                   |
| `kafka-exporter`| `danielqsj/kafka-exporter:latest`     | 將 Kafka Lag 指標轉換為 Prometheus 能理解的格式。                        |
| `prometheus`    | `prom/prometheus`                     | 監控系統，定期拉取 exporter 的 metrics。                                 |
| `grafana`       | `grafana/grafana`                     | 圖表化工具，可視化 Kafka metrics 與其他系統狀態。                         |

---

## 🧵 Kafka 預設設定與運作方式

使用的是 Bitnami Kafka 3.5 單節點架構，主要設定如下：

| 參數                        | 預設值說明                              |
|-----------------------------|------------------------------------------|
| Partition 預設數量          | **1**（預設主題建立時）                  |
| Replication Factor          | **1**（因為是單節點，無副本）             |
| Listener 設定               | `INTERNAL://kafka:9092` 使用 PLAINTEXT 傳輸 |
| Broker ID                   | `1`                                     |
| Inter-Broker 通訊           | 使用 `INTERNAL` listener                 |
| Controller/Broker 整合     | Kafka 同時作為 Controller 與 Broker     |

📌 訊息流程：

1. `app` 發送訊息至 Kafka topic。
2. Kafka 存儲訊息（預設單 partition）。
3. 消費者接收訊息，若沒有及時消費，則會在 Kafka Exporter 中觀察到 lag。

---

## 📊 Grafana 操作指南

=======
>>>>>>> architecture
1. 啟動服務後，瀏覽器前往：[http://service_ip:3000](http://service_ip:3000)
2. 預設帳密：
   - **帳號**：`admin`
   - **密碼**：`admin`
3. 登入後，點選左側側邊欄 → "Data Sources" → 選擇或新增 Prometheus（URL: `http://prometheus:9090`）
4. 匯入 Kafka Lag Dashboard：
   - 點選 "+" → "Import"
   - 可使用 dashboard id：`7589`（Kafka Exporter Overview）或匯入自訂 JSON。

---

## 🔧 如何操作與測試 Kafka

<<<<<<< HEAD
### 進入 Kafka container
=======
### 進入 Kafka container：
>>>>>>> architecture

```bash
docker exec -it kafka bash
```

### 建立 Topic：

```bash
kafka-topics.sh --create --topic my-topic --bootstrap-server kafka:9092 --partitions 3 --replication-factor 1
```

### 檢視 Topic：

```bash
kafka-topics.sh --list --bootstrap-server kafka:9092
```

### 發送訊息：

```bash
kafka-console-producer.sh --broker-list kafka:9092 --topic my-topic
```

### 接收訊息：

```bash
kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic my-topic --from-beginning
```

---

## 🔍 查看 Metrics（Kafka Exporter）

<<<<<<< HEAD
Kafka Exporter 預設監聽在 [http://service_ip/metrics](http://service_ip:9308/metrics)，包含以下重要指標：
=======
Kafka Exporter 預設監聽在 [http://service_ip:9308/metrics](http://service_ip:9308/metrics)，包含以下重要指標：
>>>>>>> architecture

| 指標名稱                                   | 說明                         |
| -------------------------------------- | -------------------------- |
| `kafka_consumergroup_lag`              | 消費者群組與 partition 的 lag 數量  |
| `kafka_topic_partition_current_offset` | 每個 partition 當前 offset     |
| `kafka_topic_partition_leader`         | partition leader 所在 broker |

---

## 🔁 啟動與重建服務

### 啟動：

```bash
docker compose up --build
```

### 關閉：

```bash
docker compose down
```

<<<<<<< HEAD
### 若 port 被占用，可嘗試查看與釋放：
=======
### 若 port 被占用，可嘗試查看與釋放：https://chatgpt.com/c/67d52a6f-d32c-800b-a15a-57316c11441a
>>>>>>> architecture

```bash
sudo lsof -i :8080
docker rm -f <container_id>
```

---

## 🕸️ Container 網路互動關係圖

```
[GRAFANA] --> [PROMETHEUS] <-- [KAFKA-EXPORTER] <-- [KAFKA] <-- [APP]
<<<<<<< HEAD
                                 
=======
                                  
>>>>>>> architecture
```

* `kafka-exporter` 定期從 Kafka 拉 Lag 資訊。
* `prometheus` 拉取 `kafka-exporter` 和 `app` 的 metrics。
* `grafana` 從 `prometheus` 可視化全部 metrics。
* `app` 寫入 Kafka。
<<<<<<< HEAD

---
=======
>>>>>>> architecture
