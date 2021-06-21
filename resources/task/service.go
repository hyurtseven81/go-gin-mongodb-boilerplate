package task

import (
	"log"

	"data-pad.app/data-api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (x TaskService) Insert(document Task) Task {
	insertedId, err := taskRepository.Insert(document)

	if err != nil {
		log.Panicf("Error when inserting the doc: %v", err)
	}

	document.ID = insertedId.(primitive.ObjectID)

	return document
}

func (x TaskService) Update(id string, document Task) (Task, *utils.DataPadError) {
	updatedDocument, err := taskRepository.Update(id, document)

	if err != nil {
		return document, err
	}

	return updatedDocument.(Task), nil
}

func (x TaskService) Delete(id string) *utils.DataPadError {
	err := taskRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
