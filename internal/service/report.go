package service

import (
	"bytes"
	"encoding/csv"
	"time"

	"github.com/4040www/NativeCloud_HR/internal/model"
	"github.com/4040www/NativeCloud_HR/internal/repository"
	"github.com/jung-kurt/gofpdf"
)

// 計算一日的工作時數，並判斷是否遲到
func calculateDailyWorkHours(logs []model.AccessLog) (float64, bool) {
	var clockIn, clockOut *time.Time
	isLate := false

	for _, log := range logs {
		if log.Direction == "IN" {
			if clockIn == nil || log.AccessTime.Before(*clockIn) {
				clockIn = &log.AccessTime
			}
			// 判斷是否超過 08:30 遲到
			if log.AccessTime.Hour() > 8 || (log.AccessTime.Hour() == 8 && log.AccessTime.Minute() > 30) {
				isLate = true
			}
		}
		if log.Direction == "OUT" {
			if clockOut == nil || log.AccessTime.After(*clockOut) {
				clockOut = &log.AccessTime
			}
		}
	}

	if clockIn != nil && clockOut != nil {
		workHours := clockOut.Sub(*clockIn).Hours()
		return workHours, isLate
	}
	return 0, isLate
}

// 取得今日某員工的刷卡紀錄
func FetchTodayRecords(employeeID string) ([]model.AccessLog, error) {
	today := time.Now()
	start := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	end := start.Add(24 * time.Hour)
	return repository.GetAccessLogsByEmployeeBetween(employeeID, start, end)
}

// 取得近一個月的刷卡紀錄
func FetchHistoryRecords(employeeID string) ([]model.AccessLog, error) {
	start := time.Now().AddDate(0, -1, 0)
	end := time.Now()
	return repository.GetAccessLogsByEmployeeBetween(employeeID, start, end)
}

// 取得自訂日期區間的刷卡紀錄
func FetchHistoryRecordsBetween(employeeID, startDate, endDate string) ([]model.AccessLog, error) {
	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)
	return repository.GetAccessLogsByEmployeeBetween(employeeID, start, end.Add(24*time.Hour))
}

// 統計部門在指定月份的總工作時數、加班時數、人數
func FetchMonthlyTeamReport(departmentID, month string) (float64, float64, int, error) {
	employees, err := repository.GetAllEmployees()
	if err != nil {
		return 0, 0, 0, err
	}

	monthTime, _ := time.Parse("2006-01", month)
	start := time.Date(monthTime.Year(), monthTime.Month(), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0)

	totalHours := 0.0
	otHours := 0.0
	employeeCount := 0

	for _, e := range employees {
		if e.OrganizationID == departmentID {
			logs, _ := repository.GetAccessLogsByEmployeeBetween(e.EmployeeID, start, end)
			workHours, _ := calculateDailyWorkHours(logs)
			totalHours += workHours
			if workHours > 8 {
				otHours += workHours - 8 // 超過 8 小時的部分視為加班
			}
			employeeCount++
		}
	}

	return totalHours, otHours, employeeCount, nil
}

// 本質上同 FetchCustomPeriodTeamReport
func FetchWeeklyTeamReport(departmentID, startDate, endDate string) (float64, float64, int, error) {
	return FetchCustomPeriodTeamReport(departmentID, startDate, endDate)
}

// 統計部門在指定日期區間的總工作時數、加班時數、人數
func FetchCustomPeriodTeamReport(departmentID, startDate, endDate string) (float64, float64, int, error) {
	employees, err := repository.GetAllEmployees()
	if err != nil {
		return 0, 0, 0, err
	}
	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)

	totalHours := 0.0
	otHours := 0.0
	employeeCount := 0

	for _, e := range employees {
		if e.OrganizationID == departmentID {
			logs, _ := repository.GetAccessLogsByEmployeeBetween(e.EmployeeID, start, end.Add(24*time.Hour))
			workHours, _ := calculateDailyWorkHours(logs)
			totalHours += workHours
			if workHours > 8 {
				otHours += workHours - 8
			}
			employeeCount++
		}
	}

	return totalHours, otHours, employeeCount, nil
}

// 匯出 CSV 格式的出勤紀錄
func GenerateAttendanceCSV(departmentID, startDate, endDate string) ([]byte, error) {
	logs, _ := FetchAttendanceFiltered(departmentID, startDate, endDate)
	var b bytes.Buffer
	w := csv.NewWriter(&b)
	w.Write([]string{"EmployeeID", "AccessTime", "Direction", "GateName"})
	for _, log := range logs {
		w.Write([]string{log.EmployeeID, log.AccessTime.Format("2006-01-02 15:04"), log.Direction, log.GateName})
	}
	w.Flush()
	return b.Bytes(), nil
}

// 匯出 PDF 格式的出勤紀錄
func GenerateAttendancePDF(departmentID, startDate, endDate string) ([]byte, error) {
	logs, _ := FetchAttendanceFiltered(departmentID, startDate, endDate)
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "Attendance Records")
	pdf.Ln(10)
	for _, log := range logs {
		pdf.Cell(0, 10, log.EmployeeID+" "+log.AccessTime.Format("2006-01-02 15:04")+" "+log.Direction+" "+log.GateName)
		pdf.Ln(8)
	}
	var b bytes.Buffer
	err := pdf.Output(&b)
	return b.Bytes(), err
}

// 撈出某部門某時間區間內所有人的出勤紀錄
func FetchAttendanceFiltered(departmentID, startDate, endDate string) ([]model.AccessLog, error) {
	employees, _ := repository.GetAllEmployees()
	var result []model.AccessLog
	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)
	for _, e := range employees {
		if e.OrganizationID == departmentID {
			logs, _ := repository.GetAccessLogsByEmployeeBetween(e.EmployeeID, start, end.Add(24*time.Hour))
			result = append(result, logs...)
		}
	}
	return result, nil
}
