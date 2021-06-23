package snippet

import "go.mongodb.org/mongo-driver/bson/primitive"

type SnippetItem struct {
	Name  string             `json:"name,omitempty" bson:"name,omitempty" binding:"required"`
	Time  primitive.DateTime `json:"time,omitempty" bson:"time,omitempty"`
	Value float64            `json:"value,omitempty" bson:"value,omitempty"`
}

type Snippet struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	ProjectID primitive.ObjectID `json:"project_id,omitempty" bson:"project_id,omitempty" binding:"required"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty" binding:"required"`
	Type      string             `json:"type,omitempty" bson:"type,omitempty"`
	ChartType string             `json:"chart_type,omitempty" bson:"chart_type,omitempty"`
	Value     float64            `json:"value,omitempty" bson:"value,omitempty"`
	Time      primitive.DateTime `json:"time,omitempty" bson:"time,omitempty"`
	Symbol    string             `json:"symbol,omitempty" bson:"symbol,omitempty"`
	Data      []SnippetItem      `json:"data,omitempty" bson:"data,omitempty" binding:"required"`
}
