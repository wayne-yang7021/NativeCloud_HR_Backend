package service

import (
	"time"

	"github.com/4040www/NativeCloud_HR/internal/model"
	"github.com/4040www/NativeCloud_HR/internal/repository"
)

func Checkin(userID uint, siteClass, site, note string) error {
	record := model.CheckinRecord{
		UserID:    userID,
		CheckinAt: time.Now(),
		SiteClass: siteClass,
		Site:      site,
		Note:      note,
	}

	return repository.CreateCheckinRecord(&record)
}
