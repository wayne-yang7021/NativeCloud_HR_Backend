package api

import (
	"time"

	"github.com/4040www/NativeCloud_HR/internal/api/routes"
	"github.com/gin-gonic/gin"
)

// modularized api router
func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "online",
			"timestamp": time.Now(),
			"version":   "1.0.0",
		})
	})

	routes.RegisterAuthRoutes(api.Group("/auth"))
	routes.RegisterClockRoutes(api.Group("/clock"))
	routes.RegisterNotifyRoutes(api.Group("/notify"))
	routes.RegisterReportRoutes(api.Group("/report"))
}
