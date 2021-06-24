package project

import (
	"net/http"

	"data-pad.app/data-api/utils"
	"github.com/gin-gonic/gin"
)

// Get Projects
// @Summary get all projects
// @Tags project
// @Accept json
// @Produce json
// @Param filter query string false "string as JSON / used for filtering the results"
// @Param select query string false "string as JSON / used for projecting the results"
// @Param sort query string false "string as JSON / used for sorting the results"
// @Param skip query int false "used for skipping the results"
// @Param limit query int false "used for limitting the results"
// @Success 200 {object} utils.DataResult "Project result"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "error details"
// @Security ApiKeyAuth
// @Router /v1/projects [get]
func List(c *gin.Context) {
	var query utils.Query
	c.BindQuery(&query)

	projection := query.GetSelect()
	filter := query.GetFilter()
	sort := query.GetSort()

	s := NewProjectService()

	dataResult := s.List(filter, projection, int64(query.Skip), int64(query.Limit), sort)

	c.JSON(http.StatusOK, dataResult)
}

// Get Project By Id
// @Summary get a project
// @Tags project
// @Accept json
// @Produce json
// @Param project_id path string true "Project Id"
// @Success 200 {object} Project "Project result"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Document not found"
// @Security ApiKeyAuth
// @Router /v1/projects/{project_id} [get]
func Get(c *gin.Context) {
	projectId := c.Param("project_id")

	s := NewProjectService()
	result, err := s.Get(projectId)

	if err != nil {
		c.JSON(err.StatusCode, err.Err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Insert new project
// @Summary insert new project
// @Tags project
// @Accept json
// @Produce json
// @Param message body Project true "Project Info"
// @Success 201 {integer} string "OK"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/projects [post]
func Insert(c *gin.Context) {
	var body Project
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s := NewProjectService()

	inserted := s.Insert(body)

	c.JSON(http.StatusCreated, inserted)
}

// Update a project
// @Summary update a project
// @Tags project
// @Accept json
// @Produce json
// @Param project_id path string true "Project Id"
// @Param message body Project true "Project Info"
// @Success 200 {object} utils.DataResult "Project result"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/projects/{project_id} [put]
func Update(c *gin.Context) {
	var body Project
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("project_id")

	s := NewProjectService()

	updated, err := s.Update(id, body)

	if err != nil {
		c.JSON(err.StatusCode, err.Err)
	}

	c.JSON(http.StatusOK, updated)
}

// Delete a project
// @Summary delete a project
// @Tags project
// @Accept json
// @Produce json
// @Param project_id path string true "Project Id"
// @Success 204 {int} string "Empty response"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/projects/{project_id} [delete]
func Delete(c *gin.Context) {
	id := c.Param("project_id")

	s := NewProjectService()

	err := s.Delete(id)

	if err != nil {
		c.JSON(err.StatusCode, err.Err)
	}

	c.String(http.StatusNoContent, "")
}
