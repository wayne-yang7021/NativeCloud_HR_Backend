package routes

import (
    "github.com/gin-gonic/gin"
    "NativeCloud_HR/api/handlers"
)

func RegisterRoutes(report *gin.RouterGroup) {

    report.GET("/me", handlers.GenerateMyReport)           // GET /report/me
    report.GET("/department", handlers.GenerateDeptReport) // GET /report/department
    report.GET("/today", handlers.GetTodayRecords)         // GET /report/today
    
}
