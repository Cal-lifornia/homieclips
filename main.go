package main

import (
	"context"
	"fmt"
	"homieclips/app"
	db "homieclips/db/models"
	"homieclips/storage"
	"homieclips/util"
	"log"
	"time"

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

	dbCtx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)

	dbClient, err := mongo.Connect(dbCtx, options.Client().ApplyURI(dbConnString))
	if err != nil {
		cancelFunc()
		log.Fatalf("ran into error connecting to mongo instance %s\n", err)
	}

	minioClient, err := setupMinio(dbCtx, &config)
	if err != nil {
		cancelFunc()
		log.Fatalf("ran into error connecting to minio: %s\n", err)
	}

	dbCtx.Done()

	models := db.New(dbClient, config.DbName)

	storageClient := storage.New(minioClient, config)

	mainApp := app.NewServer(config, models, storageClient)

	err = mainApp.Start("localhost:8080")
	if err != nil {
		return
	}
}

func setupMinio(ctx context.Context, config *util.Config) (*minio.Client, error) {
	minioClient, err := minio.New(config.MinioURL, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinioAccessKey, config.MinioSecretKey, ""),
		Secure: true,
	})
	if err != nil {
		return nil, err
	}

	exists, errBucketExists := minioClient.BucketExists(ctx, config.BucketName)
	if errBucketExists == nil && exists {
		log.Printf("bucket exists and is ours")
		return minioClient, nil
	}

	return nil, errBucketExists
}
