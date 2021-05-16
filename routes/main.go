package routes

import (
	"data-pad.app/data-api/health"
	"data-pad.app/data-api/resources/task"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes() {
	v1 := router.Group("/v1")
	task.AddTaskRoutes(v1)
	health.AddHealthRoutes(v1)
}

// Run will start the server
func Run() {
	getRoutes()
	router.Run(":5000")
}
