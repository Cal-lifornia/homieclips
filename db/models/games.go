package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Game struct {
	ApiID    string `json:"api_id" bson:"api_id"`
	Name     string `json:"name" bson:"name"`
	CoverId  string `json:"cover_id" bson:"cover_id"`
	CoverUrl string `json:"cover_url" bson:"cover_url"`
}

func (models *Models) CreateGame(game Game) error {
	_, err := gamesCollection.InsertOne(context.TODO(), game)
	if err != nil {
		return err
	}

	return nil
}

func (models *Models) GetGames() ([]Game, error) {
	results, err := gamesCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	var games []Game
	err = results.All(context.TODO(), &games)
	if err != nil {
		return nil, err
	}

	return games, err
}

func (models *Models) GetGame(gameName string) (Game, error) {
	var game Game
	filter := bson.D{{Key: "object_name", Value: gameName}}
	err := gamesCollection.FindOne(context.TODO(), filter).Decode(&game)
	if err != nil {
		return Game{}, err
	}

	return game, nil
}

func (models *Models) DeleteGame(gameName string) error {
	filter := bson.D{{Key: "object_name", Value: gameName}}
	_, err := gamesCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}
