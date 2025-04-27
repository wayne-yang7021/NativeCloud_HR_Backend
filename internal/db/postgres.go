package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

// InitDB 初始化資料庫連接
func InitDB(connStr string) error {
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("連接資料庫失敗: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("資料庫連線測試失敗: %w", err)
	}

	return nil
}

// CloseDB 關閉資料庫連接
func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
