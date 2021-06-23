package project

import (
	"log"

	"data-pad.app/data-api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectService struct {
}

var projectRepository ProjectRepository

func NewProjectService() ProjectService {
	projectRepository = NewProjectRepository()

	s := ProjectService{}

	return s
}

func list(filter interface{}, projection interface{},
	skip int64, limit int64, sort interface{}) []Project {

	var results []Project
	if err := projectRepository.List(filter, projection, skip, limit, sort, &results); err != nil {
		log.Panicf("Error while fecthing data: %v", err)
	}

	return results
}

func count(filter interface{}) int64 {

	count, err := projectRepository.Count(filter)
	if err != nil {
		log.Panicf("Error while fecthing data: %v", err)
	}

	return count
}

func (x ProjectService) List(filter interface{}, projection interface{},
	skip int64, limit int64, sort interface{}) utils.DataResult {
	items := make(chan []Project)
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

func (x ProjectService) Insert(document Project) Project {
	insertedId, err := projectRepository.Insert(document)

	if err != nil {
		log.Panicf("Error when inserting the doc: %v", err)
	}

	document.ID = insertedId.(primitive.ObjectID)

	return document
}

func (x ProjectService) Update(id string, document Project) (Project, *utils.DataPadError) {
	updatedDocument, err := projectRepository.Update(id, document)

	if err != nil {
		return document, err
	}

	return updatedDocument.(Project), nil
}

func (x ProjectService) Delete(id string) *utils.DataPadError {
	err := projectRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
