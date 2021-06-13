package task

import "go.mongodb.org/mongo-driver/bson/primitive"

// Task - Model of a basic task
type Task struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	Title string             `json:"title,omitempty" bson:"title,omitempty" binding:"required"`
	Body  string             `json:"body,omitempty" bson:"body,omitempty"`
}
