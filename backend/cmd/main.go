package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"

	"github.com/4040www/NativeCloud_HR/config"
	"github.com/4040www/NativeCloud_HR/internal/api"
	"github.com/4040www/NativeCloud_HR/internal/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("設定加載失敗: %v", err)
	}

	db.InitPostgres()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))
	api.SetupRoutes(router)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}

	// 啟動伺服器
	go func() {
		log.Printf("伺服器啟動於 %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("伺服器啟動失敗: %v", err)
		}
	}()

	// 監聽中止或中斷訊號（例如 Ctrl+C 或 SIGTERM）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 收到終止訊號，正在關閉伺服器...")

	// 優雅關閉伺服器（等待所有請求處理完畢）
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("伺服器關閉失敗: %v", err)
	}

	log.Println("✅ 伺服器已關閉")
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/4040www/NativeCloud_HR/config"
// 	"github.com/4040www/NativeCloud_HR/internal/api"
// 	"github.com/4040www/NativeCloud_HR/internal/db"

// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	// 初始化設定
// 	cfg, err := config.LoadConfig()
// 	if err != nil {
// 		log.Fatalf("設定加載失敗: %v", err)
// 	}

// 	// 連接資料庫
// 	db.InitPostgres()

// 	// 設置 API 路由
// 	router := gin.Default()

// 	router.Use(cors.New(cors.Config{
// 		AllowOrigins:     []string{"*"}, // 修正這裡
// 		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
// 		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
// 		ExposeHeaders:    []string{"Content-Length"},
// 		AllowCredentials: false,
// 		MaxAge:           12 * time.Hour,
// 	}))

// 	api.SetupRoutes(router)

// 	serverAddr := fmt.Sprintf(":%d", cfg.Server.Port)
// 	log.Printf("伺服器啟動於 %s", serverAddr)

//		// 改用 Gin 提供的啟動方式
//		if err := router.Run(serverAddr); err != nil {
//			log.Fatalf("伺服器啟動失敗: %v", err)
//		}
//	}
