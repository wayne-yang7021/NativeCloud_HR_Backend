package routes

import (
    "github.com/gin-gonic/gin"
    "NativeCloud_HR/api/routes/auth"
    "NativeCloud_HR/api/routes/clock"
    "NativeCloud_HR/api/routes/notify"
    "NativeCloud_HR/api/routes/report"
)

// modularized api router
func SetupRoutes(r *gin.Engine) {
    api := r.Group("/api")

    api.GET("/status", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "online",
            "timestamp": time.Now(),
            "version": "1.0.0",
        })
    })

    auth.RegisterRoutes(api.Group("/auth"))
    clock.RegisterRoutes(api.Group("/clock"))
    notify.RegisterRoutes(api.Group("/notify"))
    report.RegisterRoutes(api.Group("/report")) 
}
