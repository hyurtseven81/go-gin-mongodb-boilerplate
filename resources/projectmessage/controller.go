package projectmessage

import (
	"net/http"

	"data-pad.app/data-api/utils"
	"github.com/gin-gonic/gin"
)

// Get ProjectMessages
// @Summary get all projectmessages under a project
// @Tags projectmessage
// @Accept json
// @Produce json
// @Param project_id path string true "Project Id"
// @Param filter query string false "string as JSON / used for filtering the results"
// @Param select query string false "string as JSON / used for projecting the results"
// @Param sort query string false "string as JSON / used for sorting the results"
// @Param skip query int false "used for skipping the results"
// @Param limit query int false "used for limitting the results"
// @Success 200 {object} utils.DataResult "ProjectMessage result"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "error details"
// @Security ApiKeyAuth
// @Router /v1/projects/{project_id}/projectmessages [get]
func List(c *gin.Context) {
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

// Get ProjectMessage By Id
// @Summary get a projectmessage under a project
// @Tags projectmessage
// @Accept json
// @Produce json
// @Param project_id path string true "Project Id"
// @Param project_message_id path string true "ProjectMessage Id"
// @Success 200 {object} ProjectMessage "ProjectMessage result"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Document not found"
// @Security ApiKeyAuth
// @Router /v1/projects/{project_id}/projectmessages/{project_message_id} [get]
func Get(c *gin.Context) {
	projectmessageId := c.Param("project_message_id")

	s := NewProjectMessageService()
	result, err := s.Get(projectmessageId)

	if err != nil {
		c.JSON(err.StatusCode, err.Err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Insert new projectmessage
// @Summary insert new projectmessage
// @Tags projectmessage
// @Accept json
// @Produce json
// @Param project_id path string true "Project Id"
// @Param message body ProjectMessage true "ProjectMessage Info"
// @Success 201 {integer} string "OK"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/projects/{project_id}/projectmessages [post]
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

// Update a projectmessage
// @Summary update a projectmessage
// @Tags projectmessage
// @Accept json
// @Produce json
// @Param project_id path string true "Project Id"
// @Param project_message_id path string true "ProjectMessage Id"
// @Param message body ProjectMessage true "ProjectMessage Info"
// @Success 200 {object} utils.DataResult "ProjectMessage result"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/projects/{project_id}/projectmessages/{project_message_id} [put]
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

// Delete a projectmessage
// @Summary delete a projectmessage
// @Tags projectmessage
// @Accept json
// @Produce json
// @Param project_id path string true "Project Id"
// @Param project_message_id path string true "ProjectMessage Id"
// @Success 204 {int} string "Empty response"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/projects/{project_id}/projectmessages/{project_message_id} [delete]
func Delete(c *gin.Context) {
	id := c.Param("project_message_id")

	s := NewProjectMessageService()

	err := s.Delete(id)

	if err != nil {
		c.JSON(err.StatusCode, err.Err)
	}

	c.String(http.StatusNoContent, "")
}
