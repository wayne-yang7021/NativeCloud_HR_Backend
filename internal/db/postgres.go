package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/4040www/NativeCloud_HR/internal/model"
)

var DB *gorm.DB

func InitPostgres() {
	// 從環境變數讀取（建議，不要寫死）
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT") // 通常是 5432

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei",
		host, user, password, dbname, port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL")
}

func AutoMigrate() error {
	db := GetDB()

	log.Println("📦 開始自動建立資料表...")

	return db.AutoMigrate(
		&model.Employee{},
		&model.AccessLog{},
		&model.NotifyRecord{},
	)
}

func GetDB() *gorm.DB {
	return DB
}
