package snippet

import (
	"net/http"

	"data-pad.app/data-api/utils"
	"github.com/gin-gonic/gin"
)

// Get Snippets
// @Summary get all snippets under a project
// @Tags snippet
// @Accept json
// @Produce json
// @Param project_id path string true "Project Id"
// @Param filter query string false "string as JSON / used for filtering the results"
// @Param select query string false "string as JSON / used for projecting the results"
// @Param sort query string false "string as JSON / used for sorting the results"
// @Param skip query int false "used for skipping the results"
// @Param limit query int false "used for limitting the results"
// @Success 200 {object} utils.DataResult "Snippet result"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "error details"
// @Security ApiKeyAuth
// @Router /v1/projects/{project_id}/snippets [get]
func List(c *gin.Context) {
	var query utils.Query
	c.BindQuery(&query)

	projectId := c.Param("project_id")
	projection := query.GetSelect()
	filter := query.GetFilter()
	sort := query.GetSort()

	s := NewSnippetService()

	dataResult := s.List(projectId, filter, projection, int64(query.Skip), int64(query.Limit), sort)

	c.JSON(http.StatusOK, dataResult)
}

// Get Snippet By Id
// @Summary get a snippet under a project
// @Tags snippet
// @Accept json
// @Produce json
// @Param project_id path string true "Project Id"
// @Param snippet_id path string true "Snippet Id"
// @Success 200 {object} Snippet "Snippet result"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Document not found"
// @Security ApiKeyAuth
// @Router /v1/projects/{project_id}/snippets/{snippet_id} [get]
func Get(c *gin.Context) {
	snippetId := c.Param("snippet_id")

	s := NewSnippetService()
	result, err := s.Get(snippetId)

	if err != nil {
		c.JSON(err.StatusCode, err.Err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Insert new snippet
// @Summary insert new snippet
// @Tags snippet
// @Accept json
// @Produce json
// @Param project_id path string true "Project Id"
// @Param message body Snippet true "Snippet Info"
// @Success 201 {integer} string "OK"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/projects/{project_id}/snippets [post]
func Insert(c *gin.Context) {
	var body Snippet
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s := NewSnippetService()

	inserted := s.Insert(body)

	c.JSON(http.StatusCreated, inserted)
}

// Update a snippet
// @Summary update a snippet
// @Tags snippet
// @Accept json
// @Produce json
// @Param project_id path string true "Project Id"
// @Param snippet_id path string true "Snippet Id"
// @Param message body Snippet true "Snippet Info"
// @Success 200 {object} utils.DataResult "Snippet result"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/projects/{project_id}/snippets/{snippet_id} [put]
func Update(c *gin.Context) {
	var body Snippet
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("snippet_id")

	s := NewSnippetService()

	updated, err := s.Update(id, body)

	if err != nil {
		c.JSON(err.StatusCode, err.Err)
	}

	c.JSON(http.StatusOK, updated)
}

// Delete a snippet
// @Summary delete a snippet
// @Tags snippet
// @Accept json
// @Produce json
// @Param project_id path string true "Project Id"
// @Param snippet_id path string true "Snippet Id"
// @Success 204 {int} string "Empty response"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/projects/{project_id}/snippets/{snippet_id} [delete]
func Delete(c *gin.Context) {
	id := c.Param("snippet_id")

	s := NewSnippetService()

	err := s.Delete(id)

	if err != nil {
		c.JSON(err.StatusCode, err.Err)
	}

	c.String(http.StatusNoContent, "")
}
