package repository

import (
	"context"
	"errors"
	"time"

	"data-pad.app/data-api/db"
	"go.mongodb.org/mongo-driver/bson"
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
	skip int64, limit int64, sort interface{}, results interface{}) error {
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
		return err
	}
	if err = cursor.All(ctx, results); err != nil {
		return err
	}

	return nil
}

func (x MongoRepository) Count(query interface{}) (int64, error) {
	c := db.GetDB().Collection(x.Collection)

	ctx, _ := context.WithTimeout(context.Background(), timeout*time.Second)

	if query == nil {
		query = bson.M{}
	}

	result, err := c.CountDocuments(ctx, query)

	return result, err
}

func (x MongoRepository) Insert(document interface{}) (interface{}, error) {
	c := db.GetDB().Collection(x.Collection)

	ctx, _ := context.WithTimeout(context.Background(), timeout*time.Second)

	res, err := c.InsertOne(ctx, document)

	return res.InsertedID, err
}

func (x MongoRepository) Update(id interface{}, document interface{}) (interface{}, error) {
	c := db.GetDB().Collection(x.Collection)

	ctx, _ := context.WithTimeout(context.Background(), timeout*time.Second)

	updateResult, err := c.ReplaceOne(ctx, bson.D{{Key: "_id", Value: id}}, document)

	if updateResult.MatchedCount == 0 {
		return nil, errors.New("document not found")
	}

	return document, err
}

func (x MongoRepository) Delete(id interface{}) error {
	c := db.GetDB().Collection(x.Collection)

	ctx, _ := context.WithTimeout(context.Background(), timeout*time.Second)

	deleteResult, err := c.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})

	if deleteResult.DeletedCount == 0 {
		return errors.New("document not found")
	}

	return err
}
