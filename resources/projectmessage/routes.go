package projectmessage

import (
	"github.com/gin-gonic/gin"
)

func AddProjectMessageRoutes(rg *gin.RouterGroup) {
	projectMessagesRoute := rg.Group("/projects/:project_id/messages")

	projectMessagesRoute.GET("", List)
	projectMessagesRoute.GET("/:project_message_id", Get)
	projectMessagesRoute.POST("", Insert)
	projectMessagesRoute.PUT("/:project_message_id", Update)
	projectMessagesRoute.DELETE("/:project_message_id", Delete)
}
