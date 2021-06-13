package task

import "log"

type TaskService struct {
}

var taskRepository TaskRepository

func NewTaskService() TaskService {
	taskRepository = NewTaskRepository()

	s := TaskService{}

	return s
}

func (x TaskService) List(filter interface{}, projection interface{},
	skip int64, limit int64, sort interface{}, results interface{}) {

	if err := taskRepository.List(filter, projection, skip, limit, sort, results); err != nil {
		log.Panicf("Error while fecthing data: %v", err)
	}
}
