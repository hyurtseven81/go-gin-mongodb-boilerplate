package repository

import (
	"context"
	"fmt"
	"testing"

	"data-pad.app/data-api/db"
	"data-pad.app/data-api/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func initTest() {
	test.Init("../test.env")
}

func TestFindShouldReturn0Items(t *testing.T) {
	initTest()

	defer test.Clear()

	mongo_repository := MongoRepository{
		collection: "test",
	}

	items := mongo_repository.List(nil, nil, 0, 50, nil)

	assert.Equal(t, 0, len(items))
}

func TestFindShouldReturnFiltered(t *testing.T) {
	initTest()

	defer test.Clear()

	var docs bson.A

	for i := 0; i <= 510; i++ {
		doc := bson.D{
			{Key: "id", Value: int32(i)},
			{Key: "title", Value: fmt.Sprintf("Test Title %d", i)},
		}
		docs = append(docs, doc)
	}

	db.GetDB().Collection("test").InsertMany(context.Background(), docs)

	mongo_repository := MongoRepository{
		collection: "test",
	}

	items := mongo_repository.List(
		bson.D{{Key: "id", Value: bson.D{{Key: "$gt", Value: 10}}}},
		nil, 0, 10, nil,
	)

	firstItem := items[0]
	doc, _ := toDoc(firstItem)

	assert.Equal(t, 10, len(items))
	assert.Equal(t, int32(11), doc.Map()["id"])
}

func TestFindShouldReturnProjected(t *testing.T) {
	initTest()

	defer test.Clear()

	var docs bson.A

	for i := 0; i <= 510; i++ {
		doc := bson.D{
			{Key: "id", Value: int32(i)},
			{Key: "title", Value: fmt.Sprintf("Test Title %d", i)},
		}
		docs = append(docs, doc)
	}

	db.GetDB().Collection("test").InsertMany(context.Background(), docs)

	mongo_repository := MongoRepository{
		collection: "test",
	}

	items := mongo_repository.List(
		nil,
		bson.D{{Key: "id", Value: 1}},
		0, 10, nil,
	)

	firstItem := items[0]
	doc, _ := toDoc(firstItem)

	assert.Equal(t, 10, len(items))
	assert.Equal(t, int32(0), doc.Map()["id"])
	assert.Equal(t, nil, doc.Map()["title"])
}

func TestFindShouldReturnLimited(t *testing.T) {
	initTest()

	defer test.Clear()

	var docs bson.A

	for i := 0; i < 510; i++ {
		doc := bson.D{
			{Key: "id", Value: int32(i)},
			{Key: "title", Value: fmt.Sprintf("Test Title %d", i)},
		}
		docs = append(docs, doc)
	}

	db.GetDB().Collection("test").InsertMany(context.Background(), docs)

	mongo_repository := MongoRepository{
		collection: "test",
	}

	items := mongo_repository.List(nil, nil, 0, 10, nil)

	assert.Equal(t, 10, len(items))
}

func TestFindShouldReturnSkipped(t *testing.T) {
	initTest()

	defer test.Clear()

	var docs bson.A

	for i := 0; i < 510; i++ {
		doc := bson.D{
			{Key: "id", Value: int32(i)},
			{Key: "title", Value: fmt.Sprintf("Test Title %d", i)},
		}
		docs = append(docs, doc)
	}

	db.GetDB().Collection("test").InsertMany(context.Background(), docs)

	mongo_repository := MongoRepository{
		collection: "test",
	}

	items := mongo_repository.List(nil, nil, 10, 10, nil)

	firstItem := items[0]
	doc, _ := toDoc(firstItem)

	assert.Equal(t, 10, len(items))
	assert.Equal(t, int32(10), doc.Map()["id"])
}

func TestFindShouldReturnSorted(t *testing.T) {
	initTest()

	defer test.Clear()

	var docs bson.A

	for i := 0; i <= 510; i++ {
		doc := bson.D{
			{Key: "id", Value: int32(i)},
			{Key: "title", Value: fmt.Sprintf("Test Title %d", i)},
		}
		docs = append(docs, doc)
	}

	db.GetDB().Collection("test").InsertMany(context.Background(), docs)

	mongo_repository := MongoRepository{
		collection: "test",
	}

	items := mongo_repository.List(nil, nil, 0, 10, bson.D{
		{Key: "id", Value: -1},
	})

	firstItem := items[0]
	doc, _ := toDoc(firstItem)

	assert.Equal(t, 10, len(items))
	assert.Equal(t, int32(510), doc.Map()["id"])
}

func TestFindShouldReturn500ItemsWithLimitOverThan500(t *testing.T) {
	initTest()

	defer test.Clear()

	var docs bson.A

	for i := 0; i < 510; i++ {
		doc := bson.D{
			{Key: "id", Value: int32(i)},
			{Key: "title", Value: fmt.Sprintf("Test Title %d", i)},
		}
		docs = append(docs, doc)
	}

	db.GetDB().Collection("test").InsertMany(context.Background(), docs)

	mongo_repository := MongoRepository{
		collection: "test",
	}

	items := mongo_repository.List(nil, nil, 0, 530, nil)

	assert.Equal(t, 500, len(items))
}

func toDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
