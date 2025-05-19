package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/4040www/NativeCloud_HR/internal/model"
	"github.com/4040www/NativeCloud_HR/internal/service"
)

func CheckIn(c *gin.Context) {
	var req model.CheckInRequest

	// 解析使用者傳的 JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	// 呼叫 service
	if err := service.Clock(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Check-in successful",
	})
}
