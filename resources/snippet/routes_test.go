package snippet

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
	AddSnippetRoutes(v1)

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

func createSnippet(t *testing.T, router *gin.Engine) (int, map[string]interface{}) {
	_, createdProject := createProject(t, router)
	projectId := createdProject["_id"].(string)

	body := fmt.Sprintf(`{
		"title": "test title",
		"project_id": "%s",
		"data":  [{
			"name": "test name"
		}]
	}`, projectId)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("POST", fmt.Sprintf("/v1/projects/%s/snippets", projectId), strings.NewReader(body))
	router.ServeHTTP(w, req)

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	return w.Code, got
}

func TestSnippetsRouteShouldReturn200(t *testing.T) {
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
	req, _ := http.NewRequest("GET", fmt.Sprintf("/v1/projects/%s/snippets", projectId), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSnippetsRouteShouldReturn401WithAuth(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.Use(middlewares.AuthMiddleware())
	AddSnippetRoutes(v1)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/projects/123/snippets", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestCreateShouldSucceed(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := setupRouter()

	statusCode, got := createSnippet(t, router)

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
	req, _ := http.NewRequest("POST", "/v1/projects/123/snippets", strings.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateShouldSucceed(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := setupRouter()
	_, createdSnippet := createSnippet(t, router)

	projectId := createdSnippet["project_id"].(string)
	id := createdSnippet["_id"].(string)

	body := fmt.Sprintf(`{
		"title":  "new test title",
		"project_id": "%s",
		"data": [{
			"name": "test item"
		}]
	}`, projectId)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/v1/projects/%s/snippets/%s", projectId, id), strings.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateShouldRaiseValidationError(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := setupRouter()
	_, createdSnippet := createSnippet(t, router)

	projectId := createdSnippet["project_id"].(string)
	id := createdSnippet["_id"].(string)

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
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/v1/projects/%s/snippets/%s", projectId, id), strings.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateShouldRaiseNotFoundError(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := setupRouter()
	_, createdSnippet := createSnippet(t, router)

	projectId := createdSnippet["project_id"].(string)

	body := fmt.Sprintf(`{
		"title":  "new test title",
		"project_id": "%s",
		"data": [{
			"name": "test item"
		}]
	}`, projectId)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/v1/projects/%s/snippets/60d0b16f8376024923f3d409", projectId), strings.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteShouldRaiseNotFoundError(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := setupRouter()
	_, createdSnippet := createSnippet(t, router)

	projectId := createdSnippet["project_id"].(string)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/v1/projects/%s/snippets/60d0b16f8376024923f3d409", projectId), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteShouldSucceed(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := setupRouter()
	_, createdSnippet := createSnippet(t, router)

	projectId := createdSnippet["project_id"].(string)
	id := createdSnippet["_id"].(string)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/v1/projects/%s/snippets/%s", projectId, id), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
