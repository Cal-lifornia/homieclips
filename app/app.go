package app

import (
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"homieclips/util"
)

type App struct {
	config      *util.Config
	dbClient    *mongo.Client
	minioClient *minio.Client
}

func New(config *util.Config, dbClient *mongo.Client, minioClient *minio.Client) *App {
	return &App{
		config:      config,
		dbClient:    dbClient,
		minioClient: minioClient,
	}
}
