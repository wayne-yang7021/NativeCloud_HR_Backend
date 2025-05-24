// report_handlers.go
package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/4040www/NativeCloud_HR/internal/service"
	"github.com/gin-gonic/gin"
)

type AttendanceSummary struct {
	Date         string `json:"date"`
	Name         string `json:"name"`
	ClockInTime  string `json:"clock_in_time"`
	ClockOutTime string `json:"clock_out_time"`
	ClockInGate  string `json:"clock_in_gate"`
	ClockOutGate string `json:"clock_out_gate"`
	Status       string `json:"status"` // "On Time" / "Late" / "Abnormal"
}

func GetMyTodayRecords(c *gin.Context) {
	userID := c.Param("userID")

	summary, err := service.GetTodayAttendanceSummary(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if summary == nil {
		c.JSON(http.StatusOK, gin.H{"message": "no access records today"})
		return
	}
	c.JSON(http.StatusOK, summary)

}

func GetMyHistoryRecords(c *gin.Context) {
	userID := c.Param("userID")
	end := time.Now()
	start := end.AddDate(0, 0, -30)

	summaries, err := service.GetAttendanceWithEmployee(userID, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, summaries)

}

func GetMyPeriodRecords(c *gin.Context) {
	userID := c.Param("userID")
	startDate := c.Param("startDate")
	endDate := c.Param("endDate")

	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid startDate"})
		return
	}
	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid endDate"})
		return
	}

	summaries, err := service.GetAttendanceWithEmployee(userID, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, summaries)

}

func GetThisMonthTeam(c *gin.Context) {
	department := c.Param("department")
	monthsStr := c.DefaultQuery("months", "1")
	months, err := strconv.Atoi(monthsStr)
	if err != nil || months < 1 {
		months = 1
	}

	today := time.Now()
	var reports []interface{}

	for i := 0; i < months; i++ {
		targetMonth := today.AddDate(0, -i, 0).Format("2006-01")

		report, err := service.FetchMonthlyTeamReport(department, targetMonth)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		reports = append(reports, report)
	}

	c.JSON(http.StatusOK, reports)
}

func GetThisWeekTeam(c *gin.Context) {
	department := c.Param("department")
	weeksStr := c.DefaultQuery("weeks", "1")
	weeks, err := strconv.Atoi(weeksStr)
	if err != nil || weeks < 1 {
		weeks = 1
	}

	today := time.Now()
	var reports []interface{}

	for i := 0; i < weeks; i++ {
		end := today.AddDate(0, 0, -7*i)
		start := end.AddDate(0, 0, -int(end.Weekday())+1)

		report, err := service.FetchWeeklyTeamReport(department, start.Format("2006-01-02"), end.Format("2006-01-02"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		reports = append(reports, report)
	}

	c.JSON(http.StatusOK, reports)
}

func GetCustomPeriodTeam(c *gin.Context) {
	department := c.Param("department")
	startDate := c.Param("startDate")
	endDate := c.Param("endDate")
	result, err := service.FetchCustomPeriodTeamReport(department, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)

}

func GetAlertList(c *gin.Context) {
	start := c.Param("startDate")
	end := c.Param("endDate")
	list, err := service.GenerateAlertList(start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

func GetInChargeDepartments(c *gin.Context) {
	userID := c.Param("userID")
	depts := service.GetManagedDepartments(userID)
	c.JSON(http.StatusOK, depts)
}

func ExportSummaryCSV(c *gin.Context) {
	department := c.Param("department")
	start := c.Param("startDate")
	end := c.Param("endDate")
	data, err := service.GenerateAttendanceSummaryCSV(department, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Disposition", "attachment; filename=summary.csv")
	c.Data(http.StatusOK, "text/csv", data)
}

func ExportSummaryPDF(c *gin.Context) {
	department := c.Param("department")
	start := c.Param("startDate")
	end := c.Param("endDate")
	data, err := service.GenerateAttendanceSummaryPDF(department, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Disposition", "attachment; filename=summary.pdf")
	c.Data(http.StatusOK, "application/pdf", data)
}

// Page: Attendence Log
func GetMyDepartments(c *gin.Context) {
	userID := c.Param("userID")
	departments, err := service.GetManagedDepartmentsFromDB(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, departments)
}

func FilterAttendanceSummary(c *gin.Context) {
	department := c.Query("department")
	fromDate := c.Query("fromDate")
	toDate := c.Query("toDate")
	result, err := service.GetAttendanceSummaryForDepartments(department, fromDate, toDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func ExportAttendanceSummaryCSV(c *gin.Context) {
	department := c.Query("department")
	fromDate := c.Query("fromDate")
	toDate := c.Query("toDate")
	csvData, err := service.GenerateAttendanceSummaryCSV(department, fromDate, toDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Disposition", "attachment; filename=summary.csv")
	c.Data(http.StatusOK, "text/csv", csvData)
}

func ExportAttendanceSummaryPDF(c *gin.Context) {
	department := c.Query("department")
	fromDate := c.Query("fromDate")
	toDate := c.Query("toDate")
	pdfData, err := service.GenerateAttendanceSummaryPDF(department, fromDate, toDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Disposition", "attachment; filename=summary.pdf")
	c.Data(http.StatusOK, "application/pdf", pdfData)
}
