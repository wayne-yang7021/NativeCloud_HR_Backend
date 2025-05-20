package main

import (
	"fmt"
	"log"
	"time"

	"github.com/4040www/NativeCloud_HR/config"
	"github.com/4040www/NativeCloud_HR/internal/api"
	"github.com/4040www/NativeCloud_HR/internal/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化設定
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("設定加載失敗: %v", err)
	}

	// 連接資料庫
	db.InitPostgres()

	// 設置 API 路由
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8000", "http://localhost:8080"}, // 修正這裡
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api.SetupRoutes(router)

	serverAddr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("伺服器啟動於 %s", serverAddr)

	// 改用 Gin 提供的啟動方式
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("伺服器啟動失敗: %v", err)
	}
}
