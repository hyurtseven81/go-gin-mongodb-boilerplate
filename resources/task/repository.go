package task

import (
	"data-pad.app/data-api/repository"
)

type TaskRepository struct {
	repository.MongoRepository
}

func NewTaskRepository() TaskRepository {
	r := TaskRepository{}
	r.Collection = "tasks"

	return r
}
