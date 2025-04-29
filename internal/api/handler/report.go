package handlers

import (
	"net/http"
	"time"

	"github.com/4040www/NativeCloud_HR/internal/service"
	"github.com/gin-gonic/gin"
)

func GetMyTodayRecords(c *gin.Context) {
	userID := c.Param("userID")
	logs, err := service.FetchTodayRecords(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}

func GetMyHistoryRecords(c *gin.Context) {
	userID := c.Param("userID")
	logs, err := service.FetchHistoryRecords(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}

func GetMyPeriodRecords(c *gin.Context) {
	userID := c.Param("userID")
	startDate := c.Param("startDate")
	endDate := c.Param("endDate")
	logs, err := service.FetchHistoryRecordsBetween(userID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}

func GetThisMonthTeam(c *gin.Context) {
	department := c.Param("department")
	// userID := c.Param("userID")
	month := c.DefaultQuery("month", time.Now().Format("2006-01"))
	totalWork, totalOT, count, err := service.FetchMonthlyTeamReport(department, month)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_work_hours": totalWork, "total_overtime_hours": totalOT, "employee_count": count})
}

func GetThisWeekTeam(c *gin.Context) {
	department := c.Param("department")
	// userID := c.Param("userID")
	now := time.Now()
	start := now.AddDate(0, 0, -int(now.Weekday())+1)
	end := start.AddDate(0, 0, 6)
	totalWork, totalOT, count, err := service.FetchWeeklyTeamReport(department, start.Format("2006-01-02"), end.Format("2006-01-02"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_work_hours": totalWork, "total_overtime_hours": totalOT, "employee_count": count})
}

func GetCustomPeriodTeam(c *gin.Context) {
	department := c.Param("department")
	startDate := c.Param("startDate")
	endDate := c.Param("endDate")
	// userID := c.Param("userID")
	totalWork, totalOT, count, err := service.FetchCustomPeriodTeamReport(department, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_work_hours": totalWork, "total_overtime_hours": totalOT, "employee_count": count})
}

func FilterAttendance(c *gin.Context) {
	department := c.Query("department")
	startDate := c.Query("fromDate")
	endDate := c.Query("toDate")
	logs, err := service.FetchAttendanceFiltered(department, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}

func ExportAttendanceCSV(c *gin.Context) {
	department := c.Query("department")
	startDate := c.Query("fromDate")
	endDate := c.Query("toDate")
	data, err := service.GenerateAttendanceCSV(department, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Disposition", "attachment; filename=attendance.csv")
	c.Data(http.StatusOK, "text/csv", data)
}

func ExportAttendancePDF(c *gin.Context) {
	department := c.Query("department")
	startDate := c.Query("fromDate")
	endDate := c.Query("toDate")
	data, err := service.GenerateAttendancePDF(department, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Disposition", "attachment; filename=attendance.pdf")
	c.Data(http.StatusOK, "application/pdf", data)
}
