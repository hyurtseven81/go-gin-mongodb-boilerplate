package snippet

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

type SnippetItem struct {
	Name  string    `json:"name,omitempty" bson:"name,omitempty" binding:"required"`
	Time  time.Time `json:"time,omitempty" bson:"time,omitempty"`
	Value float64   `json:"value,omitempty" bson:"value,omitempty"`
}

type Snippet struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	ProjectID primitive.ObjectID `json:"project_id,omitempty" bson:"project_id,omitempty" binding:"required"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty" binding:"required"`
	Type      string             `json:"type,omitempty" bson:"type,omitempty"`
	ChartType string             `json:"chart_type,omitempty" bson:"chart_type,omitempty"`
	Value     float64            `json:"value,omitempty" bson:"value,omitempty"`
	Time      time.Time          `json:"time,omitempty" bson:"time,omitempty"`
	Symbol    string             `json:"symbol,omitempty" bson:"symbol,omitempty"`
	Data      []SnippetItem      `json:"data,omitempty" bson:"data,omitempty" binding:"required"`
	Sys       Sys                `json:"sys,omitempty" bson:"sys,omitempty"`
}
