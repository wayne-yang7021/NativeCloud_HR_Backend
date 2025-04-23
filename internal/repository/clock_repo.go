package repository

import (
	"github.com/4040www/NativeCloud_HR/internal/db"
	"github.com/4040www/NativeCloud_HR/internal/model"
)

func CreateCheckinRecord(record *model.CheckinRecord) error {
	return db.GetDB().Create(record).Error
}
