package repository

import (
	"context"
	"log"
	"time"

	"data-pad.app/data-api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Timeout operations after N seconds
	timeout = 5
)

type BaseMongoRepository interface {
	List(query interface{}, projection interface{},
		skip int64, limit int64, sort interface{}) []interface{}
}

type MongoRepository struct {
	BaseMongoRepository
	collection string
}

func (x MongoRepository) List(query interface{}, projection interface{},
	skip int64, limit int64, sort interface{}) []interface{} {
	c := db.GetDB().Collection(x.collection)

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
		log.Fatalf("Find Error: %s", err)
	}
	var items []interface{}

	if err = cursor.All(ctx, &items); err != nil {
		log.Fatal(err)
	}

	return items
}
