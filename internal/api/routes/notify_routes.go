package routes

import (
    "github.com/gin-gonic/gin"
    "NativeCloud_HR/api/handlers"
)

func RegisterRoutes(notify *gin.RouterGroup) {

    notify.GET("/warning", handlers.OvertimeOrLateCheck)         // GET /notify/warning
    notify.POST("/late", handlers.NotifyManagerTooManyLate)      // POST /notify/late
    notify.POST("/overtime", handlers.NotifyHRExceedOvertime)    // POST /notify/overtime
    
}
