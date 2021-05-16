package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var tasks []Task

	var count int64 = 0

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"items": tasks,
			"count": count,
		},
	})
}
