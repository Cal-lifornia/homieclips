package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Upload struct {
	ObjectName string    `json:"object_name,omitempty" bson:"object_name"`
	Ready      bool      `json:"ready" bson:"ready"`
	Logs       []string  `json:"logs,omitempty" bson:"logs,omitempty"`
	Errors     []string  `json:"errors,omitempty" bson:"errors,omitempty"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}

func (models *Models) CreateUpload(upload Upload) (*mongo.InsertOneResult, error) {
	result, err := uploadsCollection.InsertOne(context.Background(), upload)
	if err != nil {
		return nil, err
	}

	return result, err
}
