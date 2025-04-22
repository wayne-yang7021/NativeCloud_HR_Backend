func CreateCheckinRecord(record *model.CheckinRecord) error {
    return db.GetDB().Create(record).Error
}
