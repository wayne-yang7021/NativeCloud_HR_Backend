# 第一階段：建置 Go 二進位檔
FROM golang:1.23 as builder

# 設定工作目錄
WORKDIR /app

# 複製模組定義檔，並先下載依賴，加速建置快取
COPY go.mod go.sum ./
RUN go mod download

# 複製程式碼
COPY . .

# 編譯成靜態執行檔（適合容器部署）
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/main.go


# 第二階段：使用 distroless 作為最小化執行環境
FROM gcr.io/distroless/static-debian11

# 建立工作目錄
WORKDIR /app

# 複製編譯完成的 Go 執行檔
COPY --from=builder /app/app .

# 開放 port
EXPOSE 8080

# 啟動容器執行程式
CMD ["./app"]
