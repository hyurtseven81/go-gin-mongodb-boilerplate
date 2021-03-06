package routes

import (
	"fmt"

	"data-pad.app/data-api/config"
	"data-pad.app/data-api/health"
	"data-pad.app/data-api/middlewares"
	"data-pad.app/data-api/resources/dashboard"
	"data-pad.app/data-api/resources/project"
	"data-pad.app/data-api/resources/projectmessage"
	"data-pad.app/data-api/resources/snippet"
	"data-pad.app/data-api/resources/task"

	_ "data-pad.app/data-api/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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

	c := config.GetConfig()

	url := ginSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", c.Host))

	// The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	v1 := router.Group("/v1")

	v1.Use(middlewares.AuthMiddleware())

	task.AddTaskRoutes(v1)
	dashboard.AddDashboardRoutes(v1)
	snippet.AddSnippetRoutes(v1)
	project.AddProjectRoutes(v1)
	projectmessage.AddProjectMessageRoutes(v1)
}

// Run will start the server
func Run() {
	getRoutes()

	router.Use(gin.Recovery())

	router.Run("0.0.0.0:5000")
}
