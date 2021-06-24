package dashboard

import (
	"net/http"

	"data-pad.app/data-api/utils"
	"github.com/gin-gonic/gin"
)

// Get Dashboards
// @Summary get all dashboards under a project
// @Tags dashboard
// @Accept json
// @Produce json
// @Param project_id path string true "Project Id"
// @Param filter query string false "string as JSON / used for filtering the results"
// @Param select query string false "string as JSON / used for projecting the results"
// @Param sort query string false "string as JSON / used for sorting the results"
// @Param skip query int false "used for skipping the results"
// @Param limit query int false "used for limitting the results"
// @Success 200 {object} utils.DataResult "Dashboard result"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "error details"
// @Security ApiKeyAuth
// @Router /v1/projects/{project_id}/dashboards [get]
func List(c *gin.Context) {
	var query utils.Query
	c.BindQuery(&query)

	projectId := c.Param("project_id")
	projection := query.GetSelect()
	filter := query.GetFilter()
	sort := query.GetSort()

	s := NewDashboardService()

	dataResult := s.List(projectId, filter, projection, int64(query.Skip), int64(query.Limit), sort)

	c.JSON(http.StatusOK, dataResult)
}

// Get Dashboard By Id
// @Summary get a dashboard under a project
// @Tags dashboard
// @Accept json
// @Produce json
// @Param project_id path string true "Project Id"
// @Param dashboard_id path string true "Dashboard Id"
// @Success 200 {object} Dashboard "Dashboard result"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Document not found"
// @Security ApiKeyAuth
// @Router /v1/projects/{project_id}/dashboards/{dashboard_id} [get]
func Get(c *gin.Context) {
	dashboardId := c.Param("dashboard_id")

	s := NewDashboardService()
	result, err := s.Get(dashboardId)

	if err != nil {
		c.JSON(err.StatusCode, err.Err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Insert new dashboard
// @Summary insert new dashboard
// @Tags dashboard
// @Accept json
// @Produce json
// @Param project_id path string true "Project Id"
// @Param message body Dashboard true "Dashboard Info"
// @Success 201 {integer} string "OK"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/projects/{project_id}/dashboards [post]
func Insert(c *gin.Context) {
	var body Dashboard
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s := NewDashboardService()

	inserted := s.Insert(body)

	c.JSON(http.StatusCreated, inserted)
}

// Update a dashboard
// @Summary update a dashboard
// @Tags dashboard
// @Accept json
// @Produce json
// @Param project_id path string true "Project Id"
// @Param dashboard_id path string true "Dashboard Id"
// @Param message body Dashboard true "Dashboard Info"
// @Success 200 {object} utils.DataResult "Dashboard result"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/projects/{project_id}/dashboards/{dashboard_id} [put]
func Update(c *gin.Context) {
	var body Dashboard
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("dashboard_id")

	s := NewDashboardService()

	updated, err := s.Update(id, body)

	if err != nil {
		c.JSON(err.StatusCode, err.Err)
	}

	c.JSON(http.StatusOK, updated)
}

// Delete a dashboard
// @Summary delete a dashboard
// @Tags dashboard
// @Accept json
// @Produce json
// @Param project_id path string true "Project Id"
// @Param dashboard_id path string true "Dashboard Id"
// @Success 204 {int} string "Empty response"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/projects/{project_id}/dashboards/{dashboard_id} [delete]
func Delete(c *gin.Context) {
	id := c.Param("dashboard_id")

	s := NewDashboardService()

	err := s.Delete(id)

	if err != nil {
		c.JSON(err.StatusCode, err.Err)
	}

	c.String(http.StatusNoContent, "")
}
