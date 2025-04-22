package routes

import (
    "github.com/gin-gonic/gin"
    "NativeCloud_HR/api/handlers"
)

func RegisterRoutes(clock *gin.RouterGroup) {

    clock.POST("/", handlers.CheckIn)             // POST /clock
    clock.GET("/", handlers.GetMyRecords)         // GET /clock
    clock.GET("/summary", handlers.MonthlySummary) // GET /clock/summary
    
}
