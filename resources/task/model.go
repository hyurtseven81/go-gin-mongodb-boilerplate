package task

import "go.mongodb.org/mongo-driver/bson/primitive"

// Task - Model of a basic task
type Task struct {
	ID    primitive.ObjectID `json:"_id"`
	Title string             `json:"title"`
	Body  string             `json:"body"`
}
