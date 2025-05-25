package handlers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/4040www/NativeCloud_HR/config"
	"github.com/gin-gonic/gin"
)

var cfg *config.Config

func Init(c *config.Config) {
	cfg = c
}

func CheckIn(c *gin.Context) {
	bodyBytes, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	url := cfg.MessageQueue.URL
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "HTTP request failed"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": fmt.Sprintf("Remote service returned status: %s", resp.Status)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Forwarded successfully"})
}
