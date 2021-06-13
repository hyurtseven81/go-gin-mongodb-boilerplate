package routes

import (
	"data-pad.app/data-api/health"
	"data-pad.app/data-api/middlewares"
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
	plain := router.Group("/")
	health.AddHealthRoutes(plain)

	v1 := router.Group("/v1")

	v1.Use(middlewares.AuthMiddleware())

	task.AddTaskRoutes(v1)
}

// Run will start the server
func Run() {
	getRoutes()

	router.Use(gin.Recovery())

	router.Run("localhost:5000")
}
