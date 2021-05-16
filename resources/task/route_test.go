package task

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"data-pad.app/data-api/test"
)

func TestTasksRouteShouldReturn200(t *testing.T) {
	test.Init()

	router := gin.Default()
	v1 := router.Group("/v1")
	AddTaskRoutes(v1)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/tasks", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTasksRouteShouldInsertTask(t *testing.T) {
	test.Init()

	router := gin.Default()
	v1 := router.Group("/v1")
	AddTaskRoutes(v1)

	body := `{ 
		"title": "test title",
		"body":  "test body" 
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/tasks", strings.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
