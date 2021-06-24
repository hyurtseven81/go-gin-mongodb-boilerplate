package dashboard

import (
	"github.com/gin-gonic/gin"
)

func AddDashboardRoutes(rg *gin.RouterGroup) {
	dashboardsRoute := rg.Group("/projects/:project_id/dashboards")

	dashboardsRoute.GET("", List)
	dashboardsRoute.GET("/:dashboard_id", Get)
	dashboardsRoute.POST("", Insert)
	dashboardsRoute.PUT("/:dashboard_id", Update)
	dashboardsRoute.DELETE("/:dashboard_id", Delete)
}
