package projectmessage

import (
	"log"

	"data-pad.app/data-api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectMessageService struct {
}

var projectMessageRepository ProjectMessageRepository

func NewProjectMessageService() ProjectMessageService {
	projectMessageRepository = NewProjectMessageRepository()

	s := ProjectMessageService{}

	return s
}

func list(projectId string, filter interface{}, projection interface{},
	skip int64, limit int64, sort interface{}) []ProjectMessage {

	var results []ProjectMessage
	if err := projectMessageRepository.List(projectId, filter, projection, skip, limit, sort, &results); err != nil {
		log.Panicf("Error while fecthing data: %v", err)
	}

	return results
}

func count(projectId string, filter interface{}) int64 {

	count, err := projectMessageRepository.Count(projectId, filter)
	if err != nil {
		log.Panicf("Error while fecthing data: %v", err)
	}

	return count
}

func (x ProjectMessageService) List(projectId string, filter interface{}, projection interface{},
	skip int64, limit int64, sort interface{}) utils.DataResult {
	items := make(chan []ProjectMessage)
	countResult := make(chan int64)

	go func() {
		items <- list(projectId, filter, projection, skip, limit, sort)
	}()

	go func() {
		countResult <- count(projectId, filter)
	}()
	dataResult := utils.NewDataresult(<-items, <-countResult)

	return dataResult
}

func (x ProjectMessageService) Get(projectMessageId string) (ProjectMessage, *utils.DataPadError) {
	var result ProjectMessage
	err := projectMessageRepository.Get(projectMessageId, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (x ProjectMessageService) Insert(document ProjectMessage) ProjectMessage {
	insertedId, err := projectMessageRepository.Insert(document)

	if err != nil {
		log.Panicf("Error when inserting the doc: %v", err)
	}

	document.ID = insertedId.(primitive.ObjectID)

	return document
}

func (x ProjectMessageService) Update(id string, document ProjectMessage) (ProjectMessage, *utils.DataPadError) {
	updatedDocument, err := projectMessageRepository.Update(id, document)

	if err != nil {
		return document, err
	}

	return updatedDocument.(ProjectMessage), nil
}

func (x ProjectMessageService) Delete(id string) *utils.DataPadError {
	err := projectMessageRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
