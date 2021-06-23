package dashboard

import (
	"github.com/gin-gonic/gin"
)

func AddDashboardRoutes(rg *gin.RouterGroup) {
	dashboardsRoute := rg.Group("/projects/:project_id/dashboards")

	dashboardsRoute.GET("", Get)
	dashboardsRoute.POST("", Insert)
	dashboardsRoute.PUT("/:dashboard_id", Update)
	dashboardsRoute.DELETE("/:dashboard_id", Delete)
}
