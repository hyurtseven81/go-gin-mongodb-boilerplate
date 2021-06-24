package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags healthcheck
// @Accept */*
// @Produce json
// @Success 200 {object} string
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
