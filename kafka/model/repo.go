package model

import (
	"errors"

	"github.com/4040www/NativeCloud_HR/kafka/internal/db"
)

// 寫入 DB 的邏輯（僅供 TCP server 使用）
func CreateCheckInRecord(record *CheckInRequest) error {
	if record == nil {
		return errors.New("record is nil")
	}
	return db.GetDB().Create(record).Error
}
