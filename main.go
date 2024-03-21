package main

import (
	"context"
	"fmt"
	"homieclips/api"
	db "homieclips/db/models"
	"homieclips/util"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed loading config: %s\n", err)
	}

	dbConnString := fmt.Sprintf("mongodb://%s:%s@%s", config.MongoUsername, config.MongoPass, config.DbAddress)

	dbClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbConnString))
	if err != nil {
		log.Fatalf("ran into error connecting to mongo instance %s\n", err)
	}

	minioClient, err := setupMinio(&config)
	if err != nil {
		log.Fatalf("ran into error connecting to minio: %s\n", err)
	}

	models := db.New(dbClient, config.DbName)

	mainApp := api.NewServer(config, models, minioClient)

	mainApp.Start(":8080")
}

func setupMinio(config *util.Config) (*minio.Client, error) {
	minioClient, err := minio.New(config.MinioURL, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinioAccessKey, config.MinioSecretKey, ""),
		Secure: true,
	})
	if err != nil {
		return nil, err
	}

	exists, errBucketExists := minioClient.BucketExists(context.TODO(), config.BucketName)
	if errBucketExists == nil && exists {
		log.Printf("bucket exists and is ours")
		return minioClient, nil
	}

	return nil, errBucketExists
}
