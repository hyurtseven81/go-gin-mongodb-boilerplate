package repository

import (
	"context"
	"errors"
	"time"

	"data-pad.app/data-api/db"
	"data-pad.app/data-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Timeout operations after N seconds
	timeout = 5
)

type baseMongoRepository interface {
	List(query interface{}, projection interface{},
		skip int64, limit int64, sort interface{}, results interface{}) error
}

type MongoRepository struct {
	baseMongoRepository
	Collection string
}

func (x MongoRepository) List(query interface{}, projection interface{},
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
		query = bson.M{}
	}

	cursor, err := c.Find(ctx, query, opts)
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

func (x MongoRepository) Count(query interface{}) (int64, *utils.DataPadError) {
	c := db.GetDB().Collection(x.Collection)

	ctx, _ := context.WithTimeout(context.Background(), timeout*time.Second)

	if query == nil {
		query = bson.M{}
	}

	result, err := c.CountDocuments(ctx, query)

	if err != nil {
		return 0, &utils.DataPadError{
			StatusCode: 500,
			Err:        err,
		}
	}

	return result, nil
}

func (x MongoRepository) Insert(document interface{}) (interface{}, *utils.DataPadError) {
	c := db.GetDB().Collection(x.Collection)

	ctx, _ := context.WithTimeout(context.Background(), timeout*time.Second)

	res, err := c.InsertOne(ctx, document)

	if err != nil {
		return nil, &utils.DataPadError{
			StatusCode: 500,
			Err:        err,
		}
	}

	return res.InsertedID, nil
}

func (x MongoRepository) Update(id string, document interface{}) (interface{}, *utils.DataPadError) {
	c := db.GetDB().Collection(x.Collection)

	ctx, _ := context.WithTimeout(context.Background(), timeout*time.Second)

	objectId, objectIdParseError := primitive.ObjectIDFromHex(id)
	var (
		updateResult *mongo.UpdateResult
		err          error
	)
	if objectIdParseError == nil {
		updateResult, err = c.ReplaceOne(ctx, bson.D{{Key: "_id", Value: objectId}}, document)
	} else {
		updateResult, err = c.ReplaceOne(ctx, bson.D{{Key: "_id", Value: id}}, document)
	}

	if err != nil {
		return nil, &utils.DataPadError{
			StatusCode: 500,
			Err:        err,
		}
	}

	if updateResult.MatchedCount == 0 {
		return nil, &utils.DataPadError{
			StatusCode: 404,
			Err:        errors.New("document not found"),
		}
	}

	return document, nil
}

func (x MongoRepository) Delete(id string) *utils.DataPadError {
	c := db.GetDB().Collection(x.Collection)

	ctx, _ := context.WithTimeout(context.Background(), timeout*time.Second)

	var (
		deleteResult *mongo.DeleteResult
		err          error
	)

	objectId, objectIdParseError := primitive.ObjectIDFromHex(id)
	if objectIdParseError != nil {
		deleteResult, err = c.DeleteOne(ctx, bson.D{{Key: "_id", Value: objectId}})
	} else {
		deleteResult, err = c.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})
	}

	if err != nil {
		return &utils.DataPadError{
			StatusCode: 500,
			Err:        err,
		}
	}

	if deleteResult.DeletedCount == 0 {
		return &utils.DataPadError{
			StatusCode: 404,
			Err:        errors.New("document not found"),
		}
	}

	return nil
}
