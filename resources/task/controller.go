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

	s := NewTaskService()

	dataResult := s.List(filter, projection, int64(query.Skip), int64(query.Limit), sort)

	c.JSON(http.StatusOK, dataResult)
}
