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

func (m *Models) init(dbName string) {
	recordingsCollection = m.db.Database(dbName).Collection("recordings")
}
