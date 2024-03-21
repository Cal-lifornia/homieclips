package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Recording struct {
	ObjectName   string    `json:"object_name,omitempty" bson:"object_name"`
	FriendlyName string    `json:"friendly_name,omitempty" bson:"friendly_name"`
	GameName     string    `json:"game_name,omitempty" bson:"game_name"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}

func (models *Models) CreateRecording(recording Recording) (*mongo.InsertOneResult, error) {
	result, err := recordingsCollection.InsertOne(context.TODO(), recording)
	if err != nil {
		return nil, err
	}

	return result, err
}
