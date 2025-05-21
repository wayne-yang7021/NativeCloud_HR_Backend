package handlers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckIn(c *gin.Context) {
	bodyBytes, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	url := "http://35.229.242.253:8080/api/clock/"
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
