// scripts/migrate.go
package main

import (
	"log"

	"github.com/4040www/NativeCloud_HR/internal/db"
)

func main() {
	log.Println("📦 正在初始化資料庫連線...")

	// 初始化連線（透過環境變數）
	db.InitPostgres()

	log.Println("🧱 執行 AutoMigrate 以建立或更新資料表...")

	// 自動建表（定義於 db.AutoMigrate）
	if err := db.AutoMigrate(); err != nil {
		log.Fatalf("❌ Migration 失敗: %v", err)
	}

	log.Println("✅ 資料庫 migration 完成！")
}
