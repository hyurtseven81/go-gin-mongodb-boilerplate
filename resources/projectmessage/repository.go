package projectmessage

import (
	"context"
	"time"

	"data-pad.app/data-api/db"
	"data-pad.app/data-api/repository"
	"data-pad.app/data-api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Timeout operations after N seconds
	timeout = 5
)

type ProjectMessageRepository struct {
	repository.MongoRepository
}

func NewProjectMessageRepository() ProjectMessageRepository {
	r := ProjectMessageRepository{}
	r.Collection = "messages"

	return r
}

func (x ProjectMessageRepository) List(projectId string, query interface{}, projection interface{},
	skip int64, limit int64, sort interface{}, results interface{}) *utils.DataPadError {

	c := db.GetDB().Collection(x.Collection)

	ctx, _ := context.WithTimeout(context.Background(), timeout*time.Second)

	opts := options.Find()
	if sort != nil {
		opts.SetSort(sort)
	}

	if projection != nil {
		opts.SetProjection(projection)
	}

	if skip > 0 {
		opts.SetSkip(skip)
	}

	if limit > 0 {
		if limit <= 500 {
			opts.SetLimit(limit)
		} else {
			opts.SetLimit(500)
		}
	} else {
		opts.SetLimit(500)
	}

	if query == nil {
		query = make(map[string]interface{})
	}

	projectObjectId, objectIdParseError := primitive.ObjectIDFromHex(projectId)
	if objectIdParseError != nil {
		return &utils.DataPadError{
			StatusCode: 500,
			Err:        objectIdParseError,
		}
	}

	_query := query.(map[string]interface{})
	if _query == nil {
		_query = make(map[string]interface{})
	}
	_query["project_id"] = projectObjectId

	cursor, err := c.Find(ctx, _query, opts)
	if err != nil {
		return &utils.DataPadError{
			StatusCode: 500,
			Err:        err,
		}
	}
	if err = cursor.All(ctx, results); err != nil {
		return &utils.DataPadError{
			StatusCode: 500,
			Err:        err,
		}
	}

	return nil
}

func (x ProjectMessageRepository) Count(projectId string, query interface{}) (int64, *utils.DataPadError) {
	c := db.GetDB().Collection(x.Collection)

	ctx, _ := context.WithTimeout(context.Background(), timeout*time.Second)

	projectObjectId, objectIdParseError := primitive.ObjectIDFromHex(projectId)
	if objectIdParseError != nil {
		return 0, &utils.DataPadError{
			StatusCode: 500,
			Err:        objectIdParseError,
		}
	}

	_query := query.(map[string]interface{})
	if _query == nil {
		_query = make(map[string]interface{})
	}
	_query["project_id"] = projectObjectId

	result, err := c.CountDocuments(ctx, _query)

	if err != nil {
		return 0, &utils.DataPadError{
			StatusCode: 500,
			Err:        err,
		}
	}

	return result, nil
}
