package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Clip struct {
	ObjectName   string    `json:"object_name,omitempty" bson:"object_name"`
	FriendlyName string    `json:"friendly_name,omitempty" bson:"friendly_name"`
	GameName     string    `json:"game_name,omitempty" bson:"game_name"`
	Ready        bool      `json:"ready,omitempty" bson:"ready,omitempty"`
	UploadedBy   string    `json:"uploaded_by,omitempty" bson:"uploaded_by,omitempty"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}

func (models *Models) CreateClip(clip Clip) (*mongo.InsertOneResult, error) {
	result, err := clipsCollection.InsertOne(context.TODO(), clip)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (models *Models) GetClips() ([]Clip, error) {
	results, err := clipsCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	var clips []Clip
	err = results.All(context.TODO(), &clips)
	if err != nil {
		return nil, err
	}

	return clips, err
}

func (models *Models) GetClip(objectName string) (Clip, error) {
	var clip Clip
	filter := bson.D{{Key: "object_name", Value: objectName}}
	err := clipsCollection.FindOne(context.TODO(), filter).Decode(&clip)
	if err != nil {
		return Clip{}, err
	}

	return clip, nil
}

func (models *Models) DeleteClip(objectName string) error {
	filter := bson.D{{Key: "object_name", Value: objectName}}
	_, err := clipsCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}
