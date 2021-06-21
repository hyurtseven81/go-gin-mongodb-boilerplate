package repository

import (
	"context"
	"fmt"
	"log"
	"testing"

	"data-pad.app/data-api/db"
	"data-pad.app/data-api/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func initTest() {
	test.Init("../test.env")
}

func TestFindShouldReturn0Items(t *testing.T) {
	initTest()

	defer test.Clear()

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	var items []bson.D
	err := mongo_repository.List(nil, nil, 0, 50, nil, &items)
	if err != nil {
		log.Fatal(err)
	}

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
		Collection: "test",
	}

	var items []bson.D
	err := mongo_repository.List(
		bson.D{{Key: "id", Value: bson.D{{Key: "$gt", Value: 10}}}},
		nil, 0, 10, nil, &items)
	if err != nil {
		log.Fatal(err)
	}

	doc := items[0]

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
		Collection: "test",
	}

	var items []bson.D
	err := mongo_repository.List(
		nil,
		bson.D{{Key: "id", Value: 1}},
		0, 10, nil, &items)
	if err != nil {
		log.Fatal(err)
	}

	doc := items[0]

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
		Collection: "test",
	}

	var items []bson.D
	mongo_repository.List(nil, nil, 0, 10, nil, &items)

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
		Collection: "test",
	}

	var items []bson.D
	err := mongo_repository.List(
		nil,
		nil,
		10, 10, nil, &items)
	if err != nil {
		log.Fatal(err)
	}

	doc := items[0]

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
		Collection: "test",
	}

	var items []bson.D
	err := mongo_repository.List(
		nil,
		nil,
		0, 10, bson.D{
			{Key: "id", Value: -1},
		}, &items)
	if err != nil {
		log.Fatal(err)
	}

	doc := items[0]

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
		Collection: "test",
	}

	var items []bson.D
	err := mongo_repository.List(nil, nil, 0, 530, nil, &items)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 500, len(items))
}

func TestCountShouldSucceed(t *testing.T) {
	initTest()

	defer test.Clear()

	var docs bson.A

	for i := 0; i < 500; i++ {
		doc := bson.D{
			{Key: "id", Value: int32(i)},
			{Key: "title", Value: fmt.Sprintf("Test Title %d", i)},
		}
		docs = append(docs, doc)
	}

	db.GetDB().Collection("test").InsertMany(context.Background(), docs)

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	count, err := mongo_repository.Count(nil)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, int64(500), count)
}

func TestCountShouldSucceedWithQuery(t *testing.T) {
	initTest()

	defer test.Clear()

	var docs bson.A

	for i := 0; i < 500; i++ {
		doc := bson.D{
			{Key: "id", Value: int32(i)},
			{Key: "title", Value: fmt.Sprintf("Test Title %d", i)},
		}
		docs = append(docs, doc)
	}

	db.GetDB().Collection("test").InsertMany(context.Background(), docs)

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	count, err := mongo_repository.Count(bson.D{
		{Key: "id", Value: bson.D{
			{Key: "$lt", Value: 100},
		}},
	})
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, int64(100), count)
}

func TestInsertShouldSucceed(t *testing.T) {
	initTest()

	defer test.Clear()

	doc := bson.D{
		{Key: "title", Value: "Test Title"},
	}

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	insertedId, err := mongo_repository.Insert(doc)
	if err != nil {
		log.Fatal(err)
	}

	assert.NotEqual(t, nil, insertedId)
}

func TestUpdateShouldSucceed(t *testing.T) {
	initTest()

	defer test.Clear()

	_id := primitive.NewObjectID().Hex()

	doc := bson.D{
		{Key: "_id", Value: _id},
		{Key: "title", Value: "Test Title"},
	}
	db.GetDB().Collection("test").InsertOne(context.Background(), doc)

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	newDoc := bson.D{
		{Key: "title", Value: "New Test Title"},
	}

	assert.NotPanics(
		t,
		func() { mongo_repository.Update(_id, newDoc) },
		"Document not found!",
	)
}

func TestUpdateShouldRaiseError(t *testing.T) {
	initTest()

	defer test.Clear()

	_id := primitive.NewObjectID().Hex()

	doc := bson.D{
		{Key: "_id", Value: _id},
		{Key: "title", Value: "Test Title"},
	}
	db.GetDB().Collection("test").InsertOne(context.Background(), doc)

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	newDoc := bson.D{
		{Key: "title", Value: "New Test Title"},
	}

	_, err := mongo_repository.Update(primitive.NewObjectID().Hex(), newDoc)

	assert.NotEqual(t, nil, err)
}

func TestDeleteShouldRaiseError(t *testing.T) {
	initTest()

	defer test.Clear()

	_id := primitive.NewObjectID()

	doc := bson.D{
		{Key: "_id", Value: _id},
		{Key: "title", Value: "Test Title"},
	}
	db.GetDB().Collection("test").InsertOne(context.Background(), doc)

	mongo_repository := MongoRepository{
		Collection: "test",
	}
	err := mongo_repository.Delete(primitive.NewObjectID().Hex())

	assert.NotEqual(t, nil, err)
}

func TestDeleteShouldSucceed(t *testing.T) {
	initTest()

	defer test.Clear()

	_id := primitive.NewObjectID().Hex()

	doc := bson.D{
		{Key: "_id", Value: _id},
		{Key: "title", Value: "Test Title"},
	}
	db.GetDB().Collection("test").InsertOne(context.Background(), doc)

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	assert.NotPanics(
		t,
		func() { mongo_repository.Delete(_id) },
		"Document not found!",
	)
}
