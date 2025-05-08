package model

import "time"

type AccessLog struct {
	AccessID     string `gorm:"primaryKey"`
	EmployeeID   string
	AccessTime   time.Time
	Direction    string
	GateType     string
	GateName     string
	AccessResult string
}

// package model

// import "time"

// // DailyReport 單日打卡資料
// type DailyReport struct {
// 	ID           uint      `json:"id" gorm:"primaryKey"`
// 	Date         time.Time `json:"date"`
// 	EmployeeID   uint      `json:"employee_id"`
// 	Name         string    `json:"name"`
// 	ClockInTime  time.Time `json:"clock_in_time"`
// 	ClockOutTime time.Time `json:"clock_out_time"`
// 	ClockInGate  string    `json:"clock_in_gate"`
// 	ClockOutGate string    `json:"clock_out_gate"`
// 	Status       string    `json:"status"`
// }

// // MonthlyReport 月統計資料
// type MonthlyReport struct {
// 	ID                 uint    `json:"id" gorm:"primaryKey"`
// 	Month              string  `json:"month"`
// 	EmployeeID         uint    `json:"employee_id"`
// 	Name               string  `json:"name"`
// 	TotalWorkHours     float64 `json:"total_work_hours"`
// 	TotalOvertimeHours float64 `json:"total_overtime_hours"`
// 	LateCount          int     `json:"late_count"`
// 	EarlyLeaveCount    int     `json:"early_leave_count"`
// 	AbsenceCount       int     `json:"absence_count"`
// 	WorkingDays        int     `json:"working_days"`
// }

// // DepartmentSummary 部門彙總資料
// type DepartmentSummary struct {
// 	DepartmentName string  `json:"department_name"`
// 	TotalWorkHours float64 `json:"total_work_hours"`
// 	TotalOTHours   float64 `json:"total_ot_hours"`
// 	EmployeeCount  int     `json:"employee_count"`
// }
