package task

import (
	"net/http"

	"github.com/gin-gonic/gin"

	pb "data-pad.app/data-api/resources/task/pb"
)

func Get(c *gin.Context) {
	_id := int64(1)
	title := "test"
	body := "test test test"

	task := &pb.Task{Id: &_id, Title: &title, Body: &body}

	c.ProtoBuf(http.StatusOK, task)
}
