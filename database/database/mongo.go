package database

import (
	"context"

	"skyzar-database/constants"
	"skyzar-database/logging"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client

func StartClient() {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(constants.MongoURI),
	)
	if err != nil {
		logging.Fatal("Failed to initialize mongo client")
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		logging.Fatal("Failed to ping mongo client, error: " + err.Error())
	}
	logging.Info("Successfully initialized and pinged mongo client")
	MongoClient = client
}