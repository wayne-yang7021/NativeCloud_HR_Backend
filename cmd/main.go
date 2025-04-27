package main

import (
	"NativeCloud_HR/internal/db"
	"fmt"
	"log"
)

func main() {
	// 替換成您 PostgreSQL 的連接資訊
	connStr := "host=35.221.151.72 user=postgres password=postgres dbname=postgres sslmode=disable"

	// 初始化資料庫連接
	err := db.InitDB(connStr)
	if err != nil {
		log.Fatalf("初始化資料庫失敗: %v", err)
	}
	defer db.CloseDB() // 確保程式結束時關閉資料庫連接

	fmt.Println("成功連接到 PostgreSQL 資料庫！")

}
