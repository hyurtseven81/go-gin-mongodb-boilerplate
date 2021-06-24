package snippet

import (
	"github.com/gin-gonic/gin"
)

func AddSnippetRoutes(rg *gin.RouterGroup) {
	snippetsRoute := rg.Group("/projects/:project_id/snippets")

	snippetsRoute.GET("", List)
	snippetsRoute.GET("/:snippet_id", Get)
	snippetsRoute.POST("", Insert)
	snippetsRoute.PUT("/:snippet_id", Update)
	snippetsRoute.DELETE("/:snippet_id", Delete)
}
