package snippet

import (
	"log"

	"data-pad.app/data-api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SnippetService struct {
}

var snippetRepository SnippetRepository

func NewSnippetService() SnippetService {
	snippetRepository = NewSnippetRepository()

	s := SnippetService{}

	return s
}

func list(projectId string, filter interface{}, projection interface{},
	skip int64, limit int64, sort interface{}) []Snippet {

	var results []Snippet
	if err := snippetRepository.List(projectId, filter, projection, skip, limit, sort, &results); err != nil {
		log.Panicf("Error while fecthing data: %v", err)
	}

	return results
}

func count(projectId string, filter interface{}) int64 {

	count, err := snippetRepository.Count(projectId, filter)
	if err != nil {
		log.Panicf("Error while fecthing data: %v", err)
	}

	return count
}

func (x SnippetService) List(projectId string, filter interface{}, projection interface{},
	skip int64, limit int64, sort interface{}) utils.DataResult {
	items := make(chan []Snippet)
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

func (x SnippetService) Insert(document Snippet) Snippet {
	insertedId, err := snippetRepository.Insert(document)

	if err != nil {
		log.Panicf("Error when inserting the doc: %v", err)
	}

	document.ID = insertedId.(primitive.ObjectID)

	return document
}

func (x SnippetService) Update(id string, document Snippet) (Snippet, *utils.DataPadError) {
	updatedDocument, err := snippetRepository.Update(id, document)

	if err != nil {
		return document, err
	}

	return updatedDocument.(Snippet), nil
}

func (x SnippetService) Delete(id string) *utils.DataPadError {
	err := snippetRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
