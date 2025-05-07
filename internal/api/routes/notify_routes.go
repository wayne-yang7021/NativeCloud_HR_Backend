package routes

import (
	handlers "github.com/4040www/NativeCloud_HR/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func RegisterNotifyRoutes(router *gin.RouterGroup) {
	// notify := router.Group("/")
	{
		router.GET("/warning", handlers.OvertimeOrLateCheck)
		router.POST("/late", handlers.NotifyManagerTooManyLate)
		router.POST("/overtime", handlers.NotifyHRExceedOvertime)
	}
}

// curl -X GET http://localhost:8080/api/notify/warning \
// -H "Content-Type: application/json" \
// -d '{
// "email": "test@example.com",
// "password": "your_password"}'
