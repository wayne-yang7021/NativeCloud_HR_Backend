package api

import (
	"net/http"

	"github.com/4040www/NativeCloud_HR/kafka/internal/messageQueue"
	"github.com/4040www/NativeCloud_HR/kafka/model"
	"github.com/gin-gonic/gin"
)

func CheckIn(c *gin.Context) {
	var req model.CheckInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// 發送到 Kafka（非同步寫入）
	if err := messagequeue.SendCheckIn(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enqueue check-in"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Check-in enqueued successfully",
	})
}
