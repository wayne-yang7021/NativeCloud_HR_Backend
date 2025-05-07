// report_routes.go
package routes

import (
	handlers "github.com/4040www/NativeCloud_HR/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func RegisterReportRoutes(router *gin.RouterGroup) {
	// report := router.Group("/report")
	{
		router.GET("/myRecords/:userID", handlers.GetMyTodayRecords)
		router.GET("/historyRecords/:userID", handlers.GetMyHistoryRecords)
		router.GET("/historyRecords/:userID/:startDate/:endDate", handlers.GetMyPeriodRecords)
		router.GET("/thisMonth/:department/:userID", handlers.GetThisMonthTeam)
		router.GET("/thisWeek/:department/:userID", handlers.GetThisWeekTeam)
		router.GET("/PeriodTime/:department/:startDate/:endDate/:userID", handlers.GetCustomPeriodTeam)
		router.GET("/filterAttendence", handlers.FilterAttendance)
		router.GET("/exportAttendenceCSV", handlers.ExportAttendanceCSV)
		router.GET("/exportAttendencePDF", handlers.ExportAttendancePDF)
	}
}
