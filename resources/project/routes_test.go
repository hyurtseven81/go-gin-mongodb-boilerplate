package project

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"data-pad.app/data-api/middlewares"
	"data-pad.app/data-api/test"
)

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

func TestProjectsRouteShouldReturn200(t *testing.T) {
	test.Init("")

	defer test.Clear()

	router := gin.Default()
	v1 := router.Group("/v1")
	AddProjectRoutes(v1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("GET", "/v1/projects", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestProjectsRouteShouldReturn401WithAuth(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.Use(middlewares.AuthMiddleware())
	AddProjectRoutes(v1)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/projects", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestCreateShouldSucceed(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := gin.Default()
	v1 := router.Group("/v1")
	AddProjectRoutes(v1)

	statusCode, got := createProject(t, router)

	assert.Equal(t, http.StatusCreated, statusCode)
	assert.Equal(t, "test project", got["name"])
}

func TestCreateShouldRaiseError(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := gin.Default()
	v1 := router.Group("/v1")
	AddProjectRoutes(v1)

	body := `{
		"body":  "test body"
	}`

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("POST", "/v1/projects", strings.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateShouldSucceed(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := gin.Default()
	v1 := router.Group("/v1")
	AddProjectRoutes(v1)

	createdStatusCode, got := createProject(t, router)

	assert.Equal(t, http.StatusCreated, createdStatusCode)

	id := got["_id"].(string)

	body := `{
		"name":  "new test project"
	}`

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("PUT", "/v1/projects/"+id, strings.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateShouldRaiseValidationError(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := gin.Default()
	v1 := router.Group("/v1")
	AddProjectRoutes(v1)

	createdStatusCode, got := createProject(t, router)

	assert.Equal(t, http.StatusCreated, createdStatusCode)

	id := got["_id"].(string)

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
	req, _ := http.NewRequest("PUT", "/v1/projects/"+id, strings.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateShouldRaiseNotFoundError(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := gin.Default()
	v1 := router.Group("/v1")
	AddProjectRoutes(v1)

	createdStatusCode, _ := createProject(t, router)

	assert.Equal(t, http.StatusCreated, createdStatusCode)

	body := `{
		"name": "test project"
	}`

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("PUT", "/v1/projects/60d0b16f8376024923f3d409", strings.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteShouldRaiseNotFoundError(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := gin.Default()
	v1 := router.Group("/v1")
	AddProjectRoutes(v1)

	createdStatusCode, _ := createProject(t, router)

	assert.Equal(t, http.StatusCreated, createdStatusCode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("DELETE", "/v1/projects/60d0b16f8376024923f3d409", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteShouldSucceed(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := gin.Default()
	v1 := router.Group("/v1")
	AddProjectRoutes(v1)

	createdStatusCode, got := createProject(t, router)

	assert.Equal(t, http.StatusCreated, createdStatusCode)

	id := got["_id"].(string)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("DELETE", "/v1/projects/"+id, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
