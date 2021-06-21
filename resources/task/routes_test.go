package task

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

func TestTasksRouteShouldReturn200(t *testing.T) {
	test.Init("")

	defer test.Clear()

	router := gin.Default()
	v1 := router.Group("/v1")
	AddTaskRoutes(v1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("GET", "/v1/tasks", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTasksRouteShouldReturn401WithAuth(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.Use(middlewares.AuthMiddleware())
	AddTaskRoutes(v1)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/tasks", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestCreateShouldSucceed(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := gin.Default()
	v1 := router.Group("/v1")
	AddTaskRoutes(v1)

	body := `{
		"title": "test title",
		"body":  "test body"
	}`

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("POST", "/v1/tasks", strings.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "test title", got["title"])
}

func TestCreateShouldRaiseError(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := gin.Default()
	v1 := router.Group("/v1")
	AddTaskRoutes(v1)

	body := `{
		"body":  "test body"
	}`

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("POST", "/v1/tasks", strings.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateShouldSucceed(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := gin.Default()
	v1 := router.Group("/v1")
	AddTaskRoutes(v1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})

	body := `{
		"title": "test title",
		"body": "test body"
	}`
	insertReq, _ := http.NewRequest("POST", "/v1/tasks", strings.NewReader(body))
	router.ServeHTTP(w, insertReq)

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	id := got["_id"].(string)

	body = `{
		"title":  "new test title",
		"body":  "test body"
	}`

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("PUT", "/v1/tasks/"+id, strings.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateShouldRaiseValidationError(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := gin.Default()
	v1 := router.Group("/v1")
	AddTaskRoutes(v1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})

	body := `{
		"title": "test title",
		"body": "test body"
	}`
	insertReq, _ := http.NewRequest("POST", "/v1/tasks", strings.NewReader(body))
	router.ServeHTTP(w, insertReq)

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	id := got["_id"].(string)

	body = `{
		"body":  "test body"
	}`

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("PUT", "/v1/tasks/"+id, strings.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateShouldRaiseNotFoundError(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := gin.Default()
	v1 := router.Group("/v1")
	AddTaskRoutes(v1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})

	body := `{
		"title": "test title",
		"body": "test body"
	}`
	insertReq, _ := http.NewRequest("POST", "/v1/tasks", strings.NewReader(body))
	router.ServeHTTP(w, insertReq)

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	body = `{
		"title": "test title",
		"body":  "test body"
	}`

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("PUT", "/v1/tasks/60d0b16f8376024923f3d409", strings.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteShouldRaiseNotFoundError(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := gin.Default()
	v1 := router.Group("/v1")
	AddTaskRoutes(v1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})

	body := `{
		"title": "test title",
		"body": "test body"
	}`
	insertReq, _ := http.NewRequest("POST", "/v1/tasks", strings.NewReader(body))
	router.ServeHTTP(w, insertReq)

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("DELETE", "/v1/tasks/60d0b16f8376024923f3d409", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteShouldSucceed(t *testing.T) {
	test.Init("")
	defer test.Clear()

	router := gin.Default()
	v1 := router.Group("/v1")
	AddTaskRoutes(v1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})

	body := `{
		"title": "test title",
		"body": "test body"
	}`
	insertReq, _ := http.NewRequest("POST", "/v1/tasks", strings.NewReader(body))
	router.ServeHTTP(w, insertReq)

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	id := got["_id"].(string)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Set("User", gin.H{
		"id":       "test",
		"username": "test",
	})
	req, _ := http.NewRequest("DELETE", "/v1/tasks/"+id, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
