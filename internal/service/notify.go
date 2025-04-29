package service

import (
	"time"

	"github.com/4040www/NativeCloud_HR/internal/repository"
)

type ProblemEmployee struct {
	EmployeeID string   `json:"employee_id"`
	Problems   []string `json:"problems"`
}

func FindProblematicEmployees() ([]ProblemEmployee, error) {
	employeeIDs, err := repository.GetAllEmployeeIDs()
	if err != nil {
		return nil, err
	}

	var result []ProblemEmployee
	now := time.Now()
	monthStr := now.Format("2006-01")

	for _, eid := range employeeIDs {
		logs, err := repository.GetAccessLogsByEmployeeAndMonth(eid, monthStr)
		if err != nil {
			continue
		}
		lateCount := 0
		overtimeHours := 0.0

		for _, log := range logs {
			if log.Direction == "IN" && log.AccessTime.Hour() > 8 || (log.AccessTime.Hour() == 8 && log.AccessTime.Minute() > 30) {
				lateCount++
			}
			if log.Direction == "OUT" && log.AccessTime.Hour() >= 18 {
				overtimeHours += float64(log.AccessTime.Hour()-18) + float64(log.AccessTime.Minute())/60
			}
		}

		var problems []string
		if lateCount >= 4 {
			problems = append(problems, "TooManyLate")
		}
		if overtimeHours > 46 {
			problems = append(problems, "OvertimeExceeded")
		}

		if len(problems) > 0 {
			result = append(result, ProblemEmployee{
				EmployeeID: eid,
				Problems:   problems,
			})
		}
	}
	return result, nil
}

func NotifyManagerLate(employeeID string) string {
	return "員工 " + employeeID + " 本月遲到超過 4 次，請主管關注。"
}

func NotifyHROvertime(employeeID string) string {
	return "員工 " + employeeID + " 本月加班超過 46 小時，請 HR 檢查。"
}
