// report_routes.go
package routes

import (
	handlers "github.com/4040www/NativeCloud_HR/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func RegisterReportRoutes(router *gin.RouterGroup) {
	report := router.Group("/report")
	{
		report.GET("/myRecords/:userID", handlers.GetMyTodayRecords)
		report.GET("/historyRecords/:userID", handlers.GetMyHistoryRecords)
		report.GET("/historyRecords/:userID/:startDate-:endDate", handlers.GetMyPeriodRecords)
		report.GET("/thisMonth/:department/:userID", handlers.GetThisMonthTeam)
		report.GET("/thisWeek/:department/:userID", handlers.GetThisWeekTeam)
		report.GET("/PeriodTime/:department/:startDate-:endDate/:userID", handlers.GetCustomPeriodTeam)
		report.GET("/filterAttendence", handlers.FilterAttendance)
		report.GET("/exportAttendenceCSV", handlers.ExportAttendanceCSV)
		report.GET("/exportAttendencePDF", handlers.ExportAttendancePDF)

	}
}
