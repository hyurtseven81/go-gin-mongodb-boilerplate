package task

import (
	"log"

	"data-pad.app/data-api/utils"
)

type TaskService struct {
}

var taskRepository TaskRepository

func NewTaskService() TaskService {
	taskRepository = NewTaskRepository()

	s := TaskService{}

	return s
}

func list(filter interface{}, projection interface{},
	skip int64, limit int64, sort interface{}) []Task {

	var results []Task
	if err := taskRepository.List(filter, projection, skip, limit, sort, &results); err != nil {
		log.Panicf("Error while fecthing data: %v", err)
	}

	return results
}

func count(filter interface{}) int64 {

	count, err := taskRepository.Count(filter)
	if err != nil {
		log.Panicf("Error while fecthing data: %v", err)
	}

	return count
}

func (x TaskService) List(filter interface{}, projection interface{},
	skip int64, limit int64, sort interface{}) utils.DataResult {
	items := make(chan []Task)
	countResult := make(chan int64)

	go func() {
		items <- list(filter, projection, skip, limit, sort)
	}()

	go func() {
		countResult <- count(filter)
	}()
	dataResult := utils.NewDataresult(<-items, <-countResult)

	return dataResult
}
