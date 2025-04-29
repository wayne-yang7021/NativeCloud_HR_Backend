package handlers

import (
	"net/http"

	"github.com/4040www/NativeCloud_HR/internal/service"
	"github.com/gin-gonic/gin"
)

func OvertimeOrLateCheck(c *gin.Context) {
	problems, err := service.FindProblematicEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查詢異常員工失敗"})
		return
	}
	c.JSON(http.StatusOK, problems)
}

func NotifyManagerTooManyLate(c *gin.Context) {
	employeeID := c.Query("employee_id")
	message := service.NotifyManagerLate(employeeID)
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func NotifyHRExceedOvertime(c *gin.Context) {
	employeeID := c.Query("employee_id")
	message := service.NotifyHROvertime(employeeID)
	c.JSON(http.StatusOK, gin.H{"message": message})
}
