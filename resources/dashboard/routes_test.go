package dashboard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"data-pad.app/data-api/middlewares"
	"data-pad.app/data-api/resources/project"
	"data-pad.app/data-api/test"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	project.AddProjectRoutes(v1)
	AddDashboardRoutes(v1)

	return router
}

func createProject(t *testing.T, router *gin.Engine) (int, map[string]interface{}) {
	body := `{
		"name": "test project"
	}`

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("POST", "/v1/projects", strings.NewReader(body))
	router.ServeHTTP(w, req)

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	return w.Code, got
}

func createDashboard(t *testing.T, router *gin.Engine) (int, map[string]interface{}) {
	_, createdProject := createProject(t, router)
	projectId := createdProject["_id"].(string)

	body := fmt.Sprintf(`{
		"title": "test title",
		"project_id": "%s"
	}`, projectId)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("POST", fmt.Sprintf("/v1/projects/%s/dashboards", projectId), strings.NewReader(body))
	router.ServeHTTP(w, req)

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	return w.Code, got
}

func TestDashboardsRouteShouldReturn200(t *testing.T) {
	test.Init("")

	defer test.Clear()

	router := setupRouter()

	_, createdProject := createProject(t, router)
	projectId := createdProject["_id"].(string)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("GET", fmt.Sprintf("/v1/projects/%s/dashboards", projectId), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDashboardsRouteShouldReturn401WithAuth(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.Use(middlewares.AuthMiddleware())
	AddDashboardRoutes(v1)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/projects/123/dashboards", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestCreateShouldSucceed(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := setupRouter()

	statusCode, got := createDashboard(t, router)

	assert.Equal(t, http.StatusCreated, statusCode)
	assert.Equal(t, "test title", got["title"])
}

func TestCreateShouldRaiseError(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := setupRouter()

	body := `{
		"body":  "test body"
	}`

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("POST", "/v1/projects/123/dashboards", strings.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateShouldSucceed(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := setupRouter()
	_, createdDashboard := createDashboard(t, router)

	projectId := createdDashboard["project_id"].(string)
	id := createdDashboard["_id"].(string)

	body := fmt.Sprintf(`{
		"title":  "new test title",
		"project_id": "%s"
	}`, projectId)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/v1/projects/%s/dashboards/%s", projectId, id), strings.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateShouldRaiseValidationError(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := setupRouter()
	_, createdDashboard := createDashboard(t, router)

	projectId := createdDashboard["project_id"].(string)
	id := createdDashboard["_id"].(string)

	body := `{
		"data": [{
			"name": "test item"
		}]
	}`

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/v1/projects/%s/dashboards/%s", projectId, id), strings.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateShouldRaiseNotFoundError(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := setupRouter()
	_, createdDashboard := createDashboard(t, router)

	projectId := createdDashboard["project_id"].(string)

	body := fmt.Sprintf(`{
		"title":  "new test title",
		"project_id": "%s"
	}`, projectId)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/v1/projects/%s/dashboards/60d0b16f8376024923f3d409", projectId), strings.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteShouldRaiseNotFoundError(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := setupRouter()
	_, createdDashboard := createDashboard(t, router)

	projectId := createdDashboard["project_id"].(string)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/v1/projects/%s/dashboards/60d0b16f8376024923f3d409", projectId), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteShouldSucceed(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := setupRouter()
	_, createdDashboard := createDashboard(t, router)

	projectId := createdDashboard["project_id"].(string)
	id := createdDashboard["_id"].(string)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/v1/projects/%s/dashboards/%s", projectId, id), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
