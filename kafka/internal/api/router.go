package api

import (
	"time"

	"github.com/gin-gonic/gin"
)

// 註冊 API 所有路由
func SetupRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")

	apiGroup.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "online",
			"timestamp": time.Now(),
			"version":   "1.0.0",
		})
	})

	RegisterClockRoutes(apiGroup.Group("/clock"))
}

func RegisterClockRoutes(clock *gin.RouterGroup) {
	clock.POST("/", CheckIn) // POST /api/clock
}
