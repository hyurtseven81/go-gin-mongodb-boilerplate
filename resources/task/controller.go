package task

import (
	"net/http"

	"data-pad.app/data-api/utils"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var query utils.Query
	c.BindQuery(&query)

	projection := query.GetSelect()
	filter := query.GetFilter()
	sort := query.GetSort()

	var items []Task
	s := NewTaskService()

	s.List(filter, projection, int64(query.Skip), int64(query.Limit), sort, &items)

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"items": items,
			"count": 0,
		},
	})
}
