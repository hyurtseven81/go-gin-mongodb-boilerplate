package projectmessage

import (
	"net/http"

	"data-pad.app/data-api/utils"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var query utils.Query
	c.BindQuery(&query)

	projectId := c.Param("project_id")
	projection := query.GetSelect()
	filter := query.GetFilter()
	sort := query.GetSort()

	s := NewProjectMessageService()

	dataResult := s.List(projectId, filter, projection, int64(query.Skip), int64(query.Limit), sort)

	c.JSON(http.StatusOK, dataResult)
}

func Insert(c *gin.Context) {
	var body ProjectMessage
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s := NewProjectMessageService()

	inserted := s.Insert(body)

	c.JSON(http.StatusCreated, inserted)
}

func Update(c *gin.Context) {
	var body ProjectMessage
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("project_message_id")

	s := NewProjectMessageService()

	updated, err := s.Update(id, body)

	if err != nil {
		c.JSON(err.StatusCode, err.Err)
	}

	c.JSON(http.StatusOK, updated)
}

func Delete(c *gin.Context) {
	id := c.Param("project_message_id")

	s := NewProjectMessageService()

	err := s.Delete(id)

	if err != nil {
		c.JSON(err.StatusCode, err.Err)
	}

	c.String(http.StatusNoContent, "")
}
