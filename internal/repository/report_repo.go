package repository

import (
	"time"

	"github.com/4040www/NativeCloud_HR/internal/db"
	"github.com/4040www/NativeCloud_HR/internal/model"
)

func GetAccessLogsByEmployeeBetween(employeeID string, start, end time.Time) ([]model.AccessLog, error) {
	var logs []model.AccessLog
	err := db.DB.Where("employee_id = ? AND access_time BETWEEN ? AND ?", employeeID, start, end).Order("access_time asc").Find(&logs).Error
	return logs, err
}

func GetAllEmployees() ([]model.Employee, error) {
	var employees []model.Employee
	err := db.DB.Find(&employees).Error
	return employees, err
}

// // report_repository.go
// package repository

// import (
// 	"time"

// 	"github.com/4040www/NativeCloud_HR/internal/db"
// 	"github.com/4040www/NativeCloud_HR/internal/model"
// )

// func GetTodayRecords(userID uint) ([]model.DailyReport, error) {
// 	var records []model.DailyReport
// 	today := time.Now().Format("2006-01-02")
// 	err := db.DB.Where("employee_id = ? AND date = ?", userID, today).Find(&records).Error
// 	return records, err
// }

// func GetHistoryRecords(userID uint) ([]model.DailyReport, error) {
// 	var records []model.DailyReport
// 	yearMonth := time.Now().Format("2006-01")
// 	err := db.DB.Where("employee_id = ? AND to_char(date, 'YYYY-MM') = ?", userID, yearMonth).Find(&records).Error
// 	return records, err
// }

// func GetHistoryRecordsBetween(userID uint, start, end string) ([]model.DailyReport, error) {
// 	var records []model.DailyReport
// 	err := db.DB.Where("employee_id = ? AND date BETWEEN ? AND ?", userID, start, end).Find(&records).Error
// 	return records, err
// }

// func GetMonthlyTeamRecords(department string, month string) ([]model.DailyReport, error) {
// 	var records []model.DailyReport
// 	err := db.DB.Table("daily_reports").Joins("JOIN employees ON employees.id = daily_reports.employee_id").
// 		Where("employees.department = ? AND to_char(daily_reports.date, 'YYYY-MM') = ?", department, month).Find(&records).Error
// 	return records, err
// }

// func GetWeeklyTeamRecords(department string, weekStartDate, weekEndDate string) ([]model.DailyReport, error) {
// 	var records []model.DailyReport
// 	err := db.DB.Table("daily_reports").Joins("JOIN employees ON employees.id = daily_reports.employee_id").
// 		Where("employees.department = ? AND daily_reports.date BETWEEN ? AND ?", department, weekStartDate, weekEndDate).Find(&records).Error
// 	return records, err
// }

// func GetCustomPeriodTeamRecords(department, start, end string) ([]model.DailyReport, error) {
// 	var records []model.DailyReport
// 	err := db.DB.Table("daily_reports").Joins("JOIN employees ON employees.id = daily_reports.employee_id").
// 		Where("employees.department = ? AND daily_reports.date BETWEEN ? AND ?", department, start, end).Find(&records).Error
// 	return records, err
// }

// func GetAttendanceFiltered(department, start, end string) ([]model.DailyReport, error) {
// 	var records []model.DailyReport
// 	err := db.DB.Table("daily_reports").Joins("JOIN employees ON employees.id = daily_reports.employee_id").
// 		Where("employees.department = ? AND daily_reports.date BETWEEN ? AND ?", department, start, end).Find(&records).Error
// 	return records, err
// }
