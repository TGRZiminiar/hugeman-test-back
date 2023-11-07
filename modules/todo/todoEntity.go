package todo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Todo struct {
		Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Title       string             `json:"title" bson:"title"`
		Description string             `json:"description" bson:"description"`
		CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
		Image       string             `json:"image" bson:"image"`
		Status      string             `json:"status" bson:"status"`
	}
)
