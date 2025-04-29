package routes

import (
	handlers "github.com/4040www/NativeCloud_HR/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func RegisterNotifyRoutes(router *gin.RouterGroup) {
	notify := router.Group("/notify")
	{
		notify.GET("/warning", handlers.OvertimeOrLateCheck)
		notify.POST("/late", handlers.NotifyManagerTooManyLate)
		notify.POST("/overtime", handlers.NotifyHRExceedOvertime)
	}
}
