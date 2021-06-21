package task

import (
	"github.com/gin-gonic/gin"
)

func AddTaskRoutes(rg *gin.RouterGroup) {
	tasksRoute := rg.Group("/tasks")

	tasksRoute.GET("", Get)
	tasksRoute.POST("", Insert)
	tasksRoute.PUT("/:id", Update)
	tasksRoute.DELETE("/:id", Delete)
}
