package dashboard

import (
	"log"

	"data-pad.app/data-api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DashboardService struct {
}

var dashboardRepository DashboardRepository

func NewDashboardService() DashboardService {
	dashboardRepository = NewDashboardRepository()

	s := DashboardService{}

	return s
}

func list(projectId string, filter interface{}, projection interface{},
	skip int64, limit int64, sort interface{}) []Dashboard {

	var results []Dashboard
	if err := dashboardRepository.List(projectId, filter, projection, skip, limit, sort, &results); err != nil {
		log.Panicf("Error while fecthing data: %v", err)
	}

	return results
}

func count(projectId string, filter interface{}) int64 {

	count, err := dashboardRepository.Count(projectId, filter)
	if err != nil {
		log.Panicf("Error while fecthing data: %v", err)
	}

	return count
}

func (x DashboardService) List(projectId string, filter interface{}, projection interface{},
	skip int64, limit int64, sort interface{}) utils.DataResult {
	items := make(chan []Dashboard)
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

func (x DashboardService) Get(dashboardId string) (Dashboard, *utils.DataPadError) {
	var result Dashboard
	err := dashboardRepository.Get(dashboardId, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (x DashboardService) Insert(document Dashboard) Dashboard {
	insertedId, err := dashboardRepository.Insert(document)

	if err != nil {
		log.Panicf("Error when inserting the doc: %v", err)
	}

	document.ID = insertedId.(primitive.ObjectID)

	return document
}

func (x DashboardService) Update(id string, document Dashboard) (Dashboard, *utils.DataPadError) {
	updatedDocument, err := dashboardRepository.Update(id, document)

	if err != nil {
		return document, err
	}

	return updatedDocument.(Dashboard), nil
}

func (x DashboardService) Delete(id string) *utils.DataPadError {
	err := dashboardRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
