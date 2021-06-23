package dashboard

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dashboard struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	ProjectID primitive.ObjectID `json:"project_id,omitempty" bson:"project_id,omitempty" binding:"required"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty" binding:"required"`
}
