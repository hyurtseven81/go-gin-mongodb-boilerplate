package health

import (
	"github.com/gin-gonic/gin"
)

func AddHealthRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/ping")

	ping.GET("", Ping)
}
