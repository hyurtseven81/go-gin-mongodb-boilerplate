package project

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
type Project struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	Name string             `json:"name,omitempty" bson:"name,omitempty" binding:"required"`
	Sys  Sys                `json:"sys,omitempty" bson:"sys,omitempty"`
}
