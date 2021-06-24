package dashboard

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sys struct {
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	CreatedBy string    `json:"created_by,omitempty" bson:"created_by,omitempty"`
	UpdatedBy string    `json:"updated_by,omitempty" bson:"updated_by,omitempty"`
}

type Dashboard struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	ProjectID primitive.ObjectID `json:"project_id,omitempty" bson:"project_id,omitempty" binding:"required"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty" binding:"required"`
	Sys       Sys                `json:"sys,omitempty" bson:"sys,omitempty"`
}
