package routes

import (
	handlers "github.com/4040www/NativeCloud_HR/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func RegisterClockRoutes(clock *gin.RouterGroup) {

	clock.POST("/", handlers.CheckIn)              // POST /clock
	clock.GET("/", handlers.GetMyRecords)          // GET /clock
	clock.GET("/summary", handlers.MonthlySummary) // GET /clock/summary

}
