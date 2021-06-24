package project

import (
	"github.com/gin-gonic/gin"
)

func AddProjectRoutes(rg *gin.RouterGroup) {
	projectsRoute := rg.Group("/projects")

	projectsRoute.GET("", List)
	projectsRoute.GET("/:project_id", Get)
	projectsRoute.POST("", Insert)
	projectsRoute.PUT("/:project_id", Update)
	projectsRoute.DELETE("/:project_id", Delete)
}
