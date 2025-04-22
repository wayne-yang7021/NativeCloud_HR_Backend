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
