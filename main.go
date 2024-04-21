package main

import (
	"context"
	"fmt"
	"github.com/Cal-lifornia/homieclips/app"
	db "github.com/Cal-lifornia/homieclips/db/models"
	"github.com/Cal-lifornia/homieclips/util"
	"log"
	"time"

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

	dbCtx.Done()

	models := db.New(dbClient, config.DbName)

	mainApp := app.NewServer(config, models)

	err = mainApp.Start("localhost:8080")
	if err != nil {
		return
	}
}
