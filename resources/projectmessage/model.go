package projectmessage

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProjectMessage struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	ProjectID primitive.ObjectID `json:"project_id,omitempty" bson:"project_id,omitempty" binding:"required"`
	From      string             `json:"from,omitempty" bson:"from,omitempty" binding:"required"`
	To        string             `json:"to,omitempty" bson:"to,omitempty"`
	Text      string             `json:"text,omitempty" bson:"text,omitempty" binding:"required"`
}
