package db

import "go.mongodb.org/mongo-driver/mongo"

type Models struct {
	db *mongo.Client
}

var recordingsCollection *mongo.Collection

func New(mongoClient *mongo.Client, dbName string) *Models {
	newModels := &Models{
		db: mongoClient,
	}

	newModels.init(dbName)
	return newModels
}

func (models *Models) init(dbName string) {
	recordingsCollection = models.db.Database(dbName).Collection("recordings")
}
